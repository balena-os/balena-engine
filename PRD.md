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
