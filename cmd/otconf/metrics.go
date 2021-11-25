package main

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"syscall"
	"time"

	hostMetrics "go.opentelemetry.io/contrib/instrumentation/host"
	runtimeMetrics "go.opentelemetry.io/contrib/instrumentation/runtime"
	"go.opentelemetry.io/otel/attribute"
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
		if period, err = time.ParseDuration(c.MetricReportingPeriod); err != nil {
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

func collectMetrics(ctx context.Context) {
	// Set attributes for all metrics
	appKey := attribute.Key("fib")
	containerKey := attribute.Key(os.Getenv("HOSTNAME"))

	// 1. Declare a meter.
	meter := metricglobal.Meter("container")

	// 2. Declare specific metrics to collect
	mem, _ := meter.NewInt64Counter("mem_usage")             // metric.WithDescription("Amount of memory used."),
	disc, _ := meter.NewFloat64Counter("disk_usage")         // metric.WithDescription("Amount of disk used."),
	quota, _ := meter.NewFloat64Counter("disk_quota")        // metric.WithDescription("Amount of disk quota available."),
	goroutines, _ := meter.NewInt64Counter("num_goroutines") // metric.WithDescription("Amount of goroutines running."),

	// 3. Fetch runtime measurements that application makes available.
	var ms runtime.MemStats
	for {
		// Read memory allocator statistics.
		runtime.ReadMemStats(&ms)

		// Syscall to gather file system statistics
		// for the current directory.
		var fs syscall.Statfs_t
		wd, _ := os.Getwd()
		syscall.Statfs(wd, &fs)
		all := float64(fs.Blocks) * float64(fs.Bsize)
		free := float64(fs.Bfree) * float64(fs.Bsize)

		// Assign observed metric values to our declared meters.
		meter.RecordBatch(ctx, []attribute.KeyValue{
			appKey.String(os.Getenv("PROJECT_DOMAIN")), // TODO project id?
			containerKey.String(os.Getenv("HOSTNAME"))},
			disc.Measurement(all-free),
			quota.Measurement(all),
			mem.Measurement(int64(ms.Sys)),
			goroutines.Measurement(int64(runtime.NumGoroutine())),
		)

		// Take measurements once per minute.
		time.Sleep(time.Minute)
	}
}
