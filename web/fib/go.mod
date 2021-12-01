module github.com/tullo/otel-workshop/web/fib

go 1.17

require (
	github.com/hashicorp/go-cleanhttp v0.5.2
	github.com/newrelic/go-agent/v3 v3.15.1
	go.opentelemetry.io/contrib/instrumentation/net/http/httptrace/otelhttptrace v0.27.0
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.27.0
	go.opentelemetry.io/otel v1.2.0
	go.opentelemetry.io/otel/trace v1.2.0
)

require (
	github.com/felixge/httpsnoop v1.0.2 // indirect
	github.com/golang/protobuf v1.3.3 // indirect
	go.opentelemetry.io/otel/internal/metric v0.25.0 // indirect
	go.opentelemetry.io/otel/metric v0.25.0 // indirect
	golang.org/x/net v0.0.0-20190311183353-d8887717615a // indirect
	golang.org/x/sys v0.0.0-20190215142949-d0b11bdaac8a // indirect
	golang.org/x/text v0.3.0 // indirect
	google.golang.org/genproto v0.0.0-20190819201941-24fa4b261c55 // indirect
	google.golang.org/grpc v1.27.0 // indirect
)
