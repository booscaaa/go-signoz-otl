package util

import (
	"context"
	"log"

	"github.com/spf13/viper"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	"google.golang.org/grpc/credentials"

	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

var (
	ServiceName  = ""
	CollectorURL = ""
	Insecure     = false
)

func InitTracer() *sdktrace.TracerProvider {
	ServiceName = viper.GetString("otl.service_name")
	CollectorURL = viper.GetString("otl.otel_exporter_otlp_endpoint")
	Insecure = viper.GetBool("otl.insecure_mode")

	secureOption := otlptracegrpc.WithTLSCredentials(credentials.NewClientTLSFromCert(nil, ""))
	if Insecure {
		secureOption = otlptracegrpc.WithInsecure()
	}

	ctx := context.Background()

	exporter, err := otlptrace.New(
		ctx,
		otlptracegrpc.NewClient(
			secureOption,
			otlptracegrpc.WithEndpoint(CollectorURL),
		),
	)

	if err != nil {
		log.Fatal(err)
	}

	resources, err := resource.New(
		ctx,
		resource.WithAttributes(
			attribute.String("service.name", ServiceName),
			attribute.String("library.language", "go"),
		),
	)

	if err != nil {
		log.Printf("Could not set resources: %v", err)
	}

	provider := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resources),
	)

	otel.SetTracerProvider(
		provider,
	)
	return provider
}
