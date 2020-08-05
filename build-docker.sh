#!/bin/sh
set -e

[ -z "$HERE" ] && HERE="$(dirname "$(readlink -f -- "$0")")"

#     __            __ 
#    / /____  _____/ /_
#   / __/ _ \/ ___/ __/
#  / /_/  __(__  ) /_  
#  \__/\___/____/\__/  
#                      

[ -z "$SKIP_TEST" ] && DOCKER_BUILDKIT=1 make test-unit # test-integration

# install buildx
BUILDX_VERSION="0.4.1"
BUILDX_BIN="${HOME}/.docker/cli-plugins/docker-buildx"
if [ ! -e "${BUILDX_BIN}" ]; then
    curl -LO "https://github.com/docker/buildx/releases/download/v${BUILDX_VERSION}/buildx-v${BUILDX_VERSION}.linux-amd64"
    mkdir -vp "$(dirname "${BUILDX_BIN}")"
    mv -vf "buildx-v${BUILDX_VERSION}.linux-amd64" "${BUILDX_BIN}"
    chmod a+x "${BUILDX_BIN}"
fi

VERSION=$(cat VERSION)
export VERSION

mkdir -p "${HERE}/deploy"
pack() {
    version="v${VERSION:?}"
    [ -n "$1" ] && platform="-$(echo "$1" | sed 's/\//-/g')"
    [ -n "$2" ] && postfix="-${2}"
    packfile="balena-engine-${version}${postfix}${platform}.tar.gz"
    echo "* packing ${packfile}"
    cd "${BUNDLE_DIR}/cross/${1}" || exit 1
    md5sum "balena-engine-${VERSION}.md5"
    tar cvzf "${HERE}/deploy/${packfile}" .
    cd "${HERE}" || exit 1
}

#      __          _ __    __
#     / /_  __  __(_) /___/ /
#    / __ \/ / / / / / __  /
#   / /_/ / /_/ / / / /_/ /
#  /_.___/\__,_/_/_/\__,_/
#

BUNDLE_DIR="${HERE}/bundles"
if [ ! -e "${BUNDLE_DIR}" ]; then
    mkdir -p "${BUNDLE_DIR}"
fi

[ -z "$SKIP_CLEAN" ] && rm -rf "${BUNDLE_DIR:?}/*"

BUILDTAGS="no_btrfs no_cri no_devmapper no_zfs exclude_disk_quota exclude_graphdriver_btrfs exclude_graphdriver_devicemapper exclude_graphdriver_zfs no_buildkit"
export BUILDTAGS
docker buildx bake --progress=plain \
    --set=*.output=type=local,dest="${BUNDLE_DIR}"

pack linux/amd64
pack linux/arm64
pack linux/arm/v7
