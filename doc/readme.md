# Workshop

Slides: [Understanding Distributed Systems with OpenTelemetry](doc/OT-Workshop-Fib-example.pdf)

© [Adam Johnson](https://bit.ly/3wigdFJ), Lightstep

## SDK, Exporters and Collector Services

- SDK
  - OpenTelemetry's SDK implements `trace` & `span` creation.
- Exporter
  - Exporters can be instantiated to send the data collected by OpenTelemetry to the backend of your choice.
  - E.g. Jaeger, Lightstep, Honeycomb, Stackdriver, etc.
- Collector
  - OpenTelemetry collector proxies data between instrumented code and backend service(s).
  - The exporters
can be reconﬁgured without changing instrumented code.
  - Receive multiple trace formats and marshal them into a new one.
  - Enhance trace/metric data with resources.
  - Perform sampling, ﬁltering, or custom processors to modify attributes.
  - Run them as agents or in standalone mode.
    - A Kubernetes operator exists to aid in deployment.

## OT API
- Tracer
  - Responsible for tracking the currently active span.
- Meter
  - Responsible for accumulating a collection of statistics.

### Tracer methods

 - `tracer.Start(ctx, name, options)`
   - Returns a child of the current span, and makes it current.
 - `tracer.SpanFromContext(ctx)`
   - Access & add information to the current span.
 - `span.AddEvent(ctx, msg)`
   - Adds structured annotations (e.g. logs) about what is currently happening.
 - `span.SetAttribute(key, value)`
   - Adds a structured, typed attribute to the current span.
   - This may include an user id, an build id, an user-agent, etc.
- `span.End()`
  - Called when the unit of work is complete and the span can be sent.

```go
// Get the current span
sp := trace.SpanFromContext(ctx)

// Update the span status
sp.SetStatus(codes.OK)

// Add events
sp.AddEvent(ctx, "user approved")

// Add attributes
sp.SetAttributes(attribute.String("key", "value"))
```

## Context Propagation
- Distributed context is an abstract data type that represents collection of entries.
- Each key is associated with exactly one value.
- It is serialized for propagation across **process boundaries**.
- Passing around the context enables related spans to be associated with a single trace.
- `W3C` TraceContext is the de-facto standard.
  - `B3` is more broadly compatible with existing systems.

## Automatic Instrumentation

OpenTelemetry has wrappers around common frameworks to propagate context and make it accessible.
```go
import "go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"

// Wraps the handler and starts a trace named 'root'.
mux.Handle("/", otelhttp.NewHandler(otelhttp.WithRouteTag("/", http.HandlerFunc(someHandler)), "root", otelhttp.WithPublicEndpoint()))

func someHandler(w ResponseWriter, r *Request) {
    span := trace.SpanFromContext(r.Context())
}
```

## Manual Instrumentation

 - [resource deﬁnition](../cmd/stdout/loop/resource.go) - identiﬁes the resource or machine creating the telemetry.
 - [exporter](../cmd/stdout/loop/exporter.go) - sends the tracing data somewhere.
 - tracerProvider - registered with the SDK.
   - The TracerProvider unites:
     - the spans that are made,
     - the resource deﬁnition for all spans produced, 
     - the exporter.
    ```go
    tp := trace.NewTracerProvider(
      trace.WithBatcher(exporter()),
      trace.WithResource(resource()),
    )
    otel.SetTracerProvider(tp)
    ```
 - Instrument the code - add spans.
    ```go
    func (a *App) Poll(ctx context.Context) (int, error) {
      _, span := otel.Tracer(name).Start(ctx, "Poll")
      defer span.End()
    }
    ```

## Automatic Metrics Instrumentation

Metric instrumentation looks a little different

- Often agents are in place to collect metrics automatically
  - Just conﬁgure to export to an OT collector with a conﬁgured receiver
- OT collector can both PUSH and PULL
  - Works with existing Prometheus libraries

## Manual Metrics Instrumentation

 - [resource deﬁnition](../cmd/otconf/config.go#L127) - identiﬁes the resource or machine creating the metrics.
 - [exporter](../cmd/otconf/exporter.go#L35) - sends the metrics data somewhere.
 - [controller](../cmd/otconf/metrics.go#L36)
    ```go
    pusher := controller.New(
      // ...
      controller.WithExporter(metricExporter),
      controller.WithResource(resource),
      controller.WithCollectPeriod(period),
    )
    ```
 - [meterProvider](../cmd/otconf/metrics.go#L50) - registered with the SDK.
   - `provider := pusher.MeterProvider()`
   - `mglobal.SetMeterProvider(provider)`
 - Register metrics to your controller and start measuring!
    ```go
    if err = runtimeMetrics.Start(runtimeMetrics.WithMeterProvider(provider)); err != nil {
      return nil, fmt.Errorf("failed to start runtime metrics: %v", err)
    }

    if err = hostMetrics.Start(hostMetrics.WithMeterProvider(provider)); err != nil {
      return nil, fmt.Errorf("failed to start host metrics: %v", err)
    }
    ```
