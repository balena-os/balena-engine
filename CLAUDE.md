# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

balenaEngine is a container engine for IoT, forked from the Moby Project (Docker). Unlike Docker's multi-binary distribution, balenaEngine compiles into a single busybox-style binary that dispatches to different components based on the invoked command name.

## Build Commands

```bash
# Build statically linked binary (default)
make binary

# Build dynamically linked binary
make dynbinary

# Cross-compile for a specific platform
make cross DOCKER_CROSSPLATFORMS=linux/arm64
make cross DOCKER_CROSSPLATFORMS=linux/arm/v5

# Build and enter development shell
make dynbinary shell

# Inside development container, run the built engine:
cp bundles/dynbinary/balena* /bin
balena-engine-daemon &
balena-engine ps
```

Output binaries are placed in `bundles/binary/` or `bundles/dynbinary/`.

## Testing

```bash
# Run unit tests
make test-unit

# Run unit tests for specific directory
make test-unit TESTDIRS=./image

# Run unit tests with specific test pattern
TESTFLAGS="-test.run TestNameOrPrefix" make test-unit

# Run all integration tests (may take ~1 hour)
TIMEOUT=240m make test-integration

# Run integration tests matching a pattern
make test-integration TEST_FILTER=TestDelta

# Run only new integration tests (under integration/)
make test-integration TEST_SKIP_INTEGRATION_CLI=1

# Run only legacy integration tests (under integration-cli/)
make test-integration TEST_SKIP_INTEGRATION=1
```

## Validation/Linting

```bash
# Run all validations (includes golangci-lint, shfmt, swagger, etc.)
make validate

# Run specific validation
make validate-golangci-lint
make validate-shfmt
make validate-vendor

# Format code
gofmt -s -w file.go
```

## Architecture

### Busybox-Style Binary

The entry point is `cmd/balena-engine/main.go`, which dispatches based on the invoked binary name:
- `balena-engine` / `balena` → CLI (docker/cli fork)
- `balena-engine-daemon` / `balenad` → daemon (moby/moby)
- `balena-engine-containerd` → containerd (containerd/containerd fork)
- `balena-engine-runc` → runc (opencontainers/runc fork)
- `balena-engine-proxy` → docker-proxy

### Key Directories

- `cmd/` - Entry points for all binaries
- `daemon/` - Docker daemon implementation
- `daemon/images/` - Image management, including delta support (`image_delta.go`)
- `distribution/` - Image pull/push logic, including resilient pulls (`pull_v2.go`)
- `client/` - Docker API client library
- `api/` - REST API definitions and handlers
- `libnetwork/` - Networking stack (merged from moby/moby v22.06+)
- `integration/` - New API-based integration tests
- `integration-cli/` - Legacy CLI-based tests (deprecated, do not add new tests here)

### Balena-Specific Features

**Delta Updates** (`daemon/images/image_delta.go`, `distribution/xfer/download.go`):
Uses librsync-go for bandwidth-efficient binary diffs between images.

**Resilient Image Pulls** (`distribution/pull_v2.go`):
Continues interrupted downloads instead of failing on network errors.

**Alternative Delta Data Root**:
`--delta-data-root` and `--delta-storage-driver` flags allow using a separate image store for delta basis images (used during Host OS Updates).

## Vendoring

Uses Go modules via `vendor.mod` (not standard `go.mod` to avoid SemVer requirements). Balena forks are integrated using `replace` directives:

```bash
# Get prerelease version for a balena fork branch
GOPROXY=direct ./hack/with-go-mod.sh go get -d github.com/balena-os/balena-containerd@<branch-name>

# Add/update replace directive
./hack/with-go-mod.sh go mod edit -modfile=vendor.mod \
  -replace=github.com/containerd/containerd=github.com/balena-os/balena-containerd@<version>

# Regenerate vendor directory
./hack/vendor.sh all
```

**Important:** Do not create `go.mod` or `go.sum` in the repo root; they interfere with the vendor process.

## Build Tags

Default build tags (set in Makefile): `apparmor seccomp no_btrfs no_cri no_devmapper no_zfs exclude_disk_quota exclude_graphdriver_btrfs exclude_graphdriver_devicemapper exclude_graphdriver_zfs`

## Test Environment Notes

- Integration tests require privileged containers
- libnetwork tests cannot run in parallel (use `-p=1`)
- If userns-remap tests fail with permission errors, delete `bundles/` directory: `sudo rm -rf bundles`

## Running Integration Tests in Docker (Recommended)

For reliable integration testing, run tests inside a Docker container with a properly configured daemon:

```bash
# Quick integration test run
docker run --rm --privileged \
  -v "$(pwd):/go/src/github.com/docker/docker" \
  -w /go/src/github.com/docker/docker \
  -e DOCKER_GRAPHDRIVER=overlay2 \
  docker-dev bash -c '
    cp bundles/dynbinary-daemon/balena* /usr/local/bin/
    balena-engine-daemon --storage-driver overlay2 2>/dev/null &
    sleep 5
    export DOCKER_HOST=unix:///var/run/balena-engine.sock
    go test -v -timeout 15m ./integration/container/...
  '
```

## Current Project Status (January 2026)

**Branch**: `balena/v27-rebase` (based on moby v27.5.1)
**Status**: Feature complete, testing complete

| Test Suite | Status |
|------------|--------|
| Unit tests | ✅ 524 pass, 3 skip |
| Container integration | ✅ 135 pass |
| Delta integration | ✅ 10 pass |
| Daemon integration | ✅ All pass |

See `PRD.md` for detailed status and history.

## Local Forks

balena-engine uses local forks in `forks/` directory for busybox-style binary:

```
forks/
├── balena-runc/           # Exports Main() for runc
├── balena-containerd/     # Exports Main() for containerd components
└── balena-engine-cli/     # Exports Main() for CLI
```

These are referenced via replace directives in `vendor.mod`. See `VENDORING.md` for details.

## Key Files for This Project

| File | Purpose |
|------|---------|
| `PRD.md` | Project status, history, and detailed notes |
| `TESTING.md` | Test documentation |
| `VENDORING.md` | Vendoring and fork documentation |
| `cmd/balena-engine/main.go` | Busybox dispatcher |
| `vendor.mod` | Dependencies with replace directives |
