package model

type (
	// Config is the model for all configurations
	Config struct {
		HTTPPort  int `json:"httpPort" yaml:"http_port"`
		HTTPSPort int `json:"httpsPort" yaml:"https_port"`
	}

	// Spec is the model for spec file
	Spec struct {
		Config    Config     `json:"config" yaml:"config"`
		HTTPMocks []HTTPMock `json:"http" yaml:"http"`
		RESTMock  []RESTMock `json:"rest" yaml:"rest"`
	}
)
