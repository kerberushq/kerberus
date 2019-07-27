package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CRDConfigSpec defines the desired state of CRDConfig
// +k8s:openapi-gen=true
type CRDConfigSpec struct {
	// only AllowAll or denyAll can be set to true
	// if both/or none are set - denyAll is default

	// if allowAll is set, we allow all but deny list
	AllowAll bool     `json:"allowAll,omitempty"`
	Deny     []string `json:"deny,omitempty"`

	// if denyAll is set, we deny all but allow list
	DenyAll bool     `json:"denyAll,omitempty"`
	Allow   []string `json:"allow,omitempty"`
}

// CRDConfigStatus defines the observed state of CRDConfig
// +k8s:openapi-gen=true
type CRDConfigStatus struct {
	// List of conditions for the CR
	Conditions []Condition `json:"conditions,omitempty"`
	// State that this condition represents
	State string `json:"state"`
}

// CRDConfig is the Schema for the crdconfigs API
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +genclient
// +kubebuilder:subresource:status
type CRDConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CRDConfigSpec   `json:"spec,omitempty"`
	Status CRDConfigStatus `json:"status,omitempty"`
}

// CRDConfigList contains a list of CRDConfig
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type CRDConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CRDConfig `json:"items"`
}

//func init() {
//	SchemeBuilder.Register(&CRDConfig{}, &CRDConfigList{})
//}
