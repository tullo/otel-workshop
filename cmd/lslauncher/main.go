package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/lightstep/otel-launcher-go/launcher"
	"github.com/tullo/otel-workshop/web/fib"
	"go.opentelemetry.io/otel/attribute"
	metricglobal "go.opentelemetry.io/otel/metric/global"
)

func newLighstepLauncher() launcher.Launcher {
	/*
		Export env vars picked up by launcher
		LS_ACCESS_TOKEN=...
		LS_SERVICE_NAME=fib-workshop
		LS_SERVICE_VERSION=v0.1.0
		OTEL_EXPORTER_OTLP_SPAN_ENDPOINT=ingest.lightstep.com:443
		OTEL_LOG_LEVEL=debug
		OTEL_RESOURCE_ATTRIBUTES=host.hostname=fib.workshop.com,deployment.region=eu-west1,deployment.environment=workshop
	*/
	// return launcher.ConfigureOpentelemetry()
	otel := launcher.ConfigureOpentelemetry(
		launcher.WithServiceName("fib-workshop"),
		launcher.WithAccessToken("access-token"),
	)
	return otel
	/*
		ls := launcher.ConfigureOpentelemetry(
			launcher.WithServiceName("fib-workshop"),
			launcher.WithServiceVersion("v0.1.2"),
			launcher.WithResourceAttributes(map[string]string{
				"host.hostname":          "fib.workshop.com",
				"deployment.region":      "eu-west1",
				"deployment.environment": "workshop",
			}),
			launcher.WithAccessToken("..."),
		)
		return ls
	*/
}

func main() {
	l := log.New(os.Stdout, "", 0)

	ls := newLighstepLauncher()
	defer ls.Shutdown()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	errCh := make(chan error)

	// Start metrics collection.
	go collectMetrics(context.Background())

	// Start web server.
	s := fib.NewServer(os.Stdin, l)
	go func() {
		errCh <- s.Serve(context.Background())
	}()

	select {
	case <-sigCh:
		l.Println("\ngoodbye")
		return
	case err := <-errCh:
		if err != nil {
			l.Fatal(err)
		}
	}
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
