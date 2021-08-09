package bl

import (
	"time"
	"github.com/prometheus/client_golang/prometheus"
	// "github.com/prometheus/client_golang/prometheus/promauto"
  	_ "github.com/prometheus/client_golang/prometheus/promhttp"

	"../interfaces"
)

// diagnostic service implementation type
type diagnosticService struct {
	//logger
	_log interfaces.ILogger

	// prometheus method calls counter
	_handlersCalledRegistered *prometheus.CounterVec

	// prometheus method duration histogramms
	_handlersDurationRegistered *prometheus.HistogramVec
}

// DiagnosticServiceNew creates an instance of diagnosticService
func DiagnosticServiceNew(
	log interfaces.ILogger) interfaces.IDiagnosticService {
	service := &diagnosticService{}

	service._log = log

	service._handlersCalledRegistered = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_handlers_called",
			Help: "http handlers counter",
		},
		[]string{"path"},
	)

	service._handlersDurationRegistered = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "http_handlers_called_microseconds",
		Help: "Duration of HTTP requests.",
		Buckets: []float64{0.5, 1, 5, 10, 20, 30, 100, 1000 } ,
	}, []string{"path"})

	/*requestProcessingTimeHistogramMs := prometheus.NewHistogram(
		prometheus.HistogramOpts{
		  Name: "request_processing_time_histogram_ms",
		  Buckets: prometheus.LinearBuckets(0, 10, 20),
		})
	  prometheus.MustRegister(requestProcessingTimeHistogramMs)*/

	prometheus.MustRegister(service._handlersCalledRegistered)
	prometheus.MustRegister(service._handlersDurationRegistered)

	return service
}

// IncCount increments specified counter
func (service *diagnosticService) IncCount(counterName string) {
	handlersCalledRegistered := service._handlersCalledRegistered

	handlersCalledRegistered.WithLabelValues(counterName).Inc()
}


// SetDuration sets specified duration
func (service *diagnosticService) SetDuration(counterName string, duration time.Duration) {
	_handlersDurationRegistered := service._handlersDurationRegistered

	_handlersDurationRegistered.WithLabelValues(counterName).Observe(float64(duration) / 1e6)
}