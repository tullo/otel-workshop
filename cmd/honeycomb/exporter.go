package main

import (
	"context"
	"os"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"google.golang.org/grpc/credentials"
)

// Configure a new exporter using environment variables for sending data to
// Honeycomb over gRPC.
func newExporter(ctx context.Context) (*otlptrace.Exporter, error) {
	opts := []otlptracegrpc.Option{
		otlptracegrpc.WithEndpoint(os.Getenv("HONEYCOMB_API_ENDPOINT")),
		otlptracegrpc.WithHeaders(map[string]string{
			"x-honeycomb-team":    os.Getenv("HONEYCOMB_API_KEY"),
			"x-honeycomb-dataset": os.Getenv("HONEYCOMB_DATASET"),
		}),
		otlptracegrpc.WithTLSCredentials(credentials.NewClientTLSFromCert(nil, "")),
	}

	client := otlptracegrpc.NewClient(opts...)
	return otlptrace.New(ctx, client)
}
