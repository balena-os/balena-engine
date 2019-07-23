#!/bin/sh
set -e

# first arg is `-f` or `--some-option`
if [ "${1#-}" != "$1" ]; then
	set -- balena-engine "$@"
fi

# if our command is a valid Docker subcommand, let's invoke it through Docker instead
# (this allows for "docker run docker ps", etc)
if balena-engine help "$1" > /dev/null 2>&1; then
	set -- balena-engine "$@"
fi

# if we have "--link some-docker:docker" and not DOCKER_HOST, let's set DOCKER_HOST automatically
if [ -z "$DOCKER_HOST" -a "$DOCKER_PORT_2375_TCP" ]; then
	export DOCKER_HOST='tcp://balena:2375'
fi

if [ "$1" = 'balena-engine-daemon' -o "$1" = 'balenad' ]; then
	cat >&2 <<-'EOW'

		📎 Hey there!  It looks like you're trying to run a balenaEngine daemon.

		   You probably should use the "beind" image variant instead, something like:

		     docker run --privileged --name some-overlay-balena -d balena/engine:beind --storage-driver=overlay2

	EOW
	sleep 3
fi

exec "$@"
