package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NyanCatMinerSpec defines the desired state of NyanCatMiner
// +k8s:openapi-gen=true
type NyanCatMinerSpec struct {
	Replicas int32 `json:"replicas"`
}

// NyanCatMinerStatus defines the observed state of NyanCatMiner
// +k8s:openapi-gen=true
type NyanCatMinerStatus struct {
	PodNames []string `json:"podNames"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NyanCatMiner is the Schema for the nyancatminers API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
type NyanCatMiner struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NyanCatMinerSpec   `json:"spec,omitempty"`
	Status NyanCatMinerStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NyanCatMinerList contains a list of NyanCatMiner
type NyanCatMinerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NyanCatMiner `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NyanCatMiner{}, &NyanCatMinerList{})
}
