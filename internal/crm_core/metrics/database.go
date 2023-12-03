package metrics

import "time"

const (
	repoMethodLabel = "method"
	repoStatusLabel = "status"
)

var databaseResponseTime = newHistogramVec(
	"db_response_time_seconds",
	"Histogram of application RT for DB",
	TimeBucketsMedium,
	repoMethodLabel, repoStatusLabel,
)

func DatabaseQueryTime(method string) (ok, fail func()) {
	start := time.Now()

	completed := false

	ok = func() {
		if completed {
			return
		}

		completed = true

		databaseResponseTime.WithLabelValues(method, statusOk).Observe(time.Since(start).Seconds())
	}

	fail = func() {
		if completed {
			return
		}

		completed = true

		databaseResponseTime.WithLabelValues(method, statusError).Observe(time.Since(start).Seconds())
	}

	return ok, fail
}

// nolint: gochecknoinits
func init() {
	mustRegister(databaseResponseTime)
}
