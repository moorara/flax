package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	tests := []struct {
		name                    string
		expectedLogLevel        string
		expectedHTTPPort        string
		expectedHTTPSPort       string
		expectedJaegerAgentAddr string
		expectedJaegerLogSpans  bool
	}{
		{
			name:                    "Defauts",
			expectedLogLevel:        defaultLogLevel,
			expectedHTTPPort:        defaultHTTPPort,
			expectedHTTPSPort:       defaultHTTPSPort,
			expectedJaegerAgentAddr: defaultJaegerAgentAddr,
			expectedJaegerLogSpans:  defaultJaegerLogSpans,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedLogLevel, Config.LogLevel)
			assert.Equal(t, tc.expectedHTTPPort, Config.HTTPPort)
			assert.Equal(t, tc.expectedHTTPSPort, Config.HTTPSPort)
			assert.Equal(t, tc.expectedJaegerAgentAddr, Config.JaegerAgentAddr)
			assert.Equal(t, tc.expectedJaegerLogSpans, Config.JaegerLogSpans)
		})
	}
}
