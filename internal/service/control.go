package service

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/moorara/flax/internal/model"
	"github.com/moorara/flax/pkg/log"
	"github.com/moorara/flax/pkg/metrics"
	"github.com/opentracing/opentracing-go"
	yaml "gopkg.in/yaml.v2"
)

type (
	// ControlService is the interface for controller service
	ControlService interface {
	}

	controlService struct {
		logger  *log.Logger
		metrics *metrics.Metrics
		tracer  opentracing.Tracer
		spec    *model.Spec
	}
)

// NewControlService creates a new instance of ControlService
func NewControlService(logger *log.Logger, metrics *metrics.Metrics, tracer opentracing.Tracer) ControlService {
	return &controlService{
		logger:  logger,
		metrics: metrics,
		tracer:  tracer,
	}
}

func (s *controlService) ReadSpec(path string) error {
	spec := new(model.Spec)

	f, err := os.Open(path)
	if err != nil {
		return err
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
		return err
	}

	s.spec = spec
	return nil
}
