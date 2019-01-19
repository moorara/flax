package v1

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

type (
	// Config is the model for mock server configurations
	Config struct {
		HTTPPort  int `json:"httpPort" yaml:"http_port"`
		HTTPSPort int `json:"httpsPort" yaml:"https_port"`
	}

	// Spec is the model for specification files
	Spec struct {
		Config    Config     `json:"config" yaml:"config"`
		HTTPMocks []HTTPMock `json:"http" yaml:"http"`
		RESTMocks []RESTMock `json:"rest" yaml:"rest"`
	}
)

// ReadSpec reads and returns a Spec from a YAML/JSON file
func ReadSpec(path string) (*Spec, error) {
	spec := new(Spec)

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	ext := strings.ToLower(filepath.Ext(path))
	switch ext {
	case ".json":
		err = json.NewDecoder(f).Decode(spec)
	case ".yml", ".yaml":
		err = yaml.NewDecoder(f).Decode(spec)
	default:
		err = fmt.Errorf("unknown file format %s", ext)
	}

	if err != nil {
		return nil, err
	}

	for i := range spec.HTTPMocks {
		spec.HTTPMocks[i].SetDefaults()
	}

	for i := range spec.RESTMocks {
		spec.RESTMocks[i].SetDefaults()
	}

	return spec, nil
}
