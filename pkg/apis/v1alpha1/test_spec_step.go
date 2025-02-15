package v1alpha1

// TestSpecStep contains the test step definition used in a test spec.
type TestSpecStep struct {
	// Name of the step.
	// +optional
	Name string `json:"name,omitempty"`

	// Spec of the step.
	Spec TestStepSpec `json:"spec"`
}
