package model

type (
	// JSON is the model for json objects
	JSON map[string]interface{}

	// RESTFilters is the model for filtering RESTful requests
	RESTFilters struct {
		Headers map[string]string `json:"headers" yaml:"headers"`
	}

	// HTTPMock is the model for an http mock
	HTTPMock struct {
		HTTPExpectation `json:",inline" yaml:",inline"`
		Delay           string            `json:"delay" yaml:"delay"`
		StatusCode      int               `json:"statusCode" yaml:"status_code"`
		Headers         map[string]string `json:"headers" yaml:"headers"`
		Body            interface{}       `json:"body" yaml:"body"`
	}

	// RESTMock is the model for a RESTful resource mock
	RESTMock struct {
		Name       string            `json:"name" yaml:"name"`
		BasePath   string            `json:"basePath" yaml:"base_path"`
		Delay      string            `json:"delay" yaml:"delay"`
		Identifier string            `json:"identifier" yaml:"identifier"`
		ListHandle string            `json:"listHandle" yaml:"list_handle"`
		Filters    RESTFilters       `json:"filters" yaml:"filters"`
		Headers    map[string]string `json:"headers" yaml:"headers"`
		Store      []JSON            `json:"store" yaml:"store"`
	}
)
