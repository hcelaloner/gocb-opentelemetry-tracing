package tracing

import (
	"context"
	"fmt"
	"github.com/couchbase/gocb/v2"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"log"
	"time"
)

// openTelemetryRequestSpan an implementation of gocb.RequestSpan that wraps an OpenTelemetry span, ready to be passed in into options for each operation into the SDK as a parent.
type openTelemetryRequestSpan struct {
	ctx  context.Context
	span trace.Span
}

func NewOpenTelemetryRequestSpanFromContext(ctx context.Context) gocb.RequestSpan {
	return &openTelemetryRequestSpan{ctx: ctx, span: trace.SpanFromContext(ctx)}
}

// End completes the gocb.RequestSpan.
func (o openTelemetryRequestSpan) End() {
	o.span.End()
}

// Context returns the gocb.RequestSpanContext for gocb.RequestSpan.
func (o openTelemetryRequestSpan) Context() gocb.RequestSpanContext {
	return o.ctx
}

// AddEvent adds an event with the provided name and timestamp to the gocb.RequestSpan.
func (o openTelemetryRequestSpan) AddEvent(name string, timestamp time.Time) {
	o.span.AddEvent(name, trace.WithTimestamp(timestamp))
}

// SetAttribute sets kv as attributes of the gocb.RequestSpan.
func (o openTelemetryRequestSpan) SetAttribute(key string, value interface{}) {
	switch v := value.(type) {
	case string:
		o.span.SetAttributes(attribute.String(key, v))
	case *string:
		o.span.SetAttributes(attribute.String(key, *v))
	case bool:
		o.span.SetAttributes(attribute.Bool(key, v))
	case *bool:
		o.span.SetAttributes(attribute.Bool(key, *v))
	case int:
		o.span.SetAttributes(attribute.Int(key, v))
	case *int:
		o.span.SetAttributes(attribute.Int(key, *v))
	case int64:
		o.span.SetAttributes(attribute.Int64(key, v))
	case *int64:
		o.span.SetAttributes(attribute.Int64(key, *v))
	case float64:
		o.span.SetAttributes(attribute.Float64(key, v))
	case *float64:
		o.span.SetAttributes(attribute.Float64(key, *v))
	case []string:
		o.span.SetAttributes(attribute.StringSlice(key, v))
	case []bool:
		o.span.SetAttributes(attribute.BoolSlice(key, v))
	case []int:
		o.span.SetAttributes(attribute.IntSlice(key, v))
	case []int64:
		o.span.SetAttributes(attribute.Int64Slice(key, v))
	case []float64:
		o.span.SetAttributes(attribute.Float64Slice(key, v))
	case fmt.Stringer:
		o.span.SetAttributes(attribute.String(key, v.String()))
	default:
		// This isn't great but we should make some effort to output some sort of warning.
		log.Println("Unable to determine value as a type that we can handle")
	}
}
