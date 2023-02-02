package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/sethvargo/go-envconfig"
	//"go.opentelemetry.io/collector/translator/conventions"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/resource"
)

const (
	version                       = "1.0.0"
	DefaultSpanExporterEndpoint   = "ingest.lightstep.com:443"
	DefaultMetricExporterEndpoint = "ingest.lightstep.com:443"
)

type OTConf struct {
	config        Config
	shutdownFuncs []func() error
}

type Option func(*Config)

type Logger interface {
	Fatalf(format string, v ...interface{})
	Debugf(format string, v ...interface{})
}

func WithLogger(logger Logger) Option {
	return func(c *Config) {
		c.logger = logger
	}
}

type Config struct {
	SpanExporterEndpoint           string   `env:"OTEL_EXPORTER_OTLP_SPAN_ENDPOINT,default=ingest.lightstep.com:443"`
	SpanExporterEndpointInsecure   bool     `env:"OTEL_EXPORTER_OTLP_SPAN_INSECURE,default=false"`
	MetricExporterEndpoint         string   `env:"OTEL_EXPORTER_OTLP_METRIC_ENDPOINT,default=ingest.lightstep.com:443"`
	MetricExporterEndpointInsecure bool     `env:"OTEL_EXPORTER_OTLP_METRIC_INSECURE,default=false"`
	MetricReportingPeriod          string   `env:"OTEL_EXPORTER_OTLP_METRIC_PERIOD,default=30s"`
	LogLevel                       string   `env:"OTEL_LOG_LEVEL,default=info"`
	Propagators                    []string `env:"OTEL_PROPAGATORS,default=b3"`
	TracesSampler                  string   `env:"OTEL_TRACES_SAMPLER,default=traceidratio"`
	TracesSamplerArg               float64  `env:"OTEL_TRACES_SAMPLER_ARG,default=0.01"`
	AccessToken                    string   `env:"LS_ACCESS_TOKEN"`
	MetricsEnabled                 bool     `env:"LS_METRICS_ENABLED,default=true"`
	ServiceName                    string   `env:"LS_SERVICE_NAME"`
	ServiceVersion                 string   `env:"LS_SERVICE_VERSION,default=unknown"`
	resourceAttributes             map[string]string
	Resource                       *resource.Resource
	logger                         Logger
	errorHandler                   otel.ErrorHandler
}

func newConfig(opts ...Option) Config {
	var c Config
	envError := envconfig.Process(context.Background(), &c)
	c.logger = &DefaultLogger{}
	c.errorHandler = &defaultErrorHandler{logger: c.logger}
	var defaultOpts []Option

	for _, opt := range append(defaultOpts, opts...) {
		opt(&c)
	}
	c.Resource = newResource(&c)

	if envError != nil {
		c.logger.Fatalf("environment error: %v", envError)
	}

	return c
}

type defaultErrorHandler struct {
	logger Logger
}

func (h *defaultErrorHandler) Handle(err error) {
	h.logger.Debugf("error: %v\n", err)
}

type DefaultLogger struct {
}

func (l *DefaultLogger) Fatalf(format string, v ...interface{}) {
	log.Fatalf(format, v...)
}

func (l *DefaultLogger) Debugf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

type setupFunc func(Config, string) (func() error, error)

func ConfigureOpentelemetry(headerName string, opts ...Option) OTConf {
	c := newConfig(opts...)
	if c.LogLevel == "debug" {
		c.logger.Debugf("debug logging enabled")
		c.logger.Debugf("configuration")
		s, _ := json.MarshalIndent(c, "", "\t")
		c.logger.Debugf(string(s))
	}

	if c.errorHandler != nil {
		otel.SetErrorHandler(c.errorHandler)
	}

	/*
		otc := OTConf{config: c}
		for _, setup := range []setupFunc{setupTracing, setupMetrics} {
			shutdown, err := setup(c, headerName)
			if err != nil {
				c.logger.Fatalf("setup error: %v", err)
				continue
			}
			if shutdown != nil {
				otc.shutdownFuncs = append(otc.shutdownFuncs, shutdown)
			}
		}
	*/

	return OTConf{config: c}
}

func newResource(c *Config) *resource.Resource {
	r := resource.Environment()

	/*
		hostnameSet := false
		for iter := r.Iter(); iter.Next(); {
			if iter.Attribute().Key == conventions.AttributeHostName && len(iter.Attribute().Value.Emit()) > 0 {
				hostnameSet = true
			}
		}

		attributes := []attribute.KeyValue{
			//attribute.String(conventions.AttributeTelemetrySDKName, "otconf"),
			//attribute.String(conventions.AttributeTelemetrySDKLanguage, "go"),
			//attribute.String(conventions.AttributeTelemetrySDKVersion, version),
		}

		if len(c.ServiceName) > 0 {
			//attributes = append(attributes, attribute.String(conventions.AttributeServiceName, c.ServiceName))
		}

		if len(c.ServiceVersion) > 0 {
			//attributes = append(attributes, attribute.String(conventions.AttributeServiceVersion, c.ServiceVersion))
		}

		for key, value := range c.resourceAttributes {
			if len(value) > 0 {
				//if key == conventions.AttributeHostName {
				//	hostnameSet = true
				//}
				attributes = append(attributes, attribute.String(key, value))
			}
		}

		if !hostnameSet {
			hostname, err := os.Hostname()
			if err != nil {
				c.logger.Debugf("unable to set host.name. Set OTEL_RESOURCE_ATTRIBUTES=\"host.name=<your_host_name>\" env var or configure WithResourceAttributes in code: %v", err)
			} else {
				//attributes = append(attributes, attribute.String(conventions.AttributeHostName, hostname))
			}
		}

		attributes = append(r.Attributes(), attributes...)

		// These detectors can't actually fail, ignoring the error.
		r, _ = resource.New(
			context.Background(),
			resource.WithSchemaURL(sc.SchemaURL),
			resource.WithAttributes(attributes...),
		)

		// Note: There are new detectors we may wish to take advantage
		// of, now available in the default SDK (e.g., WithProcess(),
		// WithOSType(), ...).
	*/
	return r
}

func (otc OTConf) Shutdown() {
	for _, shutdown := range otc.shutdownFuncs {
		if err := shutdown(); err != nil {
			otc.config.logger.Fatalf("failed to stop exporter: %v", err)
		}
	}
}
