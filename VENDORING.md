# Vendoring policies

This document outlines recommended Vendoring policies for Docker repositories.
(Example, libnetwork is a Docker repo and logrus is not.)

## Vendoring using tags

Commit ID based vendoring provides little/no information about the updates
vendored. To fix this, vendors will now require that repositories use annotated
tags along with commit ids to snapshot commits. Annotated tags by themselves
are not sufficient, since the same tag can be force updated to reference
different commits.

Each tag should:
- Follow Semantic Versioning rules (refer to section on "Semantic Versioning")
- Have a corresponding entry in the change tracking document.

Each repo should:
- Have a change tracking document between tags/releases. Ex: CHANGELOG.md,
github releases file.

The goal here is for consuming repos to be able to use the tag version and
changelog updates to determine whether the vendoring will cause any breaking or
backward incompatible changes. This also means that repos can specify having
dependency on a package of a specific version or greater up to the next major
release, without encountering breaking changes.

## Semantic Versioning
Annotated version tags should follow [Semantic Versioning](http://semver.org) policies:

"Given a version number MAJOR.MINOR.PATCH, increment the:

   1. MAJOR version when you make incompatible API changes,
   2. MINOR version when you add functionality in a backwards-compatible manner, and
   3. PATCH version when you make backwards-compatible bug fixes.

Additional labels for pre-release and build metadata are available as extensions
to the MAJOR.MINOR.PATCH format."

## Vendoring cadence
In order to avoid huge vendoring changes, it is recommended to have a regular
cadence for vendoring updates. e.g. monthly.

## Pre-merge vendoring tests
All related repos will be vendored into docker/docker.
CI on docker/docker should catch any breaking changes involving multiple repos.

## Balena Engine Local Forks

balena-engine uses local forks of three key dependencies to enable the busybox-style single binary. These forks export `Main()` functions so the busybox dispatcher can call into them.

### Fork Directory Structure

```
forks/
├── balena-runc/           # Fork of opencontainers/runc
├── balena-containerd/     # Fork of containerd/containerd
└── balena-engine-cli/     # Fork of docker/cli
```

### vendor.mod Replace Directives

```
replace github.com/opencontainers/runc => ./forks/balena-runc
replace github.com/containerd/containerd => ./forks/balena-containerd
replace github.com/containerd/containerd/api => ./forks/balena-containerd/api
replace github.com/docker/cli => ./forks/balena-engine-cli
```

### Fork Modifications

Each fork is modified to export a `Main()` function:

**balena-runc** (`forks/balena-runc/main.go`):
- Package changed from `main` to `runc`
- Function changed from `main()` to `Main()`

**balena-containerd** (multiple files):
- `cmd/containerd/main.go` - Package `containerd`, exports `Main()`
- `cmd/ctr/main.go` - Package `ctr`, exports `Main()`
- `cmd/containerd-shim-runc-v2/main.go` - Package `main`, exports `Main()`
- Additional compatibility patches for API changes

**balena-engine-cli** (`cmd/docker/docker.go`):
- Already had `package docker` and exported `Main()` - no changes needed

### Updating Local Forks

1. **Pull upstream changes**:
   ```bash
   cd forks/balena-containerd
   git fetch origin
   git merge v1.7.31  # desired version
   ```

2. **Re-apply balena patches**:
   - Main() exports in cmd/*/main.go
   - API compatibility fixes if needed
   - go.mod version constraints

3. **Regenerate vendor**:
   ```bash
   ./hack/vendor.sh all
   ```

4. **Test**:
   ```bash
   make dynbinary
   make test-unit
   ```

### Publishing Forks for Production

For production releases, push local forks to GitHub:

```bash
# Create tagged release branch
cd forks/balena-containerd
git remote add balena git@github.com:balena-os/balena-containerd.git
git tag v1.7.30-balena.1
git push balena v1.7.30-balena.1

# Update vendor.mod to use remote
# replace github.com/containerd/containerd => github.com/balena-os/balena-containerd v1.7.30-balena.1
```

### Version Compatibility Notes

When updating containerd:
- API submodule versions must match (e.g., api/v1.8.0 for containerd v1.7.30)
- Check go-runc interface changes (`StartLocked()` added in v1.1.0)
- Check cri-api changes (`RuntimeConfig()` added in v0.28.3)
- golang.org/x/* packages may need downgrading for Go version compatibility
