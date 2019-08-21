# Patches on top of moby

```
87ef8e0dca4281eac019704c4fd2f1a95c06824a 2019-06-26T11:49:23+03:00 v18.9.7
178650445602475aeee1eab8076c5a83f9b6ed16 2019-06-25T14:31:02+02:00 contrib/install.sh: Improve error output
8eb3f4af0e0b1b70d77cc3e456458fb430027b76 2019-06-25T14:31:02+02:00 contrib/install.sh: Add details to the success message
44ace080a685a4c04107de18fcd430d2ada7949e 2019-06-25T14:31:02+02:00 contrib/install.sh: Rename balena to balenaEngine in ASCII art output
64cb87f1e17b195836a8722cdfd6360a65fd6a91 2019-06-25T14:31:02+02:00 contrib/install.sh: Fail on error
95c7371304f9cef494efe93f0a8ffd53a75eac21 2019-06-03T18:22:01+03:00 v18.9.6
fa1d0b6490f9ecd1d5525bd5208522ab8b6713a5 2019-06-03T14:49:33+02:00 Bump containerd/cgroups to dbea6f2bd41658b84b00417ceefa416b97
7cab3339d1c1e4ad39b5baa49f4a38d5c1eb1ad5 2019-04-29T20:02:29+03:00 v18.9.5
c88f6f82b644c6c0b0b7932d980a86542be411b1 2019-04-26T19:19:29+02:00 Add daemon flags to configure max download/upload attempts during pull/push
cf02022c246852e4503e5f63f30cea96350feaff 2019-04-26T14:16:30+02:00 v18.9.4
5e472e9eb556ac6aeb4e0346551d12b652c62a7f 2019-04-26T14:13:59+02:00 integration-tests: Skip aufs test, doesn't work with dind
4c3cafd8c56102bcf83773dcb50a597e944d382a 2019-04-26T14:13:59+02:00 integration-tests: Add image pull tests
0ef2d9dc7c6456b97e3b019f5110b6a67603c76c 2019-04-26T14:13:59+02:00 aufs,overlay2: Add driver opts for disk sync
44824fbca7f7d95f2716367bf63515dae0cd62e5 2019-04-26T14:13:53+02:00 v18.9.3
5bf4b8087e14af7d7af83402ebfa7b00fa416648 2019-04-18T15:19:39+02:00 dockerfile: Rename docker-init to balena-engine-init
beb7f70265dd873c1e674ae65fefcf9bb3e372b4 2019-04-18T15:19:39+02:00 integration-tests: Skip tests relying on swarm,plugin support
49810dc78e6865f2305969616245160e4df9460d 2019-04-09T18:10:58+02:00 Fix double locking in the event handling code of OOM events
9eca1531dae212ac85daa658c19fa652ed085364 2019-04-08T19:30:49+02:00 integration-tests: Add test for containers with memory,cpu constraints
2c29eccd53f91349fc8896d0521f5a6e9712a181 2019-04-08T19:29:22+02:00 Update Dockerfiles used for build to Go 1.10.8
2189e871a666ac5263425f831b495ec932c14ac0 2019-04-08T19:29:22+02:00 delta: Move implementation under ImageService
b43966d4a90590a78cddaf24326be0adadd1e1e9 2019-04-08T19:29:22+02:00 builder-next: Implement xfer.DownloadDescriptor for layerDescriptor
0601a3917f796ab28cb5fd34d721e291f4a51713 2019-04-08T19:29:22+02:00 Add 18.09.3 changelog from upstream
3a3129812cf193a4bd167a035fecad2892330a00 2019-04-08T19:29:22+02:00 Fixes after update to 18.09
93ba024772c0fa3f7b58966d19130907e706b162 2019-04-08T19:28:07+02:00 vendor: Update dependencies to 18.09
e284726bae2fe269d721f67e6a74171289ad551b 2019-03-18T20:07:50+02:00 v17.13.4
52245090150bb56a5b27f45666ea66964836531e 2019-03-18T17:36:05+01:00 Fix event filter filtering on "or"
fe78e2c9a69313007c53c83fff4b5525fbc2ba45 2019-02-25T15:11:03+01:00 v17.13.3
ad1ae964378edc3d61be6488ff5f973d7228edb4 2019-02-25T12:25:27+01:00 vendor: Update runc to include fix for opencontainers/runc#1766
613c2f27ed2e7d65474c2f3e786d9e24e757d99d 2019-02-25T10:44:48+01:00 Windows: Bump busybox to v1.1
ffeebb217c1da556c2dc686fbff80cbb5d74e564 2019-02-23T01:49:38+01:00 Update containerd runtime 1.2.4
c7fca75c035ba0b750f46a9676a376f8e4409f15 2019-02-23T01:49:25+01:00 Update runc to 6635b4f (fix CVE-2019-5736)
88330c9aac5556d0abc7a5afcb4d906604a07fa2 2019-02-23T01:49:12+01:00 Revert "Merge pull request #240 from seemethere/bundle_me_up_1809"
0841c61862e66c47ec735ad9de0039b806de37aa 2019-02-22T13:22:29+01:00 Fix nil pointer derefence on failure to connect to containerd
a7b3bbea5f4f775818641ec3ba522ca912f641c2 2019-02-21T15:33:03+01:00 v17.13.2
7486436688c818c256c6584baac9055bf3178bb1 2019-02-21T15:01:57+01:00 travis: Only run builds for PRs and master/version branches
fa432f08380ff5921ac75346586ec4d51bde9fb2 2019-02-21T15:01:57+01:00 travis: Build for armv7 and aarch64 as well
be3cfa585b98a425b053d3d3475fbb7a4793d26d 2019-02-21T15:01:57+01:00 travis: Use the minimal machine
a010e2e2bcd935046d6bcf20aa69e1c93de1c8fc 2019-02-21T10:21:55+01:00 v17.13.1
278f1a130b66de23f07e472792b70b640f777666 2019-02-20T13:51:17-08:00 Disabled these tests on s390x and ppc64le: - TestAPISwarmLeaderElection - TestAPISwarmRaftQuorum - TestSwarmClusterRotateUnlockKey
d6fc6068c2e5d02e0d6ddcfb5ee5b8a73ddf23f8 2019-02-20T18:27:58+01:00 vendor: Update runc to include fix for CVE-2019-5736
3744b45ba8ad93f1a21cbc80420856b04efc4593 2019-02-20T14:31:18+01:00 Graphdriver: fix "device" mode not being detected if "character-device" bit is set
a818442de73b16d8ad756c74e5e660d132e97848 2019-02-20T13:31:11+01:00 Makes a few modifications to the name generator.
19e733f89f7652f58b567b5178bacc10ef2940b5 2019-02-20T11:27:07+01:00 Fix: plugin-tests discarding current environment
e9ecd5e486c591979e34782025bb849f7faf8eba 2019-02-19T22:35:50+01:00 make test-integration: use correct dockerd binary
7b9ec00eec7ffe745ebd2f807daa50d84b3e10e7 2019-02-18T11:19:49+01:00 hack: no need to git fetch in CI
748f37022df465c39a76461c4970f4c678f629e3 2019-02-18T11:19:23+01:00 Allow overriding repository and branch in validate scripts
1d0353548a21b0b60595ad44859a7072f90ea6c6 2019-02-15T01:01:52+01:00 Delete stale containerd object on start failure
d158b9e74fffe377b6acc7794246c89f9fb26d2f 2019-02-12T00:13:23+01:00 pkg/archive: fix TestTarUntarWithXattr failure on recent kernel
317e0acc4e602f978e4d9c0130a113d179026c8e 2019-02-11T22:12:52+01:00 keep old network ids
325f6ee47a8edaf093ea9f829c26962310c83759 2019-02-09T11:05:52+01:00 [18.09] Bump Golang 1.10.8 (CVE-2019-6486)
c51d247f030051abb4b97770d49bac30343e45c8 2019-02-09T11:04:09+01:00 Ignore xattr ENOTSUP errors on copy (fixes #38155)
03dfb0ba53cc5f64b746a25aa5ed8a48763ea223 2019-02-06T00:25:54+00:00 Apply git bundles for CVE-2019-5736
9e5d8ad10bae6db6737944dbdcdf2b2486a84593 2019-01-31T14:35:39+01:00 v17.13.0
744c524d03a34ac36cb94163db89ca93e6e7d503 2019-01-14T13:00:29+01:00 Add cli for tagging delta images
594b651faad03a87ce7f95b0608f2f2fc2c38af5 2019-01-14T13:00:29+01:00 Add delta integration test
bd66b913256fa6dfee063cd4e8653d6b8fc63a4f 2019-01-14T13:00:29+01:00 Allow tagging of image deltas on creation
dceb2fc48071b78a8a828e0468a15a479515385f 2019-01-14T13:00:29+01:00 Backport fixes for hanging on low-entropy situations
377d55202abb397dd7b71034bfc87e2df2f5d414 2019-01-14T13:00:22+01:00 v17.12.1
44bda6ccd8ca4988e9ba8dedeb2df16a0c0d9cf2 2019-01-14T12:26:53+01:00 Enable travis-ci
acf966f17d0c263adf48f44b0a663b974b8ecb81 2019-01-14T12:26:53+01:00 Add repo.yml
a79fabbfe84117696a19671f4aa88b82d0f64fc1 2019-01-09T17:31:53+00:00 If url includes scheme, urlPath will drop hostname, which would not match the auth check
fc274cd2ff4cf3b48c91697fb327dd1fb95588fb 2019-01-09T17:31:53+00:00 Authz plugin security fixes for 0-length content and path validation Signed-off-by: Jameson Hyde <jameson.hyde@docker.com>
f80c6d7ae15443f15c014ddbd17e30cfac47b906 2019-01-08T02:45:06+01:00 Bump containerd to v1.2.2
e042692db1316a60be35bfdca10d7e08d20f50ad 2019-01-05T09:53:31+01:00 Skip kernel-memory tests on RHEL/CentOS daemons
24f71e39980e8a4c6eabcea16e0a9efce1660bbe 2018-12-28T09:40:26-08:00 Revert "[18.09 backport] API: fix status code on conflicting service names"
6646d0878247b1e0875da33da606283f5d16ea07 2018-12-17T12:07:32+00:00 libcontainerd: prevent exec delete locking
a9ae6c7547466f754da01a53c6be455c555e6102 2018-12-17T12:06:35+00:00 Revert "Propagate context to exec delete"
b6430ba41388f0300ceea95c10738cbe1a9a7b10 2018-12-14T22:54:46+00:00 Propagate context to exec delete
d161dfe1a36929a03ee3dfa916d296abfd4ccef1 2018-12-14T22:47:44+00:00 Update containerd to aa5e000c963756778ab3ebd1a12c6
8afe9f422dc0183ce48e1db09189ccbde634080a 2018-12-14T00:44:49+01:00 Bump Golang 1.10.6 (CVE-2018-16875)
3482a3b14a6414977bd9860c513abf06dedd6bf7 2018-12-12T18:12:01+01:00 registry: use len(via)!=0 instead of via!=nil
55a4be8cf5787f88c1ce6bc3e2ec73402c964e5d 2018-12-12T01:59:01+01:00 Swagger: fix definition of IPAM driver options
1043f40fb561ffbf23fbdde9989abcebd8e48279 2018-12-11T15:15:20+01:00 fixes display text in Multiple IDs found with provided prefix
43dedf397503532f8bc1313af82a9356ed8c3550 2018-12-10T13:03:13+00:00 vendor: update buildkit to d9f75920
a69626afb12eb9ec3e374aa563b561c0ba28f27f 2018-12-10T12:21:26+01:00 Add test for status code on conflicting service names
ad7105260f3c2ff32a375ff78dce9a96e01d87cb 2018-12-10T12:18:32+01:00 Update swarmkit to return correct error-codes on conflicting names
b66c7ad62ebff12112318db0d1ed48b14f817efa 2018-12-07T18:58:03+01:00 use empty string as cgroup path to grab first find
5cd4797c89383159219716de92775138c2dd17c5 2018-12-07T18:57:54+01:00 vndr libnetwork to adjust for updated runc
7dfd23acf1d604cde70d46e36bfe2e51df2dcf46 2018-12-07T18:45:14+01:00 update containerd to v1.2.1
2c64d7c858b5fd16d0d14d692a3765cf8d5a5d7e 2018-12-07T11:20:22+01:00 update just installer of containerd to 1.2.1
00ad8e7c5730f3c50ae2e548b47d1340202f72b2 2018-11-30T20:43:05+01:00 Bump Go to 1.10.5
5fffdb32261145b1178f571e25fbd71572769d58 2018-11-30T14:57:51+01:00 Masked /proc/asound
9c93de59da8eaa0e0e0377578a33b465b9465bb2 2018-11-30T01:38:11+01:00 Windows:Tie busybox to version
73911117b30ba74c42177b0f002bb4e98f2473f9 2018-11-29T09:15:15-08:00 builder: delete sandbox in a goroutine for performance
8fe3b4d2ec06720fedfce2c2ef1b4fd7940961a6 2018-11-29T09:15:00-08:00 builder: set externalkey option for faster hook processing
850fff5fc7f033e76ef0eec04eb98384ddf2065d 2018-11-21T14:10:01-08:00 vendor: update buildkit to v0.3.3
0d17f4099496ba3de583273eeebefa3cce4694ca 2018-11-21T14:09:31-08:00 builder: avoid unset credentials in containerd
34867646af1d1344b1f0877bb3a00a51f7700245 2018-11-21T14:08:18-08:00 builder: ignore `label` and `label!` prune filters
0b2d88d328ca88c8732dc11c72873b53be3bd2f8 2018-11-21T14:08:04-08:00 builder: deprecate prune filter `unused-for` in favor of `until`
4cc45d91eb44cbaddcfd75335c3f72ede035c440 2018-11-21T22:15:18+01:00 Ignore default address-pools on API < 1.39
67c602c3fe3d5d817fb7210e50d7ed1688b28801 2018-11-21T22:13:56+01:00 apparmor: allow receiving of signals from 'docker kill'
db7f375d6a2aaf6d79f5c690e2f302c640bdde04 2018-11-21T21:59:41+01:00 Update containerd to v1.2.1-rc.0
7d6ec38402f4e2a5e1c83a981a88bf1a5f202858 2018-11-21T21:59:33+01:00 wip: bump containerd and runc version
64a05e3d162b7234f8a7aa32d10434db4c5e6364 2018-11-21T21:59:27+01:00 Bump containerd binary to fix shim hang
262abed3d2e84756e16c73c7c241aa62918c51c8 2018-11-21T21:59:20+01:00 Update runc to 58592df56734acf62e574865fe40b9e53e967910
e137337fe6083da91fd6d83d699cff3a857f636e 2018-11-21T21:59:13+01:00 Update containerd to v1.2.0
c9c87d76d651d57d72e52c575a2c9600170b5212 2018-11-21T21:59:06+01:00 Add a note about updating runc / runc vendoring
a4decd0c4cd6033907fe85576a3d7dc8990aa758 2018-11-21T21:58:58+01:00 Update containerd to v1.1.4
25bec4665b6f011e52f7b2765ba1579c7430481d 2018-11-20T18:08:44+01:00 Add CONFIG_IP_VS_PROTO_TCP, CONFIG_IP_VS_PROTO_UDP, IP_NF_TARGET_REDIRECT to check-config.sh
56cc26f927e6a1de51731f88baf5b0af3a5688bc 2018-11-20T15:50:46+01:00 Add missing default address pool fields to swagger
8486ea11ae800a1e6d634b741dfb007ba29f6003 2018-11-12T15:51:52+01:00 runc.installer: add nokmem build tag for rhel7 kernel
5b8cee93b5b6a2449d9af225e17d85c612f64ed2 2018-11-12T15:51:44+01:00 Bump runc
49556e047059d81b64f6cd12f4c602c9a9c471c7 2018-11-12T11:44:37+01:00 client: use io.LimitedReader for reading HTTP error
02fe71843e0e45ddc986a6c8182370b042349a27 2018-11-09T23:31:49+01:00 Windows: DetachVhd attempt in cleanup
757650e8dcca87f95ba083a80639769a0b6ca1cc 2018-11-08T15:26:01+01:00 awslogs: account for UTF-8 normalization in limits
9e06a421234f0bba8392b9a8908a94ff74f0c254 2018-11-08T14:01:27+01:00 API: properly handle invalid JSON to return a 400 status
e8eb3ca4eef1c913563195787eb1f2527c9febf3 2018-11-08T14:00:20+01:00 Enable volume tests on Windows
673f04f0b1afdb0a1f739b9a4a8b41ea1e015ea4 2018-11-08T14:00:14+01:00 Integration test: use filepath.Join() to make path cross-platform
65bf95f3df84de5901479091685715f227c333ce 2018-11-08T14:00:08+01:00 Some improvements to TestVolumesInspect
9fc9c3099da40b21fed9adc758f8787dbd3cedfd 2018-11-08T13:56:04+01:00 Renamed windowsRS1.ps1 to windows.ps1
37cb9e73006acd13b9708cd594ebc25054fef666 2018-11-08T13:55:59+01:00 Enabling Windows integration tests
59be98043a02f44b63b26f159461fed08292e027 2018-11-08T13:55:48+01:00 Windows: Start of enabling tests under integration/
45654ed0126aadaf6c3293b0a32ca8cf15021626 2018-11-06T10:52:34-08:00 builder: update copy to 0.1.9
e1783a72d1b84bc3e32470c468d14445e5fba8db 2018-11-06T12:39:04+01:00 [18.09 backport] update libnetwork to fix iptables compatibility on debian
c27094289aadaad4ad4d78aefcc44e95278d3508 2018-11-06T11:03:22+01:00 update containerd client and dependencies to v1.2.0
0afe0309bd9580bc76496c9e0da75216795c1c01 2018-11-06T11:03:14+01:00 bump up runc
41f3cea42f2f45244051b7f829c0ef9c27383c26 2018-11-06T11:03:06+01:00 Vendor Microsoft/hcsshim @ v0.7.9
9cf6464b639aef09df9b16999cc340a1276d0bf8 2018-11-06T11:03:00+01:00 LCOW: ApplyDiff() use tar2ext4, not SVM
52a3c39506b883f713694ce39d1a4fd9f5638800 2018-11-05T22:59:24+00:00 builder: fix bugs when pruning buildkit cache with filters
46dfcd83bf1bb820840df91629c04c47b32d1e21 2018-10-31T16:04:51-07:00 [18.09] Vendor swarmkit to 6186e40fb04a7681e25a9101dbc7418c37ef0c8b
c40a7d393bca990e07973024e71034b4b6fc05e5 2018-10-31T14:25:40+01:00 Fix double "unix://" scheme in TestInfoAPIWarnings
6ca0546f2571cf4acdc1f541bccfac23a78cb8d2 2018-10-30T23:04:27+00:00 cluster: set bigger grpc limit for array requests
64b0c76151ceb7b26f9c7477f3044dac747d227b 2018-10-30T14:25:19+01:00 Add more API doc details on service update version.
5591f0b1ee7dec101b490228258613cd7caf64ee 2018-10-30T11:29:02+01:00 Fix incorrect spelling in error message
72368177254811e5816f03a4773deaafb9df5202 2018-10-26T12:14:01+02:00 Bump Go to 1.10.4
5853cd510c3272755ca5d6605ca8039d54a5ba15 2018-10-24T20:11:51+02:00 builder: fix duplicate mount release
7a8d0d21c9e047852f81cac8f6eeacfe565fa00c 2018-10-23T23:26:41+01:00 docs: Fix Docker capitalisation in balenaEngine docs
2fb1f862d65a82617ce42290f76c460fe079ea14 2018-10-23T20:57:22+01:00 Update balenaEngine logo in README.md
6ee7d86a12fe83953eff0efd4de5878b4ff6814d 2018-10-23T13:37:15+02:00 Add note that we use the bump_v18.09 branch for SwarmKit
ae6284a623bac86ac6ab718fa4a369dd8c0a3cfc 2018-10-23T13:20:45+02:00 testing: add case for exec closeStdin
1222a7081ac9ebb0830a6c8008142258c49800b5 2018-10-22T15:10:20-05:00 Bump swarmkit
8a1313357a3c41a33a72c97fa5cfa180918e4ba6 2018-10-19T15:41:50+02:00 Update CHANGELOG.md
fd1fe0b702571865cc77d66937e4ca570b5b9cc3 2018-10-18T10:52:57-04:00 Bump libnetwork to 6da50d19 for DSR changes
fdaf08a57b2348623f33e0b9855c488421fc7bf6 2018-10-17T17:54:13-07:00 builder: fix private pulls on buildkit
63f30e90f9bd709821dd37412a68220e3dcf495b 2018-10-17T16:59:26+02:00 Update install.sh script for v17.12.0 release
25755b07fe92e331870307de4ab3641c88abe774 2018-10-17T12:32:20+02:00 Use Balena's fork of golang.org/x/sys/unix (ARM SyncFileRange syscall)
3f8ae5574074f1e9cd6687e085c32cd77718c929 2018-10-16T17:45:26+02:00 Have the balena-engine binary accept being called as balena, balenad...
a9c160256fee050d7021a9631dd3fe11399a8620 2018-10-15T23:38:26+02:00 Run vndr tool for balena-cli and balena-libnetwork (balenaEngine rename)
fa8ac946165b8004a15e85744e774ed6ba99fd38 2018-10-12T09:29:38-07:00 btrfs: ensure graphdriver home is bind mount
2199ada691dc635cac5cdd065d909a539dd0b793 2018-10-12T09:29:38-07:00 pkg/mount: add MakeMount()
fd7611ff1f1d61d5b4b45b2c0bd83976cbccf174 2018-10-12T09:29:38-07:00 pkg/mount: simplify ensureMountedAs
c20e8dffbbb4dd328e4c00d18e04eecb80cf4d4e 2018-10-12T02:26:17+02:00 Deprecate legacy overlay storage driver, and add warning
734e7a8e55e41b064025273776dc5a76e45bf2e1 2018-10-12T02:25:39+02:00 Deprecate "devicemapper" storage driver, and add warning
dbfc648a94569d8dbc8c6468d56ec93559363bb0 2018-10-11T20:35:43+00:00 builder: treat unset keep-storage as 0
690e097fedd7362f3b2781c32ca872ad966d286e 2018-10-11T22:09:38+02:00 overlay2: use index=off if possible
dc0a4db7c9dc593a8568a8e30e4e21e118c2839d 2018-10-11T22:09:30+02:00 overlay2: use global logger instance
f58f8421433d18e0fb9a51567068a2ddc1b13a1b 2018-10-11T21:55:49+02:00 bump buildkit to c7bb575343df0cbfeab8b5b28149630b8153fcc6
40c33e33d23882936ec402b1563fca7456af001d 2018-10-11T18:33:37+01:00 Fix daemon/cluster/executor/container/ unit tests
7184074c0880c656be00645007588a00ec2266cd 2018-10-11T16:04:45+02:00 Windows:Allow process isolation
b40c26d2aec12f21f46aeb96ae8c170e96809f6a 2018-10-11T09:18:22+01:00 Rename balena to balena-engine (executables) or balenaEngine (project)
6679a5faeb724f1ad060f2fdf6d189f1005924b9 2018-10-10T20:43:14+02:00 bugfix: wait for stdin creation before CloseIO
90c72824c36369efd8be52bedd731d12b3415508 2018-10-11T03:01:18+09:00 bump up buildkit
aa6df9901d7d439f2c980ef5b25e24b815d2ee44 2018-10-08T10:10:16+01:00 Disable incompatible integration tests
7b54720ccbfa5d8242e896f27e8b36ee58612401 2018-10-05T18:01:10+00:00 Switch copy image to a docker org based one
0922d32bce74657266aff213f83dfa638e8077f4 2018-10-05T15:13:43+02:00 Fix denial of service with large numbers in cpuset-cpus and cpuset-mems
148d9f0e58bc180fefffcfc0a9e7a00b4276a67a 2018-10-05T14:53:33+02:00 Update containerd client and dependencies to v1.2.0-rc.1
5070e418b806cc96ad0f5b3ac32c8d416ff8449a 2018-10-05T14:38:34+02:00 Update containerd dependencies
054c3c2931cec5dca8bb84af97f1457c343ec02f 2018-10-05T12:35:59+02:00 Remove version-checks for containerd and runc
9406f3622d18a0d9b6c438190e8fdd8be53d3b22 2018-10-04T21:59:25+02:00 Fix for default-addr-pool-mask-length param max value check
9816bfcaf58a609d64d648043c10817c27dcfa36 2018-10-04T21:59:09+02:00 Global Default AddressPool - Update
58e51512704b6d7656952e140332472a4c37e46f 2018-10-04T21:20:54+02:00 Masking credentials from proxy URL
54bd14a3fe1d4925c6fa88b24949063d99067c07 2018-10-03T15:24:34+02:00 Fix long startup on windows, with non-hns governed Hyper-V networks
c9ddc6effc444c54def41d498b359a9a986ad79d 2018-10-03T14:01:13+02:00 gd/dm: fix error message
16836e60bc87abb3e9ab16f33c2038931c1d473b 2018-10-02T20:33:38+02:00 Move the syslog syscall to be gated by CAP_SYS_ADMIN or CAP_SYSLOG
b499acc0e834e11882909269238407c65f68f034 2018-09-28T14:35:55+02:00 Tweak bind mount errors
67541d5841e645f3408b01f189ec4339df449edc 2018-09-27T22:46:57+00:00 vendor buildkit to 8f4dff0d16ea91cb43315d5f5aa4b27f4fe4e1f2
3e2973d26934bd22c46f07764afb1ed8b11bf6a1 2018-09-26T17:58:10+00:00 mobynit: Add support to mount rootfs from a custom location
6bf8dfc4d89461228031a595d63482b9603c8899 2018-09-25T23:09:25+00:00 fix daemon tests that were using wrong containerd socket
e090646d477f2e7d00aba971bcc187f3af7948a3 2018-09-25T23:09:25+00:00 hack/make: remove 'docker-' prefix when copying binaries
b3bb2aabb8ed5a8af0a9f48fb5aba3f39af38e0d 2018-09-24T22:35:36+00:00 Remove 'docker-' prefix for containerd and runc binaries
cce1763d57b5c8fc446b0863517bb5313e7e53be 2018-09-22T01:24:11+00:00 vendor: remove boltdb dependency which is superseded by bbolt
3d67dd046539f8e04db82ce07ea56f97b832676b 2018-09-21T17:06:25-07:00 builder: vendor buildkit to 39404586a50d1b9d0fb1c578cf0f4de7bdb7afe5
73e2f72a7c5bd6d6f8306e0ffe4371e1c3b00a21 2018-09-21T17:06:25-07:00 builder: use buildkit's GC for build cache
2926a45be6b9315d2ddeec27d1193278b6bbae91 2018-09-21T17:06:25-07:00 add support of registry-mirrors and insecure-registries to buildkit
b73fd4d936864998451cdd37f45694541e43006e 2018-09-21T17:06:25-07:00 update vendor
bb2adc4496f2fd1b755fc701dbed5dab33175efd 2018-09-21T17:06:25-07:00 daemon/images: removed "found leaked image layer" warning, because it is expected now with buildkit
b501aa82d5151b8af73d6670772cc4e8ba94765f 2018-09-21T17:06:25-07:00 vendor: update bolt to bbolt
46a703bb3bfe75e99de2cc457dc0d568a1976a6b 2018-09-21T17:06:25-07:00 vendor: add bbolt v1.3.1-etcd.8
9f4cd6a7ea39ec0c1ad62a44b98f3f02b70efa78 2018-09-20T17:28:35-07:00 update containerd/console to fix race: lock Cond before Signal
66ed41aec82dbcdfbc38027e3d800e429af1cd58 2018-09-20T10:54:47-07:00 fixed the dockerd won't start bug when 'runtimes' field is defined in both daemon config file and cli flags
a5d731edecc75927f602c7f15e5ba9f5f77d3655 2018-09-18T11:19:51-07:00 create newBuildKit function separately in daemon_unix.go and daemon_windows.go for cross platform build
fc576226b24e8b5db6e95e48967d56c5808f9fe9 2018-09-18T12:34:56+02:00 Loosen permissions on /etc/docker directory
2c26eac56628527ed64c79ce9145ed97583cbeca 2018-09-17T12:28:09+02:00 pkg/progress: work around closing closed channel panic
5badfb40ebf1877bb319af3892b32a78491fb8e8 2018-09-14T17:29:47+02:00 always hornor client side to choose which builder to use with DOCKER_BUILDKIT env var regardless the server setup
f43fc6650cf5a452157fe081086098c124a426fd 2018-09-14T15:22:43+02:00 TestServiceWithDefaultAddressPoolInit: avoid panic
85361af1f749517c8bdfd3d36b0df94a92e29b2b 2018-09-14T15:20:07+02:00 Add fail fast path when containerd fails on startup
ee40a9ebcda2f46ea731ac1e2f840a2a23be0a07 2018-09-13T16:42:13-07:00 update vendor
e8620110fcbfca840e91bfb06348da8d9fd53e2e 2018-09-13T16:36:57-07:00 propagate the dockerd cgroup-parent config to buildkitd
2a82480df9ad91593d59be4b5283917dbea2da39 2018-09-06T18:39:22-07:00 TestFollowLogsProducerGone: add
84a5b528aede5579861201e869870d10fc98c07c 2018-09-06T18:39:22-07:00 daemon.ContainerLogs(): fix resource leak on follow
511741735e0aa2fe68a66d99384c00d187d1a157 2018-09-06T18:39:22-07:00 daemon/logger/loggerutils: add TestFollowLogsClose
2b8bc86679b7153bb4ace063a858637df0f16a2e 2018-09-06T18:39:22-07:00 daemon.ContainerLogs: minor debug logging cleanup
4e2dbfa1af48191126b0910b9463bf94d8371886 2018-09-06T18:39:21-07:00 pkg/filenotify/poller: fix Close()
3a3bfcbf47e98212abfc9cfed860d9e99fc41cdc 2018-09-06T18:39:21-07:00 pkg/filenotify/poller: close file asap
7be43586af6824c1e55cb502d9d2bab45c9b4505 2018-09-06T18:39:21-07:00 pkg/filenotify: poller.Add: fix fd leaks on err
d7085abec2e445630bedd3e79782c5ec33f62682 2018-09-06T17:24:39-07:00 vendor: update tar-split
fc1d808c44f77e3929c4eeba7066501890aecd4e 2018-09-06T17:24:36-07:00 integration/build: add TestBuildHugeFile
deba4bbc54f980b92f5aeab674688265486aa3b1 2018-09-05T16:00:14-07:00 delta: use chain ids to decide whether to skip a layer
f121eccf29576ce5d4b8256a71a9d32ee688ff7d 2018-09-05T06:59:52+00:00 Fix supervisor healthcheck throttling
c2d005320705c1339f8b2d5935f50f4c5b9cc6fc 2018-09-05T01:17:31+00:00 client: dial tls on Dialer if tls config is set
4c35d811471b80eec10476050b1222b653e6c5d9 2018-09-04T15:52:24+00:00 vendor buildkit to fix a couple of bugs
28150fc70cc0688c63d27879bf1c701ade47caff 2018-09-04T15:02:28+00:00 builder: implement ref checker
d2c316364225cd6bdf1581d11f370e0d159ad362 2018-09-04T15:02:28+00:00 builder: fix pruning all cache
3153708f13ceefc3e8e4d4a93778dcc9e4347f5a 2018-09-04T15:02:28+00:00 builder: add prune options to the API
2f94f103425985dc1677ee9a87f1161007949a45 2018-09-04T15:01:46+00:00 allow features option live reloadable
4032b6778df39f53fda0e6e54f0256c9a3b1d618 2018-08-30T17:34:59-07:00 Fix relabeling local volume source dir
5fa80da2d385520f2c3c91e8ce23d181e4f93097 2018-08-29T12:54:06+02:00 Fix regression when filtering container names using a leading slash
c87589c33b9974a1eeceede2b9606fbbddf3a8f5 2018-08-28T11:38:40+01:00 version: Fix balena CLI version string
9d1d910e5d293dddd6dfb96b9f50316d03697850 2018-08-28T11:35:16+01:00 version: Fix balena server version string
3685c83503cb7ebdfc0fff3f900123cec7913c73 2018-08-27T09:49:14-07:00 pkg/chrootarchive: disable memory cgroups until pending issues are fixed
85b036bd3a57a0cece7e09bc947cbafd7ba4fa4e 2018-08-27T09:40:30-07:00 vendor: update libnetwork to include stale default bridge fix
1d531ff64f99e07ac8733894416de8212a6c7278 2018-08-23T05:32:51+00:00 builder: fix bridge networking when using buildkit
b706f5daf673cde20571a611a33ae62d9fba26cb 2018-07-17T10:46:18-07:00 pkg/ioutils: implement eager writer
08b01efe225263dd8e7f7f82fd0fddc403f267b9 2018-07-05T20:33:22-07:00 Revert "vendor: update golang/x/sys to support fadvise for arm64"
60f2a21c95a6ee96b5a03367cc3c2e463be3787c 2018-07-05T20:33:21-07:00 pull: rely on memory cgroups to avoid page cache thrashing
38b223b0013390cb47276e1f4f5ddf6f44cb5db3 2018-06-28T12:18:45+03:00 pkg/stringid: don't bother seeding math/random with crypto grade seed
f08057baa72ce97af7de809d9f944bbb7f5275fe 2018-06-07T22:17:02+01:00 vendor: update btrfs dependency
519ed006c67b4637f5782dc737f78b33d53248e0 2018-05-07T20:18:00-07:00 container: remove extraneous lock leading to deadlocks
2e2f9df86df5cf35b47f2fbaa186f09a942f710b 2018-05-07T20:18:00-07:00 tests: more integration test fixes
276ee9d99df2a6ff8b1f1f3c4c8d1b46dad71bd9 2018-05-07T20:18:00-07:00 cmd/mobynit: adapt to new internal API
8e47b094929f8a7e47a96e51c883e9ff072ca01e 2018-05-07T20:18:00-07:00 build: switch the default build to be the dynamically linked binary
137b066c421a4a8390b0232648b992801356f7fc 2018-05-07T20:18:00-07:00 tests: remove plugin support in tests
64f52ee1e3ed120fc617aab21d3047983d252b24 2018-05-07T20:18:00-07:00 tests: skip swarm tests
e0e5db31fc8c0388f22dd1f664150b0b0ba738cf 2018-05-07T20:18:00-07:00 fix regression of DockerSuite.TestAPINetworkCreateCheckDuplicate
5955d382f3c9829a960dbfb9136bfb137477e5d6 2018-05-07T20:18:00-07:00 build: do not install embedded binaries separately
a466c05736c506467667623f51134a7839c20abd 2018-05-07T20:18:00-07:00 cmd/balena: exit with non-zero code if called with unknown command
3a1be7aa960835cf3f2e34f92532c4549c19cc66 2018-05-07T20:18:00-07:00 a lot of balena rename fixes for integration tests
f3b6b8a1aea4186d0d3a0e49241dc3750f97481a 2018-05-07T20:18:00-07:00 vendor: update containerd
b64eefe3a2fe7311d27f8a1a6a24565ba005e6f2 2018-05-07T20:18:00-07:00 build: switch to statically linked builds
9ed4298d6637ee54f92a7b45ce4edd23bd3d171c 2018-02-12T20:00:44-08:00 build: let the go compiler do the stripping
bd23724524ac213d3cb807048f3cdeb8f1bed203 2018-02-12T20:00:44-08:00 build: limit max go procs to avoid qemu hangs
5ead292ff82763acdc68efd8175dc7442682a9f0 2018-02-12T20:00:44-08:00 vendor: update golang/x/sys to support fadvise for arm64
0386158271fa492c85e6aacd38c683519c7d64a1 2018-02-12T20:00:44-08:00 build: add libudev dependency
fd78fe46cc62efec5f8fa45d9e217f69ae7072fb 2018-02-12T20:00:43-08:00 vendor: update containerd to non-plugin version
a1191cbdff1bd13a0cc3575ef7f63c7efc75dcef 2018-02-12T20:00:43-08:00 daemon/config: remove swarm support
ddaa8c17ff345a4ea0aa52fb59b4167e1f9a240b 2018-02-12T20:00:43-08:00 daemon: add appropriate container locks to avoid races
c24bda923c9347977b0e0a65ff224c7c026d2baa 2018-02-12T20:00:43-08:00 healthcheck: fix docker segfaulting
1cf563e3d171ef76964df10665e70ce211d9599d 2018-02-12T20:00:43-08:00 vendor: revendor everything
97505a483051f77a65f62f1800cf8dd2508365e8 2018-02-12T20:00:42-08:00 vendor: update vendor.conf with all required dependencies
8c124159175f3ad53482eccbafc3dfdbb931832c 2018-02-12T20:00:42-08:00 restartmanager: fixed the unit test
8af842e3ac86b58574c702676554ccb987afc210 2018-02-12T20:00:42-08:00 tests: renamed runc to balena-runc
55f437987735d6e4ccf356c9e5d4cfe13ada1211 2018-02-12T20:00:42-08:00 fixed balena version error
24b643b1b57d606a603b3bfd1f7f8831dff3e2d7 2018-02-12T20:00:42-08:00 daemon: experimental: restart container when they become unhealthy
b430038f36453d8f9006d380cc9f3030617076f4 2018-02-12T20:00:42-08:00 daemon: only attempt to prune local networks since swarm is disabled
eac6aa078447edb7fcda36975a5aa594ff8c5408 2018-02-12T20:00:41-08:00 Updated init scripts for Balena
062cf0e404d063811bb1d35313bf085b71f94468 2018-02-12T20:00:41-08:00 Updated github hooks for balena
07e8c0a0f9e261343ea4680c6c047d1ae16783cf 2018-02-12T20:00:41-08:00 Update website copy
5d81d5a215f5e8842ef8e999a30bc10c70988f8e 2018-02-12T20:00:41-08:00 Issue template should refer to balena throughout
a8846e22160da0ce2f9367769d92d04a0acfe1c2 2018-02-12T20:00:41-08:00 updated the mock of xfer to pass unit test
8f898bb3add828e82ec16f8514bd6f7f020326d7 2018-02-12T20:00:41-08:00 fixed integration with balena
60cb5cb2427f14ea100b77e450902a0c6115d321 2018-02-12T20:00:40-08:00 Renaming target to support balena
bce9bc78ee79cc80bfa47857dd756714dd8bdd54 2018-02-12T20:00:40-08:00 Fixed the runc version test
add016d96df3477e3f6e5b86b96e60e1c5419099 2018-02-12T20:00:40-08:00 skip tests of unsopported components
5d3045460459918738a55c181a69a1c9019fbf53 2018-02-12T20:00:40-08:00 fix addidental mention of balaena name instead of balena
5c46120af9340bb71c47befa8e39a027e627a8e2 2018-02-12T20:00:40-08:00 landr: add correct feature descriptions
189482e3725ced0d066afc44cbceb6b49ccbbaf8 2018-02-12T20:00:40-08:00 build: temporary switch to other base images
fcf3865fddb99e7a45dd7a1004f644b2e5d0a93b 2018-02-12T17:13:27-08:00 pkg/archive: sync files before issuing the fadvise syscall
090ec746e0d59c4ee38fb1f4720494c727b11233 2018-02-12T17:13:27-08:00 3 -> 3.5
b196586536c568f0f8e977a076c90c2da10a9328 2018-02-12T17:13:26-08:00 add analytics
2e78618717feb287c182240541ead3faf9b9bd9f 2018-02-12T17:13:26-08:00 docs: write a getting started guide
1003600d7200c13f1b48624ecb4d57b13dfae0ff 2018-02-12T17:13:26-08:00 build: use resin base images
3e2c3baab164afe9c82680b301d7ee9e7e25dbd9 2018-02-12T17:13:26-08:00 landr: use install.sh
b4a4651f0069902d77f05d92d82a9a4081c1a06e 2018-02-12T17:13:26-08:00 readme: Improve docker-ce section
392319c984f0d66b8c965d82f5a1d22fab80565e 2018-02-12T17:13:26-08:00 landr: fix architecture extraction from assets
f5922f9f421566d99d356c2a660b0f70fac995bb 2018-02-12T17:13:26-08:00 add post-build hook
e13fca33295c06d3f94a6e69e702250dd9c83a89 2018-02-12T17:13:25-08:00 add website data
99f6d0c5405d1d8a15f1edb02f347cb804d34e0e 2018-02-12T17:13:25-08:00 balena build scripts
db7076c18a833998c0413900013854629b458028 2018-02-12T17:13:25-08:00 add changelog for 17.06+rev1
d254fc040d5cec0b169078378241b0f6383797c5 2018-02-12T17:13:25-08:00 readme: describe what balena is
385bc86e9096a6aa62575e1b6f612ee842a7492b 2018-02-12T17:13:25-08:00 rename container engine to balena
49ddb26883b698cce2242d2b0e71febbc69079e9 2018-02-12T17:13:25-08:00 Fix typo in identifier name
3a8b3ac46396fb9af64bac83474aa435776af652 2018-02-12T17:13:24-08:00 Fix invalid API URL reference in swagger.yml
a33a1d80f721bf03b74ff09d91c1075ffb3548c9 2018-02-12T17:13:24-08:00 daemon: compute and print summary of delta efficiency
6a0936bba24099ee20f7f1e084bb82fa4c4018df 2018-02-12T17:13:24-08:00 delta: revert automatic tag generation for delta image
697bf0500a7cdecdcd30683a08af220ac04e060a 2018-02-12T17:13:24-08:00 api: move delta creation under POST images/delta
b67a57f6b44a4ea25ca6233a8ae482c1dac235bd 2018-02-12T17:13:24-08:00 image/delta: client support with progress reporting
08de939889c3046ff734f948e6f285f5554b523d 2018-02-12T17:13:24-08:00 pkg/ioutils: export SeekerSize utility
374fa017bcd6eb5ee60db134afb62e96fd117ea9 2018-02-12T17:13:24-08:00 distribution: calculate combined progress when pulling
128b42eb29f8f7378dc011494b1b09554e4b5f53 2018-02-12T17:13:23-08:00 pkg/progress: add progress sink
ec7e8dcafe3962da51692a2c98e02d3e1c963867 2018-02-12T17:13:23-08:00 daemon: return engine name as part of the version information
0c56d80d4982d41bd05e64aa05bbefebeae82eba 2018-02-12T17:13:23-08:00 cmd/mobynit: fix storage driver path
19da842517cd7fe6fe1b21297cf4f14d2307b4de 2018-02-12T17:13:23-08:00 cmd/mobynit: get graphdriver parameter from disk
2585aa19338217724fe2dac6d34741c8bb672ee1 2018-02-12T17:13:23-08:00 vendor: whitelist VERSION files
dd56c7ab7921f14a66f35f7d39d00c3f172c202d 2018-02-12T17:13:23-08:00 distribution: separate layer and image config for v1 pushes
b37c7dfd27eb9139efd32b957b87f427f2fa7cba 2018-02-12T17:13:23-08:00 build: support bind mounts during docker build
23a7f4da7ce39a3442dc8bcfdf47d82a610f7412 2018-02-12T17:13:22-08:00 rce-runc: changed the build to work only with resin-os/runc
35e9dadbdc864be2d21aae070d4e3217bf7ac45d 2018-02-12T17:13:22-08:00 rce-runc: porpogated the git commit and version
43f44b0b1e793cd8c42776c38a226426ffca2891 2018-02-12T17:13:22-08:00 delta: remove false assumption about src image
b1d10730870cfe7d767e17f05c43917156263d93 2018-02-12T17:13:22-08:00 integration: Skip plugin tests
63199860001969b8e552555f752a490d15a58605 2018-02-12T17:13:22-08:00 integration-cli: Disabled swarm tests
ea20d07becb11e6949c4fc001f0e9b9819fd9140 2018-02-12T17:13:22-08:00 Makefile: Revert cli-integration on swarm disable
860bb33b77acd6278fdb44255b417784286159f1 2018-02-12T17:13:22-08:00 hack/make Modified docker-py to pass integration w/o swarm
e58efd2e9593d086af59ecb163cbb78f2e8d1a32 2018-02-12T17:13:21-08:00 hack/make: Adjusted integration tests for rce-docker
ea1aaacf90a90f61342643075483797d8f6b32ea 2018-02-12T17:13:21-08:00 daemon: Adjusted unix_test for the container struct
6b4eb9670fa8c867ff089fd39eac3bc4b06b6975 2018-02-12T17:13:21-08:00 hack/make: Removed cluster unit tests
aab108d7020b5af19c5a3fefa9e12bf505165755 2018-02-12T17:13:21-08:00 migratev1_test: Fixed compilation errors
ca4c0ad02675de3385ec19a48036b08e9c42837e 2018-02-12T17:13:21-08:00 xfer/download_test: Fixed compilation errors
d22c96a33573d42551838c75a9472ce3e6b5ea7d 2018-02-12T17:13:21-08:00 hack/make/cross: Disabled the Windows target
29ece15118234d3be7ba6c56be504afd5ffa105d 2018-02-12T17:13:21-08:00 distribution: Added a warning when download failes
2496793a9d26e0a1015d27a3a4b51227c965f9ce 2018-02-12T17:13:20-08:00 cmd/mobynit: permanently remount root as rw
4192b6d153d88d47dabfee13026cc7863e906ac2 2018-02-12T17:13:20-08:00 cmd/mobynit: switch the pivot path to /mnt/sysroot/active
e0537fc5b06b53668fc321f5b32d6d77d6577270 2018-02-12T17:13:20-08:00 daemon: allow a secondary daemon store to be loaded
2d5b662827ffcc2c4dc700b79d3d0908246b2426 2018-02-12T17:13:18-08:00 distribution: implement delta pull
0713bcd1cb014d6fde2ecd00c0f169f8c18972b4 2018-02-12T17:12:43-08:00 distribution: serialise pulling the config and layers
0b6d9b43a9cbd62845787ec927d273d72dafa1ef 2018-02-12T17:08:00-08:00 deltas: implement image diff based on librsync
b3816803a038a44c411a95cba11db166307c5dc3 2018-02-12T17:06:55-08:00 api: skeleton for POST /deltas/create method
826321faf8a1eb42385362a017d68bbc807f1a34 2018-02-12T17:06:55-08:00 make GetTarSeekStream part of the ImageConfigStore interface
4d8f686b16efdf239e53ae28b49139002b8ad7ac 2018-02-12T17:06:55-08:00 image: implement seekable tar stream for a whole image
98d01b27a5c82ae1afc003185b3dc3457e803ec0 2018-02-12T17:06:55-08:00 layer: implement seekable tar stream
c8643e4d0d3c5dd1c4b8c9e03044857aea8d7478 2018-02-12T17:06:55-08:00 pkg/tarsplitutils: optimise reads using binary search
c2b1d1bbe130d292553d31f1dbd070d683987412 2018-02-12T17:06:55-08:00 pkg/tarsplitutils: implement random access seeker stream
c6133b90f337ea8a7c4ebbb73ee62df779bf1c9d 2018-02-12T17:06:55-08:00 ioutils: implement ReadSeekCloser concat stream
34769d6eb18be01b51225ad44e7410178b4e3520 2018-02-12T17:06:54-08:00 ioutils: implement ReadSeekCloser interface and wrapper
429ac7e5bb07f3adacb5ea61db0a7fff918e7dcc 2018-02-12T17:06:54-08:00 hack: revert binary stripping by default
bc9b3093c8707afef3a183e3be18f709404ec486 2018-02-12T17:06:54-08:00 hack: create both dynamic and static flavours of rce
caa9b94f726aa25812044a8cf58c52648d3d5c89 2018-02-12T17:06:53-08:00 distribution: check for nil before closing the download
3b2046fdc9e2d40ada9027b847e9e7898b72f503 2018-02-12T17:06:53-08:00 cmd/dockerd: remove support for docker plugins
0b51dbc886d0f534c4920dfb0ac418ff52e18317 2018-02-12T17:06:53-08:00 daemon/events: remove swarm related events
82b451966319a5fe9a697dff5a28551eadb8b1e8 2018-02-12T17:06:53-08:00 container_operations: remove swarm functionalities
77818f12ce92072b69f3097afaa63dbe8129eb03 2018-02-12T17:06:53-08:00 cmd/dockerd: drop support for swarm and checkpoint commands
98e2ea0b9506ac6211b25d765ae8ea851280c38c 2018-02-12T17:06:53-08:00 router/system: remove swarm dependency
35906f35bcbd8d969616061ec1420cb0136abbe5 2018-02-12T17:06:53-08:00 router/network: remove swarm dependency
0dcd39eeba55bca5f112fbe4e92f02a27523bb61 2018-02-12T17:06:52-08:00 cmd/mobynit: propagate initrd mounts to chroot
471e34a0182ad8e93763d831500a0606fd68dd71 2018-02-12T17:06:52-08:00 mobynit: read the containerID from /current symlink
fb65dbf59a338352b4d7e010b5e91ef697f8ee1e 2018-02-12T17:06:52-08:00 daemon: revert short circuit of volume setup for bare containers
8c0ceaba294daa7c0b1c1a7c722c77070e8b8786 2018-02-12T17:06:52-08:00 pkg/archive: use fadvise to prevent pagecache thrashing
10ae0733da29c67488126159370e87000bc018b0 2018-02-12T17:06:52-08:00 container: make sure config on disk has a valid Config
175199490819176eec7c78cdeb95466828e3eb28 2018-02-12T17:06:52-08:00 pkg/ioutils: sync parent directory too
684d8ba6109c853b355bf11ca3733c4099f14b92 2018-02-12T17:06:52-08:00 aufs,overlay: durably write layer on disk before returning
9d7480d8b6a0aaf6f35eb7de2620e23e00dbbba8 2018-02-12T17:06:51-08:00 cmd/mobynit: accept a flag for the graph driver
7191674c2c95069ef8918a063d4a6f33b283543f 2018-02-12T17:06:51-08:00 daemon: skip initLayer for bare runtime containers
60649c7a18a69dd39ba32216fe6c2a6b9f5f4f24 2018-02-12T17:06:51-08:00 cmd: add mobynit for host app booting
ad509615ba39b00d14c00f667f7444dd2619d961 2018-02-12T17:06:51-08:00 distribution: resume streaming download in case of failure
bf0b86825bc72b60ef6d74091f08089779864b7f 2018-02-12T17:06:51-08:00 ioutils: implement TeeReadCloser
fd1683789adbfb83e1e21c0b6191b95092efd5d2 2018-02-12T17:06:51-08:00 distribution: stream download directly to the layer store
55dcad3cacb620c605c4752b53b0e5156e988c81 2018-02-12T17:06:51-08:00 hack: create all appropriate symlinks after building rce
ae5453e659a263f830615ff738651212980ff0f4 2018-02-12T17:06:50-08:00 Renamed the target and binary name
c46694c87c2665c64885858b08bc1255e71f3bab 2018-02-12T17:06:50-08:00 rce: create stripped binary
7cd827f8eb2923f079cb6dd2c429ba3ad1b9dfec 2018-02-12T17:06:50-08:00 hack: allow variables to be set by the environment
e752e6aa4c219da7ba74434cd4838fcedb5c229e 2018-02-12T17:06:50-08:00 pkg/discovery: remove consul,etcd,zookeeper backends
e468a65bedc93ce942c0bcc7f7c55889270d17e2 2018-02-12T17:06:50-08:00 daemon: only support journald and jsonfile log drivers
6e257f07dfd7f928251250f99adca929e0c6197b 2018-02-12T17:06:50-08:00 cmd/rce: adapt imports to new packages
8b8b55c8a0229339848cfd11bd8967c4363f5e2c 2018-02-12T17:06:47-08:00 cmd/rce: Added the main of binary consolidation
23fd4d23c471548703d8edabf38634393fd48e73 2018-01-17T17:42:35-08:00 daemon, plugin: follow containerd namespace conventions
bf0d23511a92cd0f0de885c8ecedc658831db7bc 2018-01-17T17:40:40-08:00 Ensure containers are stopped on daemon startup
e3ffcd18f361097a522dea73eabcd90c660a1f0e 2018-01-17T17:39:23-08:00 Fix some missing synchronization in libcontainerd
8efdf3a4604a4ab5c8bbfeed8b0bb220ddd06e6f 2018-01-17T17:39:19-08:00 Fix error handling for kill/process not found
38559c69d2832ff29d6d77c7a7351b9aaf12cd37 2018-01-17T17:38:13-08:00 Remove support for referencing images by 'repository:shortid'
751a406dece4742eca27b7eece982768bce94c14 2018-01-17T17:34:32-08:00 Use commit-sha instead of tag for containerd
a95b4d549802423844de38a19818ac86e14b7d70 2018-01-17T17:33:06-08:00 Update API version-history for 1.35
```
