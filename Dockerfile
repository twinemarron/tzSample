FROM golang:latest
RUN mkdir /go/src/tzSample
WORKDIR /go/src/tzSample
ADD . /go/src/tzSample
