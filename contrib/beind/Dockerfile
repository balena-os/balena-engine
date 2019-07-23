FROM docker.io/library/alpine:3.8 as balena

ARG BALENA_VERSION=v18.9.7

RUN apk add --no-cache \
		ca-certificates

# set up nsswitch.conf for Go's "netgo" implementation (which Docker explicitly uses)
# - https://github.com/docker/docker-ce/blob/v17.09.0-ce/components/engine/hack/make.sh#L149
# - https://github.com/golang/go/blob/go1.9.1/src/net/conf.go#L194-L275
# - docker run --rm debian:stretch grep '^hosts:' /etc/nsswitch.conf
RUN [ ! -e /etc/nsswitch.conf ] && echo 'hosts: files dns' > /etc/nsswitch.conf

RUN set -eux; \
	\
	apkArch="$(apk --print-arch)"; \
	case "$apkArch" in \
		x86_64) dlArch='x86_64' ;; \
		armhf) dlArch='armv7hf' ;; \
		aarch64) dlArch='aarch64' ;; \
		*) echo >&2 "error: unsupported architecture ($apkArch)"; exit 1 ;;\
	esac; \
	\
	if ! wget -O balena-engine.tgz "https://github.com/balena-os/balena-engine/releases/download/${BALENA_VERSION}/balena-engine-${BALENA_VERSION}-${dlArch}.tar.gz"; then \
		echo >&2 "error: failed to download balenaEngine ${BALENA_VERSION} for ${dlArch}"; \
		exit 1; \
	fi; \
	\
	tar --extract \
		--file balena-engine.tgz \
		--strip-components 1 \
		--directory /usr/local/bin/ \
	; \
	rm balena-engine.tgz; \
	\
	balena-engine-daemon --version; \
	balena-engine --version

COPY modprobe.sh /usr/local/bin/modprobe
COPY balena-entrypoint.sh /usr/local/bin/

ENTRYPOINT ["/usr/local/bin/balena-entrypoint.sh"]
CMD []


FROM balena as beind

# https://github.com/docker/docker/blob/master/project/PACKAGERS.md#runtime-dependencies
RUN set -eux; \
	apk add --no-cache \
		e2fsprogs \
		e2fsprogs-extra \
		iptables \
		tini \
		xz \
# pigz: https://github.com/moby/moby/pull/35697 (faster gzip implementation)
		pigz \
	; \
	ln -sfv "$(which tini)" /usr/local/bin/balena-engine-init

# TODO aufs-tools

# set up subuid/subgid so that "--userns-remap=default" works out-of-the-box
RUN set -x \
	&& addgroup -S dockremap \
	&& adduser -S -G dockremap dockremap \
	&& echo 'dockremap:165536:65536' >> /etc/subuid \
	&& echo 'dockremap:165536:65536' >> /etc/subgid

# https://github.com/docker/docker/tree/master/hack/dind
# ENV DIND_COMMIT=37498f009d8bf25fbb6199e8ccd34bed84f2874b
ENV DIND_COMMIT=v18.9.7

RUN set -eux; \
	wget -O /usr/local/bin/beind "https://raw.githubusercontent.com/balena-os/balena-engine/${DIND_COMMIT}/hack/dind"; \
	chmod +x /usr/local/bin/beind

COPY balenad-entrypoint.sh /usr/local/bin/

VOLUME /var/lib/balena-engine
EXPOSE 2375

ENTRYPOINT ["/usr/local/bin/balenad-entrypoint.sh"]
CMD []
