/*
Copyright The Kubernetes Authors.

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

// KerberusSpec defines the desired state of Kerberus operator
// +k8s:openapi-gen=true
type KerberusSpec struct {
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

// KerberusStatus defines the observed state of Kerberus
// +k8s:openapi-gen=true
type KerberusStatus struct {
	// List of conditions for the CR
	Conditions []Condition `json:"conditions,omitempty"`
	// State that this condition represents
	State string `json:"state"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Kerberus is the Schema for the kerberus API
// +k8s:openapi-gen=true
// +genclient
// +kubebuilder:subresource:status
type Kerberus struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KerberusSpec   `json:"spec,omitempty"`
	Status KerberusStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KerberusList contains a list of Kerberus
type KerberusList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Kerberus `json:"items"`
}
