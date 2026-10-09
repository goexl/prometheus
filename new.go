package prometheus

import (
	"sync"

	"github.com/goexl/prometheus/internal/builder"
)

var once sync.Once

// New 创建普罗米修斯
// 使用单例模式，多次创建只会创建一个实例
func New() (prometheus *builder.Prometheus) {
	once.Do(func() {
		prometheus = builder.NewPromethy()
	})

	return
}
