package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// NyanCatMinerSpec defines the desired state of NyanCatMiner
// +k8s:openapi-gen=true
type NyanCatMinerSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// NyanCatMinerStatus defines the observed state of NyanCatMiner
// +k8s:openapi-gen=true
type NyanCatMinerStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
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
