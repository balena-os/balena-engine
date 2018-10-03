#!/bin/sh

set -o errexit

case "$(go env GOARCH)" in
	"arm")
		arch="armv$(go env GOARM)"
		;;
	"arm64")
		arch="aarch64"
		;;
	"386")
		arch="i386"
		;;
	"amd64")
		arch="x86_64"
		;;
esac

version=$(git describe --tags --always)

AUTO_GOPATH=1 GOMAXPROCS=1 DOCKER_LDFLAGS="-s" VERSION="$version" ./hack/make.sh binary-balena

src="bundles/latest/binary-balena"
dst="balena-engine"

rm -rf "$dst"
mkdir "$dst"

cp -L "$src/balena-engine" "$dst/balena-engine"

ln -s balena-engine "$dst/balena-engine-daemon"
ln -s balena-engine "$dst/balena-engine-containerd"
ln -s balena-engine "$dst/balena-engine-containerd-ctr"
ln -s balena-engine "$dst/balena-engine-containerd-shim"
ln -s balena-engine "$dst/balena-engine-proxy"
ln -s balena-engine "$dst/balena-engine-runc"

tar czfv "balena-engine-$version-$arch.tar.gz" "$dst"
