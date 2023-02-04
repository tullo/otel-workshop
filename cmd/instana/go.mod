module github.com/tullo/otel-workshop/cmd/instana

go 1.19

require (
	github.com/hashicorp/go-cleanhttp v0.5.2
	github.com/instana/go-sensor v1.50.0
	go.opentelemetry.io/contrib/instrumentation/net/http/httptrace/otelhttptrace v0.38.0
	go.opentelemetry.io/otel v1.12.0
	go.opentelemetry.io/otel/trace v1.12.0
)

require (
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/looplab/fsm v0.1.0 // indirect
	github.com/opentracing/opentracing-go v1.1.0 // indirect
)
