FROM golang:1.15-alpine AS builder
WORKDIR /go/src/github.com/smvfal/kubernetes/client-go/examples/in-cluster-logs-reader
COPY . .
RUN go build -o logs-reader

FROM alpine
WORKDIR /root/
COPY --from=builder /go/src/github.com/smvfal/kubernetes/client-go/examples/in-cluster-logs-reader/logs-reader .
CMD ["./logs-reader"]
