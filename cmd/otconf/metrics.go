package main

import (
	"context"
	"fmt"
	"time"

	hostMetrics "go.opentelemetry.io/contrib/instrumentation/host"
	runtimeMetrics "go.opentelemetry.io/contrib/instrumentation/runtime"
	metricglobal "go.opentelemetry.io/otel/metric/global"
	controller "go.opentelemetry.io/otel/sdk/metric/controller/basic"
	processor "go.opentelemetry.io/otel/sdk/metric/processor/basic"
	selector "go.opentelemetry.io/otel/sdk/metric/selector/simple"
)

func setupMetrics(c Config, headerName string) (func() error, error) {
	if !c.MetricsEnabled {
		c.logger.Debugf("metrics are disabled by configuration: no endpoint set")
		return nil, nil
	}
	metricExporter, err := newMetricsExporter(c, headerName)
	if err != nil {
		return nil, fmt.Errorf("failed to create metric exporter: %v", err)
	}

	period := controller.DefaultPeriod
	if c.MetricReportingPeriod != "" {
		period, err = time.ParseDuration(c.MetricReportingPeriod)
		if err != nil {
			return nil, fmt.Errorf("invalid metric reporting period: %v", err)
		}
		if period <= 0 {

			return nil, fmt.Errorf("invalid metric reporting period: %v", c.MetricReportingPeriod)
		}
	}
	pusher := controller.New(
		processor.New(
			selector.NewWithInexpensiveDistribution(),
			metricExporter,
		),
		controller.WithExporter(metricExporter),
		controller.WithResource(c.Resource),
		controller.WithCollectPeriod(period),
	)

	if err = pusher.Start(context.Background()); err != nil {
		return nil, fmt.Errorf("failed to start controller: %v", err)
	}

	provider := pusher.MeterProvider()

	if err = runtimeMetrics.Start(runtimeMetrics.WithMeterProvider(provider)); err != nil {
		return nil, fmt.Errorf("failed to start runtime metrics: %v", err)
	}

	if err = hostMetrics.Start(hostMetrics.WithMeterProvider(provider)); err != nil {
		return nil, fmt.Errorf("failed to start host metrics: %v", err)
	}

	metricglobal.SetMeterProvider(provider)
	return func() error {
		_ = pusher.Stop(context.Background())
		return metricExporter.Shutdown(context.Background())
	}, nil
}
