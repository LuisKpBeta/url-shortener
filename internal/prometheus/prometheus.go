package prometheus

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type PrometheusService struct {
	httpRequestHistogram *prometheus.HistogramVec
}
type HTTPMetric struct {
	Handler    string
	Method     string
	StatusCode string
	StartedAt  time.Time
	FinishedAt time.Time
	Duration   float64
}

func NewHTTPMetric(handler string, method string) *HTTPMetric {
	return &HTTPMetric{
		Handler: handler,
		Method:  method,
	}
}
func (h *HTTPMetric) Started() {
	h.StartedAt = time.Now()
}

func (h *HTTPMetric) Finished() {
	h.FinishedAt = time.Now()
	h.Duration = time.Since(h.StartedAt).Seconds()
}

type UseCase interface {
	SaveHTTP(h *HTTPMetric)
}

func NewPrometheusService() (*PrometheusService, error) {
	http := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "http",
		Name:      "request_duration_seconds",
		Help:      "The latency of the HTTP requests.",
		Buckets:   prometheus.DefBuckets,
	}, []string{"handler", "method", "code"})

	s := &PrometheusService{
		httpRequestHistogram: http,
	}
	err := prometheus.Register(s.httpRequestHistogram)
	if err != nil && err.Error() != "duplicate metrics collector registration attempted" {
		return nil, err
	}
	return s, nil
}

func (s *PrometheusService) SaveHTTP(h *HTTPMetric) {
	s.httpRequestHistogram.WithLabelValues(h.Handler, h.Method, h.StatusCode).Observe(h.Duration)
}
