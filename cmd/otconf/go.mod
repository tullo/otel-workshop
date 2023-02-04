module github.com/tullo/otel-workshop/cmd/otconf

go 1.19

replace github.com/tullo/otel-workshop/web/fib => /home/anda/code/otel/workshop/web/fib

require (
	github.com/sethvargo/go-envconfig v0.8.3
	github.com/tullo/otel-workshop/web/fib v0.0.0-20230202225614-77ab436b2376
	go.opentelemetry.io/contrib/propagators/b3 v1.13.0
	go.opentelemetry.io/contrib/propagators/ot v1.13.0
	go.opentelemetry.io/otel v1.12.0
	go.opentelemetry.io/otel/sdk v1.12.0
)

require (
	github.com/felixge/httpsnoop v1.0.3 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/httptrace/otelhttptrace v0.38.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.38.0 // indirect
	go.opentelemetry.io/otel/metric v0.35.0 // indirect
	go.opentelemetry.io/otel/trace v1.12.0 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/multierr v1.9.0 // indirect
	golang.org/x/sys v0.4.0 // indirect
)
