package service

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/moorara/flax/internal/model"
	"github.com/moorara/flax/pkg/log"
	"github.com/moorara/flax/pkg/metrics"
)

type (
	// ControlService is the interface for controller service
	ControlService interface {
		AddHTTPMocks(w http.ResponseWriter, r *http.Request)
		AddRESTMocks(w http.ResponseWriter, r *http.Request)
		RemoveHTTPMocks(w http.ResponseWriter, r *http.Request)
		RemoveRESTMocks(w http.ResponseWriter, r *http.Request)
	}

	controlService struct {
		logger  *log.Logger
		metrics *metrics.Metrics
		routes  map[uint64]mux.Route
	}
)

// NewControlService creates a new instance of ControlService
func NewControlService(port string, logger *log.Logger, metrics *metrics.Metrics) ControlService {
	return &controlService{
		logger:  logger,
		metrics: metrics,
		routes:  map[uint64]mux.Route{},
	}
}

func (s *controlService) AddHTTPMocks(w http.ResponseWriter, r *http.Request) {
	mocks := []model.HTTPMock{}
	if err := json.NewDecoder(r.Body).Decode(mocks); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Add Routes

	w.WriteHeader(http.StatusOK)
}

func (s *controlService) AddRESTMocks(w http.ResponseWriter, r *http.Request) {
	mocks := []model.RESTMock{}
	if err := json.NewDecoder(r.Body).Decode(mocks); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Add Routes

	w.WriteHeader(http.StatusOK)
}

func (s *controlService) RemoveHTTPMocks(w http.ResponseWriter, r *http.Request) {
	mocks := []model.HTTPMock{}
	if err := json.NewDecoder(r.Body).Decode(mocks); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Delete Routes

	w.WriteHeader(http.StatusOK)
}

func (s *controlService) RemoveRESTMocks(w http.ResponseWriter, r *http.Request) {
	mocks := []model.RESTMock{}
	if err := json.NewDecoder(r.Body).Decode(mocks); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Delete Routes

	w.WriteHeader(http.StatusOK)
}
