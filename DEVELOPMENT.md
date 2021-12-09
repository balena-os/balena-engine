
# Working on balenaEngine

## Update to a new upstream release

We need to merge the upstream release into the engine repository and update our
component forks to the new versions.

### Merge upstream changes in engine repo

Use `git merge <TARGET_VERSION>` and solve the merge conflicts. You can ignore
`vendor.conf` for now.

To make it easier you can do:
`git reset ./vendor/ && git checkout -- ./vendor/ && git clean -df ./vendor/`

### Bring components up-to-date

Looking into `vendor.conf`, you will find the new commit hashes to base our
forks on:

* github.com/balena-os/balena-runc (github.com/opencontainers/runc)
* github.com/balena-os/balena-containerd (github.com/containerd/containerd)
* github.com/balena-os/balena-libnetwork (github.com/docker/libnetwork)
* github.com/balena-os/balena-engine-cli (github.com/docker/cli)

The easiest way to do that is to copy the current version branch to
`<TARGET_VERSION>-balena` and run `git rebase --onto <TARGET_COMMIT> <FIRST_PATCH>^`.
For that just find the commit hash of the earliest balena patch (don't forget to add `^`)

There might be merge conflicts. And if `runc` or `containerd` added new files to
their `main` package, the first line needs to be changed to enable importing
as a package.

### Reconstruct vendor/

Go through the changes/merge conflicts in `vendor.conf`. We need to update the
revisions of our components above to the new `HEAD`.

There might be missing new dependencies introduced in the components that we
need to copy under the respective section at the bottom of the engine's vendor
file.

After that you can bring back the vendor directory with `make BIND_DIR=. shell`
and run `hack/vendor.sh`.

### Testing if everything works

Use `make test-unit test-integration` to confirm you were successful.

Once the tests pass we're done :tada:

### Editing the Changelog

We use [versionist](https://github.com/balena-io/versionist) to automatically
maintain our [CHANGELOG.md](./CHANGELOG.md) and expose our changelog to downstream
projects (via nested changelogs).

#### CHANGELOG.md

Copy the upstream release notes from https://docs.docker.com/engine/release-notes
and format them like so:

```markdown
# v{VERSION}
## ({DATE}) [upstream release]

<details>
<summary>Merge upstream {VERSION} [{YOUR NAME}]</summary>

{CONTENT}

</details>
```

#### .versionbot/CHANGELOG.yml

this is used to generate nested changelogs in downstream projects and needs the
changelog in YAML format, we abbreviate like so:

```yaml
- commits:
  - subject: Merge upstream v{VERSION}
    hash: {COMMIT}
    body: >-
      For full changelog see:
      {LINK TO BALENA ENGINE CHANGELOG HEADING}
    footers:
      change-type: major
      signed-off-by: {YOUR NAME} <{YOUR EMAIL}>
    author: {YOUR NAME}
    nested: []
  version: {VERSION}
  date: {DATE}
```

Finally your should bump the version found in [`VERSION`](./VERSION) to the new one.
