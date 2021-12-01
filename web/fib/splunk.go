package fib

import (
	"context"
	"fmt"
	"net/http"

	httptrace "github.com/signalfx/signalfx-go-tracing/contrib/net/http"
)

func (s *Server) ServeSplunk(ctx context.Context) error {
	mux := httptrace.NewServeMux()
	mux.Handle("/", http.HandlerFunc(rootHandler))
	mux.Handle("/favicon.ico", http.NotFoundHandler())
	mux.Handle("/fib", http.HandlerFunc(fibHandler))
	mux.Handle("/fibinternal", http.HandlerFunc(fibHandler))

	s.l.Println("Your server is live!\nTry to navigate to: http://127.0.0.1:3000/fib?n=6")
	if err := http.ListenAndServe("127.0.0.1:3000", mux); err != nil {
		return fmt.Errorf("could not start web server: %w", err)
	}

	return nil
}
