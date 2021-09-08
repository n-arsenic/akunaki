FROM golang:1.16-alpine

WORKDIR /go/src/akunaki

COPY . .
RUN go mod download

RUN go build -o /go/bin/akunaki-server

ENTRYPOINT ["/go/bin/akunaki-server"]