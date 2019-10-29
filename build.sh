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

export AUTO_GOPATH=1
export GOMAXPROCS=1
export DOCKER_LDFLAGS="-s"
export VERSION="$version"
export DOCKER_BUILDTAGS='exclude_graphdriver_btrfs exclude_graphdirver_zfs exclude_graphdriver_devicemapper no_btrfs'
./hack/make.sh binary-balena

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
