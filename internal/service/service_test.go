package service

import (
	"testing"

	"github.com/moorara/flax/internal/spec"

	"github.com/stretchr/testify/assert"
)

func TestNewRESTService(t *testing.T) {
	tests := []struct {
		name                string
		file                string
		expectedError       string
		expectedRESTService *restService
	}{
		{
			name:                "NotExist",
			file:                "./test/nofile.json",
			expectedError:       "no such file or directory",
			expectedRESTService: nil,
		},
		{
			name:                "EmptyFile",
			file:                "./test/empty.json",
			expectedError:       "EOF",
			expectedRESTService: nil,
		},
		{
			name:                "InvalidFile",
			file:                "./test/error.json",
			expectedError:       "invalid character",
			expectedRESTService: nil,
		},
		{
			name:          "Simple",
			file:          "./test/simple.json",
			expectedError: "",
			expectedRESTService: &restService{
				Spec: &spec.Spec{},
			},
		},
		{
			name:          "Full",
			file:          "./test/full.json",
			expectedError: "",
			expectedRESTService: &restService{
				Spec: &spec.Spec{
					Store: spec.Store{
						"users": []spec.JSON{
							spec.JSON{"id": "d93ce179-50f7-469e-bb36-1b3746145f00", "firstName": "John", "lastName": "Doe", "email": "john.doe@example.com"},
							spec.JSON{"id": "8cd6ef6c-2095-4c75-bc66-6f38e785299d", "firstName": "Milad", "lastName": "Irannejad", "email": "milad@example.com"},
						},
					},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			svc, err := NewRESTService(tc.file)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.Contains(t, err.Error(), tc.expectedError)
			}

			if tc.expectedRESTService == nil {
				assert.Nil(t, svc)
			} else {
				service, ok := svc.(*restService)
				assert.True(t, ok)
				assert.Equal(t, tc.expectedRESTService, service)
			}
		})
	}
}
