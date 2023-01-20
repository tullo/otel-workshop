module github.com/tullo/otel-workshop/cmd/stdout/recursive

go 1.17

//replace github.com/tullo/otel-workshop/recursive/fib => /home/anda/code/otel/workshop/recursive/fib

require (
	github.com/tullo/otel-workshop/recursive/fib v0.0.0-20230120160738-51b8f6816f33
	go.opentelemetry.io/otel v1.2.0
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.2.0
	go.opentelemetry.io/otel/sdk v1.2.0
	go.opentelemetry.io/otel/trace v1.2.0
)

require golang.org/x/sys v0.0.0-20210423185535-09eb48e85fd7 // indirect
