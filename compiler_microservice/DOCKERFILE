# syntax=docker/dockerfile:1

## Build
FROM golang:1.16-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /go-verilog-compiler

## Deploy
FROM ubuntu:latest

RUN apt update && apt install -y iverilog && rm -rf /var/lib/apt/lists/*

WORKDIR /

COPY --from=build /go-verilog-compiler /go-verilog-compiler

EXPOSE 8080

USER root

ENTRYPOINT ["/go-verilog-compiler"]