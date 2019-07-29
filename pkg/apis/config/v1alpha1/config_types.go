package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ConfigSpec defines the desired state of ConfigSpec
// +k8s:openapi-gen=true
type ConfigSpec struct {
	CRDRequest CRDRequest `json:"crdRequest,omitempty"`
	RBACConfig RBACConfig `json:"rbacConfig,omitempty"`
}

// RBACConfig defines configuration for RBAC operator
// +k8s:openapi-gen=true
type RBACConfig struct {
	Enable bool `json:"enable,omitempty"`

	WhiteList []string `json:"whiteList,omitempty"`
	BlackList []string `json:"blackList,omitempty"`
}

// CRDRequest defines configuration for CRDRequest operator
// +k8s:openapi-gen=true
type CRDRequest struct {
	Enable bool `json:"enable,omitempty"`
	// if allowAll is set, we allow all but deny list
	AllowAll bool     `json:"allowAll,omitempty"`
	Deny     []string `json:"deny,omitempty"`

	// if denyAll is set, we deny all but allow list
	DenyAll bool     `json:"denyAll,omitempty"`
	Allow   []string `json:"allow,omitempty"`
}

// ConfigSpecStatus defines the observed state of ConfigSpec
// +k8s:openapi-gen=true
type ConfigSpecStatus struct {
	// List of conditions for the CR
	Conditions []Condition `json:"conditions,omitempty"`
	// State that this condition represents
	State string `json:"state"`
}

// ConfigSpec is the Schema for the ConfigSpecs API
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +genclient
// +kubebuilder:subresource:status
type Config struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ConfigSpec       `json:"spec,omitempty"`
	Status ConfigSpecStatus `json:"status,omitempty"`
}

// ConfigList contains a list of ConfigSpec
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Config `json:"items"`
}
