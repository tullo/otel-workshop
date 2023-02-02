package main

/*
import (
	"context"

	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/encoding/gzip"
)

func newTraceExporter(c Config, headerName string) (*otlptrace.Exporter, error) {
	secureOption := otlptracegrpc.WithTLSCredentials(credentials.NewClientTLSFromCert(nil, ""))
	if c.SpanExporterEndpointInsecure {
		secureOption = otlptracegrpc.WithInsecure()
	}

	return otlptrace.New(
		context.Background(),
		otlptracegrpc.NewClient(
			secureOption,
			otlptracegrpc.WithEndpoint(c.SpanExporterEndpoint),
			otlptracegrpc.WithHeaders(
				map[string]string{
					headerName: c.AccessToken,
				},
			),
			otlptracegrpc.WithCompressor(gzip.Name),
		),
	)
}

func newMetricsExporter(c Config, headerName string) (*otlpmetric.Exporter, error) {
	secureOption := otlpmetricgrpc.WithTLSCredentials(credentials.NewClientTLSFromCert(nil, ""))
	if c.MetricExporterEndpointInsecure {
		secureOption = otlpmetricgrpc.WithInsecure()
	}

	return otlpmetric.New(
		context.Background(),
		otlpmetricgrpc.NewClient(
			secureOption,
			otlpmetricgrpc.WithEndpoint(c.MetricExporterEndpoint),
			otlpmetricgrpc.WithHeaders(
				map[string]string{
					headerName: c.AccessToken,
				},
			),
			otlpmetricgrpc.WithCompressor(gzip.Name),
		),
	)
}
*/
