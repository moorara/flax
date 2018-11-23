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

func TestFlax(t *testing.T) {
	if strings.ToLower(os.Getenv("COMPONENT_TEST")) != "true" {
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
			name:               "MockREST",
			method:             "PUT",
			endpoint:           "/api/mock/rest",
			body:               JSON{},
			expectedStatusCode: 200,
			expectedResponse:   JSON{},
		},
		{
			name:               "MockHTTPResponse",
			method:             "PUT",
			endpoint:           "/api/mock/http",
			body:               JSON{},
			expectedStatusCode: 200,
			expectedResponse:   JSON{},
		},
		{
			name:               "MockHTTPForward",
			method:             "PUT",
			endpoint:           "/api/mock/http",
			body:               JSON{},
			expectedStatusCode: 200,
			expectedResponse:   JSON{},
		},
		{
			name:               "MockHTTPResponseAndForward",
			method:             "PUT",
			endpoint:           "/api/mock/http",
			body:               JSON{},
			expectedStatusCode: 200,
			expectedResponse:   JSON{},
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
