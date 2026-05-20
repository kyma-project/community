package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type MetricData struct {
	ResourceMetrics []ResourceMetrics `json:"resourceMetrics"`
}

type ResourceMetrics struct {
	Resource     Resource       `json:"resource"`
	ScopeMetrics []ScopeMetrics `json:"scopeMetrics"`
}

type Resource struct {
	Attributes []Attribute `json:"attributes"`
}

type ScopeMetrics struct {
	Scope   Scope     `json:"scope"`
	Metrics []Metrics `json:"metrics"`
}

type Scope struct {
}

type Metrics struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Unit        string `json:"unit,omitempty"`

	Sum       Sum       `json:"sum,omitempty"`
	Gauge     Gauge     `json:"gauge,omitempty"`
	Histogram Histogram `json:"histogram,omitempty"`
	Summary   Summary   `json:"summary,omitempty"`
}

type Sum struct {
	DataPoints []DataPoint `json:"dataPoints"`
}

type Gauge struct {
	DataPoints []DataPoint `json:"dataPoints"`
}

type Histogram struct {
	DataPoints []DataPoint `json:"dataPoints"`
}

type Summary struct {
	DataPoints []DataPoint `json:"dataPoints"`
}

type DataPoint struct {
	Attributes        []Attribute `json:"attributes,omitempty"`
	StartTimeUnixNano string      `json:"startTimeUnixNano,omitempty"`
	TimeUnixNano      string      `json:"timeUnixNano,omitempty"`
	AsDouble          float64     `json:"asDouble,omitempty"`
}

type Attribute struct {
	Key   string `json:"key"`
	Value Value  `json:"value"`
}

type Value struct {
	StringValue string `json:"stringValue"`
}

type flags struct {
	file string
}

type metricType string

var (
	metricTypeGauge     metricType = "Gauge"
	metricTypeSum       metricType = "Sum"
	metricTypeHistogram metricType = "Histogram"
	metricTypeSummary   metricType = "Summary"
)

type stats struct {
	byMetricType map[metricType]metricTypeStats
}

type metricTypeStats struct {
	total                     int
	maxDataPoints             int
	maxAttributesPerDataPoint int
	totalAttributes           int
	totalDataPoints           int
}

func main() {
	var fl flags
	flag.StringVar(&fl.file, "path", "", "metric dir path")

	flag.Parse()
	if err := validate(fl); err != nil {
		fmt.Printf("Invalid flag: %v\n", err)
		os.Exit(1)
	}

	if err := run(fl); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

func validate(fl flags) error {
	if fl.file == "" {
		return errors.New("compressed path not provided")
	}
	return nil
}

func run(fl flags) error {
	s := stats{
		byMetricType: map[metricType]metricTypeStats{
			metricTypeHistogram: {},
			metricTypeGauge:     {},
			metricTypeSum:       {},
			metricTypeSummary:   {},
		},
	}

	if err := filepath.Walk(fl.file, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}

		if !info.IsDir() {
			fmt.Printf("Reading from file %s\n", path)
			return processFile(path, &s)
		}

		return nil
	}); err != nil {
		return fmt.Errorf("failed to iterate over metric data files: %v", err)
	}

	printStats(&s)

	return nil
}

func processFile(path string, s *stats) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		metricDataJSON, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Errorf("failed to read a metric data line: %v", err)
		}

		var metricData MetricData
		if err := json.Unmarshal(metricDataJSON, &metricData); err != nil {
			return fmt.Errorf("failed to unmarshal a metric data line: %v", err)
		}

		for _, resourceMetrics := range metricData.ResourceMetrics {
			for _, scopeMetrics := range resourceMetrics.ScopeMetrics {
				for _, metrics := range scopeMetrics.Metrics {
					recordStats(&metrics, s)
				}
			}
		}
	}

	return nil
}

func recordStats(m *Metrics, s *stats) {
	if len(m.Histogram.DataPoints) > 0 {
		recordStatsByType(metricTypeHistogram, m.Histogram.DataPoints, s)
	} else if len(m.Gauge.DataPoints) > 0 {
		recordStatsByType(metricTypeGauge, m.Gauge.DataPoints, s)
	} else if len(m.Sum.DataPoints) > 0 {
		recordStatsByType(metricTypeSum, m.Sum.DataPoints, s)
	} else if len(m.Summary.DataPoints) > 0 {
		recordStatsByType(metricTypeSummary, m.Summary.DataPoints, s)
	}
}

func recordStatsByType(mt metricType, dps []DataPoint, s *stats) {
	prev := s.byMetricType[mt]

	s.byMetricType[mt] = metricTypeStats{
		total:                     prev.total + 1,
		maxDataPoints:             max(prev.maxDataPoints, len(dps)),
		maxAttributesPerDataPoint: max(prev.maxAttributesPerDataPoint, maxAttributes(dps)),
		totalAttributes:           prev.totalAttributes + totalAttributes(dps),
		totalDataPoints:           prev.totalDataPoints + len(dps),
	}
}

func totalAttributes(dps []DataPoint) int {
	result := 0
	for _, dp := range dps {
		result += len(dp.Attributes)
	}
	return result
}

func maxAttributes(dps []DataPoint) int {
	result := 0
	for _, dp := range dps {
		result = max(result, len(dp.Attributes))
	}
	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func printStats(s *stats) {
	fmt.Println("Histograms")
	printStatsByType(s.byMetricType[metricTypeHistogram])

	fmt.Println("Gauges")
	printStatsByType(s.byMetricType[metricTypeGauge])

	fmt.Println("Sums")
	printStatsByType(s.byMetricType[metricTypeSum])

	fmt.Println("Summaries")
	printStatsByType(s.byMetricType[metricTypeSummary])
}

func printStatsByType(mts metricTypeStats) {
	fmt.Println()

	fmt.Printf("Total: %d\n", mts.total)
	fmt.Printf("Max data points: %d\n", mts.maxDataPoints)
	fmt.Printf("Max attributes: %d\n", mts.maxAttributesPerDataPoint)
	if mts.total > 0 {
		fmt.Printf("Average attributes: %.1f\n", float32(mts.totalAttributes)/float32(mts.totalDataPoints))
	}
}
