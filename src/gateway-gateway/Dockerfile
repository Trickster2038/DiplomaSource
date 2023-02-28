# syntax=docker/dockerfile:1

FROM golang:1.16-buster 
# FROM golang:1.9.2 
ADD . /go/src/gateway
WORKDIR /go/src/gateway
RUN go get gateway
RUN go install

EXPOSE 8082
EXPOSE 8084

ENTRYPOINT ["/go/bin/gateway"]