package prometheus

import (
	"github.com/goexl/prometheus/internal/builder"
)

func Metric() *builder.Metric {
	return builder.NewMetric()
}
