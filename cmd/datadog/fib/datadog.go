package fib

import (
	"fmt"
	"net/http"

	httptrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/net/http"
)

func (s *Server) ServeDataDog(mux *httptrace.ServeMux) error {
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
