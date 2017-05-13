FROM alpine:3.5

ENV DOCKER_VERSION 1.12.6-r0
ENV COMPOSE_VERSION 1.12.0
ENV GOSS_VERSION 0.3.2
ENV BATS_VERSION 0.4.0

RUN apk add --no-cache \
      bash \
      ca-certificates \
      curl \
      docker=${DOCKER_VERSION} \
      file \
      git \
      gzip \
      make \
      openssh-client \
      py2-pip \
      tar \
      && \
    pip install docker-compose==${COMPOSE_VERSION} && \
    cd /tmp && \
    curl -sSL -o bats_v${BATS_VERSION}.tar.gz https://github.com/sstephenson/bats/archive/v${BATS_VERSION}.tar.gz && \
    tar -xf /tmp/bats_v${BATS_VERSION}.tar.gz && \
    bats-${BATS_VERSION}/install.sh /usr/local && \
    rm -fr /tmp/bats* && \
    curl -sSL -o /usr/bin/goss https://github.com/aelsabbahy/goss/releases/download/v${GOSS_VERSION}/goss-linux-amd64 && \
    chmod 0755 /usr/bin/goss && \
    :

COPY goss.yaml /
