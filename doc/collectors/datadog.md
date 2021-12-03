# Datadog

Datadog supports a variety of open standards, including OpenTelemetry and OpenTracing.

## Datadog App

- https://app.datadoghq.eu/help/quick_start#ubuntu
- https://app.datadoghq.eu/apm/docs?architecture=host-based&language=go


## OpenTelemetry collector Datadog exporter

The OpenTelemetry Collector is a vendor-agnostic separate agent process for collecting and exporting telemetry data emitted by many processes.

## APM & Continuous Profiler

Setting up Datadog APM across hosts, containers or serverless functions takes just minutes.

- [Send traces to Datadog](https://docs.datadoghq.com/tracing/#send-traces-to-datadog)
- [Tracing Go Applications](https://docs.datadoghq.com/tracing/setup_overview/setup/go/?tab=containers)


## Universal Service Monitoring (beta-release)

- [Golden signals in seconds](https://www.datadoghq.com/blog/universal-service-monitoring-datadog/) (October 26, 2021)
- [The Four Golden Signals](https://sre.google/sre-book/monitoring-distributed-systems/#xref_monitoring_golden-signals) (Google SRE Book)

## Dash 2021

- [Guide to Datadog’s newest announcements](https://www.datadoghq.com/blog/dash-2021-new-feature-roundup/) (October 26, 2021)
- [Sensitive Data Scanner](https://docs.datadoghq.com/account_management/org_settings/sensitive_data_detection/) - credit card numbers, bank routing numbers, API keys, OAuth tokens etc.
- [Continuous Profiler](https://docs.datadoghq.com/tracing/profiler/)  - find CPU, memory, and IO bottlenecks.
  - [Enabling the Go Profiler](https://docs.datadoghq.com/tracing/profiler/enabling/go/)
  - `go get gopkg.in/DataDog/dd-trace-go.v1/profiler@latest`
- [Session Replay](https://www.datadoghq.com/blog/session-replay-datadog/) - view real-time user journeys.
- [Funnel Analysis](https://www.datadoghq.com/blog/reduce-customer-friction-funnel-analysis/) - understand and optimize key user flows.
- [Synthetic private locations](https://www.datadoghq.com/blog/private-synthetic-monitoring/) - test on-premise applications.
- Run [UDP and WebSocket API tests](https://www.datadoghq.com/blog/udp-websocket-api-tests/) to monitor latency-critical applications.
