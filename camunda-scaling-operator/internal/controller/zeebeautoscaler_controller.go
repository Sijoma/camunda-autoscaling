/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"
	"fmt"
	"time"

	appsv1 "k8s.io/api/apps/v1"
	autoscalingv1 "k8s.io/api/autoscaling/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	camundav1alpha1 "github.com/sijoma/camunda-scaling-operator/api/v1alpha1"
	"github.com/sijoma/camunda-scaling-operator/pkg/scalingclient"
)

// ZeebeAutoscalerReconciler reconciles a ZeebeAutoscaler object
type ZeebeAutoscalerReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=camunda.sijoma.dev,resources=zeebeautoscalers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=camunda.sijoma.dev,resources=zeebeautoscalers/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=camunda.sijoma.dev,resources=zeebeautoscalers/finalizers,verbs=update

// +kubebuilder:rbac:groups=core,resources=services,verbs=get;list;watch

// +kubebuilder:rbac:groups=apps,resources=statefulsets,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=apps,resources=statefulsets/status,verbs=get
// +kubebuilder:rbac:groups=apps,resources=statefulsets/scale,verbs=get;update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.18.4/pkg/reconcile
func (r *ZeebeAutoscalerReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	// populate this CRD
	zeebeAutoscalerCR := new(camundav1alpha1.ZeebeAutoscaler)

	if err := r.Get(ctx, req.NamespacedName, zeebeAutoscalerCR); err != nil {
		// do not requeue "not found" errors
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	logger = logger.WithValues("name", zeebeAutoscalerCR.Name, "namespace", zeebeAutoscalerCR.Namespace)
	logger.Info("starting reconcile")

	// 1. Lookup statefulset
	var scaleTarget appsv1.StatefulSet
	err := r.Get(ctx, types.NamespacedName{
		Name:      zeebeAutoscalerCR.Spec.ZeebeRef.Name,
		Namespace: zeebeAutoscalerCR.Namespace,
	}, &scaleTarget)
	if err != nil {
		return ctrl.Result{}, err
	}

	// 2. Lookup zeebe-gateway
	gwSvc := corev1.Service{}
	err = r.Get(ctx, types.NamespacedName{
		//Name:      zeebeAutoscalerCR.Spec.ZeebeRef.GatewayServiceName,
		Namespace: zeebeAutoscalerCR.Namespace,
		Name:      "camunda-platform-zeebe-gateway",
	}, &gwSvc)
	if err != nil {
		return ctrl.Result{}, err
	}

	// Prepare ZeebeClient
	zeebeClient := scalingclient.NewZeebeMgmtClient(gwSvc)

	// Check topology
	topology, err := zeebeClient.Topology(ctx)
	if err != nil {
		return ctrl.Result{}, err
	}

	// Check if change is already pending
	if topology.HasPendingChange() {
		logger.Info("cluster scaling in progress", "topology", topology.PendingChange.Status)
		// Todo: update status
		return ctrl.Result{}, fmt.Errorf("cluster scaling in progress")
	}

	// 2. check if we need to downscale / upscale
	// (https://docs.camunda.io/docs/next/self-managed/zeebe-deployment/operations/cluster-scaling/)
	currentReplicas := *scaleTarget.Spec.Replicas
	desiredReplicas := *zeebeAutoscalerCR.Spec.Replicas
	logger = logger.WithValues("currentReplicas", currentReplicas, "desiredReplicas", desiredReplicas)

	// Check if we already scaled down brokers, if so, we can scale down the statefulset
	// In words: "we are downscaling" && the zeebe topology has already removed the broker
	if currentReplicas > desiredReplicas && len(topology.Brokers) < int(currentReplicas) {
		if err := r.scaleStatefulSet(ctx, zeebeAutoscalerCR); err != nil {
			return ctrl.Result{}, err
		}
		return ctrl.Result{RequeueAfter: time.Second * 5}, nil
	}

	if currentReplicas < desiredReplicas {
		logger.Info("We are scaling stateful set up! ⬆️️")
		if err := r.scaleStatefulSet(ctx, zeebeAutoscalerCR); err != nil {
			return ctrl.Result{}, err
		}

		// re-enqueue until the stateful set is scaled UP
		return ctrl.Result{RequeueAfter: time.Second * 5}, nil
	}

	if len(topology.Brokers) < int(currentReplicas) {
		logger.Info("We are scaling topology up! ⬆️️")
		if err := r.scaleZeebeBrokers(ctx, zeebeClient, desiredReplicas); err != nil {
			return ctrl.Result{}, err
		}

		// re-enqueue until the topology is scaled UP
		return ctrl.Result{RequeueAfter: time.Second * 5}, nil
	}

	if currentReplicas > desiredReplicas {
		logger.Info("we are scaling down!⬇️")

		// we didnt yet request that Zeebe should scale down
		if err := r.scaleZeebeBrokers(ctx, zeebeClient, desiredReplicas); err != nil {
			return ctrl.Result{}, err
		}
		return ctrl.Result{RequeueAfter: time.Second * 5}, nil
	}

	// Refresh CR to prevent status update errors
	if err := r.Get(ctx, req.NamespacedName, zeebeAutoscalerCR); err != nil {
		// do not requeue "not found" errors
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	// Get the selector, this is important for HPA to work
	// https://book.kubebuilder.io/reference/generating-crd.html#scale
	selector, err := metav1.LabelSelectorAsSelector(scaleTarget.Spec.Selector)
	if err != nil {
		logger.Error(err, "Error retrieving statefulset selector for scale ")
		return ctrl.Result{}, err
	}
	zeebeAutoscalerCR.Status.Selector = selector.String()
	zeebeAutoscalerCR.Status.Replicas = *zeebeAutoscalerCR.Spec.Replicas

	err = r.Status().Update(ctx, zeebeAutoscalerCR)
	if err != nil {
		logger.Error(err, "Error updating ZeebeAutoscaler CR status")
		return ctrl.Result{}, err
	}
	logger.Info("reconcile success", "name", zeebeAutoscalerCR.Name)

	return ctrl.Result{}, nil
}

func (r *ZeebeAutoscalerReconciler) scaleZeebeBrokers(ctx context.Context, zeebeClient *scalingclient.ZeebeMgmtClient, desiredReplicas int32) error {
	desiredBrokerIDs := make([]int32, 0, desiredReplicas)
	for id := range desiredReplicas {
		desiredBrokerIDs = append(desiredBrokerIDs, id)
	}

	err := zeebeClient.SendScaleRequest(ctx, desiredBrokerIDs)
	if err != nil {
		return err
	}
	return nil
}

func (r *ZeebeAutoscalerReconciler) scaleStatefulSet(ctx context.Context, cr *camundav1alpha1.ZeebeAutoscaler) error {
	sts := &appsv1.StatefulSet{ObjectMeta: metav1.ObjectMeta{
		Name:      cr.Spec.ZeebeRef.Name,
		Namespace: cr.Namespace}}

	scale := &autoscalingv1.Scale{Spec: autoscalingv1.ScaleSpec{Replicas: *cr.Spec.Replicas}}
	err := r.SubResource("scale").Update(ctx, sts, client.WithSubResourceBody(scale))
	if err != nil {
		return err
	}
	return nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ZeebeAutoscalerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&camundav1alpha1.ZeebeAutoscaler{}).
		Complete(r)
}
