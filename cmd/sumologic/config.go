package main

import (
	"context"
	"log"
	"os"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/baggage"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	p "go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
)

func ConfigureOpentelemetry(ctx context.Context) func() {
	exp, err := newHTTPSExporter(ctx)
	if err != nil {
		log.Fatalf("failed to initialize exporter: %v", err)
	}

	tp := newTraceProvider(exp)
	otel.SetTracerProvider(tp)

	// Register the trace context and baggage propagators
	// so data is propagated across services/processes.
	otel.SetTextMapPropagator(
		// W3C Trace Context propagator
		p.NewCompositeTextMapPropagator(p.TraceContext{}, p.Baggage{}),
	)

	return func() {
		// Handle this error in a sensible manner where possible.
		_ = tp.Shutdown(ctx)
	}

}

func newHTTPSExporter(ctx context.Context) (*otlptrace.Exporter, error) {
	client := otlptracehttp.NewClient(otlptracehttp.WithInsecure())
	ctx = baggage.ContextWithoutBaggage(ctx)
	return otlptrace.New(ctx, client)
}

// Create a new tracer provider with a batch span processor and the otlp exporter.
func newTraceProvider(exp *otlptrace.Exporter) *sdktrace.TracerProvider {
	// service.name attribute is required.
	resource := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String(os.Getenv("SERVICE_NAME")),
	)

	return sdktrace.NewTracerProvider(
		sdktrace.WithSpanProcessor(sdktrace.NewBatchSpanProcessor(exp)),
		sdktrace.WithResource(resource),
	)
}
