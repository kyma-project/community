# metric-gen

The tool is used to generate dummy OTLP metrics and send them to a provided grpc OTLP endpoint (for example, an OpenTelemetry collector with a configured `otlp` receiver).

## Usage

### Local

```bash
go run ./ -host=telemetry-otlp-metrics -port=4317
```

### Kubernetes

You can also dockerize the tool and run it in a Kubernetes Deployment. This way, you can increase the produced load by running multiple replicas in parallel.

```bash
export IMG="{your_repository/your_image:your_tag}"
docker build -t $IMG  .
docker push $IMG
kubectl create deployment metric-gen --image=$IMG --replicas=3
```
