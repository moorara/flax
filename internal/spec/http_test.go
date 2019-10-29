package spec

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

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
			HTTPExpect{
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
			true,
		},
		{
			"NotEqual",
			HTTPExpect{
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
			HTTPExpect{
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
				Queries: nil,
				Headers: nil,
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
				Queries: nil,
				Headers: nil,
			},
		},
		{
			"NoDefaultRequired",
			HTTPExpect{
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
			HTTPExpect{
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
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.expect.SetDefaults()
			assert.Equal(t, tc.expectedExpect, tc.expect)
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
				Delay:      "",
				StatusCode: 200,
				Headers:    nil,
				Body:       nil,
			},
		},
		{
			"DefaultRequired",
			HTTPResponse{
				Delay: "10ms",
			},
			HTTPResponse{
				Delay:      "10ms",
				StatusCode: 200,
				Headers:    nil,
				Body:       nil,
			},
		},
		{
			"NoDefaultRequired",
			HTTPResponse{
				Delay:      "10ms",
				StatusCode: 201,
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
				Body: JSON{
					"id": "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa",
				},
			},
			HTTPResponse{
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
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.response.SetDefaults()
			assert.Equal(t, tc.expectedResponse, tc.response)
		})
	}
}

func TestHTTPResponseHandler(t *testing.T) {
	tests := []struct {
		name               string
		response           HTTPResponse
		expectedStatusCode int
		expectedHeaders    map[string]string
		expectedBody       JSON
	}{
		{
			name: "OK",
			response: HTTPResponse{
				Delay:      "10ms",
				StatusCode: 201,
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
				Body: JSON{
					"id": "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa",
				},
			},
			expectedStatusCode: 201,
			expectedHeaders: map[string]string{
				"Content-Type": "application/json",
			},
			expectedBody: JSON{
				"id": "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			handler := tc.response.Handler()

			res := httptest.NewRecorder()
			handler(res, nil)

			assert.Equal(t, tc.expectedStatusCode, res.Result().StatusCode)
			for key, val := range tc.expectedHeaders {
				assert.Equal(t, val, res.Header().Get(key))
			}

			resBody := JSON{}
			err := json.NewDecoder(res.Body).Decode(&resBody)
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedBody, resBody)
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
				Delay:   "",
				To:      "",
				Headers: nil,
			},
		},
		{
			"NoDefaultRequired",
			HTTPForward{
				Delay: "10ms",
				To:    "http://example.com",
				Headers: map[string]string{
					"Is-Test": "true",
				},
			},
			HTTPForward{
				Delay: "10ms",
				To:    "http://example.com",
				Headers: map[string]string{
					"Is-Test": "true",
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

func TestHTTPForwardHandler(t *testing.T) {
	tests := []struct {
		name               string
		forward            HTTPForward
		expectedStatusCode int
		expectedHeaders    map[string]string
		expectedBody       JSON
	}{
		{
			name: "OK",
			forward: HTTPForward{
				Delay: "10ms",
				To:    "http://example.com",
				Headers: map[string]string{
					"Is-Test": "true",
				},
			},
			expectedStatusCode: 501,
			expectedHeaders: map[string]string{
				"Is-Test": "true",
			},
			expectedBody: JSON{
				"message": "this functionality is not yet available!",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			handler := tc.forward.Handler()

			res := httptest.NewRecorder()
			handler(res, nil)

			assert.Equal(t, tc.expectedStatusCode, res.Result().StatusCode)
			for key, val := range tc.expectedHeaders {
				assert.Equal(t, val, res.Header().Get(key))
			}

			resBody := JSON{}
			err := json.NewDecoder(res.Body).Decode(&resBody)
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedBody, resBody)
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

func TestHTTPMockRegisterRoutes(t *testing.T) {
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

func TestDefaultHTTPMock(t *testing.T) {
	tests := []struct {
		name         string
		expectedMock HTTPMock
	}{
		{
			"OK",
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
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mock := DefaultHTTPMock()
			assert.Equal(t, tc.expectedMock, mock)
		})
	}
}
