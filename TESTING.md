# Testing

This document contains the Moby code testing guidelines. It should answer any 
questions you may have as an aspiring Moby contributor.

## Test suites

Moby has two test suites (and one legacy test suite):

* Unit tests - use standard `go test` and
  [gotest.tools/assert](https://godoc.org/gotest.tools/assert) assertions. They are located in
  the package they test. Unit tests should be fast and test only their own 
  package.
* API integration tests - use standard `go test` and
  [gotest.tools/assert](https://godoc.org/gotest.tools/assert) assertions. They are located in
  `./integration/<component>` directories, where `component` is: container,
  image, volume, etc. These tests perform HTTP requests to an API endpoint and
  check the HTTP response and daemon state after the call.

The legacy test suite `integration-cli/` is deprecated. No new tests will be 
added to this suite. Any tests in this suite which require updates should be 
ported to either the unit test suite or the new API integration test suite.

## Writing new tests

Most code changes will fall into one of the following categories.

### Writing tests for new features

New code should be covered by unit tests. If the code is difficult to test with
unit tests, then that is a good sign that it should be refactored to make it
easier to reuse and maintain. Consider accepting unexported interfaces instead
of structs so that fakes can be provided for dependencies.

If the new feature includes a completely new API endpoint then a new API 
integration test should be added to cover the success case of that endpoint.

If the new feature does not include a completely new API endpoint consider 
adding the new API fields to the existing test for that endpoint. A new 
integration test should **not** be added for every new API field or API error 
case. Error cases should be handled by unit tests.

### Writing tests for bug fixes

Bugs fixes should include a unit test case which exercises the bug.

A bug fix may also include new assertions in existing integration tests for the
API endpoint.

### Writing new integration tests

Note the `integration-cli` tests are deprecated; new tests will be rejected by
the CI.

Instead, implement new tests under `integration/`.

### Integration tests environment considerations

When adding new tests or modifying existing tests under `integration/`, testing
environment should be properly considered. `skip.If` from 
[gotest.tools/skip](https://godoc.org/gotest.tools/skip) can be used to make the 
test run conditionally. Full testing environment conditions can be found at 
[environment.go](https://github.com/moby/moby/blob/6b6eeed03b963a27085ea670f40cd5ff8a61f32e/testutil/environment/environment.go)

Here is a quick example. If the test needs to interact with a docker daemon on 
the same host, the following condition should be checked within the test code

```go
skip.If(t, testEnv.IsRemoteDaemon())
// your integration test code
```

If a remote daemon is detected, the test will be skipped.

## Running tests

### Unit Tests

To run the unit test suite:

```
make test-unit
```

or `hack/test/unit` from inside a `BINDDIR=. make shell` container or properly
configured environment.

The following environment variables may be used to run a subset of tests:

* `TESTDIRS` - paths to directories to be tested, defaults to `./...`
* `TESTFLAGS` - flags passed to `go test`, to run tests which match a pattern
  use `TESTFLAGS="-test.run TestNameOrPrefix"`

### Integration Tests

To run the integration test suite:

```
make test-integration
```

This make target runs both the "integration" suite and the "integration-cli"
suite.

You can specify which integration test dirs to build and run by specifying
the list of dirs in the TEST_INTEGRATION_DIR environment variable.

You can also explicitly skip either suite by setting (any value) in
TEST_SKIP_INTEGRATION and/or TEST_SKIP_INTEGRATION_CLI environment variables.

Flags specific to each suite can be set in the TESTFLAGS_INTEGRATION and
TESTFLAGS_INTEGRATION_CLI environment variables.

If all you want is to specify a test filter to run, you can set the
`TEST_FILTER` environment variable. This ends up getting passed directly to `go
test -run` (or `go test -check-f`, depending on the test suite). It will also
automatically set the other above mentioned environment variables accordingly.

### Go Version

You can change a version of golang used for building stuff that is being tested
by setting `GO_VERSION` variable, for example:

```
make GO_VERSION=1.12.8 test
```

## Balena-Specific Testing

### Running Tests in Docker (Recommended)

The integration tests require a running daemon with proper socket access. The recommended approach is to run tests inside a Docker container:

```bash
# Quick test run for container tests
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

### Delta Tests

Delta functionality is a key balena-engine feature. Run delta tests specifically:

```bash
# Run delta tests
make test-integration TEST_FILTER=TestDelta

# Or inside Docker container
go test -v -timeout 30m ./integration/image/... -run TestDelta
```

Expected: 10 tests pass (TestDeltaCreate, TestDeltaSize, TestDeltaCorrectness, etc.)

### Test Categories

| Category | Command | Expected |
|----------|---------|----------|
| Unit tests | `make test-unit` | 524 pass, 3 skip |
| Delta | `TEST_FILTER=TestDelta make test-integration` | 10 pass |
| Container | `go test ./integration/container/...` | 135 pass |
| Daemon | `go test ./integration/daemon/...` | All pass |

### Known Test Skips/Failures

These are expected and not bugs:

1. **Plugin tests** - balena-engine doesn't support plugins (returns 404)
2. **Swarm tests** - swarm functionality removed from balena-engine
3. **Schema1 tests** - legacy image format, not used
4. **Some legacy integration-cli tests** - rely on shared daemon infrastructure

### Debugging Test Failures

1. **EOF errors during cleanup**: The shared daemon stopped. Tests that manage their own daemons can affect the shared daemon.

2. **"page not found" for plugins**: Expected behavior. The test framework's cleanup tries to list plugins.

3. **"daemon is not started"**: Run tests inside Docker container with daemon already running.

### Manual Verification

After running automated tests, verify key functionality manually:

```bash
# Inside dev container with daemon running
balena-engine pull alpine:latest
balena-engine run --rm alpine echo "Hello World"
balena-engine images

# Test delta
balena-engine pull alpine:3.18
balena-engine pull alpine:3.19
curl -X POST --unix-socket /var/run/balena-engine.sock \
  "http://localhost/images/delta?src=alpine:3.18&dest=alpine:3.19&t=test-delta"
balena-engine images  # Should show test-delta image smaller than alpine:3.19
```
