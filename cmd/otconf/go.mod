module github.com/tullo/otel-workshop/cmd/otconf

go 1.17

replace (
	//github.com/tullo/otel-workshop/web/fib => /home/anda/code/otel/workshop/web/fib
	go.opentelemetry.io/otel/internal/metric v0.25.0 => go.opentelemetry.io/otel/internal/metric v0.23.0
	go.opentelemetry.io/otel/metric v0.25.0 => go.opentelemetry.io/otel/metric v0.23.0
)

require (
	github.com/sethvargo/go-envconfig v0.3.2
	github.com/tullo/otel-workshop/web/fib v0.0.0-20211124224815-85c39ec8e6f8
	go.opentelemetry.io/collector v0.28.0
	go.opentelemetry.io/contrib/instrumentation/host v0.23.0
	go.opentelemetry.io/contrib/instrumentation/runtime v0.23.0
	go.opentelemetry.io/contrib/propagators/b3 v0.23.0
	go.opentelemetry.io/contrib/propagators/ot v0.23.0
	go.opentelemetry.io/otel v1.2.0
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric v0.23.0
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc v0.23.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.0.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.0.0
	go.opentelemetry.io/otel/metric v0.25.0
	go.opentelemetry.io/otel/sdk v1.1.0
	go.opentelemetry.io/otel/sdk/metric v0.23.0
	google.golang.org/grpc v1.40.0
)

require (
	github.com/StackExchange/wmi v0.0.0-20210224194228-fe8f1750fd46 // indirect
	github.com/cenkalti/backoff/v4 v4.1.1 // indirect
	github.com/felixge/httpsnoop v1.0.2 // indirect
	github.com/go-ole/go-ole v1.2.5 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.16.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/shirou/gopsutil v3.21.5+incompatible // indirect
	github.com/tklauser/go-sysconf v0.3.5 // indirect
	github.com/tklauser/numcpus v0.2.2 // indirect
	go.opentelemetry.io/contrib v0.23.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/httptrace/otelhttptrace v0.27.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.27.0 // indirect
	go.opentelemetry.io/otel/internal/metric v0.25.0 // indirect
	go.opentelemetry.io/otel/sdk/export/metric v0.23.0 // indirect
	go.opentelemetry.io/otel/trace v1.2.0 // indirect
	go.opentelemetry.io/proto/otlp v0.9.0 // indirect
	golang.org/x/net v0.0.0-20210510120150-4163338589ed // indirect
	golang.org/x/sys v0.0.0-20210423185535-09eb48e85fd7 // indirect
	golang.org/x/text v0.3.6 // indirect
	google.golang.org/genproto v0.0.0-20210312152112-fc591d9ea70f // indirect
	google.golang.org/protobuf v1.27.1 // indirect
)
