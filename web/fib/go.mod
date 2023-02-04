module github.com/tullo/otel-workshop/web/fib

go 1.19

require (
	github.com/hashicorp/go-cleanhttp v0.5.2
	go.opentelemetry.io/contrib/instrumentation/net/http/httptrace/otelhttptrace v0.38.0
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.38.0
	go.opentelemetry.io/otel v1.12.0
	go.opentelemetry.io/otel/trace v1.12.0
)

require (
	github.com/felixge/httpsnoop v1.0.3 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	go.opentelemetry.io/otel/metric v0.35.0 // indirect
)
