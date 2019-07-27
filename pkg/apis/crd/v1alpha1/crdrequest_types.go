package v1alpha1

import (
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CRDRequestSpec defines the desired state of CRDRequest
// +k8s:openapi-gen=true
type CRDRequestSpec struct {
	Owned   bool                                            `json:"owned,omitempty"`
	CRDList []apiextensionsv1beta1.CustomResourceDefinition `json:"crdList,omitempty"`
}

// CRDRequestStatus defines the observed state of CRDRequest
// +k8s:openapi-gen=true
type CRDRequestStatus struct {
	// List of conditions for the CR
	Conditions []Condition `json:"conditions,omitempty"`
	// State that this condition represents
	State string `json:"state"`
}

// CRDRequest is the Schema for the crdrequests API
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +genclient
// +kubebuilder:subresource:status
type CRDRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CRDRequestSpec   `json:"spec,omitempty"`
	Status CRDRequestStatus `json:"status,omitempty"`
}

// CRDRequestList contains a list of CRDRequest
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type CRDRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CRDRequest `json:"items"`
}
