package main

import (
	main_configs "baseapplicationgo/main/configs"
	main_error "baseapplicationgo/main/configs/error"
	mainConfigsRouterHttp "baseapplicationgo/main/configs/router"
	main_configs_yml "baseapplicationgo/main/configs/yml"
	mainGatewaysWs "baseapplicationgo/main/gateways/ws"
	"context"
	"fmt"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"log"
	"net/http"
	"os"
	"time"
)

const MSG_APPLICATION_FAILED = "Application has failed to start"
const MSG_LISTENER = "Listener on port: %s"
const IDX_APPLICATION_PORT = "Application.Port"

var (
	meter = otel.Meter("rolldice")
)

func init() {
	main_configs.InitConfigBeans()
}

func main() {
	defer main_configs.TerminateConfigBeans()

	otlpEndpoint := os.Getenv("OTLP_ENDPOINT_MINE")
	fmt.Println(otlpEndpoint)
	//otlpExporterEndpoint := os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")
	//fmt.Println(otlpExporterEndpoint)
	//otlpExporterTracesEndpoint := os.Getenv("OTEL_EXPORTER_OTLP_TRACES_ENDPOINT")
	//fmt.Println(otlpExporterTracesEndpoint)

	insecureOpt := otlptracegrpc.WithInsecure()

	// Update default OTLP reciver endpoint
	endpointOpt := otlptracegrpc.WithEndpoint(otlpEndpoint)

	ctx := context.Background()
	otlptracegrpc.NewClient(endpointOpt)
	exp, err := otlptracegrpc.New(ctx, endpointOpt, insecureOpt)
	if err != nil {
		panic(err)
	}

	r, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName("myapp"),
		),
	)

	if err != nil {
		panic(err)
	}

	tracerProvider := trace.NewTracerProvider(
		trace.WithBatcher(exp),
		trace.WithResource(r),
	)
	//
	//defer func() {
	//	if err := tracerProvider.Shutdown(ctx); err != nil {
	//		panic(err)
	//	}
	//}()
	otel.SetTracerProvider(tracerProvider)

	tracer := otel.Tracer("test-tracer")

	// work begins
	ctx, span := tracer.Start(
		ctx,
		"CollectorExporter-Example")
	defer span.End()

	for i := 0; i < 10; i++ {
		_, iSpan := tracer.Start(ctx, fmt.Sprintf("Sample-%d", i))
		log.Printf("Doing really hard work (%d / 10)\n", i+1)

		<-time.After(time.Second)
		iSpan.End()
	}

	// From here, the tracerProvider can be used by instrumentation to collect
	// telemetry.

	insecureMetricOpt := otlpmetricgrpc.WithInsecure()

	// Update default OTLP reciver endpoint
	endpointMetricOpt := otlpmetricgrpc.WithEndpoint(otlpEndpoint)

	metricsExporter, err := otlpmetricgrpc.New(ctx, insecureMetricOpt, endpointMetricOpt)
	if err != nil {
		panic(err)
	}

	meterProvider := metric.NewMeterProvider(metric.WithReader(metric.NewPeriodicReader(metricsExporter)))
	defer func() {
		if err := meterProvider.Shutdown(ctx); err != nil {
			panic(err)
		}
	}()
	otel.SetMeterProvider(meterProvider)

	// From here, the meterProvider can be used by instrumentation to collect
	// telemetry.

	tracingServerName := "app-base-go"

	applicationPort := main_configs_yml.GetYmlValueByName(IDX_APPLICATION_PORT)
	routes := mainGatewaysWs.GetRoutesBean()
	router := mainGatewaysWs.ConfigRoutes(mainConfigsRouterHttp.GetRouterBean(), *routes)
	router.Use(otelmux.Middleware(tracingServerName))

	err2 := http.ListenAndServe(":"+applicationPort, router)
	if err2 != nil {
		main_error.FailOnError(err2, MSG_APPLICATION_FAILED)
	}
	log.Printf(MSG_LISTENER, applicationPort)
}
