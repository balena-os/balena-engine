#!/bin/bash

set -ex

arch=${BUILD_ARCH:-$(uname -m)}
version=${VERSION:-$(git describe --tags --always)}
target=dynbinary-balena

if [ -n "${STATIC}" ]; then
    target=binary-balena
fi

export GOMAXPROCS=1
export VERSION="$version"
export DOCKER_LDFLAGS="-s" # strip resulting binary
./hack/make.sh "$target"

src="bundles/latest/$target"
dst="balena-engine"

rm -rf "$dst" || true
mkdir "$dst"

cp --no-dereference "$src"/balena-engine* "$dst/"

tar czfv "balena-engine-$version-$arch.tar.gz" "$dst"
