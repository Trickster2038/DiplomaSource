# syntax=docker/dockerfile:1

FROM golang:1.16-buster 
# FROM golang:1.9.2 
ADD . /go/src/stats
WORKDIR /go/src/stats
RUN go get stats
RUN go install

EXPOSE 8082
EXPOSE 8085

ENTRYPOINT ["/go/bin/stats"]