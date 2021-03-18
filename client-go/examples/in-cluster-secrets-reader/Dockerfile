FROM golang:1.15-alpine AS builder
WORKDIR /go/src/github.com/smvfal/kubernetes/client-go/examples/in-cluster-secrets-reader
COPY . .
RUN go build -o secretreader

FROM alpine
WORKDIR /root/
COPY --from=builder /go/src/github.com/smvfal/kubernetes/client-go/examples/in-cluster-secrets-reader/secretreader .
CMD ["./secretreader"]
