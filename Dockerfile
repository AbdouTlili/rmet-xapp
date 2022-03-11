#syntax=docker/dockerfile:1

FROM golang:1.16-buster AS build
ARG MAKE_TARGET=build


COPY Makefile go.mod go.sum /go/src/github.com/AbdouTlili/rmet-xapp/
COPY main.go /go/src/github.com/AbdouTlili/rmet-xapp/

RUN cd /go/src/github.com/AbdouTlili/rmet-xapp/ && make ${MAKE_TARGET}

ENTRYPOINT [ "/go/src/github.com/AbdouTlili/rmet-xapp/build/out/rmet-xapp" ]

#FROM gcr.io/distroless/base-debian10