FROM golang:1.12.7-alpine as builder
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
RUN  apk add --no-cache --virtual=.run-deps ca-certificates &&\
    mkdir /app

WORKDIR /app
COPY --from=builder /app/build/commitsar ./commitsar

USER nobody
ENTRYPOINT ./commitsar