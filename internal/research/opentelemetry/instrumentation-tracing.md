# OpenTelemetry SDK Usage for Distributed Tracing in Golang Applications

OpenTelemetry is a project to handle telemetry data (logs, metrics and traces) in a vendor neutral way. It provides SDKs to instrument applications, APIs and processing tools (collector).

This document describes our findings with using the OpenTelemetry Golang SDK for distributed tracing. To evaluate the functionality of the SDK, a Golang implementation for the reviews service of the [Istio Bookinfo](https://istio.io/latest/docs/examples/bookinfo/) example was extended.

## Tracing Capabilities of the OpenTelemetry SDK

The OpenTelemetry SDK can export traces using the OTLP protocol using GRPC or HTTP. The trace context can be propagated between processes using the W3C TraceContext or B3 headers.

Traces can be exported to external backends (e.g., Jaeger) using the OpenTelemetry collector (see [deployment](deployment-tracing.md)).

## Code Instrumentation

To instrument a Golang application, a `tracer` has to be created first. The following function initializes a tracer that exports spans to an OpenTelemetry collector using the GRPC protocol.

```
func initTracer() (trace.Tracer, error) {
	ctx := context.Background()
	allOpts := []otlpgrpc.Option{
		otlpgrpc.WithEndpoint("localhost:55680"),
		otlpgrpc.WithInsecure(),
	}

	driver := otlpgrpc.NewDriver(allOpts...)
	exporter, err := otlp.NewExporter(ctx, driver)
	if err != nil {
		return nil, err
	}
	// For the demonstration, use sdktrace.AlwaysSample sampler to sample all traces.
	// In a production application, use sdktrace.ProbabilitySampler with a desired probability.
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithSyncer(exporter),
		sdktrace.WithResource(resource.NewWithAttributes(semconv.ServiceNameKey.String("my-service"))),
	)
	defer func() { _ = tp.Shutdown(ctx) }()

	pusher := controller.New(
		processor.New(
			simple.NewWithExactDistribution(),
			exporter,
		),
		controller.WithExporter(exporter),
		controller.WithCollectPeriod(5*time.Second),
	)

	err = pusher.Start(ctx)
	if err != nil {
		return nil, err
	}

	// Handle this error in a sensible manner where possible
	defer func() { _ = pusher.Stop(ctx) }()
	otel.SetTracerProvider(tp)
	global.SetMeterProvider(pusher.MeterProvider())
	propagator := propagation.NewCompositeTextMapPropagator(propagation.Baggage{}, propagation.TraceContext{}, b3.B3{})
	otel.SetTextMapPropagator(propagator)
	return tp.Tracer("reviews-go"), nil
}
```

Individual code snippets can be instrumented with spans as follows:

```
var span trace.Span
ctx := r.Context()
ctx, span = h.tracer.Start(ctx, "myRequestHandler")
defer span.End()
...
```

Spans can be enhanced with additional events. Events can contain custom attributes. E.g.:

```
...
span.AddEvent("Acquiring lock", trace.WithAttributes(label.Int("pid", 1234))
mutex.Lock()
```

## Automatic Instrumentation

OpenTelemetry provides a list of [packages](https://github.com/open-telemetry/opentelemetry-go-contrib/tree/main/instrumentation) that can instrument specific Golang packages automatically. E.g., `net/http`.

## Findings

* The trace context propagation feature lacks documentation. We were not able to propagate the B3 context that has been injected by Istio to the OpenTelemetry SDK to the own span.
* Automatic instrumentation of the `net/http` packages did not generate any traces.
* Documentation for the metrics API and logs API yet. In particular the enhancement of trace spans with metrics is not available.
* Traces were exported to the collector unreliably. Needs further investigation
* In general, many aspects of the API are not documented yet.

## Conclusions

The tracing module of the OpenTelemetry Golang SDK is still in beta state. This becomes apparent in the current state of the documentation that is limited to code examples. Since the OpenTelemetry collector supports multiple wire protocols to receive traces, during the time of our PoC the OpenTracing SDK seems to be the better choice for application instrumentation.
