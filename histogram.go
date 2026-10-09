package prometheus

type Histogram interface {
	Observe(float64)
}
