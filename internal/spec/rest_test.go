package spec

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestRESTStoreIndex(t *testing.T) {
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
				Objects:    nil,
				Directory:  map[interface{}]JSON{},
			},
		},
		{
			"WithoutIdentifier",
			RESTStore{
				Objects: []JSON{
					{"id": "aaaa", "name": "Back-end"},
					{"id": "bbbb", "name": "Front-end"},
				},
			},
			RESTStore{
				Objects: []JSON{
					{"id": "aaaa", "name": "Back-end"},
					{"id": "bbbb", "name": "Front-end"},
				},
				Directory: map[interface{}]JSON{
					"aaaa": {"id": "aaaa", "name": "Back-end"},
					"bbbb": {"id": "bbbb", "name": "Front-end"},
				},
			},
		},
		{
			"WithIdentifier",
			RESTStore{
				Identifier: "_id",
				Objects: []JSON{
					{"_id": "aaaa", "name": "Back-end"},
					{"_id": "bbbb", "name": "Front-end"},
				},
			},
			RESTStore{
				Identifier: "_id",
				Objects: []JSON{
					{"_id": "aaaa", "name": "Back-end"},
					{"_id": "bbbb", "name": "Front-end"},
				},
				Directory: map[interface{}]JSON{
					"aaaa": {"_id": "aaaa", "name": "Back-end"},
					"bbbb": {"_id": "bbbb", "name": "Front-end"},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.store.Index()
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
			"Empty",
			RESTMock{},
			RESTMock{
				RESTExpect{
					BasePath: "/",
					Headers:  nil,
				},
				RESTResponse{
					Delay:            "",
					GetStatusCode:    200,
					PostStatusCode:   201,
					PutStatusCode:    200,
					PatchStatusCode:  200,
					DeleteStatusCode: 204,
					Headers: map[string]string{
						"Content-Type": "application/json",
					},
					ListKey: "",
				},
				RESTStore{
					Identifier: "",
					Objects:    []JSON{},
					Directory:  nil,
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

func TestRESTMockkString(t *testing.T) {
	tests := []struct {
		name           string
		mock           RESTMock
		expectedString string
	}{
		{
			"Empty",
			RESTMock{},
			"",
		},
		{
			"OK",
			RESTMock{
				RESTExpect{
					BasePath: "/api/v1/teams",
					Headers: map[string]string{
						"Accept":       "application/json",
						"Content-Type": "application/json",
					},
				},
				RESTResponse{
					Delay:            "",
					GetStatusCode:    200,
					PostStatusCode:   201,
					PutStatusCode:    200,
					PatchStatusCode:  200,
					DeleteStatusCode: 204,
					Headers:          map[string]string{},
					ListKey:          "",
				},
				RESTStore{
					Identifier: "",
					Objects:    []JSON{},
				},
			},
			"/api/v1/teams",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			str := tc.mock.String()
			assert.Equal(t, tc.expectedString, str)
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
						"Accept":       "application/json",
						"Content-Type": "application/json",
					},
				},
				RESTResponse{
					Delay:            "",
					GetStatusCode:    200,
					PostStatusCode:   201,
					PutStatusCode:    200,
					PatchStatusCode:  200,
					DeleteStatusCode: 204,
					Headers:          map[string]string{},
					ListKey:          "",
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
						"Accept":       "application/json",
						"Content-Type": "application/json",
					},
				},
				RESTResponse{
					Delay:            "10ms",
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
					Identifier: "_id",
					Objects:    []JSON{},
				},
			},
			true,
		},
		{
			"DifferentVersions",
			RESTMock{
				RESTExpect{
					BasePath: "/api/v1/teams",
					Headers: map[string]string{
						"Accept":       "application/json",
						"Content-Type": "application/json",
					},
				},
				RESTResponse{
					Delay:            "",
					GetStatusCode:    200,
					PostStatusCode:   201,
					PutStatusCode:    200,
					PatchStatusCode:  200,
					DeleteStatusCode: 204,
					Headers:          map[string]string{},
					ListKey:          "",
				},
				RESTStore{
					Identifier: "",
					Objects:    []JSON{},
				},
			},
			RESTMock{
				RESTExpect{
					BasePath: "/api/v2/teams",
					Headers: map[string]string{
						"Accept":       "application/json",
						"Content-Type": "application/json",
					},
				},
				RESTResponse{
					Delay:            "10ms",
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
					Identifier: "_id",
					Objects:    []JSON{},
				},
			},
			false,
		},
		{
			"DifferentHeaders",
			RESTMock{
				RESTExpect{
					BasePath: "/api/v1/teams",
					Headers:  map[string]string{},
				},
				RESTResponse{
					Delay:            "",
					GetStatusCode:    200,
					PostStatusCode:   201,
					PutStatusCode:    200,
					PatchStatusCode:  200,
					DeleteStatusCode: 204,
					Headers:          map[string]string{},
					ListKey:          "",
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
						"Accept":       "application/json",
						"Content-Type": "application/json",
					},
				},
				RESTResponse{
					Delay:            "10ms",
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
		name                     string
		mock                     RESTMock
		reqBasePath              string
		reqHeaders               map[string]string
		expectedGetStatusCode    int
		expectedPostStatusCode   int
		expectedPutStatusCode    int
		expectedPatchStatusCode  int
		expectedDeleteStatusCode int
		expectedHeaders          map[string]string
		expectedAllBody          interface{}
	}{
		{
			name: "WithListKey",
			mock: RESTMock{
				RESTExpect{
					BasePath: "/api/v1/teams",
					Headers: map[string]string{
						"Accept":       "application/json",
						"Content-Type": "application/json",
					},
				},
				RESTResponse{
					Delay:            "10ms",
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
						{"id": "aaaa", "name": "Back-end"},
						{"id": "bbbb", "name": "Front-end"},
					},
				},
			},
			reqBasePath: "/api/v1/teams",
			reqHeaders: map[string]string{
				"Accept":       "application/json",
				"Content-Type": "application/json",
			},
			expectedGetStatusCode:    200,
			expectedPostStatusCode:   201,
			expectedPutStatusCode:    200,
			expectedPatchStatusCode:  200,
			expectedDeleteStatusCode: 204,
			expectedHeaders: map[string]string{
				"Content-Type": "application/json",
			},
			expectedAllBody: JSON{
				"data": []interface{}{
					map[string]interface{}{"id": "aaaa", "name": "Back-end"},
					map[string]interface{}{"id": "bbbb", "name": "Front-end"},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			router := mux.NewRouter()
			tc.mock.RegisterRoutes(router)

			t.Run("ALL", func(t *testing.T) {
				req, err := http.NewRequest("GET", tc.reqBasePath, nil)
				assert.NoError(t, err)

				for k, v := range tc.reqHeaders {
					req.Header.Add(k, v)
				}

				res := httptest.NewRecorder()
				router.ServeHTTP(res, req)

				assert.Equal(t, tc.expectedGetStatusCode, res.Result().StatusCode)
				for key, val := range tc.expectedHeaders {
					assert.Equal(t, val, res.Header().Get(key))
				}

				resBody := JSON{}
				err = json.NewDecoder(res.Body).Decode(&resBody)
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedAllBody, resBody)
			})

			t.Run("POST", func(t *testing.T) {
				// TODO:
			})

			t.Run("GET", func(t *testing.T) {
				// TODO:
			})

			t.Run("PUT", func(t *testing.T) {
				// TODO:
			})

			t.Run("PATCH", func(t *testing.T) {
				// TODO:
			})

			t.Run("DELETE", func(t *testing.T) {
				// TODO:
			})
		})
	}
}
