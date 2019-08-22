#!/bin/sh
#
# Add upstream as a remote: 
# $ git remote add docker https://github.com/docker/engine
# $ git fetch docker
#
# To get the list of changes since v18.9.8 run:
# $ ./generate_patchlist.sh v18.9.8..HEAD
#

set -e

UPSTREAM=${UPSTREAM:-"docker/master"}
HEAD=${1:-master}
EXPORT_PATCHES=${EXPORT_PATCHES:-}

write_patch() {
	git format-patch \
		--start-number="${patch_num}" \
		--output-directory "${EXPORT_PATCHES}" \
		"${1}^..${1}"
}

# manual_filters takes care of some minor annoyances
# with the raw results from `git log`
manual_filters() {
	case ${1} in
		# cherry-pick of 8a9d926b553345be530ddc51374c9817fcbea784
		a95b4d549802423844de38a19818ac86e14b7d70) ;;
		# cherry-pick of 751a406dece4742eca27b7eece982768bce94c14
		751a406dece4742eca27b7eece982768bce94c14) ;;
		# cherry-pick of a942c92dd77aff229680c7ae2a6de27687527b8a
		38559c69d2832ff29d6d77c7a7351b9aaf12cd37) ;;
		# cherry-pick of e55bead518e4c72cdecf7de2e49db6c477cb58eb
		8efdf3a4604a4ab5c8bbfeed8b0bb220ddd06e6f) ;;
		# cherry-pick of 647cec4324186faa3183bd6a7bc72a032a86c8c9
		e3ffcd18f361097a522dea73eabcd90c660a1f0e) ;;
		# cherry-pick of e69127bd5ba4dcf8ae1f248db93a95795eb75b93
		bf0d23511a92cd0f0de885c8ecedc658831db7bc) ;;
		# cherry-pick of 521e7eba86df25857647b93f13e5366c554e9d63
		23fd4d23c471548703d8edabf38634393fd48e73) ;;
		# revendor everything
		1cf563e3d171ef76964df10665e70ce211d9599d) ;;
		# seems to be 27d9030b2371aa4a6b167fded6b8dc25987a0af7
		# probably slipped in with the 18.09 merge
		4032b6778df39f53fda0e6e54f0256c9a3b1d618) ;;

		*) return 0 ;;
	esac
	return 1
}

patch_num=1
for commit in $(git log \
	--reverse \
	--format="%H" \
	--invert-grep -E \
		--grep='^\(cherry\spicked\sfrom\scommit\s[0-9a-f]{7,40}\)$' \
	--no-merges \
	"${HEAD}" "^${UPSTREAM}"); do

	# filter empty commits, somtimes used to add CHANGELOG entries
	[ -z "$(git --no-pager log --stat --format='' -1 "${commit}")" ] && continue

	# manual filtering:
	manual_filters "${commit}" || continue

	# output here:
	git --no-pager log --stat --format=fuller -1 "${commit}"
	printf "\n---\n\n"

	# allows writing patches out to disk
	if [ -n "$EXPORT_PATCHES" ]; then
		write_patch "${commit}"
	fi
	patch_num=$((patch_num+1))
done
