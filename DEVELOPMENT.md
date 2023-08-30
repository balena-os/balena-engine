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
    * [balena-libnetwork](https://github.com/balena-os/balena-libnetwork)
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
this more complete over time, with pointers to the relevant code.

#### Delta updates

With deltas we allow users to pull only the differences between an image they
already have (the *basis*) and one they want to have (the *target*). Spares
bandwidth from users and balena alike!

#### Resilient image pulls

In the event of network issues while pulling an image, balenaEngine will keep
trying to resume the interrupted download for some time without the need of
restarting from scratch. This is very useful for devices working with an
unstable Internet connection.

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
`bundles/dynbinary-daemon`.

#### Build and run

Using

```sh
make dynbinary shell
```

will build the Engine as above, but will also put you in a container where you
can run it. What I usually do to run the Engine inside this container is:

```sh
# Copy the binary and symlinks to somewhere in the $PATH
cp bundles/dynbinary-daemon/balena* /bin

# Run the required daemons in the background
balena-engine-containerd &
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

More recent versions of Moby use of the standard Go modules/vendoring system.
Until we update, we are using [vndr](https://github.com/LK4D4/vndr).

The safest way to vendor dependencies is this:

1. Edit `vendor.conf`, making the desired dependency point to the desired
   version or commit hash.
2. Run `make BIND_DIR=. shell` to enter into the "development environment".
   container.
3. Run `./hack/vendor.sh`. This will take a while to run, and will re-download
   all dependencies.
4. Leave the development environment (`exit` or Ctrl+D). The code under
   `vendor/` will be updated.

You probably want to stick with the steps above.

However, if you are in a hurry, really know what you are doing, and don't mind
some manual tweaking, you can ask for a single dependency to be vendored. To do
this, simply replace step 3 above with a command like `vndr
github.com/balena-os/librsync-go` (adjusting for the desired dependency). The
danger is that you'll skip some smartness built into the `vendor.sh` script. For
example, as I write this, calling `vndr` directly will *also* remove everything
under `vendor/archive/tar/` (which is needed and must be manually restored).

## Update to a new upstream release

We need to merge the upstream release into the engine repository and update our
component forks to the new versions.

### Merge upstream changes in engine repo

First, fetch the new commits and tags from upstream:
`git fetch --tags https://github.com/moby/moby.git`.

Use `git merge <TARGET_VERSION>` and solve the merge conflicts. You can ignore
`vendor.conf` for now.

You can also ignore everything under `./vendor`. To make it easier you can do:
`git reset ./vendor/ && git checkout -- ./vendor/ && git clean -df ./vendor/`

### Bring components up-to-date

This is the time to update the balena forks of some  components:

* github.com/balena-os/balena-runc (github.com/opencontainers/runc)
* github.com/balena-os/balena-containerd (github.com/containerd/containerd)
* github.com/balena-os/balena-libnetwork (github.com/docker/libnetwork)
* github.com/balena-os/balena-engine-cli (github.com/docker/cli)

The first step is to figure out what's the new commit hashes to base our forks
on:

* Normally, the desired hash is the one present in the updated `vendor.conf`.
* However, be aware that the version of containerd bundled by Moby is defined by
  the `CONTAINERD_VERSION` in `hack/dockerfile/install/containerd.installer`.
  So, you may want to use the hash of this version instead (or the newest among
  it and the one in `vendor.conf`), to make sure balenaEngine will bundle the
  same containerd version as Moby.
* We used containerd as an example above, but the same is valid for the other
  components.

Anyway, once you figure out the target commit hash for a given component, you
can proceed to update it. The easiest way to do that is to:

* Find out what is the current version branch (these are branches named
  `<VERSION>-balena`).
* Find out what is the earliest balena patch on this repo. (Look below in the
  Tips section for some help.)
* Fetch the changes and tags from upstream. For containerd, you'd use
  `git fetch --tags https://github.com/containerd/containerd.git`.
* Copy the current version branch to `<TARGET_VERSION>-balena`:
  `git checkout <CURRENT_BRANCH> && git checkout -b <TARGET_VERSION>-balena`
* Run `git rebase --onto <TARGET_COMMIT> <FIRST_PATCH>^`. Don't forget to add
  the `^`.

There might be merge conflicts.

And if any of the components added new files to their `main` package, you need
to update the `package` declaration on these new files to enable importing as a
package. (Like in [this commit](https://github.com/balena-os/balena-containerd/commit/bdc9478300894cf34bbbd975df1c11b26eb20f63).)

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

## Tips

### Random tips

* As I write this (balenaEngine 20.10.18), we support only cgroups v1.
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

For balena-libnetwork:

```text
Author: Petros Angelatos <petrosagg@gmail.com>
Date:   Tue Jul 25 16:04:43 2017 -0700

    cmd/proxy: export main package as a library

    Allows it to be used as part of a busybox-like binary

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