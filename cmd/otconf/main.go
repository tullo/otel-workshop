package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/tullo/otel-workshop/web/fib"
)

func newOTLauncher() OTConf {
	exporterHeaderName := "lightstep-access-token"
	return ConfigureOpentelemetry(exporterHeaderName)
}

func main() {
	l := log.New(os.Stdout, "", 0)

	ot := newOTLauncher()
	defer ot.Shutdown()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	errCh := make(chan error)

	// Start metrics collection.
	//go collectMetrics(context.Background())

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
