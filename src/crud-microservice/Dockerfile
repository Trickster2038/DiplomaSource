# syntax=docker/dockerfile:1

FROM golang:1.16-buster 
# FROM golang:1.9.2 
ADD . /go/src/crud
WORKDIR /go/src/crud
RUN go get crud
RUN go install

EXPOSE 8082
EXPOSE 8089

ENTRYPOINT ["/go/bin/crud"]