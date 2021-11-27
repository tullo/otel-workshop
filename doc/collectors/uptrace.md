# Uptrace

- [Uptrace](https://docs.uptrace.dev/) - all-in-one tool to optimize performance and monitor errors & logs
  - [OT tracing and metrics](https://opentelemetry.uptrace.dev/guide/) guides
  - [Span Attributes](https://opentelemetry.uptrace.dev/attributes/) contextual information
  - [OT Collector Configurator](https://opentelemetry.uptrace.dev/collector/) generates OT collector config
    - [Redis](https://opentelemetry.uptrace.dev/collector/?receiver=redis)
    - [Docker](https://opentelemetry.uptrace.dev/collector/?receiver=dockerstats)
  - [Resource detectors](https://docs.uptrace.dev/guide/go.html#resource-detectors) (host, env, aws, google cloud)
  - [Tracing API for Go](https://opentelemetry.uptrace.dev/guide/go-tracing.html#installation)
  - [Metrics API for Go](https://opentelemetry.uptrace.dev/guide/go-metrics.html#instruments)
  - [OT instrumentations](https://opentelemetry.uptrace.dev/instrumentations/) - plugins for popular frameworks and libraries that use OpenTelemetry API
  - [Timeseries metrics](https://docs.uptrace.dev/guide/metrics.html) - in-app metrics, infrastructure metrics
  - [Querying data](https://docs.uptrace.dev/guide/querying.html) - add the most useful queries to your favorites
  - [Monitor command execution](https://docs.uptrace.dev/guide/uptrace-run.html#introduction) of a script or binary that you can't instrument

## OT Go distro for Uptrace

```go
// Install
go get github.com/uptrace/uptrace-go

// Config
import "github.com/uptrace/uptrace-go/uptrace"

// https://pkg.go.dev/github.com/uptrace/uptrace-go/uptrace#Option
uptrace.ConfigureOpentelemetry(
    // .env UPTRACE_DSN=https://<key>@api.uptrace.dev/<project_id>
    uptrace.WithDSN("https://<key>@api.uptrace.dev/<project_id>"),
    uptrace.WithServiceName("myservice"),
    uptrace.WithServiceVersion("v1.0.0"),
)
```

## Blog

- [Instrumenting basic Go app using OT Tracing](https://blog.uptrace.dev/posts/opentelemetry-tracing-instrumenting-basic-go-app.html)
- [Monitoring cache stats using OT Metrics](https://blog.uptrace.dev/posts/opentelemetry-metrics-cache-stats.html)
- [Monitoring Redis with OT Collector](https://blog.uptrace.dev/posts/opentelemetry-collector-monitoring-redis.html)
- [Tips on writing JSON REST APIs in Go](https://blog.uptrace.dev/posts/go-json-rest-api.html) (segmentio/encoding)
- [Functional Options](https://blog.uptrace.dev/posts/go-functional-options-named-args.html) for e.g. flexible config options
- [Scrape OT Metrics with Prometheus](https://blog.uptrace.dev/posts/prometheus-opentelemetry-metrics.html)
