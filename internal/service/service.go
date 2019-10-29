package service

import (
	"github.com/gorilla/mux"
	"github.com/moorara/observe/log"
)

// Mock is the interface for a mock type.
type Mock interface {
	Hash() uint64
	RegisterRoutes(*mux.Router)
}

// MockService provides functionalities to manage mocks.
type MockService struct {
	logger *log.Logger
	mocks  map[uint64]Mock
}

// NewMockService creates a new instance of MockService.
func NewMockService(logger *log.Logger) *MockService {
	return &MockService{
		logger: logger,
		mocks:  map[uint64]Mock{},
	}
}

// Add registers new mocks.
// If a mock already exists, it will be replaced.
func (s *MockService) Add(mocks ...Mock) {
	for _, m := range mocks {
		key := m.Hash()
		s.mocks[key] = m
		s.logger.DebugKV("message", "mock added.")
	}
}

// Delete deregisters mocks.
func (s *MockService) Delete(mocks ...Mock) {
	for _, m := range mocks {
		key := m.Hash()
		delete(s.mocks, key)
		s.logger.DebugKV("message", "mock deleted.")
	}
}

// Router creates a new router for mocks.
func (s *MockService) Router() *mux.Router {
	router := mux.NewRouter()
	for _, m := range s.mocks {
		m.RegisterRoutes(router)
	}

	return router
}
