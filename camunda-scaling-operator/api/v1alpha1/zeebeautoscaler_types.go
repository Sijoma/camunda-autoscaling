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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ZeebeAutoscalerSpec defines the desired state of ZeebeAutoscaler
type ZeebeAutoscalerSpec struct {
	MinReplicas *int32 `json:"minReplicas,omitempty"`
	MaxReplicas *int32 `json:"maxReplicas,omitempty"`

	// +kubebuilder:validation:Required
	ZeebeRef ZeebeRef `json:"zeebeRef"`
}

// ZeebeRef references that exists in the same namespace.
type ZeebeRef struct {
	// Name of the Zeebe statefulset to scale
	// +kubebuilder:validation:MinLength=1
	Name string `json:"name,omitempty"`
}

// ZeebeAutoscalerStatus defines the observed state of ZeebeAutoscaler
type ZeebeAutoscalerStatus struct {
	// ObservedGeneration is the last observed generation by the controller.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// Conditions holds the information on the last operations on Zeebe that can be useful during scaling
	// +kubebuilder:validation:Optional
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,1,rep,name=conditions"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// ZeebeAutoscaler is the Schema for the zeebeautoscalers API
type ZeebeAutoscaler struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ZeebeAutoscalerSpec   `json:"spec,omitempty"`
	Status ZeebeAutoscalerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ZeebeAutoscalerList contains a list of ZeebeAutoscaler
type ZeebeAutoscalerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ZeebeAutoscaler `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ZeebeAutoscaler{}, &ZeebeAutoscalerList{})
}
