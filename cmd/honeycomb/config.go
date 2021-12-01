package main

import (
	"context"
	"log"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

func ConfigureOpentelemetry(ctx context.Context) func() {
	exp, err := newExporter(ctx)
	if err != nil {
		log.Fatalf("failed to initialize exporter: %v", err)
	}

	tp := newTraceProvider(exp)
	otel.SetTracerProvider(tp)

	// Register the trace context and baggage propagators
	// so data is propagated across services/processes.
	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(
			// W3C Trace Context propagator
			propagation.TraceContext{},
			propagation.Baggage{},
		),
	)

	return func() {
		// Handle this error in a sensible manner where possible.
		_ = tp.Shutdown(ctx)
	}

}
