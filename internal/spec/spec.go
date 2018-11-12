package spec

type (
	// JSON is the model for json objects
	JSON map[string]interface{}

	// Store is the model for objects store
	Store map[string][]JSON

	// Spec is the model for mocking specifications
	Spec struct {
		Store Store `json:"store"`
	}
)
