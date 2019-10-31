package spec

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	defaultSpec = &Spec{
		Config{
			HTTPPort:  8080,
			HTTPSPort: 8443,
		},
		[]HTTPMock{
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
		[]RESTMock{},
	}

	specSimple = &Spec{
		Config: Config{
			HTTPPort:  8080,
			HTTPSPort: 8443,
		},
		HTTPMocks: []HTTPMock{
			HTTPMock{
				HTTPExpect: HTTPExpect{
					Methods: []string{"GET"},
					Path:    "/health",
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
			HTTPMock{
				HTTPExpect: HTTPExpect{
					Methods: []string{"POST", "PUT"},
					Path:    "/api/v1/sendMessage",
					Prefix:  false,
					Queries: nil,
					Headers: nil,
				},
				HTTPResponse: &HTTPResponse{
					Delay:      "",
					StatusCode: 201,
					Headers:    nil,
					Body: map[string]interface{}{
						"id": "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa",
					},
				},
			},
		},
		RESTMocks: []RESTMock{
			RESTMock{
				RESTExpect{
					BasePath: "/api/v1/teams",
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
					Objects: []JSON{
						{"_id": "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa", "name": "Back-end"},
						{"_id": "bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb", "name": "Front-end"},
					},
					Directory: map[interface{}]JSON{
						"aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa": {"_id": "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa", "name": "Back-end"},
						"bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb": {"_id": "bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb", "name": "Front-end"},
					},
				},
			},
		},
	}

	specFull = &Spec{
		Config: Config{
			HTTPPort:  9080,
			HTTPSPort: 9443,
		},
		HTTPMocks: []HTTPMock{
			HTTPMock{
				HTTPExpect: HTTPExpect{
					Methods: []string{"GET"},
					Path:    "/health",
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
			HTTPMock{
				HTTPExpect: HTTPExpect{
					Methods: []string{"GET"},
					Path:    "/app",
					Prefix:  false,
					Queries: nil,
					Headers: nil,
				},
				HTTPForward: &HTTPForward{
					Delay: "",
					To:    "http://example.com",
					Headers: map[string]string{
						"Is-Test": "true",
					},
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
						"Accept":        "application/json",
						"Content-Type":  "application/json",
						"Authorization": "Bearer .*",
					},
				},
				HTTPResponse: &HTTPResponse{
					Delay:      "10ms",
					StatusCode: 201,
					Headers: map[string]string{
						"Content-Type": "application/json",
					},
					Body: map[string]interface{}{
						"id": "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa",
					},
				},
			},
		},
		RESTMocks: []RESTMock{
			RESTMock{
				RESTExpect{
					BasePath: "/api/v1/teams",
					Headers: map[string]string{
						"Accept":        "application/json",
						"Content-Type":  "application/json",
						"Authorization": "Bearer .*",
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
					Objects: []JSON{
						{"_id": "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa", "name": "Back-end"},
						{"_id": "bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb", "name": "Front-end"},
					},
					Directory: map[interface{}]JSON{
						"aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa": {"_id": "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa", "name": "Back-end"},
						"bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb": {"_id": "bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb", "name": "Front-end"},
					},
				},
			},
		},
	}
)

func TestDefaultSpec(t *testing.T) {
	tests := []struct {
		name         string
		expectedSpec *Spec
	}{
		{
			"OK",
			defaultSpec,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			spec := DefaultSpec()
			assert.Equal(t, tc.expectedSpec, spec)
		})
	}
}

func TestReadSpec(t *testing.T) {
	tests := []struct {
		name          string
		path          string
		expectedError string
		expectedSpec  *Spec
	}{
		{
			name:          "NoFile",
			path:          "./test/spec",
			expectedError: "",
			expectedSpec:  defaultSpec,
		},
		{
			name:          "UnknownFormat",
			path:          "./test/unknown",
			expectedError: "unknown spec file",
			expectedSpec:  nil,
		},
		{
			name:          "EmptyJSON",
			path:          "./test/empty.json",
			expectedError: "unknown spec file",
			expectedSpec:  nil,
		},
		{
			name:          "EmptyYAML",
			path:          "./test/empty.yaml",
			expectedError: "unknown spec file",
			expectedSpec:  nil,
		},
		{
			name:          "InvalidJSON",
			path:          "./test/invalid.json",
			expectedError: "unknown spec file",
			expectedSpec:  nil,
		},
		{
			name:          "InvalidYAML",
			path:          "./test/invalid.yaml",
			expectedError: "unknown spec file",
			expectedSpec:  nil,
		},
		{
			name:          "SimpleJSON",
			path:          "./test/simple.json",
			expectedError: "",
			expectedSpec:  specSimple,
		},
		{
			name:          "SimpleYAML",
			path:          "./test/simple.yaml",
			expectedError: "",
			expectedSpec:  specSimple,
		},
		{
			name:          "FullJSON",
			path:          "./test/full.json",
			expectedError: "",
			expectedSpec:  specFull,
		},
		{
			name:          "FullYAML",
			path:          "./test/full.yaml",
			expectedError: "",
			expectedSpec:  specFull,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			spec, err := ReadSpec(tc.path)

			if tc.expectedError == "" {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedSpec, spec)
			} else {
				assert.Contains(t, err.Error(), tc.expectedError)
				assert.Nil(t, spec)
			}
		})
	}
}
