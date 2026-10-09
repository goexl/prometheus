package builder

import (
	"github.com/goexl/log"
	"github.com/goexl/prometheus/internal/core"
	"github.com/goexl/prometheus/internal/param"
	"github.com/goexl/valuer"
)

type Prometheus struct {
	params *param.Promethy
}

func NewPromethy() *Prometheus {
	return &Prometheus{
		params: param.NewParams(),
	}
}

func (p *Prometheus) Logger(logger log.Logger) (builder *Prometheus) {
	p.params.Logger = logger
	builder = p

	return
}

func (p *Prometheus) Parser(parser *valuer.Parser) (builder *Prometheus) {
	p.params.Parser = parser
	builder = p

	return
}

func (p *Prometheus) Label(key string, value string) (builder *Prometheus) {
	p.params.Labels[key] = value
	builder = p

	return
}

func (p *Prometheus) Labels(labels map[string]string) (builder *Prometheus) {
	for key, value := range labels {
		p.params.Labels[key] = value
	}
	builder = p

	return
}

func (p *Prometheus) Build() *core.Promethy {
	return core.NewPromethy(p.params)
}
