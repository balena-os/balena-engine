#!/bin/sh

tag="17.06-rc5+testtag"

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
		arch="amd64"
		;;
	*)
		echo "Unknown machine type: $machine"
		exit 1
esac

curl -sSL "https://github.com/resin-os/balena/releases/download/${tag}/balena-${tag}-${arch}.tar.gz" | tar tz
