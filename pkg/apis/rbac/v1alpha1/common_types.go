package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// Condition defines a single condition of running the operator against an instance of the GroupPermission CR
type Condition struct {
	// LastTransitionTime is the last time this condition was active for the CR
	LastTransitionTime metav1.Time `json:"lastTransitionTime"`
	// Message related to the condition
	// +optional
	Message string `json:"message,omitempty"`
	// Flag to indicate if condition status is currently active
	Status bool `json:"status"`
	// State that this condition represents
	State State `json:"state"`
}

// State defines various states
type State string

const (
	// StateAccepted const for StateAccepted status
	StateAccepted State = "Accepted"
	// StateFailed const for Failed status
	StateFailed State = "Failed"
)
