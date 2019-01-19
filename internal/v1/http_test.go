package v1

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestHTTPExpectSetDefaults(t *testing.T) {
	tests := []struct {
		name           string
		expect         HTTPExpect
		expectedExpect HTTPExpect
	}{
		{
			"Empty",
			HTTPExpect{},
			HTTPExpect{
				Methods: []string{"GET"},
				Path:    "/",
				Prefix:  false,
				Queries: map[string]string{},
				Headers: map[string]string{},
			},
		},
		{
			"DefaultRequired",
			HTTPExpect{
				Path: "/health",
			},
			HTTPExpect{
				Methods: []string{"GET"},
				Path:    "/health",
				Prefix:  false,
				Queries: map[string]string{},
				Headers: map[string]string{},
			},
		},
		{
			"NoDefaultRequired",
			HTTPExpect{
				Methods: []string{"POST", "PUT"},
				Path:    "/sessions",
				Prefix:  true,
				Queries: map[string]string{
					"tenantId": "[0-9A-Za-z-]+",
				},
				Headers: map[string]string{
					"Accept":       "application/json",
					"Content-Type": "application/json",
				},
			},
			HTTPExpect{
				Methods: []string{"POST", "PUT"},
				Path:    "/sessions",
				Prefix:  true,
				Queries: map[string]string{
					"tenantId": "[0-9A-Za-z-]+",
				},
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

func TestHTTPExpectHash(t *testing.T) {
	tests := []struct {
		name          string
		e1            HTTPExpect
		e2            HTTPExpect
		expectedEqual bool
	}{
		{
			"Equal",
			HTTPExpect{
				Methods: []string{"POST", "PUT"},
				Path:    "/v1/sessions",
				Prefix:  true,
				Queries: map[string]string{
					"tenantId": "[0-9A-Za-z-]+",
					"deviceId": "[0-9A-Za-z-]+",
				},
				Headers: map[string]string{
					"Accept":       "application/json",
					"Content-Type": "application/json",
				},
			},
			HTTPExpect{
				Methods: []string{"POST", "PUT"},
				Path:    "/v1/sessions",
				Prefix:  true,
				Queries: map[string]string{
					"tenantId": "[0-9A-Za-z-]+",
					"deviceId": "[0-9A-Za-z-]+",
				},
				Headers: map[string]string{
					"Accept":       "application/json",
					"Content-Type": "application/json",
				},
			},
			true,
		},
		{
			"NotEqual",
			HTTPExpect{
				Methods: []string{"POST", "PUT"},
				Path:    "/v1/sessions",
				Prefix:  true,
				Queries: map[string]string{
					"tenantId": "[0-9A-Za-z-]+",
					"deviceId": "[0-9A-Za-z-]+",
				},
				Headers: map[string]string{
					"Accept":       "application/json",
					"Content-Type": "application/json",
				},
			},
			HTTPExpect{
				Methods: []string{"POST", "PUT"},
				Path:    "/v1/sessions",
				Prefix:  false,
				Queries: map[string]string{
					"tenantId": "[0-9A-Za-z-]+",
					"deviceId": "[0-9A-Za-z-]+",
				},
				Headers: map[string]string{
					"Accept":       "application/json",
					"Content-Type": "application/json",
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

func TestHTTPResponseSetDefaults(t *testing.T) {
	tests := []struct {
		name             string
		response         HTTPResponse
		expectedResponse HTTPResponse
	}{
		{
			"Empty",
			HTTPResponse{},
			HTTPResponse{
				Delay:      "0",
				StatusCode: 200,
				Headers:    map[string]string{},
				Body:       nil,
			},
		},
		{
			"DefaultRequired",
			HTTPResponse{
				Delay:      "100ms",
				StatusCode: 201,
			},
			HTTPResponse{
				Delay:      "100ms",
				StatusCode: 201,
				Headers:    map[string]string{},
				Body:       nil,
			},
		},
		{
			"NoDefaultRequired",
			HTTPResponse{
				Delay:      "100ms",
				StatusCode: 201,
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
				Body: map[string]interface{}{
					"id":    "5da8349a-0707-4064-8fad-74cedb48a8fc",
					"name":  "John Doe",
					"email": "john.doe@example.com",
				},
			},
			HTTPResponse{
				Delay:      "100ms",
				StatusCode: 201,
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
				Body: map[string]interface{}{
					"id":    "5da8349a-0707-4064-8fad-74cedb48a8fc",
					"name":  "John Doe",
					"email": "john.doe@example.com",
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

func TestHTTPResponseGetHandler(t *testing.T) {
	tests := []struct {
		name          string
		response      HTTPResponse
		resStatusCode int
		resHeaders    map[string]string
		resBody       JSON
	}{
		{
			name: "",
			response: HTTPResponse{
				Delay:      "10ms",
				StatusCode: 200,
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
				Body: JSON{
					"id": "00000000-0000-0000-0000-000000000000",
				},
			},
			resStatusCode: 200,
			resHeaders: map[string]string{
				"Content-Type": "application/json",
			},
			resBody: JSON{
				"id": "00000000-0000-0000-0000-000000000000",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			handler := tc.response.GetHandler()

			res := httptest.NewRecorder()
			handler(res, nil)

			assert.Equal(t, tc.resStatusCode, res.Result().StatusCode)
			for key, val := range tc.resHeaders {
				assert.Equal(t, val, res.Header().Get(key))
			}

			resBody := JSON{}
			err := json.NewDecoder(res.Body).Decode(&resBody)
			assert.NoError(t, err)
			assert.Equal(t, tc.resBody, resBody)
		})
	}
}

func TestHTTPForwardSetDefaults(t *testing.T) {
	tests := []struct {
		name            string
		forward         HTTPForward
		expectedForward HTTPForward
	}{
		{
			"Empty",
			HTTPForward{},
			HTTPForward{
				Delay:   "0",
				To:      "",
				Headers: map[string]string{},
			},
		},
		{
			"DefaultRequired",
			HTTPForward{
				To: "http://service:8080/api",
			},
			HTTPForward{
				Delay:   "0",
				To:      "http://service:8080/api",
				Headers: map[string]string{},
			},
		},
		{
			"NoDefaultRequired",
			HTTPForward{
				Delay: "100ms",
				To:    "http://service:8080/api",
				Headers: map[string]string{
					"Referer": "flax",
				},
			},
			HTTPForward{
				Delay: "100ms",
				To:    "http://service:8080/api",
				Headers: map[string]string{
					"Referer": "flax",
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.forward.SetDefaults()
			assert.Equal(t, tc.expectedForward, tc.forward)
		})
	}
}

func TestHTTPForwardGetHandler(t *testing.T) {
	tests := []struct {
		name          string
		forward       HTTPForward
		resStatusCode int
		resHeaders    map[string]string
		resBody       JSON
	}{
		{
			name: "",
			forward: HTTPForward{
				Delay:   "10ms",
				To:      "https://auth-service",
				Headers: map[string]string{},
			},
			resStatusCode: 501,
			resHeaders:    map[string]string{},
			resBody: JSON{
				"message": "this functionality is not yet available!",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			handler := tc.forward.GetHandler()

			res := httptest.NewRecorder()
			handler(res, nil)

			assert.Equal(t, tc.resStatusCode, res.Result().StatusCode)
			for key, val := range tc.resHeaders {
				assert.Equal(t, val, res.Header().Get(key))
			}

			resBody := JSON{}
			err := json.NewDecoder(res.Body).Decode(&resBody)
			assert.NoError(t, err)
			assert.Equal(t, tc.resBody, resBody)
		})
	}
}

func TestHTTPMockSetDefaults(t *testing.T) {
	tests := []struct {
		name         string
		httpMock     HTTPMock
		expectedMock HTTPMock
	}{
		{
			"OnlyHTTPMock",
			HTTPMock{},
			HTTPMock{
				HTTPExpect: HTTPExpect{
					Methods: []string{"GET"},
					Path:    "/",
					Prefix:  false,
					Queries: map[string]string{},
					Headers: map[string]string{},
				},
				HTTPResponse: &HTTPResponse{
					Delay:      "0",
					StatusCode: 200,
					Headers:    map[string]string{},
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
					Queries: map[string]string{},
					Headers: map[string]string{},
				},
				HTTPResponse: &HTTPResponse{
					Delay:      "0",
					StatusCode: 200,
					Headers:    map[string]string{},
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
					Queries: map[string]string{},
					Headers: map[string]string{},
				},
				HTTPForward: &HTTPForward{
					Delay:   "0",
					To:      "",
					Headers: map[string]string{},
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
					Queries: map[string]string{},
					Headers: map[string]string{},
				},
				HTTPResponse: &HTTPResponse{
					Delay:      "0",
					StatusCode: 200,
					Headers:    map[string]string{},
					Body:       nil,
				},
				HTTPForward: &HTTPForward{
					Delay:   "0",
					To:      "",
					Headers: map[string]string{},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.httpMock.SetDefaults()
			assert.Equal(t, tc.expectedMock, tc.httpMock)
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
					Path:    "/v1/sessions",
					Prefix:  true,
					Queries: map[string]string{
						"tenantId": "[0-9A-Za-z-]+",
						"deviceId": "[0-9A-Za-z-]+",
					},
					Headers: map[string]string{
						"Accept":       "application/json",
						"Content-Type": "application/json",
					},
				},
				HTTPResponse: &HTTPResponse{
					Delay:      "0",
					StatusCode: 200,
					Headers:    map[string]string{},
				},
			},
			HTTPMock{
				HTTPExpect: HTTPExpect{
					Methods: []string{"POST", "PUT"},
					Path:    "/v1/sessions",
					Prefix:  true,
					Queries: map[string]string{
						"tenantId": "[0-9A-Za-z-]+",
						"deviceId": "[0-9A-Za-z-]+",
					},
					Headers: map[string]string{
						"Accept":       "application/json",
						"Content-Type": "application/json",
					},
				},
				HTTPResponse: &HTTPResponse{
					Delay:      "100ms",
					StatusCode: 201,
					Headers: map[string]string{
						"Content-Type": "application/json",
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
					Path:    "/v1/sessions",
					Prefix:  true,
					Queries: map[string]string{
						"tenantId": "[0-9A-Za-z-]+",
						"deviceId": "[0-9A-Za-z-]+",
					},
					Headers: map[string]string{
						"Accept":       "application/json",
						"Content-Type": "application/json",
					},
				},
				HTTPForward: &HTTPForward{
					Delay:   "0",
					To:      "",
					Headers: map[string]string{},
				},
			},
			HTTPMock{
				HTTPExpect: HTTPExpect{
					Methods: []string{"POST", "PUT"},
					Path:    "/v1/sessions",
					Prefix:  true,
					Queries: map[string]string{
						"tenantId": "[0-9A-Za-z-]+",
						"deviceId": "[0-9A-Za-z-]+",
					},
					Headers: map[string]string{
						"Accept":       "application/json",
						"Content-Type": "application/json",
					},
				},
				HTTPForward: &HTTPForward{
					Delay: "100ms",
					To:    "http://service:8080/api",
					Headers: map[string]string{
						"Referer": "flax",
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
					Path:    "/v1/sessions",
					Prefix:  true,
					Queries: map[string]string{
						"tenantId": "[0-9A-Za-z-]+",
						"deviceId": "[0-9A-Za-z-]+",
					},
					Headers: map[string]string{
						"Accept":       "application/json",
						"Content-Type": "application/json",
					},
				},
				HTTPResponse: &HTTPResponse{
					Delay:      "0",
					StatusCode: 200,
					Headers:    map[string]string{},
				},
			},
			HTTPMock{
				HTTPExpect: HTTPExpect{
					Methods: []string{"POST", "PUT"},
					Path:    "/v1/sessions",
					Prefix:  false,
					Queries: map[string]string{
						"tenantId": "[0-9A-Za-z-]+",
						"deviceId": "[0-9A-Za-z-]+",
					},
					Headers: map[string]string{
						"Accept":       "application/json",
						"Content-Type": "application/json",
					},
				},
				HTTPResponse: &HTTPResponse{
					Delay:      "100ms",
					StatusCode: 201,
					Headers: map[string]string{
						"Content-Type": "application/json",
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
					Path:    "/v1/sessions",
					Prefix:  true,
					Queries: map[string]string{
						"tenantId": "[0-9A-Za-z-]+",
						"deviceId": "[0-9A-Za-z-]+",
					},
					Headers: map[string]string{
						"Accept":       "application/json",
						"Content-Type": "application/json",
					},
				},
				HTTPForward: &HTTPForward{
					Delay:   "0",
					To:      "",
					Headers: map[string]string{},
				},
			},
			HTTPMock{
				HTTPExpect: HTTPExpect{
					Methods: []string{"POST", "PUT"},
					Path:    "/v1/sessions",
					Prefix:  false,
					Queries: map[string]string{
						"tenantId": "[0-9A-Za-z-]+",
						"deviceId": "[0-9A-Za-z-]+",
					},
					Headers: map[string]string{
						"Accept":       "application/json",
						"Content-Type": "application/json",
					},
				},
				HTTPForward: &HTTPForward{
					Delay: "100ms",
					To:    "http://service:8080/api",
					Headers: map[string]string{
						"Referer": "flax",
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

func TestHTTPMockRegisterRoutes(t *testing.T) {
	tests := []struct {
		name          string
		mock          HTTPMock
		reqMehod      string
		reqURL        string
		reqQueries    map[string]string
		reqHeaders    map[string]string
		resStatusCode int
		resHeaders    map[string]string
		resBody       JSON
	}{
		{
			name: "WithHTTPResponse",
			mock: HTTPMock{
				HTTPExpect: HTTPExpect{
					Methods: []string{"POST"},
					Path:    "/v1/sessions",
					Prefix:  false,
					Queries: map[string]string{
						"tenantId": "[0-9A-Za-z-]+",
						"deviceId": "[0-9A-Za-z-]+",
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
						"id": "00000000-0000-0000-0000-000000000000",
					},
				},
			},
			reqMehod: "POST",
			reqURL:   "http://flax/v1/sessions",
			reqQueries: map[string]string{
				"tenantId": "11111111-1111-1111-1111-111111111111",
				"deviceId": "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa",
			},
			reqHeaders: map[string]string{
				"Accept":       "application/json",
				"Content-Type": "application/json",
			},
			resStatusCode: 201,
			resHeaders: map[string]string{
				"Content-Type": "application/json",
			},
			resBody: JSON{
				"id": "00000000-0000-0000-0000-000000000000",
			},
		},
		{
			name: "WithHTTPResponseAndPrefix",
			mock: HTTPMock{
				HTTPExpect: HTTPExpect{
					Methods: []string{"POST"},
					Path:    "/v1/sessions",
					Prefix:  true,
					Queries: map[string]string{
						"tenantId": "[0-9A-Za-z-]+",
						"deviceId": "[0-9A-Za-z-]+",
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
						"id": "00000000-0000-0000-0000-000000000000",
					},
				},
			},
			reqMehod: "POST",
			reqURL:   "http://flax/v1/sessions/internal",
			reqQueries: map[string]string{
				"tenantId": "11111111-1111-1111-1111-111111111111",
				"deviceId": "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa",
			},
			reqHeaders: map[string]string{
				"Accept":       "application/json",
				"Content-Type": "application/json",
			},
			resStatusCode: 201,
			resHeaders: map[string]string{
				"Content-Type": "application/json",
			},
			resBody: JSON{
				"id": "00000000-0000-0000-0000-000000000000",
			},
		},
		{
			name: "WithHTTPForward",
			mock: HTTPMock{
				HTTPExpect: HTTPExpect{
					Methods: []string{"POST"},
					Path:    "/v1/sessions",
					Prefix:  true,
					Queries: map[string]string{
						"tenantId": "[0-9A-Za-z-]+",
						"deviceId": "[0-9A-Za-z-]+",
					},
					Headers: map[string]string{
						"Accept":       "application/json",
						"Content-Type": "application/json",
					},
				},
				HTTPForward: &HTTPForward{
					Delay:   "10ms",
					To:      "https://auth-service",
					Headers: map[string]string{},
				},
			},
			reqMehod: "POST",
			reqURL:   "http://flax/v1/sessions",
			reqQueries: map[string]string{
				"tenantId": "11111111-1111-1111-1111-111111111111",
				"deviceId": "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa",
			},
			reqHeaders: map[string]string{
				"Accept":       "application/json",
				"Content-Type": "application/json",
			},
			resStatusCode: 501,
			resHeaders:    map[string]string{},
			resBody: JSON{
				"message": "this functionality is not yet available!",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			router := mux.NewRouter()
			tc.mock.RegisterRoutes(router)

			req, err := http.NewRequest(tc.reqMehod, tc.reqURL, nil)
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

			assert.Equal(t, tc.resStatusCode, res.Result().StatusCode)
			for key, val := range tc.resHeaders {
				assert.Equal(t, val, res.Header().Get(key))
			}

			resBody := JSON{}
			err = json.NewDecoder(res.Body).Decode(&resBody)
			assert.NoError(t, err)
			assert.Equal(t, tc.resBody, resBody)
		})
	}
}
