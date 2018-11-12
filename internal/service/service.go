package service

import (
	"encoding/json"
	"os"

	"github.com/moorara/flax/internal/spec"
)

type (
	// RESTService is the interface for a RESTful service
	RESTService interface {
	}

	restService struct {
		Spec *spec.Spec
	}
)

// NewRESTService creates a new instance of RESTService
func NewRESTService(file string) (RESTService, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	spec := new(spec.Spec)
	err = json.NewDecoder(f).Decode(spec)
	if err != nil {
		return nil, err
	}

	return &restService{
		Spec: spec,
	}, nil
}
