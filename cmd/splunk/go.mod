module github.com/tullo/otel-workshop/cmd/splunk

go 1.19

// replace github.com/tullo/otel-workshop/web/fib => /home/anda/code/otel/workshop/web/fib

require (
	github.com/hashicorp/go-cleanhttp v0.5.2
	github.com/signalfx/signalfx-go-tracing v1.12.0
	github.com/signalfx/signalfx-go-tracing/contrib/net/http v1.12.0
	go.opentelemetry.io/contrib/instrumentation/net/http/httptrace/otelhttptrace v0.38.0
	go.opentelemetry.io/otel v1.12.0
	go.opentelemetry.io/otel/trace v1.12.0
)

require (
	github.com/go-logfmt/logfmt v0.5.0 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/opentracing/opentracing-go v1.1.0 // indirect
	github.com/philhofer/fwd v1.1.1 // indirect
	github.com/signalfx/golib v2.5.1+incompatible // indirect
	github.com/tinylib/msgp v1.1.6 // indirect
	golang.org/x/sys v0.0.0-20201119102817-f84b799fce68 // indirect
)
