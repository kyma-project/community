FROM golang:1.20 as builder

WORKDIR /workspace
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY main.go main.go


RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o metrics-gen main.go

FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /workspace/metrics-gen .

USER 65532:65532

ENTRYPOINT ["/metrics-gen"]
