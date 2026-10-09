package prometheus

type Counter interface {
	Inc()

	Add(float64)
}
