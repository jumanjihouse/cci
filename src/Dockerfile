FROM alpine:3.11.5

ENV GOSS_VERSION 0.3.9
ENV BATS_VERSION 1.1.0
ENV SHFMT_VERSION v3.0.1

RUN apk add --no-cache \
      'bash>=4.4.19-r1' \
      ca-certificates \
      coreutils \
      curl \
      file \
      gcc \
      git \
      gzip \
      jq \
      libffi-dev \
      make \
      musl-dev \
      openssh-client \
      openssl-dev \
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

RUN apk add --no-cache \
      -X http://dl-cdn.alpinelinux.org/alpine/edge/community \
      docker \
      && \
      :

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

COPY microbadger /usr/local/bin/

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

# BUILD_DATE changes on every build, so it goes last to avoid busting the cache.
ARG BUILD_DATE
ENV BUILD_DATE ${BUILD_DATE}

ARG TAG
ENV TAG ${TAG}

ARG VCS_REF
ENV VCS_REF ${VCS_REF}

ARG CI_BUILD_URL
ENV CI_BUILD_URL ${CI_BUILD_URL}

LABEL \
    io.github.jumanjiman.ci-build-url=${CI_BUILD_URL} \
    io.github.jumanjiman.version=${TAG} \
    io.github.jumanjiman.build-date=${BUILD_DATE} \
    io.github.jumanjiman.vcs-ref=${VCS_REF} \
    io.github.jumanjiman.license="MIT" \
    io.github.jumanjiman.docker.dockerfile="/src/Dockerfile" \
    io.github.jumanjiman.vcs-type="Git" \
    io.github.jumanjiman.vcs-url="https://github.com/jumanjihouse/cci.git"
