# Splunk

- [Splunk Observability Cloud](https://www.splunk.com/en_us/download/infrastructure-monitoring.html) Free for 14 days
- [app.eu0.signalfx.com](https://app.eu0.signalfx.com/#/home) Backend web UI.
- [Integrations](https://app.eu0.signalfx.com/#/integrations) Connect to your systems and data sources.
- [Training / courses-for-signalfx](https://www.splunk.com/en_us/training/learning-path/courses-for-signalfx/overview.html)
- [Blog](https://www.splunk.com/en_us/blog)
  - [OpenTelemetry Support in Splunk APM](https://www.splunk.com/en_us/blog/conf-splunklive/announcing-native-opentelemetry-support-in-splunk-apm.html) (October 20, 2020)

## Developer Guide for Splunk Observability Cloud

- [Observability](https://docs.splunk.com/Observability/) docs / data setup
- [Splunk Observability Cloud API](https://dev.splunk.com/observability/) Developer Guide + API Reference
  - [Realms in endpoints](https://dev.splunk.com/observability/docs/realms_in_endpoints/) `https://api.<REALM>.signalfx.com`
- [net/http@v1.12.0](https://pkg.go.dev/github.com/signalfx/signalfx-go-tracing/contrib/net/http@v1.12.0) provides functions to trace the net/http package
- [signalfx-go-tracing/contrib](https://pkg.go.dev/github.com/signalfx/signalfx-go-tracing/contrib/@v1.12.0#section-readme) provide tracing on top of commonly used packages