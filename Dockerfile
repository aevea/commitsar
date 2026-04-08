FROM golang:1.26.2-alpine AS builder

RUN apk add --no-cache git make gcc musl-dev

WORKDIR /app

# This step is done separately than `COPY . /app/` in order to
# cache dependencies.
COPY go.mod go.sum Makefile /app/
RUN go mod download

COPY . /app/
# Fix git ownership issue
RUN git config --global --add safe.directory /app
RUN make build/docker

FROM alpine:3.23

LABEL repository="https://github.com/aevea/commitsar"
LABEL homepage="https://github.com/aevea/commitsar"
LABEL maintainer="Simon Prochazka <simon@fallion.net>"

LABEL com.github.actions.name="Commitsar Action"
LABEL com.github.actions.description="Check commit message compliance with conventional commits"
LABEL com.github.actions.icon="code"
LABEL com.github.actions.color="blue"

RUN apk add --no-cache ca-certificates git

WORKDIR /app
COPY --from=builder /app/build/commitsar ./commitsar
COPY entrypoint.sh /entrypoint.sh

RUN ln -s $PWD/commitsar /usr/local/bin && \
    chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
CMD ["commitsar"]
