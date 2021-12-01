package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/tullo/otel-workshop/web/fib"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"google.golang.org/grpc/credentials"
)

func newGRPCExporter(ctx context.Context) (*otlptrace.Exporter, error) {
	opts := []otlptracegrpc.Option{
		otlptracegrpc.WithEndpoint(os.Getenv("OTEL_EXPORTER_OTLP_TRACES_ENDPOINT")),
		otlptracegrpc.WithHeaders(map[string]string{
			"Authorization": os.Getenv("ASPECTO_API_KEY"),
		}),
		otlptracegrpc.WithTLSCredentials(credentials.NewClientTLSFromCert(nil, "")),
	}

	client := otlptracegrpc.NewClient(opts...)
	return otlptrace.New(ctx, client)
}

func newHTTPSExporter(ctx context.Context) (*otlptrace.Exporter, error) {
	traceExporter, err := otlptracehttp.New(
		ctx,
		otlptracehttp.WithEndpoint(os.Getenv("OTEL_EXPORTER_OTLP_TRACES_ENDPOINT")),
		otlptracehttp.WithHeaders(map[string]string{"Authorization": os.Getenv("ASPECTO_API_KEY")}),
	)

	return traceExporter, err
}

func newOTLauncher() func() {
	ctx := context.Background()
	exp, err := newGRPCExporter(ctx)
	//exp, err := newHTTPSExporter(ctx)
	if err != nil {
		log.Fatal(err)
	}

	res, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceNameKey.String("fib"),
		),
	)
	if err != nil {
		log.Fatal(err)
	}
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(res),
		sdktrace.WithBatcher(exp),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(
			// W3C Trace Context propagator
			propagation.TraceContext{},
			propagation.Baggage{},
		),
	)

	return func() {
		// Handle this error in a sensible manner where possible.
		_ = exp.Shutdown(ctx)
	}
}

func main() {
	l := log.New(os.Stdout, "", 0)

	shutdown := newOTLauncher()
	defer shutdown()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	errCh := make(chan error)

	// Start web server.
	s := fib.NewServer(os.Stdin, l)
	go func() {
		errCh <- s.Serve(context.Background())
	}()

	select {
	case <-sigCh:
		l.Println("\ngoodbye")
		return
	case err := <-errCh:
		if err != nil {
			l.Fatal(err)
		}
	}
}
