# Plan: Port balena-engine Patches to Moby v27

## Overview

Port 211 balena-specific commits from `kyle/rerun-rebase-v23.0.18` (based on moby v23.0.18) to moby v27.5.1, resulting in a fully buildable and test-passing balena-engine.

## Decisions

- **Target version**: moby v27.5.1 (chosen over v28.x due to containerd v2 migration complexity)
- **Commit preservation**: Keep all commits separate (for git bisect and clear history)
- **Fork updates**: Include balena-runc, balena-containerd, balena-engine-cli updates in this effort

## Version Decision Rationale

v28.x uses containerd v2.1.x which is a major architectural change. v27.5.1 uses containerd v1.7.24, making it a more manageable upgrade from the current v1.6.22.

## Current State

- **Branch**: `kyle/rerun-rebase-v23.0.18`
- **Base**: moby v23.0.18
- **Balena patches**: 211 commits categorized as:
  - Delta/ioutils support (~40 commits)
  - Test compatibility (~46 commits)
  - Documentation (~25 commits)
  - Build system / multicall binary (~18 commits)
  - Swarm removal (~16 commits)
  - Vendor updates (~12 commits)
  - Storage migration aufs→overlay2 (~12 commits)
  - Mobynit host app booting (~11 commits)
  - Networking customizations (~9 commits)

## Dependency Versions

| Component | Current (v23.0.18) | Target (v27.5.1) |
|-----------|-------------------|------------------|
| containerd | v1.6.22 (balena-containerd) | v1.7.24 |
| runc | v1.2.9 (balena-runc) | go-runc v1.1.0 |
| CLI | v23.0.16 (balena-engine-cli) | v27.x |

## High-Risk Areas

| Area | Risk | Reason |
|------|------|--------|
| `vendor.mod` replace directives | HIGH | v27 uses containerd 1.7.x vs 1.6.x |
| Multicall binary (`cmd/balena-engine/main.go`) | MEDIUM | containerd client library updates |
| Networking (libnetwork) | MEDIUM | v27 has networking changes (less severe than v28) |
| gRPC/BuildKit constraints | LOW | BuildKit changes are incremental |

## Implementation Plan

### Phase 1: Preparation ✅ COMPLETE

1. **Set up repository**
   ```bash
   git remote add moby https://github.com/moby/moby.git
   git fetch moby --tags
   git checkout -b balena/v27-rebase v27.5.1
   ```

2. **Extract patches for reference**
   ```bash
   mkdir -p .patches
   git format-patch v23.0.18..kyle/rerun-rebase-v23.0.18 -o .patches/
   ```

3. **Document patch categories** - Create tracking spreadsheet/issue for each patch

### Phase 2: Update Component Forks

Before applying patches, update the three critical forks to v27-compatible versions:

1. **balena-runc** (github.com/balena-os/balena-runc)
   - Current: v1.2.9-balena (already v1.2.x compatible)
   - Verify compatibility with v27.5.1 go-runc v1.1.0
   - Re-apply `Main()` export patch if needed
   - Tag as `v1.2.x-balena`

2. **balena-containerd** (github.com/balena-os/balena-containerd)
   - Current: v1.6.22-balena
   - Target: containerd/containerd v1.7.24+
   - Re-apply `Main()` export patches for containerd and shim-runc-v2
   - Verify shim plugin filtering still works
   - Tag as `v1.7.x-balena`

3. **balena-engine-cli** (github.com/balena-os/balena-engine-cli)
   - Current: v23.0.16-balena
   - Target: docker/cli v27.x
   - Re-apply `Main()` export patch
   - Port delta command if CLI-side changes exist
   - Tag as `v27.x-balena`

### Phase 3: Apply Patches (Staged)

**Strategy**: Use `git cherry-pick` to preserve individual commits. For each patch:
```bash
git cherry-pick <commit-hash>
# On conflict: resolve, then `git cherry-pick --continue`
# Document conflicts in commit message with "Adapted for v28: <reason>"
```

Apply patches in order of increasing conflict likelihood:

**Stage A: Low-conflict patches**
- Documentation (DEVELOPMENT.md updates)
- Frozen image curl improvements
- librsync-go vendor addition
- Delta API types

**Stage B: Core delta feature**
- `daemon/images/image_delta.go` (new file)
- `client/image_delta.go` (new file)
- `api/server/router/image/image.go` (add delta route)
- `distribution/xfer/download.go` (delta decorator)
- `distribution/pull_v2.go` (delta base discovery)
- `image/tarexport/load.go` (delta on load)

**Stage C: Build system**
- `cmd/balena-engine/main.go` (multicall dispatcher)
- `hack/make/binary-daemon`, `.binary-symlinks`
- `Dockerfile` modifications
- Init scripts (`contrib/init/`)

**Stage D: Daemon & config**
- `daemon/config/config.go` (balena defaults)
- `cmd/dockerd/config_unix.go` (delta flags)
- Swarm removal from routers
- Containerd shim plugin filtering

**Stage E: Networking**
- Bridge rebranding (`DefaultBridgeName = "balena0"`)
- Proxy naming (`balena-engine-proxy`)
- macvlan/overlay driver disabling (may need rewrite for v28)

**Stage F: Test fixes**
- Integration test adjustments
- CLI test compatibility
- Swarm test removal

### Phase 4: Get Tests Passing

**Order of test enablement:**

1. **Unit tests** (`make test-unit`)
   - Fast feedback (~2 min)
   - Catches compilation/import issues

2. **Integration tests** (`make test-integration`)
   - Start with delta tests: `TEST_FILTER=TestDelta`
   - Then container lifecycle tests
   - Then full suite

3. **Integration-CLI tests** (`make test-integration-cli`)
   - Depends on CLI fork being updated
   - Many CLI output changes in v28

4. **Docker-py tests** (`make test-docker-py`)
   - Python SDK compatibility

**Expected test categories to skip/remove:**
- Swarm tests (feature removed)
- Plugin tests (feature disabled)
- Schema 1 tests (removed in v28)
- Cloud logging tests (feature disabled)

### Phase 5: Verification

1. **Build verification**
   - `make binary` succeeds
   - `make cross` for ARM builds
   - Multicall binary works (all symlinks functional)

2. **Delta functionality**
   - `TestDeltaCreate` passes
   - Manual delta create/load cycle works
   - Delta sizes reasonable

3. **Runtime smoke test**
   - Container run/stop/start/remove
   - Image pull/push/tag
   - Networking (bridge, port publish)
   - Exec into container

4. **CI pipeline**
   - All GitHub Actions workflows pass
   - Coverage reports generate

## Critical Files to Modify

| File | Purpose |
|------|---------|
| `vendor.mod` | Update replace directives for v28 forks |
| `cmd/balena-engine/main.go` | Multicall binary - verify v28 imports |
| `distribution/xfer/download.go` | Delta integration point |
| `daemon/images/image_delta.go` | Delta creation (mostly additive) |
| `libnetwork/drivers/bridge/interface.go` | Bridge rebranding (v28 networking changed) |
| `Dockerfile` | Build stages, Go version |
| `.github/workflows/test.yml` | CI configuration |

## Verification Checklist

- [ ] `make binary` builds successfully
- [ ] `make test-unit` passes
- [ ] `make test-integration` passes (excluding known skips)
- [ ] `make test-integration-cli` passes (excluding swarm)
- [ ] Delta create/load cycle works manually
- [ ] Container lifecycle (run/stop/start/rm) works
- [ ] Bridge networking functional
- [ ] All symlinks work (balena-engine-daemon, balena-engine-containerd, etc.)
- [ ] GitHub Actions CI passes

## Notes

- v28 includes major networking overhaul - bridge/macvlan patches need careful review
- containerd v2 client library changes may affect multicall binary imports
- Schema 1 support removed in v28 - verify no delta functionality depends on it
- Keep `kyle/rerun-rebase-v23.0.18` branch intact as reference during porting

---

## Current State (v27 Port Progress)

**Date**: January 2026
**Branch**: `balena/v27-rebase`
**Status**: Busybox-style binary builds and dispatches correctly

### Completed Work

#### 1. Forks Directory Structure
Created `forks/` directory with local clones of the three key dependencies:
- `forks/balena-runc` - based on opencontainers/runc v1.2.4
- `forks/balena-containerd` - based on containerd/containerd v1.7.30
- `forks/balena-engine-cli` - based on docker/cli v27.4.0

#### 2. Fork Patches Applied

**balena-runc:**
- Changed `package main` → `package runc`
- Exported `func main()` → `func Main()`

**balena-containerd** (most complex, required multiple compatibility patches):
- Exported `Main()` in `cmd/containerd`, `cmd/ctr`, `cmd/containerd-shim-runc-v2`
- Added api submodule v1.8.0 (required for `api/types/runc/options` package)
- Fixed `ExpiresInSeconds` → `ExpiresIn` field naming in auth types (buildkit compat)
- Fixed `Timestamp.UnixNano()` → `protobuf.FromTimestamp().UnixNano()` (protobuf compat)
- Added `StartLocked()` method to reaper (go-runc v1.1.0 compat)
- Added `RuntimeConfig()` method to CRI service (cri-api v0.28.3 compat)
- Downgraded golang.org/x/* packages to Go 1.22 compatible versions

**balena-engine-cli:**
- Already had `package docker` and exported `Main()` - no changes needed

#### 3. vendor.mod Updates
```
replace github.com/opencontainers/runc => ./forks/balena-runc
replace github.com/containerd/containerd => ./forks/balena-containerd
replace github.com/containerd/containerd/api => ./forks/balena-containerd/api
replace github.com/docker/cli => ./forks/balena-engine-cli
```
Also added `github.com/docker/cli-docs-tool v0.8.0` to pin a Go 1.22 compatible version.

#### 4. Build System Changes
- Updated `hack/make/.binary`:
  - `GO_PACKAGE` → `github.com/docker/docker/cmd/balena-engine`
  - `BINARY_SHORT_NAME` → `balena-engine`
  - Removed version suffix for dev builds

#### 5. Build Verification
```
$ ls -la bundles/binary/balena-engine
-rwxr-xr-x 1 shaun shaun 86001464 Jan 16 12:50 balena-engine

$ file bundles/binary/balena-engine
ELF 64-bit LSB executable, x86-64, statically linked

$ ./balena-engine --version
Docker version unknown-version, build unknown-commit

$ ./balena-engine-daemon --version  # via symlink
Docker version dev, build eb0632625f

$ ./balena-runc --version  # via symlink
runc github.com/containerd/containerd unknown
spec: 1.2.0

$ ./balena-containerd --version  # via symlink
containerd github.com/containerd/containerd 1.7.30+unknown
```

### Key Learnings

#### Containerd 1.7.x Compatibility Challenges

1. **API Submodule is Separate**: In containerd 1.7.x, `github.com/containerd/containerd/api` is a separate Go module. When using a replace directive for containerd, you must also add a replace for the api submodule, otherwise the api packages won't be found.

2. **API Version Matters**: The api submodule has its own versioning (api/v1.7.19, api/v1.8.0, etc.). Different versions have different packages available:
   - api/v1.7.19 lacks `api/types/runc/options`
   - api/v1.8.0 includes `api/types/runc/options` (needed by moby v27)

3. **Field Naming Changes**: containerd 1.7.30 uses `ExpiresInSeconds` in auth types, but buildkit expects `ExpiresIn`. This required renaming the fields.

4. **Interface Changes**:
   - go-runc v1.1.0 added `StartLocked()` to `ProcessMonitor` interface
   - cri-api v0.28.3 added `RuntimeConfig()` RPC method

5. **Protobuf Timestamp Handling**: `*timestamppb.Timestamp` doesn't have `.UnixNano()` directly - must use `protobuf.FromTimestamp()` conversion.

#### Go Version Constraints

- containerd 1.7.30's go.mod specifies newer golang.org/x/* packages that require Go 1.24
- These must be downgraded to work with Go 1.22 (moby v27's Go version)
- Affected packages: x/net, x/sync, x/sys, x/crypto, x/mod, x/oauth2, x/term, x/text, x/time

#### Vendoring with Local Replace Directives

- `go mod vendor` with local replace directives works but requires the target directories to have proper go.mod files
- Nested modules (like containerd/api) need their own replace directives
- The `hack/vendor.sh` script must be run inside a container with Go installed

### Remaining Steps to Fully Working balena-engine

#### Phase A: Unit Tests (Estimated: 1-2 hours)
- [ ] Run `make test-unit` and fix any remaining failures
- [ ] Previous run showed 524 passing, 3 skipped - verify this is still the case

#### Phase B: Apply Balena-Specific Patches (Estimated: 2-4 days)
Cherry-pick the 211 balena patches from `kyle/rerun-rebase-v23.0.18`:

1. **Delta Support** (~40 commits)
   - [ ] `daemon/images/image_delta.go` - core delta creation
   - [ ] `distribution/xfer/download.go` - delta decorator
   - [ ] `distribution/pull_v2.go` - delta base discovery
   - [ ] `image/tarexport/load.go` - delta on load
   - [ ] `client/image_delta.go` - client API
   - [ ] librsync-go vendor addition

2. **Daemon Configuration** (~15 commits)
   - [ ] `daemon/config/config.go` - balena defaults
   - [ ] `cmd/dockerd/config_unix.go` - delta flags
   - [ ] Alternative delta data root support

3. **Networking** (~9 commits)
   - [ ] Bridge rebranding (`balena0`)
   - [ ] Proxy naming (`balena-engine-proxy`)
   - [ ] macvlan/overlay driver changes

4. **Storage Migration** (~12 commits)
   - [ ] AUFS to overlay2 migration support
   - [ ] `pkg/storagemigration/` package

5. **Mobynit/Host App** (~11 commits)
   - [ ] `cmd/mobynit/` for host OS booting

6. **Swarm Removal** (~16 commits)
   - [ ] Remove swarm endpoints from API router
   - [ ] Remove swarm tests

7. **Healthcheck Enhancements** (~5 commits)
   - [ ] Container restart on unhealthy
   - [ ] Healthcheck timeout handling

8. **Resilient Pulls** (~5 commits)
   - [ ] Continue interrupted downloads
   - [ ] Max retry configuration

9. **Build/CI** (~18 commits)
   - [ ] GitHub Actions workflows
   - [ ] Dockerfile adjustments
   - [ ] Init scripts

10. **Documentation** (~25 commits)
    - [ ] DEVELOPMENT.md updates
    - [ ] README rebranding

11. **Test Compatibility** (~46 commits)
    - [ ] Integration test adjustments
    - [ ] Swarm test removal

#### Phase C: Integration Tests (Estimated: 1-2 days)
- [ ] Run `make test-integration TEST_FILTER=TestDelta`
- [ ] Run full integration test suite
- [ ] Fix any test failures from API changes

#### Phase D: Manual Verification (Estimated: 1 day)
- [ ] Container lifecycle: run/stop/start/rm
- [ ] Image operations: pull/push/tag/load/save
- [ ] Delta create/apply cycle
- [ ] Networking: bridge, port publishing
- [ ] Exec into container
- [ ] All symlinks functional

#### Phase E: CI Pipeline (Estimated: 1 day)
- [ ] GitHub Actions workflows pass
- [ ] Cross-compilation for ARM works
- [ ] Release artifact generation

### Files Modified in This Session

```
cmd/balena-engine/main.go          # Busybox dispatcher (already existed)
hack/make/.binary                  # Binary naming and package
vendor.mod                         # Replace directives for forks
forks/balena-runc/main.go          # Export Main()
forks/balena-containerd/cmd/containerd/main.go
forks/balena-containerd/cmd/ctr/main.go
forks/balena-containerd/cmd/containerd-shim-runc-v2/main.go
forks/balena-containerd/api/                    # Extracted from api/v1.8.0 tag
forks/balena-containerd/remotes/docker/auth/fetch.go
forks/balena-containerd/remotes/docker/authorizer.go
forks/balena-containerd/sys/reaper/reaper_unix.go
forks/balena-containerd/pkg/cri/server/runtime_config.go
forks/balena-containerd/pkg/cri/sbserver/runtime_config.go
forks/balena-containerd/pkg/cri/sbserver/container_stats_list.go
forks/balena-containerd/pkg/cri/instrument/instrumented_service.go
forks/balena-containerd/go.mod     # Downgraded golang.org/x/* versions
```

### Risk Assessment for Remaining Work

| Task | Risk | Notes |
|------|------|-------|
| Delta patches | MEDIUM | Core functionality, may have API changes |
| Networking patches | HIGH | libnetwork changes between v23→v27 |
| Storage migration | LOW | Mostly isolated code |
| Swarm removal | LOW | Straightforward deletion |
| Test fixes | MEDIUM | Many tests may need updating for v27 API changes |
