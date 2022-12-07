![balenaEngine](./docs/static/balena-engine.svg)

**moby-based container engine for IoT**

## Highlights

- __Small footprint__: 3.5x smaller than Docker CE, packaged as a single binary
- __Multi-arch support__: Available for a wide variety of chipset architectures, supporting everything from tiny IoT devices to large industrial gateways
- __True container deltas__: Bandwidth-efficient updates with binary diffs, 10-70x smaller than pulling layers
- __Minimal wear-and-tear__: Extract layers as they arrive to prevent excessive writing to disk, protecting your storage from eventual corruption
- __Failure-resistant pulls__: Atomic and durable image pulls defend against partial container pulls in the event of power failure
- __Conservative memory use__: Prevents page cache thrashing during image pull, so your application runs undisturbed in low-memory situations

## Motivation

balenaEngine is a container engine purpose-built for embedded and IoT use cases
and compatible with Docker containers. Based on Docker’s Moby Project, balenaEngine
supports container deltas for 10-70x more efficient bandwidth usage, has 3x
smaller binaries, uses RAM and storage more conservatively, and focuses on
atomicity and durability of container pulling.

Since 2013, when we [first ported Docker to ARMv6 and the Raspberry Pi](https://www.balena.io/blog/docker-on-raspberry-pi/),
the balena team has been working in and around the Docker codebase.
Meanwhile, having seen IoT devices used in production for tens of millions of
hours, we’ve become intimately acquainted with the unique needs of the embedded world.
So we built a container engine that runs Docker containers just as well,
shares the Docker components that are needed for our use case, and is augmented
with the IoT-specific features that we’ve built out over time.

## Transitioning from Docker CE

We left out Docker features that we saw as most needed in cloud deployments and
therefore not warranting inclusion in a lightweight IoT-focused container
engine. Specifically, we’ve excluded:

- Docker Swarm
- Cloud logging drivers
- Plugin support
- Overlay networking drivers
- Non-boltdb discovery backends (consul, zookeeper, etcd, etc.)
- Buildkit (although support can be enabled using a build tag)

Unless you depend on one of the features in Docker that balenaEngine omits, using
balenaEngine should be a drop-in replacement.

## License

balenaEngine is licensed under the Apache License, Version 2.0. See
[LICENSE](https://github.com/balena-os/balena-engine/blob/master/LICENSE) for the full
license text.
