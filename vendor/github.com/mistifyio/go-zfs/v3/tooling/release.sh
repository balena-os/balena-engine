#!/usr/bin/env bash

set -euo pipefail

tag=$1
if [[ -z ${tag:-} ]] || (($# != 1)); then
	echo 'tag must be provided as the only cli arg' >&2
	exit 1
fi
if ! git tag | grep -q "^$tag$"; then
	echo "$tag is not a git tag" >&2
	exit 1
fi

ghnotes=$(
	gh api \
		--method POST \
		--header "Accept: application/vnd.github+json" \
		--header "X-GitHub-Api-Version: 2022-11-28" \
		--field "tag_name=$tag" \
		--jq .body \
		repos/{owner}/{repo}/releases/generate-notes
)

range=$(awk -F/ 'END {print $NF}' <<<"$ghnotes")
fullchangelog=$(tail -n1 <<<"$ghnotes")
ghnotes=$(head -n -2 <<<"$ghnotes)")
diffstat=$(git diff --stat "$range")
shortlog=$(git shortlog -w0 --no-merges "$range")

gh release create "$tag" --notes "
$ghnotes

$diffstat

$shortlog

$fullchangelog
"
