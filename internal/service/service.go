package service

import (
	"github.com/gorilla/mux"
	"github.com/moorara/log"
)

// Mock is the interface for a mock type.
type Mock interface {
	String() string
	Hash() uint64
	RegisterRoutes(*mux.Router)
}

// MockService provides functionalities to manage mocks.
type MockService struct {
	logger log.Logger
	mocks  map[uint64]Mock
}

// NewMockService creates a new instance of MockService.
func NewMockService(logger log.Logger) *MockService {
	return &MockService{
		logger: logger,
		mocks:  map[uint64]Mock{},
	}
}

// Add registers a new mock.
// If a mock already exists, it will be replaced.
func (s *MockService) Add(m Mock) {
	key := m.Hash()
	s.mocks[key] = m

	s.logger.Debug("message", "mock added", "mock", m.String())
}

// Delete deregisters an existing mock.
func (s *MockService) Delete(m Mock) {
	key := m.Hash()
	delete(s.mocks, key)

	s.logger.Debug("message", "mock deleted", "mock", m.String())
}

// Router creates a new router for mocks.
func (s *MockService) Router() *mux.Router {
	router := mux.NewRouter()
	for _, m := range s.mocks {
		m.RegisterRoutes(router)
	}

	return router
}
