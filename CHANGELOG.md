# Changelog

## 17.06+rev1 (2017-10-13)

### Builder

- Allow bind-mounting a volume in the build context [#27](https://github.com/balena-os/balena-engine/pull/27)

### CLI

- Add command for generating image deltas [#35](https://github.com/balena-os/balena-engine/pull/35)

### Daemon

- Add utility that can boot a system into a container [#10](https://github.com/balena-os/balena-engine/pull/10)
- Add the ability to create binary delta between two images [#11](https://github.com/balena-os/balena-engine/pull/11)
- Include engine name in version information [#32](https://github.com/balena-os/balena-engine/pull/32)
- Minimize page cache usage during pull [de0993b](https://github.com/balena-os/balena-engine/pull/12/commits/de0993b9c1ab9408f2716b147dc544b37bb1deb7)

### Plugins

- Disable plugin support [#14](https://github.com/balena-os/balena-engine/pull/14)

### Logging

- Disable awslogs, fluentd, gcplogs, gelf, logentries, splunk, and syslog logging drivers [fe4d45c](https://github.com/balena-os/balena-engine/pull/7/commits/fe4d45c5dcbaddff51aebfe584a68e8fb9f44449)

### Runtime

- Disable consul, etcd, and zookeeper discovery backends [380ba69](https://github.com/balena-os/balena-engine/pull/7/commits/380ba69d00cb56278b65b28571d4d7e000392ec3)

### Swarm Mode

- Disable swarm mode [#14](https://github.com/balena-os/balena-engine/pull/14)

### Distribution

- On-the-fly layer extraction during docker pull [#8](https://github.com/balena-os/balena-engine/pull/8)
- Improve layer reuse when pushing to v1 registries [#28](https://github.com/balena-os/balena-engine/pull/28)
- Include a combined progress bar of all the layers in the output [#31](https://github.com/balena-os/balena-engine/pull/31)
