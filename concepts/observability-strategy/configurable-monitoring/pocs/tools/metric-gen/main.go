package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/open-telemetry/opentelemetry-collector-contrib/testbed/testbed"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

type flags struct {
	host string
	port int
}

func main() {
	var fl flags
	flag.StringVar(&fl.host, "host", "telemetry-otlp-metrics", "host to send metrics to")
	flag.IntVar(&fl.port, "port", 4317, "host's port to send metrics to")

	flag.Parse()

	if err := run(fl); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

func run(fl flags) error {
	sender := testbed.NewOTLPMetricDataSender(fl.host, fl.port)

	fmt.Println("Starting")

	if err := sender.Start(); err != nil {
		return fmt.Errorf("failed to start sender: %v", err)
	}

	fmt.Println("Sending metrics")

	workerCount := runtime.GOMAXPROCS(0)
	var wg sync.WaitGroup
	wg.Add(workerCount)

	ch := make(chan pmetric.Metrics, workerCount)
	for i := 0; i < workerCount; i++ {
		go func() {
			defer wg.Done()

			for m := range ch {
				if err := sender.ConsumeMetrics(context.TODO(), m); err != nil {
					fmt.Printf("Warning: failed to send metrics: %v\n", err)
				}
			}
		}()
	}

	metricSuffix := randSeq(5)
	metricID := 0
	for {
		ch <- generateMetric(metricSuffix, metricID)
		metricID++
	}

	sender.Flush()
	return nil
}

func generateMetric(suffix string, id int) pmetric.Metrics {
	totalResourceMetrics := 5
	totalAttributes := 7
	totalPts := 2
	startTime := time.Now()

	md := pmetric.NewMetrics()
	rms := md.ResourceMetrics()
	rms.EnsureCapacity(totalResourceMetrics)
	for i := 0; i < totalResourceMetrics; i++ {
		metric := rms.AppendEmpty().ScopeMetrics().AppendEmpty().Metrics().AppendEmpty()

		metric.SetName(fmt.Sprintf("metric_%s_%d", suffix, id))
		metric.SetDescription("my-md-description")
		metric.SetUnit("my-md-units")

		gauge := metric.SetEmptyGauge()
		pts := gauge.DataPoints()
		for i := 0; i < totalPts; i++ {
			pt := pts.AppendEmpty()
			pt.SetStartTimestamp(pcommon.NewTimestampFromTime(startTime))
			pt.SetTimestamp(pcommon.NewTimestampFromTime(time.Now()))
			pt.SetDoubleValue(1.0)

			for i := 0; i < totalAttributes; i++ {
				k := fmt.Sprintf("pt-label-key-%d", i)
				v := fmt.Sprintf("pt-label-val-%d", i)
				pt.Attributes().PutStr(k, v)
			}
		}
	}
	return md
}

var letters = []rune("abcdefghijklmnopqrstuvwxyz")

func randSeq(n int) string {
	runes := make([]rune, n)
	for i := range runes {
		runes[i] = letters[rand.Intn(len(letters))]
	}
	return string(runes)
}
