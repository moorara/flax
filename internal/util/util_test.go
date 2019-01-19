package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTempFile(t *testing.T) {
	tests := []struct {
		name          string
		prefix        string
		content       string
		expectedError string
	}{
		{
			name:          "NoContent",
			prefix:        "test-",
			content:       "",
			expectedError: "",
		},
		{
			name:          "WithContent",
			prefix:        "test-",
			content:       "something",
			expectedError: "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			filepath, remove, err := CreateTempFile(tc.prefix, tc.content)
			defer remove()

			if tc.expectedError != "" {
				assert.Contains(t, err.Error(), tc.expectedError)
				assert.Empty(t, filepath)
				assert.Nil(t, remove)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, filepath)
				assert.NotNil(t, remove)
			}
		})
	}
}
