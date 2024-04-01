FROM golang:1.22 AS exporter

ENV GOBIN=/go/bin
ENV GOPATH=/go
ENV CGO_ENABLED=0
ENV GOOS=linux

WORKDIR /exporter
COPY *.go go.sum go.mod ./
RUN go build -o /rpc_exporter .

FROM debian:buster-slim

RUN useradd -ms /bin/bash exporter && chown -R exporter /usr
RUN apt-get update && apt-get install -y ca-certificates

EXPOSE 9300

COPY --from=exporter rpc_exporter /usr/bin/rpc_exporter

USER exporter