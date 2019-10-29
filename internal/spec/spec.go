package spec

import (
	"encoding/json"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

// Config has the specifications for mock server configurations.
type Config struct {
	HTTPPort  uint16 `json:"httpPort" yaml:"http_port"`
	HTTPSPort uint16 `json:"httpsPort" yaml:"https_port"`
}

// SetDefaults set default values for empty fields.
func (c *Config) SetDefaults() {
	if c.HTTPPort == 0 {
		c.HTTPPort = 8080
	}

	if c.HTTPSPort == 0 {
		c.HTTPSPort = 8443
	}
}

// Spec has all the specifications.
type Spec struct {
	Config    Config     `json:"config" yaml:"config"`
	HTTPMocks []HTTPMock `json:"http" yaml:"http"`
	RESTMocks []RESTMock `json:"rest" yaml:"rest"`
}

// DefaultSpec returns a default Spec.
func DefaultSpec() *Spec {
	config := Config{}
	config.SetDefaults()

	return &Spec{
		config,
		[]HTTPMock{
			DefaultHTTPMock(),
		},
		[]RESTMock{},
	}
}

// ReadSpec reads and returns a Spec from a JSON or YAML file.
// It returns a default spec if no spec file found.
func ReadSpec(path string) (*Spec, error) {
	f, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return DefaultSpec(), nil
		}
		return nil, err
	}
	defer f.Close()

	spec := new(Spec)
	if err := json.NewDecoder(f).Decode(spec); err != nil {
		f.Seek(0, 0) // Reset file offset
		if err := yaml.NewDecoder(f).Decode(spec); err != nil {
			return nil, fmt.Errorf("unknown spec file: %s", err)
		}
	}

	spec.Config.SetDefaults()

	for i := range spec.HTTPMocks {
		spec.HTTPMocks[i].SetDefaults()
	}

	for i := range spec.RESTMocks {
		spec.RESTMocks[i].SetDefaults()
		spec.RESTMocks[i].RESTStore.Index()
	}

	return spec, nil
}
