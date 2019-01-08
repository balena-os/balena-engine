#!/bin/sh

set -eo pipefail

tag="v17.12.0"
tag=$(echo "$tag" | sed 's|+|.|g')

machine=$(uname -m)

case "$machine" in
	"armv5"*)
		arch="armv5"
		;;
	"armv6"*)
		arch="armv6"
		;;
	"armv7"*)
		arch="armv7"
		;;
	"armv8"*)
		arch="aarch64"
		;;
	"aarch64"*)
		arch="aarch64"
		;;
	"i386")
		arch="i386"
		;;
	"i686")
		arch="i386"
		;;
	"x86_64")
		arch="x86_64"
		;;
	*)
		echo "Unknown machine type: $machine"
		exit 1
esac

url="https://github.com/balena-os/balena-engine/releases/download/${tag}/balena-engine-${tag}-${arch}.tar.gz"

curl -sL "$url" | sudo tar xzv -C /usr/local/bin --strip-components=1

cat <<EOF


   Installation successful!
     __          __                 ______            _
    / /_  ____ _/ /__  ____  ____ _/ ____/___  ____ _(_)___  ___
   / __ \\/ __ \`/ / _ \\/ __ \\/ __ \`/ __/ / __ \\/ __ \`/ / __ \\/ _ \\
  / /_/ / /_/ / /  __/ / / / /_/ / /___/ / / / /_/ / / / / /  __/
 /_.___/\\__,_/_/\\___/_/ /_/\\__,_/_____/_/ /_/\\__, /_/_/ /_/\\___/
                                            /____/
  the container engine for the IoT
EOF
