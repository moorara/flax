package middleware

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	kitLog "github.com/go-kit/kit/log"
	"github.com/moorara/flax/pkg/log"
	"github.com/moorara/flax/pkg/metrics"
	"github.com/stretchr/testify/assert"
)

func getMetric(metrics *metrics.Metrics, name string) (output string) {
	mfs, _ := metrics.Registry.Gather()
	for _, mf := range mfs {
		if *mf.Name == name {
			for _, m := range mf.Metric {
				output += m.String() + "\n"
			}
		}
	}
	return output
}

func TestMonitorMiddleware(t *testing.T) {
	tests := []struct {
		name                string
		reqMethod           string
		reqURL              string
		resStatusCode       int
		expectedMethod      string
		expectedURL         string
		expectedStatusCode  int
		expectedStatusClass string
	}{
		{
			name:                "200",
			reqMethod:           "POST",
			reqURL:              "/path",
			resStatusCode:       200,
			expectedMethod:      "POST",
			expectedURL:         "/path",
			expectedStatusCode:  200,
			expectedStatusClass: "2xx",
		},
		{
			name:                "404",
			reqMethod:           "POST",
			reqURL:              "/path",
			resStatusCode:       404,
			expectedMethod:      "POST",
			expectedURL:         "/path",
			expectedStatusCode:  404,
			expectedStatusClass: "4xx",
		},
		{
			name:                "500",
			reqMethod:           "POST",
			reqURL:              "/path",
			resStatusCode:       500,
			expectedMethod:      "POST",
			expectedURL:         "/path",
			expectedStatusCode:  500,
			expectedStatusClass: "5xx",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Logger with pipe to read from
			rd, wr, _ := os.Pipe()
			dec := json.NewDecoder(rd)
			logger := &log.Logger{
				Logger: kitLog.NewJSONLogger(wr),
			}

			// Mock metrics
			metrics := metrics.New("test_service")

			middleware := NewMonitorMiddleware(logger, metrics)
			handler := middleware.Wrap(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tc.resStatusCode)
			})

			r := httptest.NewRequest(tc.reqMethod, tc.reqURL, nil)
			w := httptest.NewRecorder()

			handler(w, r)

			// Verify logging
			var log map[string]interface{}
			dec.Decode(&log)
			assert.Equal(t, tc.expectedMethod, log["req.method"])
			assert.Equal(t, tc.expectedURL, log["req.url"])
			assert.Equal(t, float64(tc.expectedStatusCode), log["res.statusCode"])
			assert.Equal(t, tc.expectedStatusClass, log["res.statusClass"])

			// Verify metrics
			assert.NotEmpty(t, getMetric(metrics, "http_requests_duration_seconds"))
			assert.NotEmpty(t, getMetric(metrics, "http_requests_duration_quantiles_seconds"))
		})
	}
}