FROM golang:1.15-alpine AS builder
WORKDIR /go/src/github.com/smvfal/kubernetes/client-go/examples/in-cluster-top/
COPY . .
RUN go build -o top

FROM alpine
WORKDIR /root/
COPY --from=builder /go/src/github.com/smvfal/kubernetes/client-go/examples/in-cluster-top/top .
CMD ["./top"]
