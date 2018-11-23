package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	gotoConfig "github.com/moorara/goto/config"
	"github.com/stretchr/testify/assert"
)

type (
	JSON map[string]interface{}
)

var config = struct {
	SUTAddress string
}{
	SUTAddress: "http://localhost:10000",
}

func init() {
	gotoConfig.Pick(&config)
}

func isComponentTest() bool {
	val := os.Getenv("COMPONENT_TEST")
	val = strings.ToLower(val)
	return val == "true"
}

func TestMOCKHTTP(t *testing.T) {
	if !isComponentTest() {
		t.Skip()
	}

	tests := []struct {
		name               string
		method             string
		endpoint           string
		body               JSON
		expectedStatusCode int
		expectedResponse   JSON
	}{
		{
			name:     "MockHTTPResponse",
			method:   "PUT",
			endpoint: "/api/mock/http",
			body: JSON{
				"methods": []string{"GET"},
				"path":    "/me",
				"response": JSON{
					"delay":  "10ms",
					"status": 200,
					"body": JSON{
						"id":    "11111111-1111-1111-1111-111111111111",
						"name":  "John Doe",
						"email": "john.doe@example.com",
					},
				},
			},
			expectedStatusCode: 200,
			expectedResponse: JSON{
				"handle":  "aaaaaaaa",
				"methods": []string{"GET"},
				"path":    "/me",
			},
		},
		{
			name:     "MockHTTPForward",
			method:   "PUT",
			endpoint: "/api/mock/http",
			body: JSON{
				"methods": []string{"GET"},
				"path":    "/me",
				"forward": JSON{
					"delay": "10ms",
					"to":    "http://user-service:9000/me",
				},
			},
			expectedStatusCode: 200,
			expectedResponse: JSON{
				"handle":  "bbbbbbbb",
				"methods": []string{"GET"},
				"path":    "/me",
			},
		},
		{
			name:     "MockHTTPResponseAndForward",
			method:   "PUT",
			endpoint: "/api/mock/http",
			body: JSON{
				"methods": []string{"GET"},
				"path":    "/me",
				"response": JSON{
					"delay":  "10ms",
					"status": 200,
					"body": JSON{
						"id":    "11111111-1111-1111-1111-111111111111",
						"name":  "John Doe",
						"email": "john.doe@example.com",
					},
				},
				"forward": JSON{
					"delay": "10ms",
					"to":    "http://user-service:9000/me",
				},
			},
			expectedStatusCode: 200,
			expectedResponse: JSON{
				"handle":  "cccccccc",
				"methods": []string{"GET"},
				"path":    "/me",
			},
		},
		{
			name:               "VerifyHTTPByHandle",
			method:             "GET",
			endpoint:           "/api/mock/verify/aaaaaaaa",
			body:               JSON{},
			expectedStatusCode: 200,
			expectedResponse: JSON{
				"matches": []JSON{},
			},
		},
		{
			name:     "VerifyHTTPByRequest",
			method:   "GET",
			endpoint: "/api/mock/verify",
			body: JSON{
				"method": "GET",
				"path":   "/me",
			},
			expectedStatusCode: 200,
			expectedResponse: JSON{
				"matches": []JSON{},
			},
		},
	}

	client := &http.Client{
		Transport: &http.Transport{},
		Timeout:   2 * time.Second,
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			buff := new(bytes.Buffer)
			err := json.NewEncoder(buff).Encode(tc.body)
			assert.NoError(t, err)

			url := filepath.Join(config.SUTAddress, tc.endpoint)
			req, err := http.NewRequest(tc.method, url, buff)
			assert.NoError(t, err)

			req.Header.Set("Content-Type", "application/json")

			res, err := client.Do(req)
			assert.NoError(t, err)

			assert.Equal(t, tc.expectedStatusCode, res.StatusCode)

			response := JSON{}
			err = json.NewDecoder(res.Body).Decode(response)
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedResponse, response)
		})
	}
}

func TestMOCKREST(t *testing.T) {
	if !isComponentTest() {
		t.Skip()
	}

	tests := []struct {
		name               string
		method             string
		endpoint           string
		body               JSON
		expectedStatusCode int
		expectedResponse   JSON
	}{
		{
			name:     "MockREST",
			method:   "PUT",
			endpoint: "/api/mock/rest",
			body: JSON{
				"basePath": "/api/v1/teams",
				"response": JSON{
					"delay":        "10ms",
					"listProperty": "data",
				},
				"store": JSON{
					"objects": []JSON{
						JSON{"id": "11111111-1111-1111-1111-111111111111", "name": "Back-end", "tags": []string{"cloud", "go"}},
						JSON{"id": "22222222-2222-2222-2222-222222222222", "name": "Front-end", "tags": []string{"react", "redux"}},
					},
				},
			},
			expectedStatusCode: 200,
			expectedResponse: JSON{
				"handle":   "aaaaaaaa",
				"basePath": "/api/v1/teams",
			},
		},
		{
			name:               "VerifyRESTByHandle",
			method:             "GET",
			endpoint:           "/api/mock/verify/aaaaaaaa",
			body:               JSON{},
			expectedStatusCode: 200,
			expectedResponse: JSON{
				"matches": []JSON{},
			},
		},
		{
			name:     "VerifyRESTByGET",
			method:   "GET",
			endpoint: "/api/mock/verify",
			body: JSON{
				"method": "GET",
				"path":   "/api/v1/teams",
			},
			expectedStatusCode: 200,
			expectedResponse: JSON{
				"matches": []JSON{},
			},
		},
		{
			name:     "VerifyRESTByPOST",
			method:   "GET",
			endpoint: "/api/mock/verify",
			body: JSON{
				"method": "POST",
				"path":   "/api/v1/teams",
			},
			expectedStatusCode: 200,
			expectedResponse: JSON{
				"matches": []JSON{},
			},
		},
		{
			name:     "VerifyRESTByGET",
			method:   "GET",
			endpoint: "/api/mock/verify",
			body: JSON{
				"method": "GET",
				"path":   "/api/v1/teams/11111111-1111-1111-1111-111111111111",
			},
			expectedStatusCode: 200,
			expectedResponse: JSON{
				"matches": []JSON{},
			},
		},
		{
			name:     "VerifyRESTByPUT",
			method:   "GET",
			endpoint: "/api/mock/verify",
			body: JSON{
				"method": "PUT",
				"path":   "/api/v1/teams/11111111-1111-1111-1111-111111111111",
			},
			expectedStatusCode: 200,
			expectedResponse: JSON{
				"matches": []JSON{},
			},
		},
		{
			name:     "VerifyRESTByPATCH",
			method:   "GET",
			endpoint: "/api/mock/verify",
			body: JSON{
				"method": "PATCH",
				"path":   "/api/v1/teams/11111111-1111-1111-1111-111111111111",
			},
			expectedStatusCode: 200,
			expectedResponse: JSON{
				"matches": []JSON{},
			},
		},
		{
			name:     "VerifyRESTByDELETE",
			method:   "GET",
			endpoint: "/api/mock/verify",
			body: JSON{
				"method": "DELETE",
				"path":   "/api/v1/teams/11111111-1111-1111-1111-111111111111",
			},
			expectedStatusCode: 200,
			expectedResponse: JSON{
				"matches": []JSON{},
			},
		},
	}

	client := &http.Client{
		Transport: &http.Transport{},
		Timeout:   2 * time.Second,
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			buff := new(bytes.Buffer)
			err := json.NewEncoder(buff).Encode(tc.body)
			assert.NoError(t, err)

			url := filepath.Join(config.SUTAddress, tc.endpoint)
			req, err := http.NewRequest(tc.method, url, buff)
			assert.NoError(t, err)

			req.Header.Set("Content-Type", "application/json")

			res, err := client.Do(req)
			assert.NoError(t, err)

			assert.Equal(t, tc.expectedStatusCode, res.StatusCode)

			response := JSON{}
			err = json.NewDecoder(res.Body).Decode(response)
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedResponse, response)
		})
	}
}
