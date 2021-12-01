package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/tullo/otel-workshop/web/fib"
)

func main() {
	otshutdown := ConfigureOpentelemetry(context.Background())
	defer otshutdown()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	errCh := make(chan error)

	// Start web server.
	l := log.New(os.Stdout, "", 0)
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
