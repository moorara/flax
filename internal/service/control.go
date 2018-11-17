package service

import (
	"encoding/json"
	"net/http"

	"github.com/moorara/flax/internal/model"
	"github.com/moorara/flax/pkg/log"
	"github.com/moorara/flax/pkg/metrics"
)

type (
	// ControlService is the interface for controller service
	ControlService interface {
		AddHTTPMock(w http.ResponseWriter, r *http.Request)
		AddRESTMock(w http.ResponseWriter, r *http.Request)
		RemoveHTTPMock(w http.ResponseWriter, r *http.Request)
		RemoveRESTMock(w http.ResponseWriter, r *http.Request)
	}

	controlService struct {
		logger       *log.Logger
		metrics      *metrics.Metrics
		expectations map[uint64]model.HTTPExpectation
	}
)

// NewControlService creates a new instance of ControlService
func NewControlService(port string, logger *log.Logger, metrics *metrics.Metrics) ControlService {
	return &controlService{
		logger:       logger,
		metrics:      metrics,
		expectations: map[uint64]model.HTTPExpectation{},
	}
}

func (s *controlService) AddHTTPMock(w http.ResponseWriter, r *http.Request) {
	mock := new(model.HTTPMock)
	if err := json.NewDecoder(r.Body).Decode(mock); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for _, he := range mock.HTTPExpectations() {
		hash := he.GetHash()
		s.expectations[hash] = he
	}

	w.WriteHeader(http.StatusOK)
}

func (s *controlService) AddRESTMock(w http.ResponseWriter, r *http.Request) {
	mock := new(model.RESTMock)
	if err := json.NewDecoder(r.Body).Decode(mock); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for _, he := range mock.HTTPExpectations() {
		hash := he.GetHash()
		s.expectations[hash] = he
	}

	w.WriteHeader(http.StatusOK)
}

func (s *controlService) RemoveHTTPMock(w http.ResponseWriter, r *http.Request) {
	mock := new(model.HTTPMock)
	if err := json.NewDecoder(r.Body).Decode(mock); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for _, he := range mock.HTTPExpectations() {
		hash := he.GetHash()
		delete(s.expectations, hash)
	}

	w.WriteHeader(http.StatusOK)
}

func (s *controlService) RemoveRESTMock(w http.ResponseWriter, r *http.Request) {
	mock := new(model.RESTMock)
	if err := json.NewDecoder(r.Body).Decode(mock); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for _, he := range mock.HTTPExpectations() {
		hash := he.GetHash()
		delete(s.expectations, hash)
	}

	w.WriteHeader(http.StatusOK)
}
