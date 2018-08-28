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
dst="balena"

rm -rf "$dst"
mkdir "$dst"

cp -L "$src/balena" "$dst/balena"

ln -s balena "$dst/balenad"
ln -s balena "$dst/balena-containerd"
ln -s balena "$dst/balena-containerd-ctr"
ln -s balena "$dst/balena-containerd-shim"
ln -s balena "$dst/balena-proxy"
ln -s balena "$dst/balena-runc"

tar czfv "balena-$version-$arch.tar.gz" "$dst"
