// utility/otel.go
package utility

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/devopscorner/golang-bedrock/src/config"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"go.opentelemetry.io/otel/trace"
)

// InitTracer initializes the OpenTelemetry tracer
func InitTracer(cfg *config.Config) func() {
	if cfg.OtelTraceEnable != "true" {
		log.Println("Tracing is not enabled")
		return func() {}
	}

	var exporter sdktrace.SpanExporter
	var err error

	ctx := context.Background()

	switch strings.ToLower(cfg.OtelTraceName) {
	case "jaeger":
		log.Println("Jaeger tracing is not implemented in this example")
		return func() {}
	case "xray":
		log.Println("X-Ray tracing is not implemented in this example")
		return func() {}
	default:
		// OTLP exporter setup
		if strings.HasPrefix(cfg.OtelOtlpEndpoint, "http") {
			// HTTP exporter
			exporter, err = otlptracehttp.New(ctx,
				otlptracehttp.WithEndpoint(fmt.Sprintf("%s:%d", cfg.OtelOtlpEndpoint, cfg.OtelOtlpPort)),
				otlptracehttp.WithInsecure(),
			)
		} else {
			// gRPC exporter
			exporter, err = otlptrace.New(
				ctx,
				otlptracegrpc.NewClient(
					otlptracegrpc.WithEndpoint(fmt.Sprintf("%s:%d", cfg.OtelOtlpEndpoint, cfg.OtelOtlpPort)),
					otlptracegrpc.WithInsecure(),
				),
			)
		}

		if err != nil {
			log.Fatalf("Failed to create exporter: %v", err)
		}
	}

	res, err := resource.New(
		ctx,
		resource.WithAttributes(
			semconv.ServiceNameKey.String(cfg.OtelServiceName),
			attribute.String("environment", cfg.OtelEnvironment),
		),
		resource.WithFromEnv(),
		resource.WithProcess(),
	)
	if err != nil {
		log.Fatalf("Failed to create resource: %v", err)
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(res),
	)

	otel.SetTracerProvider(tp)

	return func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := tp.Shutdown(ctx); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}
}

// AddSpanAttributes adds attributes to the current span
func AddSpanAttributes(ctx context.Context, attributes ...attribute.KeyValue) {
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attributes...)
}

// StartSpan starts a new span and returns the context with the span
func StartSpan(ctx context.Context, name string) (context.Context, trace.Span) {
	return otel.Tracer("").Start(ctx, name)
}

// EndSpan ends the span in the given context
func EndSpan(span trace.Span) {
	span.End()
}

// RecordError records an error in the current span
func RecordError(ctx context.Context, err error) {
	span := trace.SpanFromContext(ctx)
	span.RecordError(err)
}

// GetTracer returns a named tracer
func GetTracer(name string) trace.Tracer {
	return otel.Tracer(name)
}

// WithSpan wraps a function with a new span
func WithSpan(ctx context.Context, name string, fn func(context.Context) error) error {
	ctx, span := StartSpan(ctx, name)
	defer EndSpan(span)

	err := fn(ctx)
	if err != nil {
		RecordError(ctx, err)
	}

	return err
}
