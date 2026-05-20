package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
	"go.opentelemetry.io/otel/trace"
)

const (
	exitCodeInterrupt       = 2
	letterBytes             = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	childSpanTotal          = 10
	childSpanValueLength    = 64
	childSpanAttributeCount = 40
)

var commonAttrs = []attribute.KeyValue{
	attribute.String("attrA", "chocolate"),
	attribute.String("attrB", "raspberry"),
	attribute.String("attrC", "vanilla"),
}

var targetURL = "localhost:4317"
var maxParallelism = 2

func main() {
	flag.StringVar(&targetURL, "t", "localhost:4317", "Specify OpenTelemetry collector url")
	flag.IntVar(&maxParallelism, "c", 2, "Specify max concurrent span generator")

	flag.Usage = func() {
		fmt.Printf("Usage of loadgenerator: \n")
		fmt.Printf("./loadgenerator -t telemerty-collector:4317 -c 5\n")
	}

	flag.Parse()

	log.Print("Waiting for connection...")

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	defer func() {
		signal.Stop(signalChan)
		cancel()
	}()

	shutdown, err := initProvider()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := shutdown(ctx); err != nil {
			log.Fatal("Failed to shutdown TracerProvider: %w", err)
		}
	}()

	go func() {
		select {
		case <-signalChan: // first signal, cancel context
			cancel()
		case <-ctx.Done():
		}
		<-signalChan // second signal, hard exit
		os.Exit(exitCodeInterrupt)
	}()

	start := time.Now()
	log.Printf("Load Generator started for %s with %d concurent producers", targetURL, maxParallelism)
	tracer := otel.Tracer("otlp-load-tester")
	run(ctx, tracer)
	elapsed := time.Since(start)
	log.Printf("Took %s", elapsed)
}

func initProvider() (func(context.Context) error, error) {
	ctx := context.Background()

	res, err := resource.New(ctx,
		resource.WithAttributes(
			// the service name used to display traces in backends
			semconv.ServiceNameKey.String("otel-load-generator"),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create resource: %w", err)
	}

	rand.Seed(time.Now().UnixNano())

	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	conn, err := grpc.DialContext(ctx, targetURL, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		return nil, fmt.Errorf("failed to create gRPC connection to collector: %w", err)
	}

	// Set up a trace exporter
	traceExporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(conn))
	if err != nil {
		return nil, fmt.Errorf("failed to create trace exporter: %w", err)
	}

	// Register the trace exporter with a TracerProvider, using a batch
	// span processor to aggregate spans before export.
	bsp := sdktrace.NewBatchSpanProcessor(traceExporter, sdktrace.WithMaxExportBatchSize(512), sdktrace.WithMaxQueueSize(2048))
	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(res),
		sdktrace.WithSpanProcessor(bsp),
	)
	otel.SetTracerProvider(tracerProvider)

	// set global propagator to trace context (the default is no-op).
	otel.SetTextMapPropagator(propagation.TraceContext{})

	// Shutdown will flush any remaining spans and shut down the exporter.
	return tracerProvider.Shutdown, nil
}

func run(ctx context.Context, tracer trace.Tracer) {
	jobsCh := make(chan struct{}, maxParallelism)
	var wg sync.WaitGroup
	wg.Add(maxParallelism)

	for i := 0; i < maxParallelism; i++ {
		go func() {
			defer wg.Done()
			for range jobsCh {
				produceTrace(ctx, tracer)
				for i := 0; i < childSpanTotal; i++ {
					produceSpan(ctx, tracer, i)
				}
			}
		}()
	}

	jobsTotal := 0
	for {
		select {
		case jobsCh <- struct{}{}:
			jobsTotal++
		case <-ctx.Done():
			log.Print("Context cancelled, closing the jobs channel...")
			close(jobsCh)
			log.Print("Closed the jobs channel")
			wg.Wait()
			log.Printf("Total traces: %d, spans: %d produced", jobsTotal, jobsTotal*(childSpanTotal+1))
			return
		}
	}
}

func produceTrace(ctx context.Context, tracer trace.Tracer) {
	ctx, span := tracer.Start(ctx, "root", trace.WithAttributes(commonAttrs...))
	defer span.End()
}

func produceSpan(ctx context.Context, tracer trace.Tracer, i int) {
	_, iSpan := tracer.Start(ctx, fmt.Sprintf("child-%d", i))
	iSpan.SetAttributes(generateRandomAttributes()...)
	defer iSpan.End()
}

func generateRandomAttributes() []attribute.KeyValue {

	var spanAttrs []attribute.KeyValue
	for k := 0; k < childSpanAttributeCount; k++ {
		b := make([]byte, childSpanValueLength)
		for i := range b {
			b[i] = letterBytes[rand.Intn(len(letterBytes))]
		}
		spanAttrs = append(spanAttrs, attribute.String(fmt.Sprintf("attribute-%d", k), string(b)))
	}

	return spanAttrs
}
