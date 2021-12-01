# Lightstep

> Increase developer productivity

- https://lightstep.com/increase-developer-productivity
- [Go Launcher](https://github.com/lightstep/otel-launcher-go), a Distro for OpenTelemetry.

```go
import "github.com/lightstep/otel-launcher-go/launcher"

func main() {
    otel := launcher.ConfigureOpentelemetry(
        launcher.WithServiceName("service-name"),
        launcher.WithAccessToken("access-token"),
    )
    defer otel.Shutdown()
}
```
