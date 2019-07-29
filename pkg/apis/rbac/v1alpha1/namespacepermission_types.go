package v1alpha1

import (
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NamespacePermissionSpec defines the desired state of NamespacePermission
// +k8s:openapi-gen=true
type NamespacePermissionSpec struct {
	Subjects []rbacv1.Subject `json:"subjects,omitempty"`
	RoleRef  rbacv1.RoleRef   `json:"roleRef,omitempty"`
}

// NamespacePermissionStatus defines the observed state of NamespacePermission
// +k8s:openapi-gen=true
type NamespacePermissionStatus struct {
	// List of conditions for the CR
	Conditions []Condition `json:"conditions,omitempty"`
	// State that this condition represents
	State string `json:"state"`
}

// NamespacePermission is the Schema for the namespacepermissions API

// CRDRequest is the Schema for the crdrequests API
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +genclient
// +kubebuilder:subresource:status
type NamespacePermission struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NamespacePermissionSpec   `json:"spec,omitempty"`
	Status NamespacePermissionStatus `json:"status,omitempty"`
}

// NamespacePermissionList contains a list of NamespacePermission
// CRDRequestList contains a list of CRDRequest
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type NamespacePermissionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NamespacePermission `json:"items"`
}
