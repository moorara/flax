package v1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	specSimpleYAML = &Spec{
		HTTPMocks: []HTTPMock{
			HTTPMock{
				HTTPExpect: HTTPExpect{
					Methods: []string{"GET"},
					Path:    "/health",
					Prefix:  false,
					Queries: map[string]string{},
					Headers: map[string]string{},
				},
				HTTPResponse: &HTTPResponse{
					Delay:      "0",
					StatusCode: 200,
					Headers:    map[string]string{},
				},
			},
			HTTPMock{
				HTTPExpect: HTTPExpect{
					Methods: []string{"GET"},
					Path:    "/me",
					Prefix:  false,
					Queries: map[string]string{},
					Headers: map[string]string{},
				},
				HTTPResponse: &HTTPResponse{
					Delay:      "0",
					StatusCode: 200,
					Headers:    map[string]string{},
					Body: map[interface{}]interface{}{
						"id":    "5da8349a-0707-4064-8fad-74cedb48a8fc",
						"name":  "John Doe",
						"email": "john.doe@example.com",
					},
				},
			},
			HTTPMock{
				HTTPExpect: HTTPExpect{
					Methods: []string{"GET"},
					Path:    "/v1/sessions",
					Prefix:  true,
					Queries: map[string]string{},
					Headers: map[string]string{},
				},
				HTTPForward: &HTTPForward{
					Delay:   "0",
					To:      "http://session-manager:8800/api/v1/sessions",
					Headers: map[string]string{},
				},
			},
		},
		RESTMocks: []RESTMock{
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
					Headers:          map[string]string{},
				},
				RESTStore{
					Objects: []JSON{
						{"id": "ad2bd67b-172e-4778-a8a3-7cfb626685b9", "make": "Mazda", "model": "CX-5"},
						{"id": "26ee9c87-fdbb-48cf-be0a-add9a3d87189", "make": "Hyundai", "model": "Sonata"},
					},
					Directory: map[interface{}]JSON{
						"ad2bd67b-172e-4778-a8a3-7cfb626685b9": {"id": "ad2bd67b-172e-4778-a8a3-7cfb626685b9", "make": "Mazda", "model": "CX-5"},
						"26ee9c87-fdbb-48cf-be0a-add9a3d87189": {"id": "26ee9c87-fdbb-48cf-be0a-add9a3d87189", "make": "Hyundai", "model": "Sonata"},
					},
				},
			},
			RESTMock{
				RESTExpect{
					BasePath: "/api/v1/teams",
					Headers:  map[string]string{},
				},
				RESTResponse{
					Delay:            "0",
					GetStatusCode:    200,
					PostStatusCode:   201,
					PutStatusCode:    200,
					PatchStatusCode:  200,
					DeleteStatusCode: 204,
					Headers:          map[string]string{},
				},
				RESTStore{
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
		},
	}

	specSimpleJSON = &Spec{
		HTTPMocks: []HTTPMock{
			HTTPMock{
				HTTPExpect: HTTPExpect{
					Methods: []string{"GET"},
					Path:    "/health",
					Prefix:  false,
					Queries: map[string]string{},
					Headers: map[string]string{},
				},
				HTTPResponse: &HTTPResponse{
					Delay:      "0",
					StatusCode: 200,
					Headers:    map[string]string{},
				},
			},
			HTTPMock{
				HTTPExpect: HTTPExpect{
					Methods: []string{"GET"},
					Path:    "/me",
					Prefix:  false,
					Queries: map[string]string{},
					Headers: map[string]string{},
				},
				HTTPResponse: &HTTPResponse{
					Delay:      "0",
					StatusCode: 200,
					Headers:    map[string]string{},
					Body: map[string]interface{}{
						"id":    "5da8349a-0707-4064-8fad-74cedb48a8fc",
						"name":  "John Doe",
						"email": "john.doe@example.com",
					},
				},
			},
			HTTPMock{
				HTTPExpect: HTTPExpect{
					Methods: []string{"GET"},
					Path:    "/v1/sessions",
					Prefix:  true,
					Queries: map[string]string{},
					Headers: map[string]string{},
				},
				HTTPForward: &HTTPForward{
					Delay:   "0",
					To:      "http://session-manager:8800/api/v1/sessions",
					Headers: map[string]string{},
				},
			},
		},
		RESTMocks: []RESTMock{
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
					Headers:          map[string]string{},
				},
				RESTStore{
					Objects: []JSON{
						{"id": "ad2bd67b-172e-4778-a8a3-7cfb626685b9", "make": "Mazda", "model": "CX-5"},
						{"id": "26ee9c87-fdbb-48cf-be0a-add9a3d87189", "make": "Hyundai", "model": "Sonata"},
					},
					Directory: map[interface{}]JSON{
						"ad2bd67b-172e-4778-a8a3-7cfb626685b9": {"id": "ad2bd67b-172e-4778-a8a3-7cfb626685b9", "make": "Mazda", "model": "CX-5"},
						"26ee9c87-fdbb-48cf-be0a-add9a3d87189": {"id": "26ee9c87-fdbb-48cf-be0a-add9a3d87189", "make": "Hyundai", "model": "Sonata"},
					},
				},
			},
			RESTMock{
				RESTExpect{
					BasePath: "/api/v1/teams",
					Headers:  map[string]string{},
				},
				RESTResponse{
					Delay:            "0",
					GetStatusCode:    200,
					PostStatusCode:   201,
					PutStatusCode:    200,
					PatchStatusCode:  200,
					DeleteStatusCode: 204,
					Headers:          map[string]string{},
				},
				RESTStore{
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
		},
	}

	specFullYAML = &Spec{
		Config: Config{
			HTTPPort:  8080,
			HTTPSPort: 8443,
		},
		HTTPMocks: []HTTPMock{
			HTTPMock{
				HTTPExpect: HTTPExpect{
					Methods: []string{"GET"},
					Path:    "/me",
					Prefix:  false,
					Queries: map[string]string{
						"tenantId": "\\w+",
					},
					Headers: map[string]string{
						"Authorization": "Bearer .*",
					},
				},
				HTTPResponse: &HTTPResponse{
					Delay:      "10ms",
					StatusCode: 200,
					Headers: map[string]string{
						"Content-Type": "application/json",
					},
					Body: map[interface{}]interface{}{
						"id":    "5da8349a-0707-4064-8fad-74cedb48a8fc",
						"name":  "John Doe",
						"email": "john.doe@example.com",
					},
				},
			},
			HTTPMock{
				HTTPExpect: HTTPExpect{
					Methods: []string{"POST", "PUT"},
					Path:    "/v1/sessions",
					Prefix:  true,
					Queries: map[string]string{
						"tenantId": "\\w+",
					},
					Headers: map[string]string{
						"Accept":       "application/json",
						"Content-Type": "application/json",
					},
				},
				HTTPForward: &HTTPForward{
					Delay: "10ms",
					To:    "http://session-manager:8800/api/v1/sessions",
					Headers: map[string]string{
						"Referer": "flax",
					},
				},
			},
		},
		RESTMocks: []RESTMock{
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
					Objects: []JSON{
						{"id": "ad2bd67b-172e-4778-a8a3-7cfb626685b9", "make": "Mazda", "model": "CX-5"},
						{"id": "26ee9c87-fdbb-48cf-be0a-add9a3d87189", "make": "Hyundai", "model": "Sonata"},
					},
					Directory: map[interface{}]JSON{
						"ad2bd67b-172e-4778-a8a3-7cfb626685b9": {"id": "ad2bd67b-172e-4778-a8a3-7cfb626685b9", "make": "Mazda", "model": "CX-5"},
						"26ee9c87-fdbb-48cf-be0a-add9a3d87189": {"id": "26ee9c87-fdbb-48cf-be0a-add9a3d87189", "make": "Hyundai", "model": "Sonata"},
					},
				},
			},
			RESTMock{
				RESTExpect{
					BasePath: "/api/v1/teams",
					Headers: map[string]string{
						"Authorization": "Bearer .*",
					},
				},
				RESTResponse{
					Delay:            "10ms",
					GetStatusCode:    206,
					PostStatusCode:   202,
					PutStatusCode:    202,
					PatchStatusCode:  202,
					DeleteStatusCode: 202,
					ListProperty:     "data",
					Headers: map[string]string{
						"Content-Type": "application/json",
					},
				},
				RESTStore{
					Identifier: "_id",
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
		},
	}

	specFullJSON = &Spec{
		Config: Config{
			HTTPPort:  8080,
			HTTPSPort: 8443,
		},
		HTTPMocks: []HTTPMock{
			HTTPMock{
				HTTPExpect: HTTPExpect{
					Methods: []string{"GET"},
					Path:    "/me",
					Prefix:  false,
					Queries: map[string]string{
						"tenantId": "\\w+",
					},
					Headers: map[string]string{
						"Authorization": "Bearer .*",
					},
				},
				HTTPResponse: &HTTPResponse{
					Delay:      "10ms",
					StatusCode: 200,
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
			HTTPMock{
				HTTPExpect: HTTPExpect{
					Methods: []string{"POST", "PUT"},
					Path:    "/v1/sessions",
					Prefix:  true,
					Queries: map[string]string{
						"tenantId": "\\w+",
					},
					Headers: map[string]string{
						"Accept":       "application/json",
						"Content-Type": "application/json",
					},
				},
				HTTPForward: &HTTPForward{
					Delay: "10ms",
					To:    "http://session-manager:8800/api/v1/sessions",
					Headers: map[string]string{
						"Referer": "flax",
					},
				},
			},
		},
		RESTMocks: []RESTMock{
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
					Objects: []JSON{
						{"id": "ad2bd67b-172e-4778-a8a3-7cfb626685b9", "make": "Mazda", "model": "CX-5"},
						{"id": "26ee9c87-fdbb-48cf-be0a-add9a3d87189", "make": "Hyundai", "model": "Sonata"},
					},
					Directory: map[interface{}]JSON{
						"ad2bd67b-172e-4778-a8a3-7cfb626685b9": {"id": "ad2bd67b-172e-4778-a8a3-7cfb626685b9", "make": "Mazda", "model": "CX-5"},
						"26ee9c87-fdbb-48cf-be0a-add9a3d87189": {"id": "26ee9c87-fdbb-48cf-be0a-add9a3d87189", "make": "Hyundai", "model": "Sonata"},
					},
				},
			},
			RESTMock{
				RESTExpect{
					BasePath: "/api/v1/teams",
					Headers: map[string]string{
						"Authorization": "Bearer .*",
					},
				},
				RESTResponse{
					Delay:            "10ms",
					GetStatusCode:    206,
					PostStatusCode:   202,
					PutStatusCode:    202,
					PatchStatusCode:  202,
					DeleteStatusCode: 202,
					ListProperty:     "data",
					Headers: map[string]string{
						"Content-Type": "application/json",
					},
				},
				RESTStore{
					Identifier: "_id",
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
		},
	}
)

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
			expectedError: "no such file or directory",
			expectedSpec:  nil,
		},
		{
			name:          "UnknownFormat",
			path:          "./test/unknown",
			expectedError: "unknown file format",
			expectedSpec:  nil,
		},
		{
			name:          "EmptyYAML",
			path:          "./test/empty.yaml",
			expectedError: "EOF",
			expectedSpec:  nil,
		},
		{
			name:          "EmptyJSON",
			path:          "./test/empty.json",
			expectedError: "EOF",
			expectedSpec:  nil,
		},
		{
			name:          "InvalidYAML",
			path:          "./test/error.yaml",
			expectedError: "cannot unmarshal",
			expectedSpec:  nil,
		},
		{
			name:          "InvalidJSON",
			path:          "./test/error.json",
			expectedError: "invalid character",
			expectedSpec:  nil,
		},
		{
			name:          "SimpleYAML",
			path:          "./test/simple.yaml",
			expectedError: "",
			expectedSpec:  specSimpleYAML,
		},
		{
			name:          "SimpleJSON",
			path:          "./test/simple.json",
			expectedError: "",
			expectedSpec:  specSimpleJSON,
		},
		{
			name:          "FullYAML",
			path:          "./test/full.yaml",
			expectedError: "",
			expectedSpec:  specFullYAML,
		},
		{
			name:          "FullJSON",
			path:          "./test/full.json",
			expectedError: "",
			expectedSpec:  specFullJSON,
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
