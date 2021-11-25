package main

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/trace"
)

func setupTracing(c Config, headerName string) (func() error, error) {
	if c.SpanExporterEndpoint == "" {
		c.logger.Debugf("tracing is disabled by configuration: no endpoint set")
		return nil, nil
	}
	spanExporter, err := newTraceExporter(c, headerName)
	if err != nil {
		return nil, fmt.Errorf("failed to create span exporter: %v", err)
	}

	bsp := trace.NewBatchSpanProcessor(spanExporter)
	tp := trace.NewTracerProvider(
		trace.WithSampler(trace.AlwaysSample()),
		trace.WithSpanProcessor(bsp),
		trace.WithResource(c.Resource),
	)

	// Default propagator.
	ps := Propagators{DefaultPropagator}
	if err = configurePropagators(ps); err != nil {
		return nil, err
	}

	otel.SetTracerProvider(tp)

	return func() error {
		_ = bsp.Shutdown(context.Background())
		return spanExporter.Shutdown(context.Background())
	}, nil
}
