#syntax=docker/dockerfile:1

FROM golang:1.16-buster AS build
ARG MAKE_TARGET=build


COPY Makefile go.mod go.sum /go/src/github.com/AbdouTlili/rmet-xapp/
COPY main.go /go/src/github.com/AbdouTlili/rmet-xapp/

RUN cd /go/src/github.com/AbdouTlili/rmet-xapp/ && make ${MAKE_TARGET}

## build stage two 

FROM gcr.io/distroless/base-debian10

COPY --from=build  /go/src/github.com/AbdouTlili/rmet-xapp/build/out/rmet-xapp /rmet-xapp

#TODO add the ports here later

ENTRYPOINT [ "/rmet-xapp" ]

