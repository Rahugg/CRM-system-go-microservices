package metrics

const (
	httpHandlerLabel = "handler"
	httpCodeLabel    = "code"
	httpMethodLabel  = "method"
)

var (
	HttpResponseTime = newHistogramVec(
		"response_time_seconds",
		"Histogram of application RT for any kind of requests seconds",
		TimeBucketsMedium,
		httpHandlerLabel, httpCodeLabel, httpMethodLabel,
	)

	HttpRequestsTotalCollector = newCounterVec(
		"requests_total",
		"Counter of HTTP requests for any HTTP-based requests",
		httpHandlerLabel, httpCodeLabel, httpMethodLabel,
	)
)

// nolint: gochecknoinits
func init() {
	mustRegister(HttpResponseTime, HttpRequestsTotalCollector)
}
