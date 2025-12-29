# Working on balenaEngine

## Crash course for beginners

### Project structure

* The balenaEngine repo is a fork of the [Moby
Project](https://github.com/moby/moby/) repo.
* From a high-level, architectural perspective the main difference between them
  is this:
    * Moby/Docker is distributed as a number of separate binaries (`docker`,
      `dockerd`, `containerd`, `runc`, `docker-proxy`, etc).
    * balenaEngine is compiled into a single [busybox-style
      binary](https://www.busybox.net/FAQ.html#design).
* To achieve this, we also maintain forks of the projects from where all the
  other Docker/Moby binaries come from:
    * [balena-containerd](https://github.com/balena-os/balena-containerd/)
    * [balena-engine-cli](https://github.com/balena-os/balena-engine-cli)
    * [balena-runc](https://github.com/balena-os/balena-runc/)
* Note: libnetwork was merged into moby/moby in v22.06, so balena-libnetwork is
  no longer needed for newer versions.
* Each of these forks contains a commit allowing them to be used as a library.
  These commits rename the package `main` and export the main function by
  renaming it from `main()` to `Main()`. These changes enable the busybox-style
  usage we want.
    * For example, [here's how we do this for for
      containerd](https://github.com/balena-os/balena-containerd/commit/bdc9478300894cf34bbbd975df1c11b26eb20f63).
    * And [this is balenaEngine's `main()`
      function](https://github.com/balena-os/balena-engine/blob/ad3f3a029cd911d4919e079df16e97922c3c437a/cmd/balena-engine/main.go#L25),
      where we dispatch the execution to the appropriate `Main()`.

### Unique features

This is an incomplete list of features unique to balenaEngine. I hope to make
this more complete over time.

#### Delta updates

With deltas we allow users to pull only the differences between an image they
already have (the *basis*) and one they want to have (the *target*). Spares
bandwidth from users and balena alike!

Relevant code:

* The delta algorithms themselves are implemented in balena's [librsync-go
  library](https://github.com/balena-os/librsync-go). This is the library that supports
* On the Engine side, delta creation is implemented in the
  [`ImageService.DeltaCreate()`](https://github.com/balena-os/balena-engine/blob/2cd17c44b267813dfc1153f7d08306024d1fc032/daemon/images/image_delta.go#L29)
  function (at `daemon/images/image_delta.go`). This code is pretty much
  self-contained.
* Applying deltas is a bit more complicated, as our code is "mixed" with Moby's
  code. The main point of interest is the
  `LayerDownloadManager.makeDownloadFunc()` function (at
  `distribution/xfer/download.go`), particularly the code around the [call to
  `DecorateWithDeltaPatcher()`](https://github.com/balena-os/balena-engine/blob/2cd17c44b267813dfc1153f7d08306024d1fc032/distribution/xfer/download.go#L364C55-L364C55).
  In a nutshell, what we have here is a pipeline of operations: downloading the
  layer data, decompressing it, etc. What we do is adding our own step into this
  pipeline. This step takes the delta itself on the input and produces the
  target layer on the output.

#### Resilient image pulls

In the event of network issues while pulling an image, balenaEngine will keep
trying to resume the interrupted download without the need of restarting from
scratch. This is very useful for devices working with an unstable Internet
connection.

Relevant code: Our changes have been to the [`v2LayerDescriptor.Read()`
function](https://github.com/balena-os/balena-engine/blob/2cd17c44b267813dfc1153f7d08306024d1fc032/distribution/pull_v2.go#L196)
(`distribution/pull_v2.go`), which basically implements Go's `Reader` interface
with data coming from an HTTP source. The idea behind our changes is simple:
instead of returning an `error` when a download error happens, we return `nil`.
This will cause the caller to keep trying until the network connectivity is
reestablished.

#### Alternative delta data root

TL;DR: Enables the use of deltas for Host OS Updates (HUPs).

Docker stores images in what is unsurprisingly called an Image Store. All images
you pull or build are placed in a single Image Store. If you are familiar with
that, the Image Store data is normally placed (along with other things) under
`/var/lib/docker/` (or `/mnt/data/docker/` in the case of balenaOS).

With balenaEngine we offer two command-line options, `--delta-data-root` and
`--delta-storage-driver`, that allow to configure a second Image Store which is
used exclusively when looking for the basis images for deltas.

balenaEngine on balenaOS will normally not use these options: just like with
Docker, a single Image Store is used. When we do a delta update of a user
container, the basis will be in this Image Store.

The only situation we use these options is during Host OS Updates (HUPs). In
this case, the basis image (i.e., the old balenaOS version) is on [a different
partition](https://os-docs.balena.io/architecture#image-partition-layout) than
the target image. So, we use `--delta-data-root` and `--delta-storage-driver` to
make sure we can find the basis image on that other partition.

### Day-to-day tasks / Cheat sheet

Unless otherwise is specified, all commands described below are to be executed
directly in your development computer.

#### Build

To build the Engine you can run

```sh
make dynbinary
```

This will place the generated binary and symlinks into
`bundles/dynbinary`.

#### Build and run

Using

```sh
make dynbinary shell
```

will build the Engine as above, but will also put you in a container where you
can run it. What I usually do to run the Engine inside this container is:

```sh
# Copy the binary and symlinks to somewhere in the $PATH
cp bundles/dynbinary/balena* /bin

# Run the required daemons in the background.
# The engine daemon also starts the balena-engine-containerd daemon
balena-engine-daemon &

# Now you can run balena-engine as you wish
balena-engine ps
```

#### Cross-compiling

Sometimes you may want to try your freshly built balenaEngine on a device. For
these cases, cross-compiling is the way to go:

```sh
# Use the platform corresponding to your device, for example:
make cross DOCKER_CROSSPLATFORMS=linux/arm64
make cross DOCKER_CROSSPLATFORMS=linux/arm/v5
```

This will place the generated binary iton `bundles/cross/...`.

Tip: You should replace your device's `/usr/bin/balena-engine` with the one you
compiled. However, the root partition of balenaOS is pretty short of space and
thus this operation may fail. So, you can copy your binary to the data partition
(`/mnt/data`) and replace `/usr/bin/balena-engine` with a symlink to it.

#### Debugging

There's no official support for running balenaEngine (or Moby, for that matter)
under a debugger in the current release. This shall be possible with the next
Moby release (22.06), which hopefully will be out soon.

Anyway, the [lmbarros/debug](/balena-engine/tree/lmbarros/debug) branch provides
a quick-and-dirty debugging support for the time being. [Check the instructions
here](https://github.com/balena-os/balena-engine/blob/lmbarros/debug/DEBUGGING.md).

#### Running automated tests

Running the **unit tests** is simple enough:

```sh
make test-unit
```

The whole suite runs in two minutes on my laptop. Anyway, you can specify a
directory and run only the tests defined there:

```sh
make test-unit TESTDIRS=./image
```

Running all **integration tests** is similar, but it's a good idea to increase
the timeout, as running the whole suite can take about an hour:

```sh
TIMEOUT=240m make test-integration
```

You can also run only a subset of the integration tests. For example, to run
only the tests containing `TestDelta` in their names, you'd use:

```sh
make test-integration TEST_FILTER=TestDelta
```

The Moby project has **two different sets of integration tests**. The new one is
under the `integration` directory and has tests that perform calls to the API.
The older set of tests is under `integration-cli` and is based on calls to the
Docker (balenaEngine, in our case) binary. This old CLI suite is still relevant,
despite being considered deprecated. Deprecation only means that, when needed,
Docker devs should not update an old test case, but instead move them to new
suite while changing them to make use of the API.

#### Vendoring

Starting with v23, Moby uses Go modules for vendoring via `vendor.mod` (not a
standard `go.mod` to avoid SemVer requirements since Moby uses CalVer).

The balena forks are integrated using `replace` directives in `vendor.mod`:

```
replace github.com/opencontainers/runc => github.com/balena-os/balena-runc <version>
replace github.com/containerd/containerd => github.com/balena-os/balena-containerd <version>
replace github.com/docker/cli => github.com/balena-os/balena-engine-cli <version>
```

##### Getting prerelease version strings for forks

Go modules require pseudo-version strings for branches/commits. To get the
correct version string for a balena fork branch:

```sh
GOPROXY=direct ./hack/with-go-mod.sh go get -d github.com/balena-os/balena-containerd@<branch-name>
```

This will fail with a module path mismatch error, but will print the resolved
version string, e.g., `v1.6.23-0.20251204223054-27af7c297d34`.

##### Adding/updating replace directives

Use the hack wrapper to add or update replace directives:

```sh
./hack/with-go-mod.sh go mod edit -modfile=vendor.mod \
  -replace=github.com/containerd/containerd=github.com/balena-os/balena-containerd@<version>
```

##### Running the vendor process

After updating `vendor.mod` with the replace directives:

```sh
./hack/vendor.sh all
```

This runs `go mod tidy` followed by `go mod vendor` to update `vendor.mod`,
`vendor.sum`, and the `vendor/` directory.

**Important:** Make sure there is no `go.mod` or `go.sum` file in the repository
root before running `vendor.sh`, as these will interfere with the process.

## Update to a new upstream release

We need to merge the upstream release into the engine repository and update our
component forks to the new versions.

### Merge upstream changes in engine repo

First, fetch the new commits and tags from upstream:
`git fetch --tags https://github.com/moby/moby.git`.

Use `git merge <TARGET_VERSION>` and solve the merge conflicts. You can ignore
`vendor.mod` for now.

You can also ignore everything under `./vendor`. To make it easier you can do:
`git reset ./vendor/ && git checkout -- ./vendor/ && git clean -df ./vendor/`

### Bring components up-to-date

This is the time to update the balena forks of some components:

| Balena Fork | Upstream |
|-------------|----------|
| <https://github.com/balena-os/balena-runc> | <https://github.com/opencontainers/runc> |
| <https://github.com/balena-os/balena-containerd> | <https://github.com/containerd/containerd> |
| <https://github.com/balena-os/balena-engine-cli> | <https://github.com/docker/cli> |

Note: libnetwork was merged into moby/moby in v22.06, so balena-libnetwork is
no longer needed.

#### Determine target versions

Check `vendor.mod` for the version of each dependency that upstream uses.

**Important:** Always use the version from `vendor.mod`, not from installer
scripts like `hack/dockerfile/install/containerd.installer`. Using a different
version can cause dependency conflicts (e.g., mismatched OpenTelemetry versions
between containerd and buildkit).

For docker/cli, check the tags in the upstream cli repo - note that cli versions
may not match moby versions exactly (e.g., moby v23.0.18 uses cli v23.0.15).

#### Rebase the balena forks

For each fork, rebase the balena patches onto the new upstream version:

1. Clone the fork and fetch upstream tags:

   ```sh
   git clone https://github.com/balena-os/balena-containerd.git
   cd balena-containerd
   git fetch --tags https://github.com/containerd/containerd.git
   ```

2. Find the current balena branch and earliest balena patch. Balena branches are
   named `<VERSION>-balena` (e.g., `1.6.22-balena`). See the "Earliest balena
   patches" section below to identify the first balena commit.

3. Create a new branch and rebase onto the target version:

   ```sh
   git checkout <CURRENT_VERSION>-balena
   git checkout -b <TARGET_VERSION>-balena
   git rebase --onto v<TARGET_VERSION> <FIRST_BALENA_PATCH>^
   ```

   Don't forget the `^` after the first patch hash.

4. Resolve any merge conflicts.

5. If the component added new files to their `main` package, update the
   `package` declaration on these new files to enable importing as a package.
   (See https://github.com/balena-os/balena-containerd/commit/bdc9478300894cf34bbbd975df1c11b26eb20f63
   for an example.)

6. Push the new branch and the upstream tag to the balena fork:

   ```sh
   git push origin <TARGET_VERSION>-balena
   git push origin v<TARGET_VERSION>
   ```

   **Important:** Pushing the upstream tag (e.g., `v1.6.22`) ensures that Go's
   pseudo-version calculation works correctly. The pseudo-version is based on
   the most recent tag in the commit history.

### Reconstruct vendor/

1. Start with the upstream `vendor.mod` from the merged release.

2. Get the prerelease version strings for each balena fork. The branch names
   follow the pattern `<TARGET_VERSION>-balena` (e.g., `1.6.22-balena` for
   containerd, `1.1.12-balena` for runc, `23.0.15-balena` for cli):

   ```sh
   GOPROXY=direct ./hack/with-go-mod.sh go get -d github.com/balena-os/balena-runc@1.1.12-balena
   GOPROXY=direct ./hack/with-go-mod.sh go get -d github.com/balena-os/balena-containerd@1.6.22-balena
   GOPROXY=direct ./hack/with-go-mod.sh go get -d github.com/balena-os/balena-engine-cli@23.0.15-balena
   ```

   Each command will fail with a module path mismatch error but will print the
   resolved pseudo-version string (e.g., `v1.6.23-0.20251204223054-27af7c297d34`).

3. Add the replace directives to `vendor.mod` using the resolved versions:

   ```sh
   ./hack/with-go-mod.sh go mod edit -modfile=vendor.mod \
     -replace=github.com/opencontainers/runc=github.com/balena-os/balena-runc@<version>
   ./hack/with-go-mod.sh go mod edit -modfile=vendor.mod \
     -replace=github.com/containerd/containerd=github.com/balena-os/balena-containerd@<version>
   ./hack/with-go-mod.sh go mod edit -modfile=vendor.mod \
     -replace=github.com/docker/cli=github.com/balena-os/balena-engine-cli@<version>
   ```

4. Run the vendor process:

   ```sh
   ./hack/vendor.sh all
   ```

   This will run `go mod tidy` to resolve transitive dependencies and then
   `go mod vendor` to populate the vendor directory.

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

## Tips

### Integration test failures with userns-remap

If `TestBuildUserNamespaceValidateCapabilitiesAreV2` (or other userns tests) fail with:

```
a subdirectory in your graphroot path (...) restricts access to the remapped root uid/gid
```

The likely cause is stale `bundles/` directory permissions. When running locally,
`bundles/` is bind-mounted from the host. If it was created with restrictive
permissions (e.g., `drwx------`), the remapped uid (165536) cannot traverse it.

**Fix:** Delete the bundles directory and let it be recreated:

```sh
sudo rm -rf bundles
```

This doesn't affect CI because GitHub Actions uses an anonymous Docker volume
for `bundles/` instead of a bind mount.

### Random tips

* This is something we need to look deeper, but I have seen some errors in
  automated tests when using very recent kernel versions. This happens because
  of changes in some kernel interface. AFAIR, this was fixed upstream, but yet
  brought to balenaEngine.
    * I know this is a very vage tip -- just be aware that things like this can
      happen.
    * FWIW, in my case (mid-2022), kernel 5.15.x was fine; 5.19 wasn't.

### Earliest balena patches

To make it easier to locate them, here's a list of the earliest balena patches
for each of the balena forks. Since commit hashes will change as we rebase, I am
not including them here.

For balena-runc:

```text
Author: Petros Angelatos <petrosagg@gmail.com>
Date:   Tue Jul 25 15:55:23 2017 -0700

    runc: export main package as a library

    Allows runc to be used as part of a busybox-like binary

    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>
```

Watch out! Don't be confused by an earlier commit by Petros, which is [merged
upstream](https://github.com/opencontainers/runc/commit/809882868011fbe1d011f23deb034f3541d556c8).

For balena-containerd:

```text
Author: Petros Angelatos <petrosagg@gmail.com>
Date:   Wed Jan 17 19:06:48 2018 -0800

    export all commands as packages

    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>
```

For balena-engine-cli:

```text
Author: Petros Angelatos <petrosagg@gmail.com>
Date:   Tue Jul 25 16:46:51 2017 -0700

    cmd/docker: export main package as a library

    Allows it to be used as part of a busybox-like binary

    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>
```

<!--
TODO:

* Add some complete workflows for the basic tasks.
    * For example, how to copy a newly built Engine to a device (including all
      the scp commands and whatever else is needed).

-->