package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRESTExpectWithDefaults(t *testing.T) {
	tests := []struct {
		name           string
		expect         RESTExpect
		expectedExpect RESTExpect
	}{
		{
			"Empty",
			RESTExpect{},
			RESTExpect{
				BasePath: "/",
				Headers:  map[string]string{},
			},
		},
		{
			"DefaultRequired",
			RESTExpect{
				BasePath: "/cars",
			},
			RESTExpect{
				BasePath: "/cars",
				Headers:  map[string]string{},
			},
		},
		{
			"NoDefaultRequired",
			RESTExpect{
				BasePath: "/teams",
				Headers: map[string]string{
					"Accept":       "application/json",
					"Content-Type": "application/json",
				},
			},
			RESTExpect{
				BasePath: "/teams",
				Headers: map[string]string{
					"Accept":       "application/json",
					"Content-Type": "application/json",
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedExpect, tc.expect.WithDefaults())
		})
	}
}

func TestRESTResponseWithDefaults(t *testing.T) {
	tests := []struct {
		name     string
		response RESTResponse
	}{}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NotNil(t, tc)
		})
	}
}

func TestRESTStoreWithDefaults(t *testing.T) {
	tests := []struct {
		name  string
		store RESTExpect
	}{}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NotNil(t, tc)
		})
	}
}

func TestRESTMockWithDefaults(t *testing.T) {
	tests := []struct {
		name string
		mock RESTMock
	}{}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NotNil(t, tc)
		})
	}
}

func TestRESTMockHash(t *testing.T) {
	tests := []struct {
		name string
		m1   RESTMock
		m2   RESTMock
	}{}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NotNil(t, tc)
		})
	}
}
