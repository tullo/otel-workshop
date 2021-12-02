module github.com/tullo/otel-workshop/cmd/elastic

go 1.17

replace github.com/tullo/otel-workshop/web/fib => /home/anda/code/otel/workshop/web/fib

require (
	github.com/tullo/otel-workshop/web/fib v0.3.2
	go.opentelemetry.io/otel v1.2.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.2.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.2.0
	go.opentelemetry.io/otel/sdk v1.2.0
)

require (
	github.com/armon/go-radix v1.0.0 // indirect
	github.com/cenkalti/backoff/v4 v4.1.1 // indirect
	github.com/elastic/go-licenser v0.3.1 // indirect
	github.com/elastic/go-sysinfo v1.1.1 // indirect
	github.com/elastic/go-windows v1.0.0 // indirect
	github.com/felixge/httpsnoop v1.0.2 // indirect
	github.com/go-logfmt/logfmt v0.5.0 // indirect
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.16.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/jcchavezs/porto v0.1.0 // indirect
	github.com/joeshaw/multierror v0.0.0-20140124173710-69b34d4ec901 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/newrelic/go-agent/v3 v3.15.1 // indirect
	github.com/opentracing/opentracing-go v1.1.0 // indirect
	github.com/philhofer/fwd v1.1.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/prometheus/procfs v0.0.8 // indirect
	github.com/santhosh-tekuri/jsonschema v1.2.4 // indirect
	github.com/signalfx/golib v2.5.1+incompatible // indirect
	github.com/signalfx/signalfx-go-tracing v1.12.0 // indirect
	github.com/signalfx/signalfx-go-tracing/contrib/net/http v1.12.0 // indirect
	github.com/tinylib/msgp v1.1.6 // indirect
	go.elastic.co/apm v1.14.0 // indirect
	go.elastic.co/apm/module/apmhttp v1.14.0 // indirect
	go.elastic.co/fastjson v1.1.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/httptrace/otelhttptrace v0.27.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.27.0 // indirect
	go.opentelemetry.io/otel/internal/metric v0.25.0 // indirect
	go.opentelemetry.io/otel/metric v0.25.0 // indirect
	go.opentelemetry.io/otel/trace v1.2.0 // indirect
	go.opentelemetry.io/proto/otlp v0.10.0 // indirect
	golang.org/x/lint v0.0.0-20201208152925-83fdc39ff7b5 // indirect
	golang.org/x/mod v0.3.0 // indirect
	golang.org/x/net v0.0.0-20210226172049-e18ecbb05110 // indirect
	golang.org/x/sys v0.0.0-20210423185535-09eb48e85fd7 // indirect
	golang.org/x/text v0.3.3 // indirect
	golang.org/x/tools v0.0.0-20201022035929-9cf592e881e9 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/grpc v1.42.0 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	howett.net/plist v0.0.0-20181124034731-591f970eefbb // indirect
)
