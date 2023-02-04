module github.com/tullo/otel-workshop/cmd/stdout/loop

go 1.19

replace github.com/tullo/otel-workshop/loop/fib => /home/anda/code/otel/workshop/loop/fib

require (
	github.com/tullo/otel-workshop/loop/fib bb4270506960
	go.opentelemetry.io/otel v1.12.0
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.12.0
	go.opentelemetry.io/otel/sdk v1.12.0
	go.opentelemetry.io/otel/trace v1.12.0
)

require (
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	golang.org/x/sys v0.4.0 // indirect
)
