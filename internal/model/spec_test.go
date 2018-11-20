package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	specSimple = &Spec{
		HTTPMocks: []HTTPMock{
			HTTPMock{
				HTTPExpect{},
				&HTTPResponse{},
				&HTTPForward{},
			},
			HTTPMock{
				HTTPExpect{},
				&HTTPResponse{},
				&HTTPForward{},
			},
			HTTPMock{
				HTTPExpect{},
				&HTTPResponse{},
				&HTTPForward{},
			},
		},
		RESTMocks: []RESTMock{
			RESTMock{
				RESTExpect{},
				RESTResponse{},
				RESTStore{},
			},
			RESTMock{
				RESTExpect{},
				RESTResponse{},
				RESTStore{},
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
				HTTPExpect{},
				&HTTPResponse{},
				&HTTPForward{},
			},
			HTTPMock{
				HTTPExpect{},
				&HTTPResponse{},
				&HTTPForward{},
			},
		},
		RESTMocks: []RESTMock{
			RESTMock{
				RESTExpect{},
				RESTResponse{},
				RESTStore{},
			},
			RESTMock{
				RESTExpect{},
				RESTResponse{},
				RESTStore{},
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
				HTTPExpect{},
				&HTTPResponse{},
				&HTTPForward{},
			},
			HTTPMock{
				HTTPExpect{},
				&HTTPResponse{},
				&HTTPForward{},
			},
		},
		RESTMocks: []RESTMock{
			RESTMock{
				RESTExpect{},
				RESTResponse{},
				RESTStore{},
			},
			RESTMock{
				RESTExpect{},
				RESTResponse{},
				RESTStore{},
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
		/* {
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
		}, */
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
