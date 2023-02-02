module github.com/tullo/otel-workshop/cmd/stdout/recursive

go 1.19

replace github.com/tullo/otel-workshop/recursive/fib => /home/anda/code/otel/workshop/recursive/fib

require (
	github.com/tullo/otel-workshop/recursive/fib v0.0.0-00010101000000-000000000000
	go.opentelemetry.io/otel v1.12.0
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.12.0
	go.opentelemetry.io/otel/sdk v1.12.0
	go.opentelemetry.io/otel/trace v1.12.0
)

require (
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	golang.org/x/sys v0.0.0-20220919091848-fb04ddd9f9c8 // indirect
)
