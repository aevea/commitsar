FROM golang:1.12.8-alpine as builder
RUN mkdir /app
WORKDIR /app

RUN apk add --no-cache make git gcc musl-dev

# This step is done separately than `COPY . /app/` in order to
# cache dependencies.
COPY go.mod go.sum Makefile /app/
RUN make install_deps

COPY . /app/
RUN make build/docker

FROM alpine:3.10.1
RUN  apk add --no-cache --virtual=.run-deps ca-certificates git &&\
    mkdir /app

WORKDIR /app
COPY --from=builder /app/build/commitsar ./commitsar

RUN ln -s $PWD/commitsar /usr/local/bin

ENTRYPOINT ./commitsar
