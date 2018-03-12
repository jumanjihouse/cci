FROM alpine:3.7

ENV DOCKER_VERSION 18.02.0-r0
ENV COMPOSE_VERSION 1.19.0
ENV GOSS_VERSION 0.3.5
ENV BATS_VERSION 0.4.0

RUN apk add --no-cache \
      'bash>=4.4.12-r2' \
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
      xz \
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

# Install latest statically-linked version of
# https://github.com/koalaman/shellcheck
RUN cd /tmp &&                                                                 \
    tarball=shellcheck-latest.linux.x86_64.tar.xz &&                           \
    curl -L -ssL -O https://storage.googleapis.com/shellcheck/${tarball} &&    \
    tar xvJf ${tarball} &&                                                     \
    cp /tmp/shellcheck-latest/shellcheck /usr/local/bin/ &&                    \
    rm -fr /tmp/shellcheck* &&                                                 \
    :

COPY goss.yaml /
COPY microbadger /usr/local/bin/
