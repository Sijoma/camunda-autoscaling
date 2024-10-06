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
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ZeebeAutoscalerSpec defines the desired state of ZeebeAutoscaler
type ZeebeAutoscalerSpec struct {
	// Replicas the number of Zeebe brokers to deploy
	Replicas *int32 `json:"replicas,omitempty"`

	// +kubebuilder:default={}
	ZeebeRef ZeebeRef `json:"zeebeRef,omitempty"`
}

// ZeebeRef references that exists in the same namespace.
type ZeebeRef struct {
	// Name of the Zeebe statefulset to scale
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:default=camunda-platform-zeebe
	Name string `json:"name,omitempty"`
	// +kubebuilder:default={}
	Gateway Gateway `json:"gateway,omitempty"`
}

type Gateway struct {
	// +kubebuilder:default=camunda-platform-zeebe-gateway
	// ServiceName of the zeebe-gateway, this is used to trigger scaling operations & request topology information
	ServiceName string `json:"serviceName,omitempty"`
	// +kubebuilder:default=9600
	// Port of the zeebe-gateway, needs to expose the management API
	Port int32 `json:"port,omitempty"`
}

// ZeebeAutoscalerStatus defines the observed state of ZeebeAutoscaler
type ZeebeAutoscalerStatus struct {
	// ObservedGeneration is the last observed generation by the controller.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// Conditions holds the information on the last operations on Zeebe that can be useful during scaling
	// +kubebuilder:validation:Optional
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,1,rep,name=conditions"`

	Replicas int32  `json:"replicas"`
	Selector string `json:"selector"`
}

// ConditionReason defines the reason why a certain condition changed
type ConditionReason string

const (
	ConditionZeebeTopologyFound    ConditionReason = "ZeebeTopologyFound"
	ConditionZeebePendingOperation ConditionReason = "ZeebePendingOperation"
)

type ScalingCondition string

const ReadyToScale ScalingCondition = "ReadyToScale"

func ZeebePendingOperations(brokerCount int) metav1.Condition {
	return metav1.Condition{
		Type:    string(ReadyToScale),
		Status:  metav1.ConditionTrue,
		Reason:  string(ConditionZeebeTopologyFound),
		Message: fmt.Sprintf("Zeebe Topology queried. Found %d Brokers.", brokerCount),
	}
}

func ZeebePendingTopologyChange(status string) metav1.Condition {
	return metav1.Condition{
		Type:    string(ReadyToScale),
		Status:  metav1.ConditionFalse,
		Reason:  string(ConditionZeebePendingOperation),
		Message: fmt.Sprintf("Topology change pending: Status: %s", status),
	}
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:subresource:scale:specpath=.spec.replicas,statuspath=.status.replicas,selectorpath=.status.selector
// +kubebuilder:printcolumn:name="Desired Replicas",type=string,JSONPath=`.spec.replicas`
// +kubebuilder:printcolumn:name="Current Replicas",type=string,JSONPath=`.status.replicas`
// +kubebuilder:printcolumn:name="Ready To Scale",type=string,JSONPath=`.status.conditions[?(@.type=='ReadyToScale')].status`
// +kubebuilder:printcolumn:name="Target",type=string,JSONPath=`.spec.zeebeRef.name`

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
