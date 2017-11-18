FROM alpine:3.6

ENV DOCKER_VERSION 17.07.0-r0
ENV COMPOSE_VERSION 1.17.1
ENV GOSS_VERSION 0.3.5
ENV BATS_VERSION 0.4.0

RUN apk add --no-cache \
      bash \
      ca-certificates \
      coreutils \
      curl \
      file \
      git \
      gzip \
      make \
      openssh-client \
      py2-pip \
      tar \
      && \
      :

RUN apk add --no-cache \
      -X http://dl-cdn.alpinelinux.org/alpine/edge/community \
      docker=${DOCKER_VERSION} \
      && \
      :

RUN pip install -Iv --compile --no-cache-dir \
      docker-compose==${COMPOSE_VERSION} \
      pypi-cli \
      && \
      :

RUN cd /tmp && \
    curl -sSL -o bats_v${BATS_VERSION}.tar.gz https://github.com/sstephenson/bats/archive/v${BATS_VERSION}.tar.gz && \
    tar -xf /tmp/bats_v${BATS_VERSION}.tar.gz && \
    bats-${BATS_VERSION}/install.sh /usr/local && \
    rm -fr /tmp/bats* && \
    curl -sSL -o /usr/bin/goss https://github.com/aelsabbahy/goss/releases/download/v${GOSS_VERSION}/goss-linux-amd64 && \
    chmod 0755 /usr/bin/goss && \
    :

COPY goss.yaml /
COPY microbadger /usr/local/bin/
