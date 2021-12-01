package fib

import (
	"context"
	"fmt"
	"net/http"

	"github.com/newrelic/go-agent/v3/newrelic"
	"go.opentelemetry.io/otel"
)

func (s *Server) ServeNewRelic(ctx context.Context, app *newrelic.Application) error {
	_, span := otel.Tracer(name).Start(ctx, "ServeNewRelic")
	defer span.End()

	mux := http.NewServeMux()
	mux.Handle(newrelic.WrapHandle(app, "/", http.HandlerFunc(rootHandler)))
	mux.Handle("/favicon.ico", http.NotFoundHandler())
	mux.Handle(newrelic.WrapHandle(app, "/fib", http.HandlerFunc(fibHandler)))
	mux.Handle(newrelic.WrapHandle(app, "/fibinternal", http.HandlerFunc(fibHandler)))

	s.l.Println("Your server is live!\nTry to navigate to: http://127.0.0.1:3000/fib?n=6")
	if err := http.ListenAndServe("127.0.0.1:3000", mux); err != nil {
		return fmt.Errorf("could not start web server: %w", err)
	}

	return nil
}
