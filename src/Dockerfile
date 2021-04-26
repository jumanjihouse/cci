FROM alpine:3.13.5

ENV GOSS_VERSION 0.3.11
ENV BATS_VERSION 1.2.0
ENV SHFMT_VERSION v3.1.1

RUN apk add --no-cache \
      -X http://dl-cdn.alpinelinux.org/alpine/edge/main \
      -X http://dl-cdn.alpinelinux.org/alpine/edge/community \
      'bash>=4.4.19-r1' \
      bind-tools \
      ca-certificates \
      coreutils \
      curl \
      'docker>=19.03.13-r0' \
      file \
      gcc \
      git \
      go \
      gzip \
      jq \
      libffi-dev \
      make \
      musl-dev \
      npm \
      openssh-client \
      openssl-dev \
      py3-pip \
      python3 \
      python3-dev \
      ruby \
      ruby-bundler \
      ruby-dev \
      ruby-etc \
      ruby-irb \
      ruby-json \
      tar \
      xz \
      yaml-dev \
      && \
      :

RUN ln -s /usr/bin/python3 /usr/bin/python

COPY requirements.txt /cci/
RUN pip3 install -Iv --compile --no-cache-dir -r /cci/requirements.txt

RUN cd /tmp && \
    curl -sSL -o bats_v${BATS_VERSION}.tar.gz https://github.com/bats-core/bats-core/archive/v${BATS_VERSION}.tar.gz && \
    tar -xf /tmp/bats_v${BATS_VERSION}.tar.gz && \
    bats-core-${BATS_VERSION}/install.sh /usr/local && \
    rm -fr /tmp/bats* && \
    curl -sSL -o /usr/bin/goss https://github.com/aelsabbahy/goss/releases/download/v${GOSS_VERSION}/goss-linux-amd64 && \
    chmod 0755 /usr/bin/goss && \
    :

# Install latest statically-linked version of
# https://github.com/koalaman/shellcheck
RUN cd /tmp &&                                                                 \
    version="latest" &&                                                        \
    tarball=shellcheck-latest.linux.x86_64.tar.xz &&                           \
    url=https://github.com/koalaman/shellcheck/releases/download/${version} && \
    curl -L -ssL -O "${url}/${tarball}" &&                                     \
    tar xvJf ${tarball} &&                                                     \
    cp /tmp/shellcheck-latest/shellcheck /usr/local/bin/ &&                    \
    rm -fr /tmp/shellcheck* &&                                                 \
    :

# Install statically-linked version of shfmt from
# https://github.com/mvdan/sh
RUN curl -L -ssL -o /usr/local/bin/shfmt "https://github.com/mvdan/sh/releases/download/${SHFMT_VERSION}/shfmt_${SHFMT_VERSION}_linux_amd64" && \
    chmod 0755 /usr/local/bin/shfmt

# Cache a bunch of pre-commit hook environments
# to reduce runtimes on circleci
# for tests that use this image.
# The pre-commit framework uses env vars as coded at
# https://github.com/pre-commit/pre-commit/blob/master/pre_commit/store.py
ENV XDG_CACHE_HOME /root/.local/
COPY .pre-commit-config.yaml /root/empty-repo/
COPY bootstrap /usr/local/sbin/
RUN cd /root/empty-repo &&                                                    \
    git config --global user.email "you@example.com" &&                       \
    git config --global user.name "Your Name" &&                              \
    git init &&                                                               \
    git commit --allow-empty -m 'initial git repo' &&                         \
    /usr/local/sbin/bootstrap

COPY python-path.sh /etc/profile.d/

# Keep this near last to avoid busting the cache when adding tests.
COPY goss.yaml /
