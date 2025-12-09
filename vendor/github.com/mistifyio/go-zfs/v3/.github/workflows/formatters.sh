#!/usr/bin/env nix-shell
#!nix-shell -i bash ../../shell.nix
#shellcheck shell=bash

set -eux

failed=0

if ! git ls-files '*.md' '*.yaml' '*.yml' | xargs prettier --list-different --write; then
	failed=1
fi

if ! shfmt -f . | xargs shfmt -l -d; then
	failed=1
fi

if ! nixfmt shell.nix; then
	failed=1
fi

if ! rufo Vagrantfile; then
	failed=1
fi

if ! git diff | (! grep .); then
	failed=1
fi

exit "$failed"
