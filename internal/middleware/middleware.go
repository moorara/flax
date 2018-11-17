package middleware

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	xhttp "github.com/moorara/flax/pkg/http"
	"github.com/moorara/flax/pkg/log"
	"github.com/moorara/flax/pkg/metrics"
)

type (
	// Middleware is a http middleware
	Middleware interface {
		Wrap(http.HandlerFunc) http.HandlerFunc
	}

	monitorMiddleware struct {
		logger  *log.Logger
		metrics *metrics.Metrics
	}
)

// NewMonitorMiddleware creates a new middleware for logging
func NewMonitorMiddleware(logger *log.Logger, metrics *metrics.Metrics) Middleware {
	return &monitorMiddleware{
		logger:  logger,
		metrics: metrics,
	}
}

func (m *monitorMiddleware) Wrap(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Next http handler
		start := time.Now()
		rw := xhttp.NewResponseWriter(w)
		next(rw, r)
		duration := time.Now().Sub(start).Seconds()

		method := r.Method
		url := r.URL.Path
		headers := r.Header
		statusCode := uint16(rw.StatusCode())
		statusClass := rw.StatusClass()

		logs := []interface{}{
			"req.method", method,
			"req.url", url,
			"req.headers", headers,
			"res.statusCode", statusCode,
			"res.statusClass", statusClass,
			"responseTime", duration,
			"message", fmt.Sprintf("%s %s %d %f", method, url, statusCode, duration),
		}

		// Logging
		switch {
		case statusCode >= 500:
			m.logger.Error(logs...)
		case statusCode >= 400:
			m.logger.Warn(logs...)
		case statusCode >= 100:
			m.logger.Info(logs...)
		}

		// Metrics
		sc := strconv.Itoa(rw.StatusCode())
		m.metrics.HTTPDurationHist.WithLabelValues(method, url, sc, statusClass).Observe(duration)
		m.metrics.HTTPDurationSumm.WithLabelValues(method, url, sc, statusClass).Observe(duration)
	}
}
