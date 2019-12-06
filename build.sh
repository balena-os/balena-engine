#!/bin/bash

set -ex

arch=${BUILD_ARCH:-$(uname -m)}
dest="${BUILD_DEST:-balena-engine}"
debug="${BUILD_DEBUG:-}"
static="${BUILD_STATIC:-}"
version=${VERSION:-$(git describe --tags --always)}

BUILDTAGS="journald" # engine features
BUILDTAGS="$BUILDTAGS exclude_graphdriver_btrfs exclude_graphdirver_zfs exclude_graphdriver_devicemapper" # engine graphdrivers
BUILDTAGS="$BUILDTAGS no_btrfs no_cri no_devicemapper" # containerd
BUILDTAGS="$BUILDTAGS seccomp apparmor" # runc

# allow setting additional go build tags
extra_buildtags="${BUILD_EXTRA_BUILDTAGS:-}"
BUILDTAGS="$BUILDTAGS $extra_buildtags"

BUILDTIME=$(date -u -d "@${SOURCE_DATE_EPOCH:-$(date +%s)}" --rfc-3339 ns 2> /dev/null | sed -e 's/ /T/')

LDFLAGS=""
if [ -z "$debug" ]; then
    LDFLAGS="$LDFLAGS -w" # strip DWARF debug data
fi
if [ -n "$static" ]; then
    LDFLAGS="$LDFLAGS -extldflags \"-fno-PIC -static\""
    BUILDTAGS="$BUILDTAGS static_build"
    IAMSTATIC="true"
fi

BUILDFLAGS=(
    -tags "netgo osusergo $BUILDTAGS"
    -installsuffix netgo
    -buildmode pie
)

COMMIT_ENGINE=${COMMIT_ENGINE:-$(git rev-parse --short HEAD)}
VERSION_ENGINE=${VERSION_ENGINE:-$version}
COMMIT_CLI=${COMMIT_CLI:-$COMMIT_ENGINE}
VERSION_CLI=${VERSION_CLI:-$VERSION_ENGINE}
COMMIT_CONTAINERD=${COMMIT_CONTAINERD:-$COMMIT_ENGINE}
VERSION_CONTAINERD=${VERSION_CONTAINERD:-$VERSION_ENGINE}
COMMIT_RUNC=${COMMIT_RUNC:-$COMMIT_ENGINE}
VERSION_RUNC=${VERSION_RUNC:-$VERSION_ENGINE}

# run the build
(
    rm -rf "$dest/balena-engine" || true
    mkdir -p "$dest/balena-engine"

    export GOMAXPROCS=1
    # export CGO_ENABLED

# allow setting additional go build tags
extra_buildtags="${BUILD_EXTRA_BUILDTAGS:-}"
BUILDTAGS="$BUILDTAGS $extra_buildtags"

# hash the binary
(
    cd "$dest/balena-engine" || exit 1
    sha256sum balena-engine > balena-engine.sha256
)

cp --no-dereference "$src"/balena-engine* "$dst/"

# pack the release artifacts
(
    archive="balena-engine-$version-$arch.tar.gz"
    tar -czvf "/tmp/$archive" -C "$dest" ./balena-engine
    mv "/tmp/$archive" "$dest/$archive"
)
