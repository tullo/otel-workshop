package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/lightstep/otel-launcher-go/launcher"
	"github.com/tullo/otel-workshop/web/fib"
)

func newLighstepLauncher() launcher.Launcher {
	/*
		Export env vars picked up by launcher
		LS_ACCESS_TOKEN=...
		LS_SERVICE_NAME=fib-workshop
		LS_SERVICE_VERSION=v0.1.0
		OTEL_EXPORTER_OTLP_SPAN_ENDPOINT=ingest.lightstep.com:443
		OTEL_LOG_LEVEL=debug
		OTEL_RESOURCE_ATTRIBUTES=host.hostname=fib.workshop.com,deployment.region=eu-west1,deployment.environment=workshop
	*/
	return launcher.ConfigureOpentelemetry()
	/*
		ls := launcher.ConfigureOpentelemetry(
			launcher.WithServiceName("fib-workshop"),
			launcher.WithServiceVersion("v0.1.2"),
			launcher.WithResourceAttributes(map[string]string{
				"host.hostname":          "fib.workshop.com",
				"deployment.region":      "eu-west1",
				"deployment.environment": "workshop",
			}),
			launcher.WithAccessToken("..."),
		)
		return ls
	*/
}

func main() {
	l := log.New(os.Stdout, "", 0)

	ls := newLighstepLauncher()
	defer ls.Shutdown()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	errCh := make(chan error)

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
