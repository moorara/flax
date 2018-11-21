package model

type (
	// JSON is the type for json objects
	JSON map[string]interface{}

	// Pair is a name-value pair
	Pair struct {
		Name  string `json:"name" yaml:"name"`
		Value string `json:"value" yaml:"value"`
	}
)
