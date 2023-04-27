# metric-gen

The tool is used to analyze OTLP JSON metrics files, created by the [file exporter](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/exporter/fileexporter).
It prints out the following statistics:
* Types of the metrics that the files contain, along with the total number per metric type
* Max data point count 
* Max and average data point attribute count

## Usage

You can extract the files created by the `file exporter` in several ways. For example, add a sidecar with a shell to the OpenTelemetry Collector pod and share a volume with it, then execute `kubectl copy` to copy the files on your local machine.
To analyze the files, run the following command:
```bash
go run ./ -path=path_to_dir_containing_metric_files
```

