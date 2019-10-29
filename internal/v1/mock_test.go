package v1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMockService(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "OK"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ms := NewMockService()

			assert.NotNil(t, ms)
			assert.NotNil(t, ms.mocks)
		})
	}
}

func TestAdd(t *testing.T) {
	h1 := &HTTPMock{
		HTTPExpect: HTTPExpect{
			Methods: []string{"GET"},
			Path:    "/api/v1/sessions",
			Headers: map[string]string{
				"Accept":       "application/json",
				"Content-Type": "application/json",
			},
		},
		HTTPResponse: &HTTPResponse{
			StatusCode: 200,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		},
	}

	r1 := &RESTMock{
		RESTExpect{
			BasePath: "/api/v1/teams",
			Headers: map[string]string{
				"Accept":       "application/json",
				"Content-Type": "application/json",
			},
		},
		RESTResponse{
			GetStatusCode:    200,
			PostStatusCode:   201,
			PutStatusCode:    200,
			PatchStatusCode:  200,
			DeleteStatusCode: 204,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
			ListKey: "data",
		},
		RESTStore{
			Identifier: "id",
			Objects: []JSON{
				{"id": "d93ce179-50f7-469e-bb36-1b3746145f00", "name": "Back-end", "tags": []interface{}{"cloud", "go"}},
				{"id": "8cd6ef6c-2095-4c75-bc66-6f38e785299d", "name": "Front-end", "tags": []interface{}{"react", "redux"}},
			},
		},
	}

	tests := []struct {
		name          string
		svc           *MockService
		mocks         []Mock
		expectedMocks []Mock
	}{
		{
			name: "HTTPMocks",
			svc: &MockService{
				mocks: map[uint64]Mock{
					h1.Hash(): h1,
					r1.Hash(): r1,
				},
			},
		},
		{
			name: "RESTMocks",
			svc: &MockService{
				mocks: map[uint64]Mock{
					h1.Hash(): h1,
					r1.Hash(): r1,
				},
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
				mocks: map[uint64]Mock{},
			},
		},
		{
			name: "RESTMocks",
			svc: &MockService{
				mocks: map[uint64]Mock{},
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
