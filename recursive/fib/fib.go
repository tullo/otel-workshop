package fib

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

const name = "otel-workshop"

type Request struct {
	ctx context.Context
	n   uint
	f   func(*Request) (uint, error)
}

func NewRequest(ctx context.Context, n uint) Request {
	return Request{ctx, n, FibonacciRec}
}

// FibonacciRec returns the n-th fibonacci number
// using recursive func calls.
func FibonacciRec(r *Request) (uint, error) {
	// Access the current active span via the context object.
	// span := trace.SpanFromContext(r.ctx)
	ctx, span := otel.Tracer(name).Start(r.ctx, "FibonacciRec")
	defer span.End()

	if r.n > 30 {
		return 0, fmt.Errorf("unsupported fibonacci number %d: too large", r.n)
	}

	switch r.n {
	case 0:
		return 0, nil
	case 1:
		return 1, nil
	case 2:
		return 2, nil
	}

	c1 := make(chan uint)
	c2 := make(chan uint)
	errCh := make(chan error)

	go func(ctx context.Context, n uint) {
		ictx, span := otel.Tracer(name).Start(
			ctx, "fibClient",
			trace.WithSpanKind(trace.SpanKindClient))
		defer span.End()
		req := Request{ictx, n - 1, r.f}
		span.SetAttributes(attribute.Int("request.n", int(req.n)))
		fib, err := req.f(&req)
		if err != nil {
			span.RecordError(err)
			errCh <- err
			return
		}
		c1 <- uint(fib)
		span.AddEvent("sent result on channel", trace.WithAttributes(
			attribute.Int("result", int(req.n)),
			attribute.String("chan", "c1"),
		))
	}(ctx, r.n)

	go func(ctx context.Context, n uint) {
		ictx, span := otel.Tracer(name).Start(
			ctx, "fibClient",
			trace.WithSpanKind(trace.SpanKindClient))
		defer span.End()
		req := Request{ictx, n - 2, r.f}
		span.SetAttributes(attribute.Int("request.n", int(req.n)))
		fib, err := req.f(&req)
		if err != nil {
			span.RecordError(err)
			errCh <- err
			return
		}
		c2 <- uint(fib)
		span.AddEvent("sent result on channel", trace.WithAttributes(
			attribute.Int("result", int(req.n)),
			attribute.String("chan", "c2"),
		))
	}(ctx, r.n)

	var result uint
	for i := 0; i < 2; i++ {
		select {
		case a := <-c1:
			result += a
		case b := <-c2:
			result += b
		case err := <-errCh:
			return 0, err
		}
	}

	span.SetAttributes(attribute.Int("result.f", int(result)))
	return result, nil
}
