#!/bin/sh
set -e

# no arguments passed
# or first arg is `-f` or `--some-option`
if [ "$#" -eq 0 ] || [ "${1#-}" != "$1" ]; then
	# add our default arguments
	set -- balena-engine-daemon \
		--host=unix:///var/run/balena-engine.sock \
		--host=tcp://0.0.0.0:2375 \
		"$@"
fi

if [ "$1" = 'balena-engine-daemon' -o "$1" = 'balenad' ]; then
	# if we're running Docker, let's pipe through dind
	set -- "$(which beind)" "$@"

	# explicitly remove Docker's default PID file to ensure that it can start properly if it was stopped uncleanly (and thus didn't clean up the PID file)
	find /run /var/run -iname 'balena*.pid' -delete
fi

exec "$@"
