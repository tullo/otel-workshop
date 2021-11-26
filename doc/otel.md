# OpenTelemetry

- [Concepts](https://opentelemetry.io/docs/concepts/) multiple components each made available as a single implementation.
- [Collector](https://opentelemetry.io/docs/collector/) receive, process and export telemetry data.
- [OpenTelemetry Specification](https://opentelemetry.io/docs/reference/specification/) v1.0
  - [ Client Design Principles](https://opentelemetry.io/docs/reference/specification/library-guidelines/)
  - [Environment Variables](https://opentelemetry.io/docs/reference/specification/sdk-environment-variables/)
  - [Error handling](https://opentelemetry.io/docs/reference/specification/error-handling/)
  - [Logging](https://opentelemetry.io/docs/reference/specification/logs/overview/)
  - [Metrics](https://opentelemetry.io/docs/reference/specification/metrics/)
  - [Protocol](https://opentelemetry.io/docs/reference/specification/protocol/)
    - Mandatory [Exporter Configuration Options](https://opentelemetry.io/docs/reference/specification/protocol/exporter/)
    - [Protocol Specification](https://opentelemetry.io/docs/reference/specification/protocol/otlp/)
  - [Resource SDK](https://opentelemetry.io/docs/reference/specification/resource/sdk/)
  - [Tracing API](https://opentelemetry.io/docs/reference/specification/trace/api/)
- [Tracing API](https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/trace/api.md) (GH)
- [Tracing SDK](https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/trace/sdk.md) (GH)

## OpenTelemetry Go API and SDK

[opentelemetry-go](https://github.com/open-telemetry/opentelemetry-go/releases/tag/v1.2.0) Release v1.2.0/v0.25.0 (November 2021)

## OpenTelemetry Blog

- [OT auto-instrumentation/agents in Kubernetes](https://medium.com/opentelemetry/using-opentelemetry-auto-instrumentation-agents-in-kubernetes-869ec0f42377) (2021-11-22)

## Status - ultimo 2021

OpenTelemetry [Tracing Specification v1.0](https://opentelemetry.io/status/):

- Our goal is to provide a generally available, production quality release for the **tracing data** source across most OpenTelemetry components in the first half of 2021.
- Several components have already reached this milestone!
- We expect **metrics** to reach the same status in the second half of 2021 and are targeting **logs** in 2022.
