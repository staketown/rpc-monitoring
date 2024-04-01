FROM golang:1.20 AS exporter

ENV GOBIN=/go/bin
ENV GOPATH=/go
ENV CGO_ENABLED=0
ENV GOOS=linux

WORKDIR /exporter
COPY *.go go.sum go.mod ./
RUN go build -o /rpc_exporter .

FROM debian:buster-slim

RUN useradd -ms /bin/bash exporter && chown -R exporter /usr

EXPOSE 9300

COPY --from=exporter rpc_exporter /usr/bin/rpc_exporter

USER exporter