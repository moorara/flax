package service

import (
	"testing"

	"github.com/moorara/log"
	"github.com/stretchr/testify/assert"
)

func TestNewMockService(t *testing.T) {
	tests := []struct {
		name   string
		logger log.Logger
	}{
		{
			name:   "OK",
			logger: log.NewNopLogger(),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			service := NewMockService(tc.logger)

			assert.NotNil(t, service)
			assert.NotNil(t, service.logger)
			assert.NotNil(t, service.mocks)
		})
	}
}

func TestMockServiceAdd(t *testing.T) {
	tests := []struct {
		name          string
		service       *MockService
		mock          Mock
		expectedMocks []Mock
	}{
		{
			name: "HTTPMock",
			service: &MockService{
				logger: log.NewNopLogger(),
				mocks:  map[uint64]Mock{},
			},
		},
		{
			name: "RESTMock",
			service: &MockService{
				logger: log.NewNopLogger(),
				mocks:  map[uint64]Mock{},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NotNil(t, tc)
		})
	}
}

func TestMockServiceDelete(t *testing.T) {
	tests := []struct {
		name          string
		service       *MockService
		mock          Mock
		expectedMocks []Mock
	}{
		{
			name: "HTTPMock",
			service: &MockService{
				logger: log.NewNopLogger(),
				mocks:  map[uint64]Mock{},
			},
		},
		{
			name: "RESTMock",
			service: &MockService{
				logger: log.NewNopLogger(),
				mocks:  map[uint64]Mock{},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NotNil(t, tc)
		})
	}
}

func TestMockServiceRouter(t *testing.T) {
	tests := []struct {
		name    string
		service *MockService
	}{}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NotNil(t, tc)
		})
	}
}
