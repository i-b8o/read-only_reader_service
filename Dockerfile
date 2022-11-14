FROM golang:latest

RUN go version

ENV GOPATH=/
ENV CONFIG-PATH=/go/configs/config.local.yaml

COPY ./ ./

RUN go mod tidy
RUN go build -o read-only-service ./cmd/main.go

CMD ["./read-only-service"]
