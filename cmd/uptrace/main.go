package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/tullo/otel-workshop/web/fib"
	"github.com/uptrace/uptrace-go/uptrace"
	"go.opentelemetry.io/otel"
)

func main() {
	l := log.New(os.Stdout, "", 0)

	ctx := context.Background()

	// Configure OpenTelemetry with sensible defaults.
	uptrace.ConfigureOpentelemetry(
		// copy your project DSN here or use UPTRACE_DSN env var
		//uptrace.WithDSN("https://<key>@api.uptrace.dev/<project_id>"),
		uptrace.WithServiceName("fib"),
		uptrace.WithServiceVersion("1.0.0"),
	)
	// Send buffered spans and free resources.
	defer uptrace.Shutdown(ctx)

	tracer := otel.Tracer("fib-workshop")
	ctx, main := tracer.Start(ctx, "main-func")
	defer main.End()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	errCh := make(chan error)

	// Start web server.
	s := fib.NewServer(os.Stdin, l)
	go func() {
		errCh <- s.Serve(ctx)
	}()

	l.Printf("trace: %s\n", uptrace.TraceURL(main))

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
