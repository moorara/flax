// +build component

package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/moorara/konfig"
	"github.com/stretchr/testify/assert"
)

type JSON map[string]interface{}

var config = struct {
	SUTAddress string
}{
	SUTAddress: "http://localhost:10000",
}

func TestMain(m *testing.M) {
	_ = konfig.Pick(&config)
	os.Exit(m.Run())
}

func TestMockHTTP(t *testing.T) {
	tests := []struct {
		name               string
		method             string
		endpoint           string
		body               JSON
		expectedStatusCode int
		expectedResponse   JSON
	}{}

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
			err = json.NewDecoder(res.Body).Decode(&response)
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedResponse, response)
		})
	}
}

func TestMockREST(t *testing.T) {
	tests := []struct {
		name               string
		method             string
		endpoint           string
		body               JSON
		expectedStatusCode int
		expectedResponse   JSON
	}{}

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
			err = json.NewDecoder(res.Body).Decode(&response)
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedResponse, response)
		})
	}
}
