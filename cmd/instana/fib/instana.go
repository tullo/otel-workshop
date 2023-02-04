package fib

import (
	"fmt"
	"net/http"

	instana "github.com/instana/go-sensor"
)

func (s *Server) ServeInstana(sensor *instana.Sensor) error {
	span := sensor.Tracer().StartSpan("ServeInstana")
	defer span.Finish()

	mux := http.NewServeMux()
	mux.Handle("/", instana.TracingHandlerFunc(sensor, "/", rootHandler))
	mux.Handle("/favicon.ico", http.NotFoundHandler())
	mux.Handle("/fib", instana.TracingHandlerFunc(sensor, "/fib/{n}", fibHandler))
	mux.Handle("/fibinternal", instana.TracingHandlerFunc(sensor, "/fibinternal/{n}", fibHandler))

	s.l.Println("Your server is live!\nTry to navigate to: http://127.0.0.1:3000/fib?n=6")
	if err := http.ListenAndServe("127.0.0.1:3000", mux); err != nil {
		return fmt.Errorf("could not start web server: %w", err)
	}

	return nil
}
