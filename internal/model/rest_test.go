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

func TestRESTExpectHash(t *testing.T) {
	tests := []struct {
		name          string
		e1            RESTExpect
		e2            RESTExpect
		expectedEqual bool
	}{
		{
			"NotEqual",
			RESTExpect{
				BasePath: "/api/v1/cars",
				Headers:  map[string]string{},
			},
			RESTExpect{
				BasePath: "/api/v1/teams",
				Headers:  map[string]string{},
			},
			false,
		},
		{
			"Equal",
			RESTExpect{
				BasePath: "/api/v1/teams",
				Headers: map[string]string{
					"tags": "go",
				},
			},
			RESTExpect{
				BasePath: "/api/v1/teams",
				Headers: map[string]string{
					"tags": "javascript",
				},
			},
			true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.expectedEqual {
				assert.Equal(t, tc.e1.Hash(), tc.e2.Hash())
			} else {
				assert.NotEqual(t, tc.e1.Hash(), tc.e2.Hash())
			}
		})
	}
}

func TestRESTResponseWithDefaults(t *testing.T) {
	tests := []struct {
		name             string
		response         RESTResponse
		expectedResponse RESTResponse
	}{
		{
			"Empty",
			RESTResponse{},
			RESTResponse{
				Delay:            "0",
				PostStatusCode:   201,
				PutStatusCode:    200,
				PatchStatusCode:  200,
				DeleteStatusCode: 204,
				ListProperty:     "",
				Headers:          map[string]string{},
			},
		},
		{
			"DefaultRequired",
			RESTResponse{
				Delay:        "100ms",
				ListProperty: "data",
			},
			RESTResponse{
				Delay:            "100ms",
				PostStatusCode:   201,
				PutStatusCode:    200,
				PatchStatusCode:  200,
				DeleteStatusCode: 204,
				ListProperty:     "data",
				Headers:          map[string]string{},
			},
		},
		{
			"NoDefaultRequired",
			RESTResponse{
				Delay:            "100ms",
				PostStatusCode:   202,
				PutStatusCode:    204,
				PatchStatusCode:  204,
				DeleteStatusCode: 202,
				ListProperty:     "data",
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
			},
			RESTResponse{
				Delay:            "100ms",
				PostStatusCode:   202,
				PutStatusCode:    204,
				PatchStatusCode:  204,
				DeleteStatusCode: 202,
				ListProperty:     "data",
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedResponse, tc.response.WithDefaults())
		})
	}
}

func TestRESTStoreWithDefaults(t *testing.T) {
	tests := []struct {
		name          string
		store         RESTStore
		expectedStore RESTStore
	}{
		{
			"Empty",
			RESTStore{},
			RESTStore{
				Identifier: "",
				Objects:    []JSON{},
			},
		},
		{
			"DefaultRequired",
			RESTStore{
				Identifier: "_id",
			},
			RESTStore{
				Identifier: "_id",
				Objects:    []JSON{},
			},
		},
		{
			"NoDefaultRequired",
			RESTStore{
				Identifier: "_id",
				Objects: []JSON{
					{"_id": "d93ce179-50f7-469e-bb36-1b3746145f00", "name": "Back-end", "tags": []interface{}{"cloud", "go"}},
					{"_id": "8cd6ef6c-2095-4c75-bc66-6f38e785299d", "name": "Front-end", "tags": []interface{}{"react", "redux"}},
				},
			},
			RESTStore{
				Identifier: "_id",
				Objects: []JSON{
					{"_id": "d93ce179-50f7-469e-bb36-1b3746145f00", "name": "Back-end", "tags": []interface{}{"cloud", "go"}},
					{"_id": "8cd6ef6c-2095-4c75-bc66-6f38e785299d", "name": "Front-end", "tags": []interface{}{"react", "redux"}},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedStore, tc.store.WithDefaults())
		})
	}
}

func TestRESTMockWithDefaults(t *testing.T) {
	tests := []struct {
		name         string
		mock         RESTMock
		expectedMock RESTMock
	}{
		{
			"DefaultsRequired",
			RESTMock{},
			RESTMock{
				RESTExpect{
					BasePath: "/",
					Headers:  map[string]string{},
				},
				RESTResponse{
					Delay:            "0",
					PostStatusCode:   201,
					PutStatusCode:    200,
					PatchStatusCode:  200,
					DeleteStatusCode: 204,
					ListProperty:     "",
					Headers:          map[string]string{},
				},
				RESTStore{
					Identifier: "",
					Objects:    []JSON{},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedMock, tc.mock.WithDefaults())
		})
	}
}

func TestRESTMockHash(t *testing.T) {
	tests := []struct {
		name          string
		m1            RESTMock
		m2            RESTMock
		expectedEqual bool
	}{
		{
			"NotEqual",
			RESTMock{
				RESTExpect{
					BasePath: "/api/v1/cars",
					Headers:  map[string]string{},
				},
				RESTResponse{
					Delay:            "0",
					PostStatusCode:   201,
					PutStatusCode:    200,
					PatchStatusCode:  200,
					DeleteStatusCode: 204,
					ListProperty:     "",
					Headers:          map[string]string{},
				},
				RESTStore{
					Identifier: "",
					Objects:    []JSON{},
				},
			},
			RESTMock{
				RESTExpect{
					BasePath: "/api/v1/teams",
					Headers:  map[string]string{},
				},
				RESTResponse{
					Delay:            "100ms",
					PostStatusCode:   201,
					PutStatusCode:    200,
					PatchStatusCode:  200,
					DeleteStatusCode: 204,
					ListProperty:     "data",
					Headers:          map[string]string{},
				},
				RESTStore{
					Identifier: "_id",
					Objects:    []JSON{},
				},
			},
			false,
		},
		{
			"Equal",
			RESTMock{
				RESTExpect{
					BasePath: "/api/v1/teams",
					Headers: map[string]string{
						"tags": "go",
					},
				},
				RESTResponse{
					Delay:            "0",
					PostStatusCode:   201,
					PutStatusCode:    200,
					PatchStatusCode:  200,
					DeleteStatusCode: 204,
					ListProperty:     "",
					Headers:          map[string]string{},
				},
				RESTStore{
					Identifier: "",
					Objects:    []JSON{},
				},
			},
			RESTMock{
				RESTExpect{
					BasePath: "/api/v1/teams",
					Headers: map[string]string{
						"tags": "javascript",
					},
				},
				RESTResponse{
					Delay:            "100ms",
					PostStatusCode:   201,
					PutStatusCode:    200,
					PatchStatusCode:  200,
					DeleteStatusCode: 204,
					ListProperty:     "data",
					Headers:          map[string]string{},
				},
				RESTStore{
					Identifier: "_id",
					Objects:    []JSON{},
				},
			},
			true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.expectedEqual {
				assert.Equal(t, tc.m1.Hash(), tc.m2.Hash())
			} else {
				assert.NotEqual(t, tc.m1.Hash(), tc.m2.Hash())
			}
		})
	}
}
