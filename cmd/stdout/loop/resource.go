package main

import (
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
	sc "go.opentelemetry.io/otel/semconv/v1.7.0"
)

// newResource returns a resource describing this application.
func newResource() *resource.Resource {
	r, _ := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			sc.SchemaURL,
			sc.ServiceNameKey.String("fib"),
			sc.ServiceVersionKey.String("v0.1.0"),
			attribute.String("environment", "workshop"),
		),
	)

	return r
}
