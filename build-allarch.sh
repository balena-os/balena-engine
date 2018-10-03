#!/bin/sh

set -o errexit

echo "Building x86_64.."
docker build -f Dockerfile.build.x86_64 -t balena-engine-build-x86_64 .
docker run --rm -v "$(pwd):/balena-engine" balena-engine-build-x86_64 ./build.sh

echo "Building i386.."
docker build -f Dockerfile.build.i386 -t balena-engine-build-i386 .
docker run --rm -v "$(pwd):/balena-engine" balena-engine-build-i386 ./build.sh

echo "Building armv5.."
docker build -f Dockerfile.build.arm -t balena-engine-build-arm .
docker run --rm -e GOARM=5 -v "$(pwd):/balena-engine" balena-engine-build-arm /bin/sh build.sh

echo "Building armv6.."
docker run --rm -e GOARM=6 -v "$(pwd):/balena-engine" balena-engine-build-arm /bin/sh build.sh

echo "Building armv7.."
docker run --rm -e GOARM=7 -v "$(pwd):/balena-engine" balena-engine-build-arm /bin/sh build.sh

echo "Building aarch64.."
docker build -f Dockerfile.build.aarch64 -t balena-engine-build-aarch64 .
docker run --rm -v "$(pwd):/balena-engine" balena-engine-build-aarch64 /bin/sh build.sh
