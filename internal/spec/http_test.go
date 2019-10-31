package spec

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestHTTPMockSetDefaults(t *testing.T) {
	tests := []struct {
		name         string
		mock         HTTPMock
		expectedMock HTTPMock
	}{
		{
			"Empty",
			HTTPMock{},
			HTTPMock{
				HTTPExpect: HTTPExpect{
					Methods: []string{"GET"},
					Path:    "/",
					Prefix:  false,
					Queries: nil,
					Headers: nil,
				},
				HTTPResponse: &HTTPResponse{
					Delay:      "",
					StatusCode: 200,
					Headers:    nil,
					Body:       nil,
				},
			},
		},
		{
			"WithHTTPResponse",
			HTTPMock{
				HTTPResponse: &HTTPResponse{},
			},
			HTTPMock{
				HTTPExpect: HTTPExpect{
					Methods: []string{"GET"},
					Path:    "/",
					Prefix:  false,
					Queries: nil,
					Headers: nil,
				},
				HTTPResponse: &HTTPResponse{
					Delay:      "",
					StatusCode: 200,
					Headers:    nil,
					Body:       nil,
				},
			},
		},
		{
			"WithHTTPForward",
			HTTPMock{
				HTTPForward: &HTTPForward{},
			},
			HTTPMock{
				HTTPExpect: HTTPExpect{
					Methods: []string{"GET"},
					Path:    "/",
					Prefix:  false,
					Queries: nil,
					Headers: nil,
				},
				HTTPForward: &HTTPForward{
					Delay:   "",
					To:      "",
					Headers: nil,
				},
			},
		},
		{
			"WithHTTPResponseAndHTTPForward",
			HTTPMock{
				HTTPResponse: &HTTPResponse{},
				HTTPForward:  &HTTPForward{},
			},
			HTTPMock{
				HTTPExpect: HTTPExpect{
					Methods: []string{"GET"},
					Path:    "/",
					Prefix:  false,
					Queries: nil,
					Headers: nil,
				},
				HTTPResponse: &HTTPResponse{
					Delay:      "",
					StatusCode: 200,
					Headers:    nil,
					Body:       nil,
				},
				HTTPForward: &HTTPForward{
					Delay:   "",
					To:      "",
					Headers: nil,
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

func TestHTTPMockHash(t *testing.T) {
	tests := []struct {
		name          string
		m1            HTTPMock
		m2            HTTPMock
		expectedEqual bool
	}{
		{
			"EqualWithHTTPResponse",
			HTTPMock{
				HTTPExpect: HTTPExpect{
					Methods: []string{"POST", "PUT"},
					Path:    "/api/v1/sendMessage",
					Prefix:  false,
					Queries: map[string]string{
						"tenantId": "[0-9A-Fa-f-]+",
						"groupId":  "[0-9A-Fa-f-]+",
					},
					Headers: map[string]string{
						"Accept":       "application/json",
						"Content-Type": "application/json",
					},
				},
				HTTPResponse: &HTTPResponse{
					Delay:      "",
					StatusCode: 200,
					Headers:    map[string]string{},
				},
			},
			HTTPMock{
				HTTPExpect: HTTPExpect{
					Methods: []string{"POST", "PUT"},
					Path:    "/api/v1/sendMessage",
					Prefix:  false,
					Queries: map[string]string{
						"tenantId": "[0-9A-Fa-f-]+",
						"groupId":  "[0-9A-Fa-f-]+",
					},
					Headers: map[string]string{
						"Accept":       "application/json",
						"Content-Type": "application/json",
					},
				},
				HTTPResponse: &HTTPResponse{
					Delay:      "10ms",
					StatusCode: 201,
					Headers: map[string]string{
						"Content-Type": "application/json",
					},
					Body: JSON{
						"id": "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa",
					},
				},
			},
			true,
		},
		{
			"EqualWithHTTPForward",
			HTTPMock{
				HTTPExpect: HTTPExpect{
					Methods: []string{"POST", "PUT"},
					Path:    "/api/v1/sendMessage",
					Prefix:  false,
					Queries: map[string]string{
						"tenantId": "[0-9A-Fa-f-]+",
						"groupId":  "[0-9A-Fa-f-]+",
					},
					Headers: map[string]string{
						"Accept":       "application/json",
						"Content-Type": "application/json",
					},
				},
				HTTPForward: &HTTPForward{
					Delay:   "",
					To:      "",
					Headers: map[string]string{},
				},
			},
			HTTPMock{
				HTTPExpect: HTTPExpect{
					Methods: []string{"POST", "PUT"},
					Path:    "/api/v1/sendMessage",
					Prefix:  false,
					Queries: map[string]string{
						"tenantId": "[0-9A-Fa-f-]+",
						"groupId":  "[0-9A-Fa-f-]+",
					},
					Headers: map[string]string{
						"Accept":       "application/json",
						"Content-Type": "application/json",
					},
				},
				HTTPForward: &HTTPForward{
					Delay: "10ms",
					To:    "http://example.com",
					Headers: map[string]string{
						"Is-Test": "true",
					},
				},
			},
			true,
		},
		{
			"NotEqualWithHTTPResponse",
			HTTPMock{
				HTTPExpect: HTTPExpect{
					Methods: []string{"POST", "PUT"},
					Path:    "/api/v1/sendMessage",
					Prefix:  false,
					Queries: map[string]string{
						"tenantId": "[0-9A-Fa-f-]+",
						"groupId":  "[0-9A-Fa-f-]+",
					},
					Headers: map[string]string{
						"Accept":       "application/json",
						"Content-Type": "application/json",
					},
				},
				HTTPResponse: &HTTPResponse{
					Delay:      "",
					StatusCode: 200,
					Headers:    map[string]string{},
				},
			},
			HTTPMock{
				HTTPExpect: HTTPExpect{
					Methods: []string{"POST"},
					Path:    "/api/v1/sendMessage",
					Prefix:  false,
					Queries: map[string]string{
						"tenantId": "[0-9A-Fa-f-]+",
						"groupId":  "[0-9A-Fa-f-]+",
					},
					Headers: map[string]string{
						"Accept":       "application/json",
						"Content-Type": "application/json",
					},
				},
				HTTPResponse: &HTTPResponse{
					Delay:      "10ms",
					StatusCode: 201,
					Headers: map[string]string{
						"Content-Type": "application/json",
					},
					Body: JSON{
						"id": "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa",
					},
				},
			},
			false,
		},
		{
			"NotEqualWithHTTPForward",
			HTTPMock{
				HTTPExpect: HTTPExpect{
					Methods: []string{"POST", "PUT"},
					Path:    "/api/v1/sendMessage",
					Prefix:  false,
					Queries: map[string]string{
						"tenantId": "[0-9A-Fa-f-]+",
						"groupId":  "[0-9A-Fa-f-]+",
					},
					Headers: map[string]string{
						"Accept":       "application/json",
						"Content-Type": "application/json",
					},
				},
				HTTPForward: &HTTPForward{
					Delay:   "",
					To:      "",
					Headers: map[string]string{},
				},
			},
			HTTPMock{
				HTTPExpect: HTTPExpect{
					Methods: []string{"POST"},
					Path:    "/api/v1/sendMessage",
					Prefix:  false,
					Queries: map[string]string{
						"tenantId": "[0-9A-Fa-f-]+",
						"groupId":  "[0-9A-Fa-f-]+",
					},
					Headers: map[string]string{
						"Accept":       "application/json",
						"Content-Type": "application/json",
					},
				},
				HTTPForward: &HTTPForward{
					Delay: "10ms",
					To:    "http://example.com",
					Headers: map[string]string{
						"Is-Test": "true",
					},
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

func TestHTTPMockRoute(t *testing.T) {
	tests := []struct {
		name               string
		mock               HTTPMock
		reqMethod          string
		reqURL             string
		reqQueries         map[string]string
		reqHeaders         map[string]string
		expectedStatusCode int
		expectedHeaders    map[string]string
		expectedBody       JSON
	}{
		{
			name: "WithHTTPResponse",
			mock: HTTPMock{
				HTTPExpect: HTTPExpect{
					Methods: []string{"POST", "PUT"},
					Path:    "/api/v1/sendMessage",
					Prefix:  false,
					Queries: map[string]string{
						"tenantId": "[0-9A-Fa-f-]+",
						"groupId":  "[0-9A-Fa-f-]+",
					},
					Headers: map[string]string{
						"Accept":       "application/json",
						"Content-Type": "application/json",
					},
				},
				HTTPResponse: &HTTPResponse{
					Delay:      "10ms",
					StatusCode: 200,
					Headers: map[string]string{
						"Content-Type": "application/json",
					},
					Body: JSON{
						"id": "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa",
					},
				},
			},
			reqMethod: "POST",
			reqURL:    "http://example.com/api/v1/sendMessage",
			reqQueries: map[string]string{
				"tenantId": "11111111-1111-1111-1111-111111111111",
				"groupId":  "22222222-2222-2222-2222-222222222222",
			},
			reqHeaders: map[string]string{
				"Accept":       "application/json",
				"Content-Type": "application/json",
			},
			expectedStatusCode: 200,
			expectedHeaders: map[string]string{
				"Content-Type": "application/json",
			},
			expectedBody: JSON{
				"id": "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa",
			},
		},
		{
			name: "WithHTTPResponseAndPrefix",
			mock: HTTPMock{
				HTTPExpect: HTTPExpect{
					Methods: []string{"POST", "PUT"},
					Path:    "/api/v1/sendMessage",
					Prefix:  true,
					Queries: map[string]string{
						"tenantId": "[0-9A-Fa-f-]+",
						"groupId":  "[0-9A-Fa-f-]+",
					},
					Headers: map[string]string{
						"Accept":       "application/json",
						"Content-Type": "application/json",
					},
				},
				HTTPResponse: &HTTPResponse{
					Delay:      "10ms",
					StatusCode: 200,
					Headers: map[string]string{
						"Content-Type": "application/json",
					},
					Body: JSON{
						"id": "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa",
					},
				},
			},
			reqMethod: "POST",
			reqURL:    "http://example.com/api/v1/sendMessage",
			reqQueries: map[string]string{
				"tenantId": "11111111-1111-1111-1111-111111111111",
				"groupId":  "22222222-2222-2222-2222-222222222222",
			},
			reqHeaders: map[string]string{
				"Accept":       "application/json",
				"Content-Type": "application/json",
			},
			expectedStatusCode: 200,
			expectedHeaders: map[string]string{
				"Content-Type": "application/json",
			},
			expectedBody: JSON{
				"id": "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa",
			},
		},
		{
			name: "WithHTTPForward",
			mock: HTTPMock{
				HTTPExpect: HTTPExpect{
					Methods: []string{"POST", "PUT"},
					Path:    "/api/v1/sendMessage",
					Prefix:  false,
					Queries: map[string]string{
						"tenantId": "[0-9A-Fa-f-]+",
						"groupId":  "[0-9A-Fa-f-]+",
					},
					Headers: map[string]string{
						"Accept":       "application/json",
						"Content-Type": "application/json",
					},
				},
				HTTPForward: &HTTPForward{
					Delay: "10ms",
					To:    "http://example.com",
					Headers: map[string]string{
						"Is-Test": "true",
					},
				},
			},
			reqMethod: "POST",
			reqURL:    "http://example.com/api/v1/sendMessage",
			reqQueries: map[string]string{
				"tenantId": "11111111-1111-1111-1111-111111111111",
				"groupId":  "22222222-2222-2222-2222-222222222222",
			},
			reqHeaders: map[string]string{
				"Accept":       "application/json",
				"Content-Type": "application/json",
			},
			expectedStatusCode: 501,
			expectedHeaders:    map[string]string{},
			expectedBody: JSON{
				"message": "this functionality is not yet available!",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			router := mux.NewRouter()
			tc.mock.RegisterRoutes(router)

			req, err := http.NewRequest(tc.reqMethod, tc.reqURL, nil)
			assert.NoError(t, err)

			q := req.URL.Query()
			for k, v := range tc.reqQueries {
				q.Add(k, v)
			}
			req.URL.RawQuery = q.Encode()

			for k, v := range tc.reqHeaders {
				req.Header.Add(k, v)
			}

			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)

			assert.Equal(t, tc.expectedStatusCode, res.Result().StatusCode)
			for key, val := range tc.expectedHeaders {
				assert.Equal(t, val, res.Header().Get(key))
			}

			resBody := JSON{}
			err = json.NewDecoder(res.Body).Decode(&resBody)
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedBody, resBody)
		})
	}
}
