package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/segmentio/encoding/json"
	"github.com/uptrace/uptrace-go/uptrace"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

var tracer = otel.Tracer("go-example-app")

func main() {
	ctx := context.Background()

	uptrace.ConfigureOpentelemetry(
		uptrace.WithServiceName("go-example-app"),
		uptrace.WithServiceVersion("1.0.0"),
	)
	defer uptrace.Shutdown(ctx)

	ctx, span := tracer.Start(ctx, "fetchCountry")
	defer span.End()

	countryInfo, err := fetchCountryInfo(ctx)
	if err != nil {
		span.RecordError(err)
		return
	}

	countryCode, countryName, err := parseCountryInfo(ctx, countryInfo)
	if err != nil {
		span.RecordError(err)
		return
	}

	span.SetAttributes(
		attribute.String("country.code", countryCode),
		attribute.String("country.name", countryName),
	)

	fmt.Println("trace URL", uptrace.TraceURL(span))
}

func fetchCountryInfo(ctx context.Context) ([]byte, error) {
	_, span := tracer.Start(ctx, "fetchCountryInfo")
	defer span.End()

	apiClient := http.Client{
		Timeout: time.Second * 2,
	}
	req, err := http.NewRequest(http.MethodGet, "https://ipapi.co/json/", nil)
	if err != nil {
		span.RecordError(err)
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "curl/7.74.0")
	resp, err := apiClient.Do(req)
	if err != nil {
		span.RecordError(err)
		return nil, err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		span.RecordError(err)
		return nil, err
	}

	span.SetAttributes(
		attribute.String("ip", "self"),
		attribute.Int("resp_len", len(b)),
	)

	return b, nil
}

func parseCountryInfo(ctx context.Context, countryInfo []byte) (code, country string, _ error) {
	_, span := tracer.Start(ctx, "parseCountryInfo")
	defer span.End()

	type geoloc struct {
		CountryName string `json:"country_name"`
		CountryCode string `json:"country_code"`
	}
	geo := geoloc{}
	jsonErr := json.Unmarshal(countryInfo, &geo)
	if jsonErr != nil {
		span.RecordError(jsonErr)
		return "", "", fmt.Errorf("ipapi: can't parse response: %q", string(countryInfo))
	}

	return geo.CountryCode, geo.CountryName, nil
}
