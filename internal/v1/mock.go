package v1

import (
	"github.com/gorilla/mux"
)

// Mock is the interface for a mock type.
type Mock interface {
	Hash() uint64
	RegisterRoutes(*mux.Router)
}

// MockService provides functionalities to manage mocks.
type MockService struct {
	mocks map[uint64]Mock
}

// NewMockService creates a new instance of MockService.
func NewMockService() *MockService {
	return &MockService{
		mocks: map[uint64]Mock{},
	}
}

// Add registers new mocks.
// If a mock already exists, it will be replaced.
func (s *MockService) Add(mocks ...Mock) {
	for _, m := range mocks {
		key := m.Hash()
		s.mocks[key] = m
	}
}

// Delete deregisters mocks.
func (s *MockService) Delete(mocks ...Mock) {
	for _, m := range mocks {
		key := m.Hash()
		delete(s.mocks, key)
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
