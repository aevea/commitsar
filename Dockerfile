FROM golang:1.13.0-alpine as builder
RUN mkdir /app
WORKDIR /app

RUN apk add --no-cache make git gcc musl-dev

# This step is done separately than `COPY . /app/` in order to
# cache dependencies.
COPY go.mod go.sum Makefile /app/
RUN make install_deps

COPY . /app/
RUN make build/docker

FROM alpine:3.10.2
ENV FLAGS = "-v"

LABEL repository="https://github.com/commitsar-app/commitsar"
LABEL homepage="https://github.com/commitsar-app/commitsar"
LABEL maintainer="Simon Prochazka <simon@fallion.net>"

LABEL com.github.actions.name="Commitsar Action"
LABEL com.github.actions.description="Check commit message compliance with conventional commits"
LABEL com.github.actions.icon="code"
LABEL com.github.actions.color="blue"

RUN  apk add --no-cache --virtual=.run-deps ca-certificates git &&\
    mkdir /app

WORKDIR /app
COPY --from=builder /app/build/commitsar ./commitsar

RUN ln -s $PWD/commitsar /usr/local/bin

CMD ["commitsar"]
