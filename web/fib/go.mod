module github.com/tullo/otel-workshop/web/fib

go 1.23.0

toolchain go1.24.5

require (
	github.com/hashicorp/go-cleanhttp v0.5.2
	go.opentelemetry.io/contrib/instrumentation/net/http/httptrace/otelhttptrace v0.62.0
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.62.0
	go.opentelemetry.io/otel v1.37.0
	go.opentelemetry.io/otel/trace v1.37.0
)

require (
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/go-logr/logr v1.4.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	go.opentelemetry.io/auto/sdk v1.1.0 // indirect
	go.opentelemetry.io/otel/metric v1.37.0 // indirect
)
