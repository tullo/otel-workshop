package main

import (
	"context"
	"log"
	"os"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	otgrpc "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	p "go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdk "go.opentelemetry.io/otel/sdk/trace"
	sc "go.opentelemetry.io/otel/semconv/v1.7.0"
)

const serviceVersion = "v1.0.0"

func ConfigureOpentelemetry(ctx context.Context) func() {
	exp, err := newGRPCExporter(ctx)
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

func newGRPCExporter(ctx context.Context) (*otlptrace.Exporter, error) {
	opts := []otgrpc.Option{
		otgrpc.WithEndpoint(os.Getenv("OTEL_EXPORTER_OTLP_TRACES_ENDPOINT")),
		otgrpc.WithInsecure(),
		otgrpc.WithTimeout(5 * time.Second),
	}

	client := otgrpc.NewClient(opts...)
	return otlptrace.New(ctx, client)
}

// Create a new tracer provider with a batch span processor and the otlp exporter.
func newTraceProvider(exp *otlptrace.Exporter) *sdk.TracerProvider {
	// service.name attribute is required.
	res := resource.NewWithAttributes(
		sc.SchemaURL,
		sc.ServiceNameKey.String(os.Getenv("SERVICE_NAME")),
		sc.ServiceVersionKey.String(serviceVersion),
		sc.TelemetrySDKVersionKey.String("v1.0.1"),
		sc.TelemetrySDKLanguageGo,
	)

	return sdk.NewTracerProvider(
		sdk.WithSampler(sdk.AlwaysSample()),
		sdk.WithResource(res),
		sdk.WithSpanProcessor(sdk.NewBatchSpanProcessor(exp)))
}
