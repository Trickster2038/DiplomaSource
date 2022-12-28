# syntax=docker/dockerfile:1

## Build
FROM golang:1.16-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /go-wavedrom

## Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /go-wavedrom /go-wavedrom

EXPOSE 8081

ENTRYPOINT ["/go-wavedrom"]