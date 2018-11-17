package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	specSimple = &Spec{
		HTTPMocks: []HTTPMock{
			HTTPMock{
				Methods:    []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
				Path:       "/health",
				Queries:    map[string]string{},
				ReqHeaders: map[string]string{},
				Delay:      "0",
				StatusCode: 200,
				ResHeaders: map[string]string{},
			},
		},
		RESTMock: []RESTMock{
			RESTMock{
				BasePath:   "/cars",
				ReqHeaders: map[string]string{},
				Delay:      "0",
				ResHeaders: map[string]string{},
				Identifier: "",
				ListHandle: "",
				Store: []JSON{
					JSON{"id": "ad2bd67b-172e-4778-a8a3-7cfb626685b9", "make": "Mazda", "model": "CX-5"},
					JSON{"id": "26ee9c87-fdbb-48cf-be0a-add9a3d87189", "make": "Hyundai", "model": "Sonata"},
				},
			},
			RESTMock{
				BasePath:   "/teams",
				ReqHeaders: map[string]string{},
				Delay:      "0",
				ResHeaders: map[string]string{},
				Identifier: "",
				ListHandle: "",
				Store: []JSON{
					JSON{"_id": "d93ce179-50f7-469e-bb36-1b3746145f00", "name": "Back-end", "tags": []interface{}{"go", "cloud"}},
					JSON{"_id": "8cd6ef6c-2095-4c75-bc66-6f38e785299d", "name": "Front-end", "tags": []interface{}{"react", "redux"}},
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
				Methods:    []string{"GET"},
				Path:       "/health",
				Queries:    map[string]string{},
				ReqHeaders: map[string]string{},
				Delay:      "0",
				StatusCode: 200,
				ResHeaders: map[string]string{},
			},
			HTTPMock{
				Methods:    []string{"GET"},
				Path:       "/current/user",
				Queries:    map[string]string{"type": "\\w+"},
				ReqHeaders: map[string]string{"Authorization": "Bearer .*"},
				Delay:      "100ms",
				StatusCode: 200,
				ResHeaders: map[string]string{"Token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"},
				Body: map[interface{}]interface{}{
					"id":    "5da8349a-0707-4064-8fad-74cedb48a8fc",
					"name":  "John Doe",
					"email": "john.doe@example.com",
				},
			},
		},
		RESTMock: []RESTMock{
			RESTMock{
				BasePath:   "/api/v1/cars",
				ReqHeaders: map[string]string{},
				Delay:      "0",
				ResHeaders: map[string]string{},
				Identifier: "",
				ListHandle: "",
				Store: []JSON{
					JSON{"id": "ad2bd67b-172e-4778-a8a3-7cfb626685b9", "make": "Mazda", "model": "CX-5"},
					JSON{"id": "26ee9c87-fdbb-48cf-be0a-add9a3d87189", "make": "Hyundai", "model": "Sonata"},
				},
			},
			RESTMock{
				BasePath:   "/api/v1/teams",
				ReqHeaders: map[string]string{"Authorization": "Bearer .*"},
				Delay:      "100ms",
				ResHeaders: map[string]string{"Token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"},
				Identifier: "_id",
				ListHandle: "data",
				Store: []JSON{
					JSON{"_id": "d93ce179-50f7-469e-bb36-1b3746145f00", "name": "Back-end", "tags": []interface{}{"go", "cloud"}},
					JSON{"_id": "8cd6ef6c-2095-4c75-bc66-6f38e785299d", "name": "Front-end", "tags": []interface{}{"react", "redux"}},
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
				Methods:    []string{"GET"},
				Path:       "/health",
				Queries:    map[string]string{},
				ReqHeaders: map[string]string{},
				Delay:      "0",
				StatusCode: 200,
				ResHeaders: map[string]string{},
			},
			HTTPMock{
				Methods:    []string{"GET"},
				Path:       "/current/user",
				Queries:    map[string]string{"type": "\\w+"},
				ReqHeaders: map[string]string{"Authorization": "Bearer .*"},
				Delay:      "100ms",
				StatusCode: 200,
				ResHeaders: map[string]string{"Token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"},
				Body: map[string]interface{}{
					"id":    "5da8349a-0707-4064-8fad-74cedb48a8fc",
					"name":  "John Doe",
					"email": "john.doe@example.com",
				},
			},
		},
		RESTMock: []RESTMock{
			RESTMock{
				BasePath:   "/api/v1/cars",
				ReqHeaders: map[string]string{},
				Delay:      "0",
				ResHeaders: map[string]string{},
				Identifier: "",
				ListHandle: "",
				Store: []JSON{
					JSON{"id": "ad2bd67b-172e-4778-a8a3-7cfb626685b9", "make": "Mazda", "model": "CX-5"},
					JSON{"id": "26ee9c87-fdbb-48cf-be0a-add9a3d87189", "make": "Hyundai", "model": "Sonata"},
				},
			},
			RESTMock{
				BasePath:   "/api/v1/teams",
				ReqHeaders: map[string]string{"Authorization": "Bearer .*"},
				Delay:      "100ms",
				ResHeaders: map[string]string{"Token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"},
				Identifier: "_id",
				ListHandle: "data",
				Store: []JSON{
					JSON{"_id": "d93ce179-50f7-469e-bb36-1b3746145f00", "name": "Back-end", "tags": []interface{}{"go", "cloud"}},
					JSON{"_id": "8cd6ef6c-2095-4c75-bc66-6f38e785299d", "name": "Front-end", "tags": []interface{}{"react", "redux"}},
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
			path:          "./test/unknown.toml",
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
			expectedSpec:  specSimple,
		},
		{
			name:          "SimpleJSON",
			path:          "./test/simple.json",
			expectedError: "",
			expectedSpec:  specSimple,
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
