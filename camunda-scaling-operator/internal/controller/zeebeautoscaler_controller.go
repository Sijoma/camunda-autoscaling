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

	appsv1 "k8s.io/api/apps/v1"
	autoscalingv1 "k8s.io/api/autoscaling/v1"
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
		// Should we actually continue to reconcile when the statefulset does not yet exist?
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// 2. check if it matches the desired replicas
	if scaleTarget.Spec.Replicas != zeebeAutoscalerCR.Spec.Replicas {

		// Todo: Should we really use scale?
		sts := &appsv1.StatefulSet{ObjectMeta: metav1.ObjectMeta{
			Name:      zeebeAutoscalerCR.Spec.ZeebeRef.Name,
			Namespace: zeebeAutoscalerCR.Namespace}}
		scale := &autoscalingv1.Scale{}
		err = r.SubResource("scale").Get(ctx, sts, scale)
		if err != nil {
			return ctrl.Result{}, err
		}
		scale = &autoscalingv1.Scale{Spec: autoscalingv1.ScaleSpec{Replicas: *zeebeAutoscalerCR.Spec.Replicas}}
		r.SubResource("scale").Update(ctx, sts, client.WithSubResourceBody(scale))

		// 3. Request change on broker
		// Todo:

	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ZeebeAutoscalerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&camundav1alpha1.ZeebeAutoscaler{}).
		Complete(r)
}
