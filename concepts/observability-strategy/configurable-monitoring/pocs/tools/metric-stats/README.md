# metric-gen

The tool is used to analyze OTLP JSON metrics files, created by the [file exporter](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/exporter/fileexporter).
It prints out the following statistics:
* types of the metrics that the files contain, along with the total number per metric type
* max data point count 
* max and average data point attribute count

## Usage

One should extract the files created by the `file exporter`. There are multiple ways to do it. For example, add a sidecar with a shell to the OpenTelemetry Collector pod and share a volume with it, then execute `kubectl copy` to copy the files on your local machine.
Then run the following command to analyze them:
```bash
go run ./ -path=path_to_dir_containing_metric_files
```

