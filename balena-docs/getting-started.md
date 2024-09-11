# Getting started with balenaEngine

Using balenaEngine (`balena-engine`) should feel very natural to docker users
since it shares the same command structure.

**Pulling images**

```bash
balena-engine pull resin/rpi-raspbian:jessie
```

**Running containers**

```bash
balena-engine run -it --rm resin/rpi-raspbian:jessie /bin/bash
```

## Container deltas

balenaEngine comes with support for container deltas, a way of computing a binary
description of what changed between two images. A delta can be pushed to the
standard docker registry, it has the same format as a docker image! Let's see
how this works.

Deltas are computed for a **target** image against a **base** image. For a
balenaEngine instance to be able pull the **target** using deltas it must have the
**base** already. balenaEngine will make sure this is the case and bail out if it
doesn't have the appropriate data to apply the delta.

Unlike layer based pulling, with container deltas if any part of any file of
any layer is shared between the two images it will be referenced instead of
being transferred. This amounts to huge bandwidth savings that range between
10-70x for most cases.

### Using deltas

**Creating a container delta**

```bash
balena-engine image delta resin/raspberrypi3-node:6 resin/raspberrypi3-node:7
```

The image ID of the delta will be printed at the end of the command. You can
now manipulate the delta image like a normal image. Of course, you won't be
able to run it since it doesn't contain a complete filesystem. Let's push our
delta to the registry! But first, let's give a name to our delta.

```bash
balena-engine tag <delta image id> resin/raspberrypi3-node:delta-6-7
```

**Pushing with deltas**

```bash
balena-engine push resin/raspberrypi3-node:delta-6-7
```
That's it! Deltas are like normal images as far as the registry is concerned.

**Pulling with deltas**

Now, let's delete the delta and target from the local daemon and try to pull using deltas.
```bash
balena-engine rmi -f resin/raspberrypi3-node:delta-6-7 resin/raspberrypi3-node:7
```
And now let's pull the delta.
```bash
balena-engine pull resin/raspberrypi3-node:delta-6-7
```

That's it! balenaEngine understands that the image it tries to download is a delta
image and switches to delta application mode automatically.

After pulling is complete you should end up with an image with the same image
id as `resin/raspberrypi3-node:7`. Let's verify that by also pulling
`resin/raspberrypi3-node:7`. It should be a no-op.

```bash
balena-engine pull resin/raspberrypi3-node:7
```
