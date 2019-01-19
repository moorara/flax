package v1

import (
	"testing"

	"github.com/moorara/flax/pkg/log"
	"github.com/moorara/flax/pkg/metrics"
	"github.com/stretchr/testify/assert"
)

func TestNewControlService(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			"Simple",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			logger := log.NewNopLogger()
			metrics := metrics.New("unit-test")
			service := NewControlService(logger, metrics)
			assert.NotNil(t, service)
		})
	}
}

func TestCreateRouter(t *testing.T) {
	tests := []struct {
		name      string
		httpMocks []*HTTPMock
		restMocks []*RESTMock
	}{}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NotNil(t, tc)
		})
	}
}

func TestAddHTTPMocks(t *testing.T) {
	tests := []struct {
		name      string
		httpMocks []*HTTPMock
	}{}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NotNil(t, tc)
		})
	}
}

func TestAddRESTMocks(t *testing.T) {
	tests := []struct {
		name      string
		restMocks []*RESTMock
	}{}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NotNil(t, tc)
		})
	}
}

func TestRemoveHTTPMocks(t *testing.T) {
	tests := []struct {
		name      string
		httpMocks []*HTTPMock
	}{}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NotNil(t, tc)
		})
	}
}

func TestRemoveRESTMocks(t *testing.T) {
	tests := []struct {
		name      string
		restMocks []*RESTMock
	}{}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NotNil(t, tc)
		})
	}
}
