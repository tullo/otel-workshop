package main

import (
	"context"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/tullo/otel-workshop/cmd/signoz/bookstore/controllers"
	"github.com/tullo/otel-workshop/cmd/signoz/bookstore/models"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

var (
	serviceName = os.Getenv("SERVICE_NAME")
	// signozToken  = os.Getenv("SIGNOZ_ACCESS_TOKEN")
	// collectorURL = os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")
	// insecure     = os.Getenv("INSECURE_MODE")
)

func initTracer() func(context.Context) error {

	return nil

	/* headers := map[string]string{
		"signoz-access-token": signozToken,
	}

	secureOption := otlpgrpc.WithTLSCredentials(credentials.NewClientTLSFromCert(nil, ""))
	if len(insecure) > 0 {
		secureOption = otlpgrpc.WithInsecure()
	}

	exporter, err := otlp.NewExporter(
		context.Background(),
		otlpgrpc.NewDriver(
			secureOption,
			otlpgrpc.WithEndpoint(collectorURL),
			otlpgrpc.WithHeaders(headers),
		),
	)

	if err != nil {
		log.Fatal(err)
	}
	resources, err := resource.New(
		context.Background(),
		resource.WithAttributes(
			attribute.String("service.name", serviceName),
			attribute.String("library.language", "go"),
		),
	)
	if err != nil {
		log.Printf("Could not set resources: ", err)
	}

	otel.SetTracerProvider(
		trace.NewTracerProvider(
			trace.WithConfig(trace.Config{DefaultSampler: trace.AlwaysSample()}),
			trace.WithSpanProcessor(trace.NewBatchSpanProcessor(exporter)),
			trace.WithSyncer(exporter),
			trace.WithResource(resources),
		),
	)
	return exporter.Shutdown */
}

func main() {

	cleanup := initTracer()
	defer cleanup(context.Background())

	r := gin.Default()
	r.Use(otelgin.Middleware(serviceName))

	// Connect to database
	models.ConnectDatabase()

	// Routes
	r.GET("/books", controllers.FindBooks)
	r.GET("/books/:id", controllers.FindBook)
	r.POST("/books", controllers.CreateBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	// Run the server
	r.Run(":8090")
}
