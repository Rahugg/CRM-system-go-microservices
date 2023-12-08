package metrics

var (
	// TimeBucketsFast suits if expected response time is very high: 1ms..100ms
	// This buckets suits for cache storages (in-memory cache, Memcache).
	TimeBucketsFast = []float64{0.001, 0.003, 0.007, 0.015, 0.05, 0.1, 0.2, 0.5, 1, 2, 5, 10, 15, 20}

	// TimeBucketsMedium suits for most of GO APIs, where response time is between 50ms..500ms.
	// Works for wide range of systems because provides near-logarithmic buckets distribution.
	//nolint:lll
	TimeBucketsMedium = []float64{0.001, 0.005, 0.015, 0.05, 0.1, 0.25, 0.5, 0.75, 1, 1.5, 2, 3.5, 5, 10, 15, 20, 30, 40, 50, 60}

	// TimeBucketsSlow suits for relatively slow services, where expected response time is > 500ms.
	TimeBucketsSlow = []float64{0.05, 0.1, 0.2, 0.5, 0.75, 1, 1.25, 1.5, 1.75, 2, 2.5, 3, 4, 5, 10, 15, 20}

	CBOpenDurationBuckets = []float64{15, 30, 60, 120, 180, 240, 300, 360, 420, 480, 540, 600, 900, 1200}
)
