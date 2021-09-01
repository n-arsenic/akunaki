FROM golang:1.16-alpine

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

ENV GOOS=linux
ENV GOARCH=amd64

RUN go build -o /go/bin/grpc-service

ENTRYPOINT ["/go/bin/grpc-service"]
CMD ["client"]