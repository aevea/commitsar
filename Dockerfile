FROM debian:13-slim as builder
RUN mkdir /app
WORKDIR /app

# Install mise and dependencies
RUN apt-get update && \
    apt-get -y --no-install-recommends install \
        sudo curl git ca-certificates build-essential make gcc && \
    rm -rf /var/lib/apt/lists/*

# Set up mise environment variables (before installation)
SHELL ["/bin/bash", "-o", "pipefail", "-c"]
ENV MISE_DATA_DIR="/mise"
ENV MISE_CONFIG_DIR="/mise"
ENV MISE_CACHE_DIR="/mise/cache"
ENV MISE_INSTALL_PATH="/usr/local/bin/mise"
ENV PATH="/mise/shims:/usr/local/bin:$PATH"

# Install mise
RUN curl https://mise.run | sh

# Copy mise config
COPY .mise.toml /app/

# Trust the mise config file
RUN mise trust

# Install Go via mise
RUN mise install

# This step is done separately than `COPY . /app/` in order to
# cache dependencies.
COPY go.mod go.sum Makefile /app/
RUN mise exec -- go mod download

COPY . /app/
# Fix git ownership issue
RUN git config --global --add safe.directory /app
RUN mise exec -- make build/docker

FROM debian:13-slim

LABEL repository="https://github.com/aevea/commitsar"
LABEL homepage="https://github.com/aevea/commitsar"
LABEL maintainer="Simon Prochazka <simon@fallion.net>"

LABEL com.github.actions.name="Commitsar Action"
LABEL com.github.actions.description="Check commit message compliance with conventional commits"
LABEL com.github.actions.icon="code"
LABEL com.github.actions.color="blue"

RUN apt-get update && \
    apt-get -y --no-install-recommends install ca-certificates git && \
    rm -rf /var/lib/apt/lists/* && \
    mkdir /app && \
    git config --system --add safe.directory '*'

WORKDIR /app
COPY --from=builder /app/build/commitsar ./commitsar
COPY entrypoint.sh /entrypoint.sh

RUN ln -s $PWD/commitsar /usr/local/bin && \
    chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
CMD ["commitsar"]
