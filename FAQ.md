# FAQ

## What is the Moby Project?

As described by Docker, the Moby Project is “an open framework to assemble specialized container systems.” You can read more at [the moby project website](https://mobyproject.org/) and [moby project introduction post](https://blog.docker.com/2017/04/introducing-the-moby-project/).

## Why containers for IoT?

We think that containers are essential to bringing modern development and deployment capabilities to connected devices. Linux containers, particularly Docker, offer, for the first time, a practical path to using virtualization on embedded devices, without heavy overhead or hardware abstraction layers that get in the way.

## Why not just use Docker’s container engine on my device?

Docker was primarily designed for datacenters with large, homogenous, well-networked servers. As such it makes tradeoffs that in some cases come in conflict with the need of small, heterogenous, remotely distributed, and differentiated devices, as found in IoT and embedded Linux use cases.

With its small footprint and purpose-built features, balenaEngine was made specifically for IoT devices, or any scenario where footprint, bandwidth, power, storage, etc. are a concern.

## What specific Docker features are excluded from balenaEngine?

We left out Docker features that we saw as most needed in cloud deployments and therefore not warranting inclusion in a lightweight IoT-focused container engine. Specifically, we’ve excluded:
- Docker Swarm
- Cloud logging drivers
- Plugin support
- Overlay networking drivers
- Non-boltdb backed stores (consul, zookeeper, etcd, etc.)
