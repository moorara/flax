package v1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRESTExpectSetDefaults(t *testing.T) {
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
			tc.expect.SetDefaults()
			assert.Equal(t, tc.expectedExpect, tc.expect)
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
			"Equal",
			RESTExpect{
				BasePath: "/api/v1/teams",
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
			},
			RESTExpect{
				BasePath: "/api/v1/teams",
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
			},
			true,
		},
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
			"NotEqual",
			RESTExpect{
				BasePath: "/api/v1/teams",
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
			},
			RESTExpect{
				BasePath: "/api/v1/teams",
				Headers: map[string]string{
					"Content-Type": "application/ld+json",
				},
			},
			false,
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

func TestRESTResponseSetDefaults(t *testing.T) {
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
				GetStatusCode:    200,
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
				GetStatusCode:    200,
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
				GetStatusCode:    206,
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
				GetStatusCode:    206,
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
			tc.response.SetDefaults()
			assert.Equal(t, tc.expectedResponse, tc.response)
		})
	}
}

func TestRESTStoreSetDefaults(t *testing.T) {
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
				Directory:  map[interface{}]JSON{},
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
				Directory:  map[interface{}]JSON{},
			},
		},
		{
			"WithIdentifier",
			RESTStore{
				Identifier: "id",
				Objects: []JSON{
					{"id": "d93ce179-50f7-469e-bb36-1b3746145f00", "name": "Back-end", "tags": []interface{}{"cloud", "go"}},
					{"id": "8cd6ef6c-2095-4c75-bc66-6f38e785299d", "name": "Front-end", "tags": []interface{}{"react", "redux"}},
				},
			},
			RESTStore{
				Identifier: "id",
				Objects: []JSON{
					{"id": "d93ce179-50f7-469e-bb36-1b3746145f00", "name": "Back-end", "tags": []interface{}{"cloud", "go"}},
					{"id": "8cd6ef6c-2095-4c75-bc66-6f38e785299d", "name": "Front-end", "tags": []interface{}{"react", "redux"}},
				},
				Directory: map[interface{}]JSON{
					"d93ce179-50f7-469e-bb36-1b3746145f00": {"id": "d93ce179-50f7-469e-bb36-1b3746145f00", "name": "Back-end", "tags": []interface{}{"cloud", "go"}},
					"8cd6ef6c-2095-4c75-bc66-6f38e785299d": {"id": "8cd6ef6c-2095-4c75-bc66-6f38e785299d", "name": "Front-end", "tags": []interface{}{"react", "redux"}},
				},
			},
		},
		{
			"WithoutIdentifier",
			RESTStore{
				Identifier: "",
				Objects: []JSON{
					{"_id": "d93ce179-50f7-469e-bb36-1b3746145f00", "name": "Back-end", "tags": []interface{}{"cloud", "go"}},
					{"_id": "8cd6ef6c-2095-4c75-bc66-6f38e785299d", "name": "Front-end", "tags": []interface{}{"react", "redux"}},
				},
			},
			RESTStore{
				Identifier: "",
				Objects: []JSON{
					{"_id": "d93ce179-50f7-469e-bb36-1b3746145f00", "name": "Back-end", "tags": []interface{}{"cloud", "go"}},
					{"_id": "8cd6ef6c-2095-4c75-bc66-6f38e785299d", "name": "Front-end", "tags": []interface{}{"react", "redux"}},
				},
				Directory: map[interface{}]JSON{
					"d93ce179-50f7-469e-bb36-1b3746145f00": {"_id": "d93ce179-50f7-469e-bb36-1b3746145f00", "name": "Back-end", "tags": []interface{}{"cloud", "go"}},
					"8cd6ef6c-2095-4c75-bc66-6f38e785299d": {"_id": "8cd6ef6c-2095-4c75-bc66-6f38e785299d", "name": "Front-end", "tags": []interface{}{"react", "redux"}},
				},
			},
		},
		{
			"CustomIdentifier",
			RESTStore{
				Identifier: "key",
				Objects: []JSON{
					{"key": "d93ce179-50f7-469e-bb36-1b3746145f00", "name": "Back-end", "tags": []interface{}{"cloud", "go"}},
					{"key": "8cd6ef6c-2095-4c75-bc66-6f38e785299d", "name": "Front-end", "tags": []interface{}{"react", "redux"}},
				},
			},
			RESTStore{
				Identifier: "key",
				Objects: []JSON{
					{"key": "d93ce179-50f7-469e-bb36-1b3746145f00", "name": "Back-end", "tags": []interface{}{"cloud", "go"}},
					{"key": "8cd6ef6c-2095-4c75-bc66-6f38e785299d", "name": "Front-end", "tags": []interface{}{"react", "redux"}},
				},
				Directory: map[interface{}]JSON{
					"d93ce179-50f7-469e-bb36-1b3746145f00": {"key": "d93ce179-50f7-469e-bb36-1b3746145f00", "name": "Back-end", "tags": []interface{}{"cloud", "go"}},
					"8cd6ef6c-2095-4c75-bc66-6f38e785299d": {"key": "8cd6ef6c-2095-4c75-bc66-6f38e785299d", "name": "Front-end", "tags": []interface{}{"react", "redux"}},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.store.SetDefaults()
			assert.Equal(t, tc.expectedStore, tc.store)
		})
	}
}

func TestRESTMockSetDefaults(t *testing.T) {
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
					GetStatusCode:    200,
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
					Directory:  map[interface{}]JSON{},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.mock.SetDefaults()
			assert.Equal(t, tc.expectedMock, tc.mock)
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
			"Equal",
			RESTMock{
				RESTExpect{
					BasePath: "/api/v1/teams",
					Headers: map[string]string{
						"Content-Type": "application/json",
					},
				},
				RESTResponse{
					Delay:            "0",
					GetStatusCode:    200,
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
						"Content-Type": "application/json",
					},
				},
				RESTResponse{
					Delay:            "100ms",
					GetStatusCode:    200,
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
		{
			"NotEqual",
			RESTMock{
				RESTExpect{
					BasePath: "/api/v1/cars",
					Headers:  map[string]string{},
				},
				RESTResponse{
					Delay:            "0",
					GetStatusCode:    200,
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
					GetStatusCode:    200,
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
			"NotEqual",
			RESTMock{
				RESTExpect{
					BasePath: "/api/v1/teams",
					Headers: map[string]string{
						"Content-Type": "application/json",
					},
				},
				RESTResponse{
					Delay:            "0",
					GetStatusCode:    200,
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
						"Content-Type": "application/ld+json",
					},
				},
				RESTResponse{
					Delay:            "100ms",
					GetStatusCode:    200,
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

func TestRESTMockRegisterRoutes(t *testing.T) {
	tests := []struct {
		name          string
		mock          RESTMock
		reqMehod      string
		reqURL        string
		reqQueries    map[string]string
		reqHeaders    map[string]string
		resStatusCode int
		resHeaders    map[string]string
		resBody       JSON
	}{}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NotNil(t, tc)
		})
	}
}
