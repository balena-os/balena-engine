# balena-engine v27 Port Guide

This document explains the work done to port balena-engine from moby v23.0.18 to v27.5.1, intended for reviewers and maintainers.

## Executive Summary

**What**: Port of 150 balena-specific commits from moby v23.0.18 to v27.5.1
**Why**: Keep balena-engine up to date with upstream security fixes, performance improvements, and new features
**Status**: Complete - all balena features ported, tested on device

### Quick Stats

| Metric | Value |
|--------|-------|
| Base version | moby v27.5.1 |
| Balena commits on top | 150+ |
| Original v23 balena commits | 211 |
| Commits merged upstream | ~60 |
| Binary size (static, stripped) | ~55MB (amd64) |
| Tested platforms | linux/amd64, linux/arm64 |

## Why v27 Instead of v28?

v28.x uses containerd v2.1.x which is a major architectural change requiring significant rework of the busybox-style binary integration. v27.5.1 uses containerd v1.7.24, making it a more manageable upgrade from the current v1.6.22.

## Key Changes from v23 to v27

### Architecture Changes

| Component | v23 | v27 | Impact |
|-----------|-----|-----|--------|
| containerd | v1.6.22 | v1.7.30 | New shim API, plugin changes |
| runc | v1.2.8 | v1.2.4 | Minor API changes |
| CLI | v23.0.16 | v27.4.0 | New flags, updated help text |
| Go version | 1.21.x | 1.22.x | Language features, stdlib changes |
| libnetwork | Separate module | Merged into moby | Import path changes |
| AUFS | Supported | Removed | Must use overlay2 |

### API Changes Affecting Balena Code

1. **Context parameters**: Many daemon functions now take `context.Context` as first parameter
2. **Logging**: Changed from `logrus` to `containerd/log` (`log.G(ctx)`)
3. **Runtime config**: `daemon.getRuntime()` → `daemonCfg.Runtimes.Get()`
4. **Layer store**: `containerfs.ContainerFS` → `string` for mount paths
5. **Plugin system**: Restructured, affects shim initialization

## Balena-Specific Features Preserved

All balena-specific features have been preserved in the v27 port:

### Core Features

| Feature | Files | Status |
|---------|-------|--------|
| **Delta updates** | `daemon/images/image_delta.go`, `distribution/xfer/` | ✅ Working |
| **Resilient pulls** | `distribution/pull_v2.go` | ✅ Working |
| **Container ID env var** | `daemon/create.go`, CLI `--cidenv` | ✅ Working |
| **Healthcheck restart** | `daemon/health.go` | ✅ Working |
| **Storage migration** | `pkg/storagemigration/` | ✅ Working |
| **Mobynit/hostapp** | `cmd/mobynit/` | ✅ Working |
| **Bare runtime** | `daemon/daemon_unix.go` | ✅ Working |

### Build System

| Feature | Description | Status |
|---------|-------------|--------|
| **Busybox binary** | Single binary dispatches to all components | ✅ Working |
| **Symlinks** | `balena-engine`, `balena-runc`, etc. | ✅ Working |
| **Static linking** | Fully static binary for portability | ✅ Working |
| **Cross-compilation** | ARM64, ARMv7, etc. | ✅ Working |

### Removed/Disabled Features

| Feature | Reason |
|---------|--------|
| Swarm mode | Not used in balenaOS |
| Plugin system | Not used in balenaOS |
| macvlan driver | Not typically used in IoT |
| overlay network driver | Not typically used in IoT (not overlay2 filesystem) |
| AUFS | Removed upstream in v27 |

## Commits Requiring Special Attention

### Busybox Binary Integration

The busybox-style binary required vendored patches to containerd:

**File**: `vendor/github.com/containerd/containerd/runtime/v2/runc/manager/manager_linux.go`
```go
// Changed os.Executable() to os.Args[0] to preserve symlink name
// Required for shim dispatch in busybox binary
self := os.Args[0]  // was: self, err := os.Executable()
```

**File**: `vendor/github.com/containerd/containerd/runtime/v2/shim/shim.go`
```go
// Filter plugins to only load shim-required ones
// Prevents nil pointer panic from uninitialized plugins
shimPlugins := map[string]bool{
    "io.containerd.internal.v1.shutdown": true,
    "io.containerd.event.v1.publisher":   true,
    "io.containerd.ttrpc.v1.task":        true,
}
```

### Seccomp Build Tags

**Files**: `Dockerfile`, `docker-bake.hcl`

The default `DOCKER_BUILDTAGS` must include `seccomp`:
```dockerfile
ARG DOCKER_BUILDTAGS="apparmor seccomp no_btrfs no_cri no_devmapper no_zfs ..."
```

Without this, containers fail with "seccomp: config provided but seccomp not supported".

### Logging Migration

Many files changed from logrus to containerd/log:
```go
// Before (v23):
logrus.WithField("key", value).Debug("message")

// After (v27):
log.G(context.TODO()).WithField("key", value).Debug("message")
```

## Testing

### Unit Tests
```bash
make test-unit  # 524 tests pass, 3 skipped
```

### Integration Tests
```bash
make test-integration TEST_FILTER=TestDelta  # Delta tests pass
make test-integration TEST_SKIP_INTEGRATION_CLI=1  # API tests pass
```

### Device Testing

Tested on Raspberry Pi 5 running balenaOS:
- Daemon starts correctly
- Existing containers continue running
- New containers can be created
- Seccomp works (no `--security-opt seccomp=unconfined` needed)
- Networking works (`balena0` bridge)
- Delta operations work

## Files Modified

### Key Files Changed for v27 Compatibility

```
cmd/balena-engine/main.go           # Busybox dispatcher
daemon/create.go                    # Container ID env var, logging
daemon/daemon_unix.go               # Bare runtime validation
daemon/health.go                    # Healthcheck restart
libnetwork/drivers_linux.go         # Disabled macvlan/overlay
api/types/container/hostconfig.go   # ContainerIDEnv field
vendor/github.com/docker/cli/...    # CLI --cidenv flag
```

### Vendored Patches (containerd)

```
vendor/github.com/containerd/containerd/runtime/v2/runc/manager/manager_linux.go
vendor/github.com/containerd/containerd/runtime/v2/shim/shim.go
```

### Build System

```
Dockerfile                          # DOCKER_BUILDTAGS default
docker-bake.hcl                     # DOCKER_BUILDTAGS default
hack/make/.binary                   # Binary naming
```

## Migration Notes for Users

### From v23 balena-engine

1. **AUFS users must migrate to overlay2** before upgrading
   - The `pkg/storagemigration/` code handles this automatically on first boot
   - Set `BALENA_MIGRATE_OVERLAY=1` environment variable

2. **API compatibility**: The Docker API version is now 1.47 (was 1.43)
   - Most clients will work unchanged
   - Very old clients may need updates

3. **Container ID env var**: Now available via CLI
   ```bash
   balena run --cidenv CONTAINER_ID alpine sh -c 'echo $CONTAINER_ID'
   ```

## Known Differences from v23

### Functional Parity Achieved

All v23 balena-specific features work in v27. The only differences are:

1. **AUFS removed**: Must use overlay2 (upstream change)
2. **Network drivers reduced**: macvlan/overlay disabled (intentional for IoT)
3. **API version**: 1.47 vs 1.43 (upstream change)

### Not Ported (Intentional)

| Commit | Description | Reason |
|--------|-------------|--------|
| `448ee8958a` | Container locks | Was regressive (introduced bug) |
| Various | AUFS-specific patches | AUFS removed in v27 |
| Various | v23 API compatibility | Not needed for v27 |

## Commit History

The v27 port preserves commit history for git bisect. Key commit groups:

1. **Original balena features** (eb0632625f and earlier): Delta, resilient pulls, healthcheck, etc.
2. **Busybox binary port** (3acc97c8b9): Initial v27 build working
3. **Test fixes** (170bb4b458): Integration test compatibility
4. **Device fixes** (7fe10db39b): Shim dispatcher, plugin filtering
5. **Seccomp fix** (4ce245eedc): Build tag restoration
6. **Cherry-picked patches** (8f7bd4cd86 - b420d8fd3f): Missing v23 patches
7. **CLI cidenv** (978ba95047): CLI flag for container ID env var
8. **Bare runtime** (9c87fc4b9a): HostConfig validation fix

## For Reviewers

### What to Focus On

1. **Vendored containerd patches**: These are critical for busybox binary
2. **Build tag changes**: Ensure seccomp is enabled
3. **API type changes**: `ContainerIDEnv` field addition
4. **Test compatibility**: Layer store test fixes

### What's Safe to Skim

1. **Documentation updates**: PRD.md, V27_PORT.md
2. **Import path changes**: Mostly mechanical
3. **Context parameter additions**: Required by v27 API

### Build Verification

```bash
# Quick build test
docker buildx bake --set '*.platform=linux/amd64' binary

# Full test
make test-unit
make test-integration TEST_FILTER=TestDelta
```

## Current State: Swarm Code Removal

### Why Swarm Was Removed

v27 upstream includes full swarm support which pulls in `github.com/moby/swarmkit/v2` (~5MB of vendor code). Since balena-engine doesn't use swarm mode, this code was stubbed/removed to reduce binary size, matching the approach taken in v23.

### What Was Changed

#### 1. Daemon Swarm Code Stubbed

**Files removed from `daemon/cluster/`:**
- All files except `provider/network.go` (type definitions only)
- Entire `executor/container/` directory
- Entire `convert/` directory
- Entire `controllers/` directory

**Files stubbed:**
- `daemon/cluster/executor/backend.go` - Reduced to only `ImageBackend` interface (no swarmkit import)
- `daemon/secrets.go` - Removed `SetContainerDependencyStore()` function and swarmkit import

**Directory deleted:**
- `libnetwork/cnmallocator/` - Container Networking Model allocator for swarm networking

#### 2. CLI Swarm Commands Removed

**Files modified:**
- `forks/balena-engine-cli/cli/command/commands/commands.go`
- `vendor/github.com/docker/cli/cli/command/commands/commands.go`

**Commands removed from CLI:**
- `config` - Swarm config management
- `node` - Swarm node management
- `secret` - Swarm secret management
- `service` - Swarm service management
- `stack` - Swarm stack management
- `swarm` - Swarm cluster management
- `checkpoint` - Container checkpoint (depends on swarm)
- `plugin` - Plugin management (not used)
- `context` - Docker context (not needed for IoT)

#### 3. Test Files Deleted

**Swarm-related test files removed:**
- `integration-cli/docker_cli_swarm_test.go`
- `integration-cli/docker_api_swarm_test.go`
- `integration-cli/docker_cli_service_health_test.go`

### Size Impact

| State | Binary Size (amd64) |
|-------|---------------------|
| Before stripping | ~70MB |
| After stripping (`-s -w`) | ~61MB |
| After swarm removal | ~55MB |

Total reduction: ~15MB (~21% smaller)

### Files Modified Summary

```
daemon/secrets.go                                    # Removed swarmkit import
daemon/cluster/executor/backend.go                   # Stubbed to ImageBackend only
daemon/cluster/provider/network.go                   # Kept (type definitions)
forks/balena-engine-cli/cli/command/commands/commands.go  # Removed swarm commands
vendor/github.com/docker/cli/cli/command/commands/commands.go  # Same
```

### Files/Directories Deleted

```
daemon/cluster/cluster.go
daemon/cluster/configs.go
daemon/cluster/errors.go
daemon/cluster/filters.go
daemon/cluster/helpers.go
daemon/cluster/listen_addr*.go
daemon/cluster/networks.go
daemon/cluster/noderunner.go
daemon/cluster/nodes.go
daemon/cluster/secrets.go
daemon/cluster/services.go
daemon/cluster/swarm.go
daemon/cluster/tasks.go
daemon/cluster/utils.go
daemon/cluster/volumes.go
daemon/cluster/convert/
daemon/cluster/controllers/
daemon/cluster/executor/container/
libnetwork/cnmallocator/
integration-cli/docker_cli_swarm_test.go
integration-cli/docker_api_swarm_test.go
integration-cli/docker_cli_service_health_test.go
```

### Verification

After swarm removal, verify no swarmkit in binary:
```bash
strings bundles/binary/balena-engine | grep -c "swarmkit"
# Should return 0
```

Verify CLI doesn't have swarm commands:
```bash
./bundles/binary/balena-engine --help | grep -E "swarm|service|stack|node|secret|config"
# Should return no matches
```

## References

- [Moby v27.5.1 Release Notes](https://github.com/moby/moby/releases/tag/v27.5.1)
- [containerd v1.7.30 Release Notes](https://github.com/containerd/containerd/releases/tag/v1.7.30)
- Original v23 branch: `kyle/rerun-rebase-v23.0.18`
- This port: `balena/v27-rebase`
