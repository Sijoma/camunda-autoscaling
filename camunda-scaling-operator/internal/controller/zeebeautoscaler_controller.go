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
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
)

// ZeebeAutoscalerReconciler reconciles a ZeebeAutoscaler object
type ZeebeAutoscalerReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=camunda.sijoma.dev,resources=zeebeautoscalers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=camunda.sijoma.dev,resources=zeebeautoscalers/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=camunda.sijoma.dev,resources=zeebeautoscalers/finalizers,verbs=update

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
	logger.Info("starting reconcile", "name", zeebeAutoscalerCR.Name)

	// 1. Lookup statefulset
	var scaleTarget appsv1.StatefulSet
	err := r.Get(ctx, types.NamespacedName{
		Name:      zeebeAutoscalerCR.Spec.ZeebeRef.Name,
		Namespace: zeebeAutoscalerCR.Namespace,
	}, &scaleTarget)
	if err != nil {
		return ctrl.Result{}, err
	}

	// Todo: maybe check the topology
	//currentTopology, err := QueryTopology(port)
	//ensureNoError(err)
	//if currentTopology.PendingChange != nil {
	//	return fmt.Errorf("cluster is already scaling")
	//}

	// 2. check if it matches the desired replicas
	// TODO: you cannot scale less than your replicationFactor
	if scaleTarget.Spec.Replicas != zeebeAutoscalerCR.Spec.Replicas {
		sts := &appsv1.StatefulSet{ObjectMeta: metav1.ObjectMeta{
			Name:      zeebeAutoscalerCR.Spec.ZeebeRef.Name,
			Namespace: zeebeAutoscalerCR.Namespace}}

		scale := &autoscalingv1.Scale{}
		err = r.SubResource("scale").Get(ctx, sts, scale)
		if err != nil {
			return ctrl.Result{}, err
		}
		scale = &autoscalingv1.Scale{Spec: autoscalingv1.ScaleSpec{Replicas: *zeebeAutoscalerCR.Spec.Replicas}}
		err = r.SubResource("scale").Update(ctx, sts, client.WithSubResourceBody(scale))
		if err != nil {
			return ctrl.Result{}, err
		}

		// 3. Request change to a gateway (only standalone gateway supported)
		// Get GW service
		gwSvc := corev1.Service{}
		err = r.Get(ctx, types.NamespacedName{
			Name:      zeebeAutoscalerCR.Spec.ZeebeRef.GatewayServiceName,
			Namespace: zeebeAutoscalerCR.Namespace,
		}, &gwSvc)
		if err != nil {
			return ctrl.Result{}, err
		}

		// Make list of broker IDs
		// Todo: get from state / cluster
		brokerIDs := []int32{1, 3, 3}
		// if no operation ongoing: Call send scale request
		err = sendScaleRequest(ctx, gwSvc, brokerIDs)
		if err != nil {
			return ctrl.Result{}, err
		}

		// if operation ongoing: Get current status

		return ctrl.Result{RequeueAfter: time.Second * 30}, nil

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

	return ctrl.Result{
		RequeueAfter: time.Second * 20,
	}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ZeebeAutoscalerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&camundav1alpha1.ZeebeAutoscaler{}).
		Complete(r)
}

func sendScaleRequest(ctx context.Context, gwSvc corev1.Service, brokerIds []int32) error {
	logger := log.FromContext(ctx)

	url := fmt.Sprintf("http://%s.%s:%d/actuator/cluster/brokers", gwSvc.Name, gwSvc.Namespace, gwSvc.Spec.Ports[1].Port)

	request, err := json.Marshal(brokerIds)
	if err != nil {
		return err
	}
	resp, err := http.Post(url, "application/json", bytes.NewReader(request))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 400 {
		return fmt.Errorf("sendScaleRequest: scaling failed with code %d", resp.StatusCode)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	logger.Info("sendScaleRequest: scaling succeeded", respBody)

	//var changeResponse ChangeResponse
	//err = json.Unmarshal(response, &changeResponse)
	//if err != nil {
	//	return nil, err
	//}
	return nil
}
