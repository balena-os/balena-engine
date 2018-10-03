#!/bin/sh

tag="17.06-rev1"
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
 _           _
| |__   __ _| | ___ _ __   __ _
| '_ \\ / _\` | |/ _ \\ '_ \ / _\` |
| |_) | (_| | |  __/ | | | (_| |
|_.__/ \__,_|_|\___|_| |_|\__,_|

the container engine for the IoT
EOF
