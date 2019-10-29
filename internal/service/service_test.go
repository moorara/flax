package service

import (
	"testing"

	"github.com/moorara/observe/log"
	"github.com/stretchr/testify/assert"
)

func TestNewMockService(t *testing.T) {
	tests := []struct {
		name   string
		logger *log.Logger
	}{
		{
			name:   "OK",
			logger: log.NewVoidLogger(),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ms := NewMockService(tc.logger)

			assert.NotNil(t, ms)
			assert.NotNil(t, ms.logger)
			assert.NotNil(t, ms.mocks)
		})
	}
}

func TestAdd(t *testing.T) {
	tests := []struct {
		name          string
		svc           *MockService
		mocks         []Mock
		expectedMocks []Mock
	}{
		{
			name: "HTTPMocks",
			svc: &MockService{
				logger: log.NewVoidLogger(),
				mocks:  map[uint64]Mock{},
			},
		},
		{
			name: "RESTMocks",
			svc: &MockService{
				logger: log.NewVoidLogger(),
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

func TestDelete(t *testing.T) {
	tests := []struct {
		name          string
		svc           *MockService
		mocks         []Mock
		expectedMocks []Mock
	}{
		{
			name: "HTTPMocks",
			svc: &MockService{
				logger: log.NewVoidLogger(),
				mocks:  map[uint64]Mock{},
			},
		},
		{
			name: "RESTMocks",
			svc: &MockService{
				logger: log.NewVoidLogger(),
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

func TestRouter(t *testing.T) {
	tests := []struct {
		name string
		svc  *MockService
	}{}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NotNil(t, tc)
		})
	}
}
