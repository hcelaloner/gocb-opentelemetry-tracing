package tracing

import (
	"context"
	"github.com/couchbase/gocb/v2"
	"go.opentelemetry.io/otel/trace"
)

// openTelemetryRequestTracer is an implementation of the gocb.RequestTracer interface that wraps an OpenTelemetry tracer.
type openTelemetryRequestTracer struct {
	tracer trace.Tracer
}

func NewOpenTelemetryRequestTracer(provider trace.TracerProvider) gocb.RequestTracer {
	return &openTelemetryRequestTracer{
		tracer: provider.Tracer("com.couchbase.client/go"),
	}
}

// RequestSpan creates a gocb.RequestSpan that wraps an OpenTelemetry span
func (t *openTelemetryRequestTracer) RequestSpan(parentContext gocb.RequestSpanContext, operationName string) gocb.RequestSpan {
	parentCtx := context.Background()
	if ctx, ok := parentContext.(context.Context); ok {
		parentCtx = ctx
	}

	ctx, _ := t.tracer.Start(parentCtx, operationName)
	return NewOpenTelemetryRequestSpanFromContext(ctx)
}
