package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/tullo/otel-workshop/cmd/datadog/fib"
	httptrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/net/http"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"gopkg.in/DataDog/dd-trace-go.v1/profiler"
)

func main() {
	// Send 100% of the traces to the backend.
	rules := []tracer.SamplingRule{tracer.RateRule(1)}

	tracer.Start(
		tracer.WithAgentAddr(os.Getenv("DD_AGENT_ADDRESS")),
		tracer.WithSamplingRules(rules),
		tracer.WithService(os.Getenv("SERVICE_NAME")),
		tracer.WithServiceVersion(os.Getenv("DD_DEPLOYMENT")),
		tracer.WithEnv(os.Getenv("DD_ENV")),
		tracer.WithDebugMode(false),
	)
	defer tracer.Stop()

	// The profiles below are disabled by default to keep
	// overhead low, but can be enabled as needed.
	if err := profiler.Start(
		profiler.WithAgentAddr(os.Getenv("DD_AGENT_ADDRESS")),
		profiler.WithService(os.Getenv("SERVICE_NAME")),
		profiler.WithVersion(os.Getenv("DD_DEPLOYMENT")),
		profiler.WithEnv(os.Getenv("DD_ENV")),
		profiler.WithProfileTypes(
			profiler.CPUProfile,
			profiler.HeapProfile,
			// profiler.BlockProfile,
			// profiler.MutexProfile,
			// profiler.GoroutineProfile,
		),
	); err != nil {
		log.Fatal(err)
	}
	defer profiler.Stop()

	// Traced mux
	mux := httptrace.NewServeMux()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	errCh := make(chan error)

	// Start web server.
	l := log.New(os.Stdout, "", 0)
	s := fib.NewServer(os.Stdin, l)
	go func() {
		errCh <- s.ServeDataDog(mux)
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
