package fib

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	cleanhttp "github.com/hashicorp/go-cleanhttp"
	"go.opentelemetry.io/contrib/instrumentation/net/http/httptrace/otelhttptrace"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	p "go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

const name = "otel-workshop"

// Server is a Fibonacci computation service.
type Server struct {
	l *log.Logger
}

// NewServer returns a new Server.
func NewServer(r io.Reader, l *log.Logger) *Server {
	return &Server{l: l}
}

func (s *Server) Serve(ctx context.Context) error {
	// Because we are going to be passing Context between HTTP calls,
	// we will need to make sure the text propagator is declared:
	// - otel.SetTextMapPropagator
	otel.SetTextMapPropagator(p.NewCompositeTextMapPropagator(p.TraceContext{}, p.Baggage{}))

	_, span := otel.Tracer(name).Start(ctx, "Serve")
	defer span.End()

	mux := http.NewServeMux()

	// WithPublicEndpoint() option marks handlers as a trace boundary.
	mux.Handle("/", otelhttp.NewHandler(otelhttp.WithRouteTag("/", http.HandlerFunc(RootHandler)), "root", otelhttp.WithPublicEndpoint()))
	mux.Handle("/favicon.ico", http.NotFoundHandler())
	mux.Handle("/fib", otelhttp.NewHandler(otelhttp.WithRouteTag("/fib", http.HandlerFunc(FibHandler)), "fibonacci", otelhttp.WithPublicEndpoint()))
	mux.Handle("/fibinternal", http.HandlerFunc(FibHandler))

	s.l.Println("Your server is live!\nTry to navigate to: http://127.0.0.1:3000/fib?n=6")
	if err := http.ListenAndServe("127.0.0.1:3000", mux); err != nil {
		return fmt.Errorf("could not start web server: %w", err)
	}

	return nil
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	// context propagation
	_ = database(r.Context(), "darkgrey")
}

func database(ctx context.Context, color string) int {
	_, span := otel.Tracer(name).Start(ctx, "database")
	defer span.End()

	span.SetAttributes(attribute.String("color", color))

	return 0
}

func FibHandler(w http.ResponseWriter, r *http.Request) {
	span := trace.SpanFromContext(r.Context())
	key := "n"
	if len(r.URL.Query()[key]) != 1 {
		err := fmt.Errorf("wrong number of arguments")
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprint(w, err.Error())

		return
	}

	nth, err := strconv.Atoi(r.URL.Query()[key][0])
	if err != nil {
		err := fmt.Errorf("couldn't parse index '%s'", r.URL.Query()[key][0])
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		w.WriteHeader(http.StatusServiceUnavailable) // 503
		fmt.Fprint(w, err.Error())

		return
	}

	if nth < 0 {
		err := fmt.Errorf("rong fibonacci number request '%d'", nth)
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprint(w, err.Error())

		return
	}

	if nth > 30 {
		err := fmt.Errorf("unsupported fibonacci number '%d': too large", nth)
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		w.WriteHeader(http.StatusServiceUnavailable) // 503
		fmt.Fprint(w, err.Error())

		return
	}

	span.SetAttributes(attribute.Int("request.n", nth))

	switch nth {
	case 0:
		fmt.Fprint(w, 0)
		return
	case 1:
		fmt.Fprint(w, 1)
		return
	case 2:
		fmt.Fprint(w, 2)
		return
	}
	c1 := make(chan uint)
	c2 := make(chan uint)
	errCh := make(chan error)

	tr := otel.Tracer("fibHandler")

	go func(ctx context.Context, n uint) {
		ictx, sp := tr.Start(ctx, "fibClient")
		defer sp.End()

		url := fmt.Sprintf("http://127.0.0.1:3000/fibinternal?n=%d", n-1)
		sp.SetAttributes(
			attribute.Int("request.n", int(n-1)),
			attribute.String("url", url),
		)

		req, _ := http.NewRequestWithContext(ictx, http.MethodGet, url, nil)
		// W3C client
		ictx, req = otelhttptrace.W3C(ictx, req)
		// Inject span context to the new request
		otelhttptrace.Inject(ictx, req)

		res, err := cleanhttp.DefaultClient().Do(req)
		if err != nil {
			sp.SetStatus(codes.Error, "failure sending HTTP request")
			sp.RecordError(err)
			errCh <- err
			return
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			res.Body.Close()
			sp.SetStatus(codes.Error, err.Error())
			sp.RecordError(err)
			errCh <- err
			return
		}

		fib, err := strconv.Atoi(string(body))
		if err != nil {
			sp.RecordError(err)
			sp.SetStatus(codes.Error, "failure parsing result")
			errCh <- err
			return
		}

		c1 <- uint(fib)
		sp.AddEvent("sent result on channel", trace.WithAttributes(
			attribute.Int("result", int(fib)),
			attribute.String("chan", "c1"),
		))
	}(r.Context(), uint(nth))

	go func(ctx context.Context, n uint) {
		ictx, sp := tr.Start(ctx, "fibClient")
		defer sp.End()

		url := fmt.Sprintf("http://127.0.0.1:3000/fibinternal?n=%d", n-2)
		sp.SetAttributes(
			attribute.Int("request.n", int(n-1)),
			attribute.String("url", url),
		)

		req, _ := http.NewRequestWithContext(ictx, http.MethodGet, url, nil)
		// W3C client
		ictx, req = otelhttptrace.W3C(ictx, req)
		// Inject your span context to the new request
		otelhttptrace.Inject(ictx, req)

		res, err := cleanhttp.DefaultClient().Do(req)
		if err != nil {
			sp.SetStatus(codes.Error, "failure sending HTTP request")
			sp.RecordError(err)
			errCh <- err
			return
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			res.Body.Close()
			sp.SetStatus(codes.Error, err.Error())
			sp.RecordError(err)
			errCh <- err
			return
		}
		res.Body.Close()

		fib, err := strconv.Atoi(string(body))
		if err != nil {
			sp.RecordError(err)
			sp.SetStatus(codes.Error, "failure parsing result")
			errCh <- err
			return
		}

		c2 <- uint(fib)
		sp.AddEvent("sent result on channel", trace.WithAttributes(
			attribute.Int("result", int(fib)),
			attribute.String("chan", "c2"),
		))
	}(r.Context(), uint(nth))

	var result uint
	for i := 0; i < 2; i++ {
		select {
		case a := <-c1:
			result += a
			span.SetAttributes(attribute.Int("result.a", int(a)))
		case b := <-c2:
			result += b
			span.SetAttributes(attribute.Int("result.b", int(b)))
		case err := <-errCh:
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			w.WriteHeader(http.StatusInternalServerError) // 500
			fmt.Fprint(w, err.Error())
			return
		}
	}

	span.SetAttributes(attribute.Int("result.f", int(result)))
	fmt.Fprint(w, result)
}
