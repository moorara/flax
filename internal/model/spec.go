package model

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
		RESTMock  []RESTMock `json:"rest" yaml:"rest"`
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

	if ext == ".json" {
		err = json.NewDecoder(f).Decode(spec)
	} else if ext == ".yaml" || ext == ".yml" {
		err = yaml.NewDecoder(f).Decode(spec)
	} else {
		err = fmt.Errorf("unknown file format %s", ext)
	}

	if err != nil {
		return nil, err
	}

	for i := range spec.HTTPMocks {
		spec.HTTPMocks[i] = spec.HTTPMocks[i].withDefaults()
	}

	for i := range spec.RESTMock {
		spec.RESTMock[i] = spec.RESTMock[i].withDefaults()
	}

	return spec, nil
}
