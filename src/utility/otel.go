// utility/otel.go
package utility

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/devopscorner/golang-restfulapi-bedrock/src/config"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

func InitTracer(cfg *config.Config) func() {
	var exporter sdktrace.SpanExporter
	var err error

	ctx := context.Background()

	if cfg.OtelTraceEnable == "true" {
		switch strings.ToLower(cfg.OtelTraceName) {
		case "jaeger":
			log.Println("Jaeger tracing is not implemented in this example")
		case "xray":
			log.Println("X-Ray tracing is not implemented in this example")
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

	return func() {}
}

// AddSpanAttributes adds attributes to the current span
func AddSpanAttributes(ctx context.Context, attributes ...attribute.KeyValue) {
	span := otel.GetTracerProvider().Tracer("").SpanFromContext(ctx)
	span.SetAttributes(attributes...)
}

// StartSpan starts a new span and returns the context with the span
func StartSpan(ctx context.Context, name string) (context.Context, sdktrace.Span) {
	return otel.GetTracerProvider().Tracer("").Start(ctx, name)
}

// EndSpan ends the span in the given context
func EndSpan(span sdktrace.Span) {
	span.End()
}

// RecordError records an error in the current span
func RecordError(ctx context.Context, err error) {
	span := otel.GetTracerProvider().Tracer("").SpanFromContext(ctx)
	span.RecordError(err)
}
