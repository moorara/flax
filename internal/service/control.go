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
		mocks   map[uint64]model.Mock
	}
)

// NewControlService creates a new instance of ControlService
func NewControlService(port string, logger *log.Logger, metrics *metrics.Metrics) ControlService {
	return &controlService{
		logger:  logger,
		metrics: metrics,
		mocks:   map[uint64]model.Mock{},
	}
}

func (s *controlService) getRouter() *mux.Router {
	router := mux.NewRouter()
	for _, m := range s.mocks {
		m.RegisterRoute(router, s.logger, s.metrics)
	}

	return router
}

func (s *controlService) AddHTTPMocks(w http.ResponseWriter, r *http.Request) {
	httpMocks := []model.HTTPMock{}
	if err := json.NewDecoder(r.Body).Decode(httpMocks); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for _, m := range httpMocks {
		s.mocks[m.Hash()] = m
	}

	w.WriteHeader(http.StatusOK)
}

func (s *controlService) AddRESTMocks(w http.ResponseWriter, r *http.Request) {
	restMocks := []model.RESTMock{}
	if err := json.NewDecoder(r.Body).Decode(restMocks); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for _, m := range restMocks {
		s.mocks[m.Hash()] = m
	}

	w.WriteHeader(http.StatusOK)
}

func (s *controlService) RemoveHTTPMocks(w http.ResponseWriter, r *http.Request) {
	httpMocks := []model.HTTPMock{}
	if err := json.NewDecoder(r.Body).Decode(httpMocks); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for _, m := range httpMocks {
		delete(s.mocks, m.Hash())
	}

	w.WriteHeader(http.StatusOK)
}

func (s *controlService) RemoveRESTMocks(w http.ResponseWriter, r *http.Request) {
	restMocks := []model.RESTMock{}
	if err := json.NewDecoder(r.Body).Decode(restMocks); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for _, m := range restMocks {
		delete(s.mocks, m.Hash())
	}

	w.WriteHeader(http.StatusOK)
}
