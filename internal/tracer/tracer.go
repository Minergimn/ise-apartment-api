package tracer

import (
	"context"
	"fmt"
	"io"

	"github.com/ozonmp/ise-apartment-api/internal/logger"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"

	"github.com/ozonmp/ise-apartment-api/internal/config"

	jaegercfg "github.com/uber/jaeger-client-go/config"
)

// NewTracer - returns new tracer.
func NewTracer(ctx context.Context, cfg *config.Config) (io.Closer, error) {
	cfgTracer := &jaegercfg.Configuration{
		ServiceName: cfg.Project.ServiceName,
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: cfg.Jaeger.Host + cfg.Jaeger.Port,
		},
	}
	tracer, closer, err := cfgTracer.NewTracer(jaegercfg.Logger(jaeger.StdLogger))
	if err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("failed init jaeger: %v", err))

		return nil, err
	}
	opentracing.SetGlobalTracer(tracer)
	logger.InfoKV(ctx, "Traces started")

	return closer, nil
}
