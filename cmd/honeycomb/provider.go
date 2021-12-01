package main

import (
	"os"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
)

// Create a new tracer provider with a batch span processor and the otlp exporter.
func newTraceProvider(exp *otlptrace.Exporter) *sdktrace.TracerProvider {
	// service.name attribute is required.
	resource := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String(os.Getenv("SERVICE_NAME")),
	)

	return sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(resource),
	)
}
