### 9. Working with balenaEngine

Service: `balena.service`

balenaEngine is balena's fork of Docker, offering a range of added features
including real delta downloads, minimal disk writes (for improved media wear)
and other benefits for edge devices, in a small resource footprint.

balenaEngine is responsible for the fetching and storage of service
images, as well as their execution as service containers, the creation of
defined networks and creation and storage of persistent data to service volumes.
As such, it is an extremely important part of balenaOS.

Additionally, as the Supervisor is also executed as a container, it is required
for its operation. This means that should balenaEngine fail for some reason,
it is likely that the Supervisor will also fail.

Issues with balenaEngine themselves are rare, although it can be initially
tempting to attribute them to balenaEngine instead of the actual underlying
issue. A couple of examples
of issues which are misattributed to :

- Failure to download release service updates - usually because there is an
  underlying network problem, or possibly issues with free space
- Failure to start service containers - most commonly customer services exit
  abnormally (or don't have appropriate error checking) although a full
  data partition can also occur, as can corrupt images

Reading the journal for balenaEngine is similar to all other `systemd` services.
Log into your device and then execute the following:

```shell
root@dee2945:~# journalctl --follow -n 100 --unit balena.service
...
```

What you'll first notice here is that there's Supervisor output here. This is
because balenaEngine is running the Supervisor and it pipes all Supervisor
logs to its own service output. This comes in particularly useful if you need to
examine the journal, because it will show both balenaEngine and Supervisor
output in the same logs chronologically.

Assuming your device is still running the pushed multicontainer app,
we can also see additionally logging for all the service containers.
To do so, we'll restart balenaEngine, so that the services are started again.
This output shows the last 50 lines of the balenaEngine journal after a restart.

```shell
root@debug-device:~# systemctl restart balena.service
root@debug-device:~# journalctl --all --follow -n 50 --unit balena.service
-- Journal begins at Fri 2022-08-19 18:08:10 UTC. --
Aug 19 18:24:36 debug-device balenad[4566]: time="2022-08-19T18:24:36.794621631Z" level=info msg="shim balena-engine-containerd-shim started" address=/containerd-shim/d7b6a623d687aea38a743fc79898e57bc3b43da3dbe144499bea2a068ad6700f.sock debug=false pid=5157
Aug 19 18:24:38 debug-device e593ab6439fe[4543]: [info]    Supervisor v14.0.8 starting up...
Aug 19 18:24:39 debug-device e593ab6439fe[4543]: [info]    Setting host to discoverable
Aug 19 18:24:39 debug-device e593ab6439fe[4543]: [warn]    Invalid firewall mode: . Reverting to state: off
Aug 19 18:24:39 debug-device e593ab6439fe[4543]: [info]    Applying firewall mode: off
Aug 19 18:24:39 debug-device e593ab6439fe[4543]: [debug]   Starting systemd unit: avahi-daemon.service
Aug 19 18:24:39 debug-device e593ab6439fe[4543]: [debug]   Starting systemd unit: avahi-daemon.socket
Aug 19 18:24:39 debug-device e593ab6439fe[4543]: [debug]   Starting logging infrastructure
Aug 19 18:24:39 debug-device e593ab6439fe[4543]: [info]    Starting firewall
Aug 19 18:24:39 debug-device e593ab6439fe[4543]: [debug]   Performing database cleanup for container log timestamps
Aug 19 18:24:39 debug-device e593ab6439fe[4543]: [success] Firewall mode applied
Aug 19 18:24:39 debug-device e593ab6439fe[4543]: [debug]   Starting api binder
Aug 19 18:24:39 debug-device e593ab6439fe[4543]: [info]    Previous engine snapshot was not stored. Skipping cleanup.
Aug 19 18:24:39 debug-device e593ab6439fe[4543]: [debug]   Handling of local mode switch is completed
Aug 19 18:24:39 debug-device e593ab6439fe[4543]: (node:1) [DEP0005] DeprecationWarning: Buffer() is deprecated due to security and usability issues. Please use the Buffer.alloc(), Buffer.allocUnsafe(), or Buffer.from() methods instead.
Aug 19 18:24:39 debug-device e593ab6439fe[4543]: [debug]   Spawning journald with: chroot  /mnt/root journalctl -a -S 2022-08-19 18:19:08 -o json CONTAINER_ID_FULL=4ce5bebc27c67bc198662cc38d7052e9d7dd3bfb47869ba83a88a74aa498c51f
Aug 19 18:24:39 debug-device e593ab6439fe[4543]: [info]    API Binder bound to: https://api.balena-cloud.com/v6/
Aug 19 18:24:39 debug-device e593ab6439fe[4543]: [event]   Event: Supervisor start {}
Aug 19 18:24:39 debug-device e593ab6439fe[4543]: [debug]   Spawning journald with: chroot  /mnt/root journalctl -a -S 2022-08-19 18:19:07 -o json CONTAINER_ID_FULL=2e2a7fcfe6f6416e32b9fa77b3a01265c9fb646387dbf4410ca90147db73a4ff
Aug 19 18:24:39 debug-device e593ab6439fe[4543]: [debug]   Connectivity check enabled: true
Aug 19 18:24:39 debug-device e593ab6439fe[4543]: [debug]   Starting periodic check for IP addresses
Aug 19 18:24:39 debug-device e593ab6439fe[4543]: [info]    Reporting initial state, supervisor version and API info
Aug 19 18:24:39 debug-device e593ab6439fe[4543]: [debug]   Skipping preloading
Aug 19 18:24:39 debug-device e593ab6439fe[4543]: [info]    Starting API server
Aug 19 18:24:39 debug-device e593ab6439fe[4543]: [info]    Supervisor API successfully started on port 48484
Aug 19 18:24:39 debug-device e593ab6439fe[4543]: [debug]   VPN status path exists.
Aug 19 18:24:39 debug-device e593ab6439fe[4543]: [info]    Applying target state
Aug 19 18:24:39 debug-device e593ab6439fe[4543]: [debug]   Ensuring device is provisioned
Aug 19 18:24:39 debug-device e593ab6439fe[4543]: [info]    VPN connection is active.
Aug 19 18:24:39 debug-device e593ab6439fe[4543]: [info]    Waiting for connectivity...
Aug 19 18:24:39 debug-device e593ab6439fe[4543]: [debug]   Starting current state report
Aug 19 18:24:39 debug-device e593ab6439fe[4543]: [debug]   Starting target state poll
Aug 19 18:24:39 debug-device e593ab6439fe[4543]: [debug]   Spawning journald with: chroot  /mnt/root journalctl -a --follow -o json _SYSTEMD_UNIT=balena.service
Aug 19 18:24:40 debug-device e593ab6439fe[4543]: [debug]   Finished applying target state
Aug 19 18:24:40 debug-device e593ab6439fe[4543]: [success] Device state apply success
Aug 19 18:24:40 debug-device e593ab6439fe[4543]: [info]    Reported current state to the cloud
Aug 19 18:24:40 debug-device e593ab6439fe[4543]: [info]    Applying target state
Aug 19 18:24:40 debug-device e593ab6439fe[4543]: [debug]   Finished applying target state
Aug 19 18:24:40 debug-device e593ab6439fe[4543]: [success] Device state apply success
Aug 19 18:24:49 debug-device e593ab6439fe[4543]: [info]    Internet Connectivity: OK
```

#### 9.1 Service Image, Container and Volume Locations

balenaEngine stores all its writeable data in the `/var/lib/docker` directory,
which is part of the data partition. We can see this by using the `mount`
command:

```shell
root@debug-device:~# mount | grep lib/docker
/dev/mmcblk0p6 on /var/lib/docker type ext4 (rw,relatime)
/dev/mmcblk0p6 on /var/volatile/lib/docker type ext4 (rw,relatime)
```

All balenaEngine state is stored in here, include images, containers and
volume data. Let's take a brief look through the most important directories
and explain the layout, which should help with investigations should they be
required.

Run `balena images` on your device:

```shell
root@debug-device:~# balena images
REPOSITORY                                                       TAG       IMAGE ID       CREATED        SIZE
registry2.balena-cloud.com/v2/8f425c77879116f77e6c8fcdebb37210   latest    f0735c857f39   32 hours ago   250MB
registry2.balena-cloud.com/v2/9664d653a6ecc4fe2c7e0cb18e64a3d2   latest    3128dae78199   32 hours ago   246MB
balena_supervisor                                                v14.0.8   936d20a463f5   7 weeks ago    75.7MB
registry2.balena-cloud.com/v2/04a158f884a537fc1bd11f2af797676a   latest    936d20a463f5   7 weeks ago    75.7MB
balena-healthcheck-image                                         latest    46331d942d63   5 months ago   9.14kB
```

Each image has an image ID. These identify each image uniquely for operations
upon it, such as executing it as a container, removal, etc. We can inspect an
image to look at how it's made up:

```shell
root@debug-device:~# balena inspect f0735c857f39
[
    {
        "Id": "sha256:f0735c857f39ebb303c5e908751f8ac51bbe0f999fe06b96d8bfc1a562e0f5ad",
        "RepoTags": [
            "registry2.balena-cloud.com/v2/8f425c77879116f77e6c8fcdebb37210:latest"
        ],
        "RepoDigests": [
            "registry2.balena-cloud.com/v2/8f425c77879116f77e6c8fcdebb37210@sha256:45c002b1bb325c1b93ff333a82ff401c9ba55ca7d00118b31a1c992f6fc5a4a4"
        ],
        "Parent": "",
        "Comment": "",
        "Created": "2022-08-18T10:17:31.25910477Z",
        "Container": "312d5c69d19954518ae932398ab9afa144feabff83906216759cffba925d9128",
        "ContainerConfig": {
            "Hostname": "faf98ca57090",
            "Domainname": "",
            "User": "",
            "AttachStdin": false,
            "AttachStdout": false,
            "AttachStderr": false,
            "Tty": false,
            "OpenStdin": false,
            "StdinOnce": false,
            "Env": [
                "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin",
                "LC_ALL=C.UTF-8",
                "DEBIAN_FRONTEND=noninteractive",
                "UDEV=off",
                "NODE_VERSION=10.22.0",
                "YARN_VERSION=1.22.4"
            ],
            "Cmd": [
                "/bin/sh",
                "-c",
                "#(nop) ",
                "CMD [\"npm\" \"start\"]"
            ],
            "Image": "sha256:c6072db33ab31c868fbce36621d57ddad8cf29b679027b7007d37ac40beea58c",
            "Volumes": null,
            "WorkingDir": "/usr/src/app",
            "Entrypoint": [
                "/usr/bin/entry.sh"
            ],
            "OnBuild": [],
            "Labels": {
                "io.balena.architecture": "aarch64",
                "io.balena.device-type": "jetson-tx2",
                "io.balena.qemu.version": "4.0.0+balena2-aarch64"
            }
        },
        "DockerVersion": "v19.03.12",
        "Author": "",
        "Config": {
            "Hostname": "faf98ca57090",
            "Domainname": "",
            "User": "",
            "AttachStdin": false,
            "AttachStdout": false,
            "AttachStderr": false,
            "Tty": false,
            "OpenStdin": false,
            "StdinOnce": false,
            "Env": [
                "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin",
                "LC_ALL=C.UTF-8",
                "DEBIAN_FRONTEND=noninteractive",
                "UDEV=off",
                "NODE_VERSION=10.22.0",
                "YARN_VERSION=1.22.4"
            ],
            "Cmd": [
                "npm",
                "start"
            ],
            "Image": "sha256:c6072db33ab31c868fbce36621d57ddad8cf29b679027b7007d37ac40beea58c",
            "Volumes": null,
            "WorkingDir": "/usr/src/app",
            "Entrypoint": [
                "/usr/bin/entry.sh"
            ],
            "OnBuild": [],
            "Labels": {
                "io.balena.architecture": "aarch64",
                "io.balena.device-type": "jetson-tx2",
                "io.balena.qemu.version": "4.0.0+balena2-aarch64"
            }
        },
        "Architecture": "amd64",
        "Os": "linux",
        "Size": 249802248,
        "VirtualSize": 249802248,
        "GraphDriver": {
            "Data": {
                "LowerDir": "/var/lib/docker/overlay2/0bc07f1279e00affc344c22f64b9f3985225fe3e06f13103e24b983b1a9fdd0e/diff:/var/lib/docker/overlay2/27229498851db6327bf134ba8ab869655a20bf5e602a435a58e03684088baf7a/diff:/var/lib/docker/overlay2/b0d957f3e2636fe4dc5b10855ad276dc0e53fa95a5c3e2ed3fa56dcee23dc50f/diff:/var/lib/docker/overlay2/4489255833d1236ba41a2d15761f065eba992069569075d0a9edf2cd53e8415f/diff:/var/lib/docker/overlay2/f77ebf3b4836d289c2515c82537cd774354b7342c2a4899fcffb51ac23e9e9b7/diff:/var/lib/docker/overlay2/c5a96f1657a4073293c3364eab567a1f62b5d7761b8dbc3617369ffbd516c8f0/diff:/var/lib/docker/overlay2/168a1333d7a784f1b1ecbe6202f8a8189e592407195349f7d2dad943084876e6/diff:/var/lib/docker/overlay2/dd17e88dd38b700214d8f13c3d820d1f808c4b1a138f91dafd729f6369d95110/diff:/var/lib/docker/overlay2/641c77f5a0c9f25154fbd868c253c8a8e894e3cebd9ba5b96cebb9384c6283d7/diff:/var/lib/docker/overlay2/8d428280199e4dc05de155f1f3b0ef63fdeef14e09662ce5676d8b1d790bdf5d/diff:/var/lib/docker/overlay2/bcc97249ce05979fc9aa578c976f770083e9948a1c1d64c05444591a7aad35a9/diff:/var/lib/docker/overlay2/41773d1c239c8a0bf31096d43ae7e17b5ca48f10530c13d965259ed386cb20d9/diff:/var/lib/docker/overlay2/697de96abdf1ae56449b0b01ce99ef867c821404d876f2b55eac8ccb760a1bc1/diff:/var/lib/docker/overlay2/4518e2c4e4ca2b5b6b26ed40a8e3f130296625c9a8c6a47c61763bf789e8df12/diff:/var/lib/docker/overlay2/662d75e19e4a9a98442111f551f48bd8841de472f16c755405276462122e1969/diff:/var/lib/docker/overlay2/338c3b3a4977e96ed25d1f2b9fbc65562a8c62a185df6baabe64dd3e40632331/diff:/var/lib/docker/overlay2/88069ce00e7cf154912b583ac49e2eb655af35777a22f68d65e560a8f8f72fb0/diff:/var/lib/docker/overlay2/7fa0af02f451bafb29410b8d6712bb84cfead585aca227dd899631b357eba12c/diff:/var/lib/docker/overlay2/fbaf99abadbb408297dae5626b85415bf62fdc40dce4ec4df11ff5f8043306b3/diff:/var/lib/docker/overlay2/f871aa611f42e273ef69b8ee3a686e084b6eddaf9d20c36e10904e4e44013fe1/diff:/var/lib/docker/overlay2/d21067d6f38aebbe391336a2251aac02e0bda38db313de18d417c12f20eba067/diff:/var/lib/docker/overlay2/daa9d5a342f3234c8f7dc2905dbbaee1823821b828f3e92b0092c9e83e56cbde/diff:/var/lib/docker/overlay2/c31888f5276a034a32c1017906f32a79000971067dee4f85c3ef87717c00fe94/diff",
                "MergedDir": "/var/lib/docker/overlay2/2b7065b1b4192e5232b470d308e4992a47d7ab4786a7fcc9356682512c69d2ec/merged",
                "UpperDir": "/var/lib/docker/overlay2/2b7065b1b4192e5232b470d308e4992a47d7ab4786a7fcc9356682512c69d2ec/diff",
                "WorkDir": "/var/lib/docker/overlay2/2b7065b1b4192e5232b470d308e4992a47d7ab4786a7fcc9356682512c69d2ec/work"
            },
            "Name": "overlay2"
        },
        "RootFS": {
            "Type": "layers",
            "Layers": [
                "sha256:af0df278bec194015fd6f217c017da2cc48c7c0dfc55974a4f1da49f1fd9c643",
                "sha256:1d3545625acc456a3b1ffe0452dc7b21b912414514f6636776d1df5f7fc5b761",
                "sha256:d50c820f06af93966310405459feaa06ae25099493e924827ecc2b87d59b81bc",
                "sha256:c73bea5e51f02ae807581a4308b93bb15f2a4e8eff77980f04ca16ada3e5d8fe",
                "sha256:30781fcde1e029490e5e31142c6f87eea28f44daa1660bf727149255095aeb25",
                "sha256:06333a8766d373657969d2c5a1785426b369f0eb67d3131100562e814944bb61",
                "sha256:afe83b1715ba56fe1075387e987414f837ad1153e0dc83983b03b21af1fff524",
                "sha256:b5b49d315a0b1d90a0e7f5bf0c9f349b64b83868aa4965846b05fb1c899b4d31",
                "sha256:40e07aec657d65d0b3c2fd84983c3f725cf342a11ac8f694d19d8162556389ca",
                "sha256:1ec17697083e5ab6e6657cb2fd3ac213bfb621497196a9b4dd01be49a05fd0ba",
                "sha256:8cdee574cf2ff2bff80961902b11f378bd228d11f4e52135fd404f79ca7e1a63",
                "sha256:b4cb326982763715403b5101026ad8333e2b5d59868cce0ddf21d1a231715758",
                "sha256:94893f94e0143c27cfd1367699776035e144b4c3b3cff1c1423253f6e2b39723",
                "sha256:8a743afad01d015d44b64fd0bedf583145d04796a429aa261596e0e6943bda7f",
                "sha256:faa977113a35d9b622af3bb1a481a7ee5bdbc3f6c35dc8ff4ff5adb8b4a95641",
                "sha256:59461c292cd4bd0f3fbd808371d37845e1850ed7b9c2e676b484c60950bdd3ba",
                "sha256:285f1fb0f99ea936e0eeb5f78c83b050c3b8f334a956a40f9ec547ac29b3a58d",
                "sha256:d653d2b6a0bde68b3194e62baec92a2ef2223cd9c56e3ea4844b38886b63798e",
                "sha256:234dc1e17bed2bae6e5368c2667c30124dca94f0b584f5cd8b0f2be249311820",
                "sha256:c2630cbbb8b9413cc6d8d5bd4fcdebf54d987d62e0a2c68cf8dadb5cc831209d",
                "sha256:a5311ef02278680ab6c2cf1a5843845f5802ed024fce4f69eb8e8ea53b7a5b4e",
                "sha256:188bd1bf502d426d5e897e31773fa885399fd69ceef850609cdaa4a45f330c71",
                "sha256:43cb5b1fb08f5faa3ae526d509b5faa014ce9b8f1099b27e87f8cc3992a973c5",
                "sha256:896037576e604880800e50afde6184d54e3f50b3cede0f564dcdd3e3243bba5a"
            ]
        },
        "Metadata": {
            "LastTagTime": "2022-08-19T18:24:39.169388977Z"
        }
    }
]
```

Of particular interest here is the `"RootFS"` section, which shows all of the
layers that make up an image. Without going into detail (there are plenty of
easily Google-able pages describing this), each balena image is made up of a
series of layers, usually associated with a `COPY`, `ADD`, `RUN`, etc.
Dockerfile command at build time. Each layer makes up part of the images
filing system, and when a service container is created from an image, it uses
these layers 'merged' together for the underlying filing system.

We can look at these individual layers by making a note of each SHA256 hash ID
and then finding this hash in the `/var/lib/docker/image/overlay2/layerdb/sha256` directory
This directory holds a set of directories, each named after each unique layer using
the SHA256 associated with it. Let's look at the layer DB directory:

```shell
root@debug-device:~# ls -lah /var/lib/docker/image/overlay2/layerdb/sha256
total 176K
drwxr-xr-x 44 root root 4.0K Aug 19 10:25 .
drwx------  5 root root 4.0K Aug 19 10:24 ..
drwx------  2 root root 4.0K Aug 19 10:25 09b78bc523987759f021e0a0e83e69c8084b1e3c20f14b4bb9534f3cdcc6ac3c
drwx------  2 root root 4.0K Jul  8 18:39 1036152b568007807de32d686a009c3f4f8adbb4f472e82ac274a27d92326d80
drwx------  2 root root 4.0K Jul  8 18:39 106a406f45c345ecbc3c42525606f80b084daf9f3423615a950c4cf5a696caa7
drwx------  2 root root 4.0K Aug 19 10:25 144d2c5112d686b5a4b14f3d4d9d8426981debc440ae533e3c374148089a66d3
drwx------  2 root root 4.0K Jul  8 18:39 1f9ad706e17e7786d85618534db2a36bb51b8aaaadd9a7e32d1e7080054ff620
drwx------  2 root root 4.0K Aug 19 10:25 21331e0e0fe982279deb041181d6af17bcd8ac70dc7dc023c225d2cfb3c33b7f
drwx------  2 root root 4.0K Jul  8 18:39 253123a3e5d5904ceeedc9b7f22f95baa93228cf7eeb8a659b2e7893e2206d32
drwx------  2 root root 4.0K Jul  8 18:39 294ac687b5fcac6acedb9a20cc756ffe39ebc87e8a0214d3fb8ef3fc3189ee2a
drwx------  2 root root 4.0K Jul  8 18:39 2ef1f0e36f419227844367aba4ddfa90df1294ab0fe4993e79d56d9fe3790362
drwx------  2 root root 4.0K Aug 19 10:24 35e3a8f4d7ed3009f84e29161552f627523e42aea57ac34c92502ead41691ce9
drwx------  2 root root 4.0K Jul  8 18:39 3eb45474328f8742c57bd7ba683431fe5f0a5154e12e382814a649cc0c4115b4
drwx------  2 root root 4.0K Jul  8 18:39 4acac926cb8671f570045c0a1dc1d73c1ca4fbfeee82b8b69b26b095a53bd9b7
drwx------  2 root root 4.0K Jul  8 18:39 4c345896c7e7169034e1cd5ae7a6ded46623c904503e37cfe0f590356aed869a
drwx------  2 root root 4.0K Aug 19 10:25 4c4fb86d946143f592f295993169aa9d48e51a2583d0905874a3a122843d6ef1
drwx------  2 root root 4.0K Jul  8 18:39 4eb65237b6a403037e5a629c93f0efd25c8cf6bc78a0c4c849e6a7d4f76892bc
drwx------  2 root root 4.0K Aug 19 10:25 521a7e6b1a2ca6e5417331c9bb9330fd94b48ec80d80099fc5cbffce33f0b871
drwx------  2 root root 4.0K Jul  8 18:39 5a5de1543b2bc574adae2f477ae308ea21d8bfcdd3828b01b1cf1b3e54e757bf
drwx------  2 root root 4.0K Aug 19 10:25 60ef73a52c9699144533715efa317f7e4ff0c066697ae0bb5936888ee4097664
drwx------  2 root root 4.0K Aug 19 10:25 7754099c5f93bb0d0269ea8193d7f81e515c7709b48c5a0ca5d7682d2f15def2
drwx------  2 root root 4.0K Aug 19 10:25 7a9f44ed69912ac134baabbe16c4e0c6293750ee023ec22488b3e3f2a73392a6
drwx------  2 root root 4.0K Aug 19 10:25 81713bf891edc214f586ab0d02359f3278906b2e53e034973300f8c7deb72ca2
drwx------  2 root root 4.0K Aug 19 10:25 8515f286995eb469b1816728f2a7dc140d1a013601f3a942d98441e49e6a38e9
drwx------  2 root root 4.0K Aug 19 10:25 8c0972b6057fafdf735562e0a92aa3830852b462d8770a9301237689dbfa6036
drwx------  2 root root 4.0K Aug 19 10:25 a1659a5774b2cd9f600b9810ac3fa159d1ab0b6c532ff22851412fe8ff21c45e
drwx------  2 root root 4.0K Aug 19 10:25 aa293c70cb6bdafc80e377b36c61368f1d418b33d56dcbc60b3088e945c24784
drwx------  2 root root 4.0K Aug 19 10:25 abaae3be7599181738152e68cfd2dcf719068b3e23de0f68a85f1bfe0a3ebe6e
drwx------  2 root root 4.0K Aug 19 10:25 acd72d7857fe303e90cd058c8b48155a0c2706ff118045baedf377934dcd5d63
drwx------  2 root root 4.0K Aug 19 10:24 af0df278bec194015fd6f217c017da2cc48c7c0dfc55974a4f1da49f1fd9c643
drwx------  2 root root 4.0K Aug 19 10:25 b9c29dbb49463dcc18dbf10d5188453a3c3d3159dd42b9fb403d3846649b8c1f
drwx------  2 root root 4.0K Aug 19 10:25 bc60d12607d916904e332fc94997d59b5619885fbb51af0e885ca21e88faec7f
drwx------  2 root root 4.0K Aug 19 10:25 c3414c33510eabf6a77f68f6714967484e7a937681d30f4578ed4415a889abbf
drwx------  2 root root 4.0K Aug 19 10:25 d89860b1b17fbfc46a8c09473504697efc768895910a383c21182c073caa249d
drwx------  2 root root 4.0K Aug 19 10:25 e665e5030545a1d1f8bb3ad1cda5a5d0bad976a23005313157e242e6a3a5932e
drwx------  2 root root 4.0K Aug 19 10:25 e80af4b706db8380d5bb3d88e4262147edfaad362fe4e43cf96658709c21195a
drwx------  2 root root 4.0K Jul  8 18:39 eb8f1151aa650015a9e6517542b193760795868a53b7644c54b5ecdac5453ede
drwx------  2 root root 4.0K Jul  8 18:39 efb53921da3394806160641b72a2cbd34ca1a9a8345ac670a85a04ad3d0e3507
drwx------  2 root root 4.0K Aug 19 10:25 f77763335291f524fee886c06f0f81f23b38638ba40c4695858cd8cca3c92c46
drwx------  2 root root 4.0K Aug 19 10:25 f92d2ac421bffffd2811552a43b88a0cc4b4fe2faa26ec85225d68e580447dc4
drwx------  2 root root 4.0K Aug 19 10:25 fae75b52757886e2f56dd7f8fbf11f5aa30ff848745b2ebf169477941282f9f9
drwx------  2 root root 4.0K Aug 19 10:25 fc810c085b1a4492e6191b546c9708b3777cbb8976744d5119ed9e00e57e7bc6
drwx------  2 root root 4.0K Jul  8 18:39 fd22d21dd7203b414e4d4a8d61f1e7feb80cbef515c7d38dc98f091ddaa0d4a2
drwx------  2 root root 4.0K Jul  8 18:39 fd6efabf36381d1f93ba9485e6b4d3d9d9bdb4eff7d52997362e4d29715eb18a
```

Note that there are layers named here that are _not_ named in the `"RootFS"`
object in the image description (although base image layers usually are). This
is because of the way balenaEngine describes layers and naming internally
(which we will not go into here). However, each layer is described by a
`cache-id` which is a randomly generated UUID. You can find the `cache-id`
for a layer by searching for it in the layer DB directory:

```shell
root@debug-device:~# cd /var/lib/docker/image/overlay2/layerdb/sha256/
root@debug-device:/var/lib/docker/image/overlay2/layerdb/sha256# grep -r 09b78bc523987759f021e0a0e83e69c8084b1e3c20f14b4bb9534f3cdcc6ac3c *
7754099c5f93bb0d0269ea8193d7f81e515c7709b48c5a0ca5d7682d2f15def2/parent:sha256:09b78bc523987759f021e0a0e83e69c8084b1e3c20f14b4bb9534f3cdcc6ac3c
```

In this case, we find two entries, but we're only interested in the directory
with the `diff` file result, as this describes the diff for the layer:

```shell
root@debug-device:/var/lib/docker/image/overlay2/layerdb/sha256# cat 09b78bc523987759f021e0a0e83e69c8084b1e3c20f14b4bb9534f3cdcc6ac3c/cache-id
f77ebf3b4836d289c2515c82537cd774354b7342c2a4899fcffb51ac23e9e9b7
```

We now have the corresponding `cache-id` for the layer's directory layout,
and we can now examine the file system for this layer (all the diffs are
store in the `/var/lib/docker/overlay2/<ID>/diff` directory):

```shell
roroot@debug-device:~# du -hc /var/lib/docker/overlay2/f278e81229574468df2f798e3ffbe576a51c2ad0c752c0b1997fdb33314130ae/diff/
8.0K	/var/lib/docker/overlay2/f278e81229574468df2f798e3ffbe576a51c2ad0c752c0b1997fdb33314130ae/diff/root/.config/configstore
16K	/var/lib/docker/overlay2/f278e81229574468df2f798e3ffbe576a51c2ad0c752c0b1997fdb33314130ae/diff/root/.config
24K	/var/lib/docker/overlay2/f278e81229574468df2f798e3ffbe576a51c2ad0c752c0b1997fdb33314130ae/diff/root
4.0K	/var/lib/docker/overlay2/f278e81229574468df2f798e3ffbe576a51c2ad0c752c0b1997fdb33314130ae/diff/tmp/balena
4.0K	/var/lib/docker/overlay2/f278e81229574468df2f798e3ffbe576a51c2ad0c752c0b1997fdb33314130ae/diff/tmp/resin
12K	/var/lib/docker/overlay2/f278e81229574468df2f798e3ffbe576a51c2ad0c752c0b1997fdb33314130ae/diff/tmp
4.0K	/var/lib/docker/overlay2/f278e81229574468df2f798e3ffbe576a51c2ad0c752c0b1997fdb33314130ae/diff/run/mount
8.0K	/var/lib/docker/overlay2/f278e81229574468df2f798e3ffbe576a51c2ad0c752c0b1997fdb33314130ae/diff/run
48K	/var/lib/docker/overlay2/f278e81229574468df2f798e3ffbe576a51c2ad0c752c0b1997fdb33314130ae/diff/
48K	total
```

You can find the diffs for subsequent layers in the same way.

However, whilst this allows you to examine all the layers for an image, the
situation changes slightly when an image is used to create a container. At this
point, a container can also bind to volumes (persistent data directories across
container restarts) and writeable layers that are used only for that container
(which are _not_ persistent across container restarts). Volumes are described in
a later section dealing with media storage. However, we will show an example
here of creating a writeable layer in a container and finding it in the
appropriate `/var/lib/docker` directory.

Assuming you're running the source code that goes along with this masterclass,
SSH into your device:

```shell
root@debug-device:~# balena ps
CONTAINER ID   IMAGE                                                            COMMAND                  CREATED          STATUS                    PORTS     NAMES
4ce5bebc27c6   3128dae78199                                                     "/usr/bin/entry.sh n…"   27 minutes ago   Up 22 minutes                       backend_5298819_2265201_0a9d4b0e8c1ff1202773ac2104a2bb48
2e2a7fcfe6f6   f0735c857f39                                                     "/usr/bin/entry.sh n…"   27 minutes ago   Up 22 minutes                       frontend_5298818_2265201_0a9d4b0e8c1ff1202773ac2104a2bb48
e593ab6439fe   registry2.balena-cloud.com/v2/04a158f884a537fc1bd11f2af797676a   "/usr/src/app/entry.…"   27 minutes ago   Up 22 minutes (healthy)             balena_supervisor
```

You should see something similar. Let's pick the `backend` service, which in
this instance is container `4ce5bebc27c6`. We'll `exec` into it via a `bash`
shell, and create a new file. This will create a new writeable layer for the
container:

```shell
root@debug-device:~# balena exec -ti 4ce5bebc27c6 /bin/bash
root@4ce5bebc27c6:/usr/src/app# echo 'This is a new, container-only writeable file!' > /mynewfile.txt
root@4ce5bebc27c6:/usr/src/app# cat /mynewfile.txt
This is a new, container-only writeable file!
root@4ce5bebc27c6:/usr/src/app# exit
```

Now we'll determine where this new file has been stored by balenaEngine.
Similarly to the images, any writeable layer ends up in the
`/var/lib/docker/overlay2/<ID>/diff` directory, but to determine the correct layer ID
we need to examine the layer DB for it. We do this by looking in the
`/var/lib/docker/image/overlay2/layerdb/mounts` directory, which lists all the
currently created containers:

```shell
root@debug-device:~# cd /var/lib/docker/image/overlay2/layerdb/mounts
root@debug-device:~# ls -lh
total 12K
drwxr-xr-x 2 root root 4.0K Aug 19 18:19 2e2a7fcfe6f6416e32b9fa77b3a01265c9fb646387dbf4410ca90147db73a4ff
drwxr-xr-x 2 root root 4.0K Aug 19 18:19 4ce5bebc27c67bc198662cc38d7052e9d7dd3bfb47869ba83a88a74aa498c51f
drwxr-xr-x 2 root root 4.0K Aug 19 18:19 e593ab6439fee4f5003e68e616fbaf3c3dfd7e37838b1e27d9773ecb65fb26c6
```

As you can see, there's a list of all the container IDs of those container
that have been created. If we look for the `mount-id` file in the directory
for the `backend` container, that will include the layer ID of the layer that
has been created. From there, we simply look in the appropriate diff layer
directory to find our newly created file (the `awk` command below is to add
a newline to the end of the discovered value for clarity reasons):

```shell
root@debug-device:~# cd /var/lib/docker/image/overlay2/layerdb/mounts
root@debug-device:/var/lib/docker/image/overlay2/layerdb/mounts# cat 4ce5bebc27c67bc198662cc38d7052e9d7dd3bfb47869ba83a88a74aa498c51f/mount-id | awk '{ print $1 }'
18d634420eceb9792f57554a5451510c1a3e38efe15552045d9b074c5120ef3c
root@debug-device:/var/lib/docker/image/overlay2/layerdb/mounts# cat /var/lib/docker/overlay2/18d634420eceb9792f57554a5451510c1a3e38efe15552045d9b074c5120ef3c/diff/mynewfile.txt
This is a new, container-only writeable file!
```

#### 9.2 Restarting balenaEngine

As with the Supervisor, it's very rare to actually need to carry this out.
However, for completeness, should you need to, this again is as simple as
carrying out a `systemd` restart with `systemctl restart balena.service`:

```shell
root@debug-device:~# systemctl restart balena.service
root@debug-device:~# systemctl status balena.service
● balena.service - Balena Application Container Engine
     Loaded: loaded (/lib/systemd/system/balena.service; enabled; vendor preset: enabled)
    Drop-In: /etc/systemd/system/balena.service.d
             └─storagemigration.conf
     Active: active (running) since Fri 2022-08-19 19:00:03 UTC; 13s ago
TriggeredBy: ● balena-engine.socket
       Docs: https://www.balena.io/docs/getting-started
   Main PID: 8721 (balenad)
      Tasks: 60 (limit: 1878)
     Memory: 130.2M
     CGroup: /system.slice/balena.service
             ├─4544 /proc/self/exe --healthcheck /usr/lib/balena/balena-healthcheck --pid 4543
             ├─8721 /usr/bin/balenad --experimental --log-driver=journald --storage-driver=overlay2 -H fd:// -H unix:///var/run/balena.so>
             ├─8722 /proc/self/exe --healthcheck /usr/lib/balena/balena-healthcheck --pid 8721
             ├─8744 balena-engine-containerd --config /var/run/balena-engine/containerd/containerd.toml --log-level info
             ├─8922 balena-engine-containerd-shim -namespace moby -workdir /var/lib/docker/containerd/daemon/io.containerd.runtime.v1.lin>
             ├─8941 balena-engine-containerd-shim -namespace moby -workdir /var/lib/docker/containerd/daemon/io.containerd.runtime.v1.lin>
             ├─9337 balena-engine-containerd-shim -namespace moby -workdir /var/lib/docker/containerd/daemon/io.containerd.runtime.v1.lin>
             └─9343 balena-engine-runc --root /var/run/balena-engine/runtime-balena-engine-runc/moby --log /run/balena-engine/containerd/>

Aug 19 19:00:06 debug-device e593ab6439fe[8721]: [info]    Applying target state
Aug 19 19:00:06 debug-device e593ab6439fe[8721]: [debug]   Finished applying target state
Aug 19 19:00:06 debug-device e593ab6439fe[8721]: [success] Device state apply success
Aug 19 19:00:06 debug-device e593ab6439fe[8721]: [info]    Reported current state to the cloud
Aug 19 19:00:14 debug-device balenad[8721]: time="2022-08-19T19:00:14.615648023Z" level=info msg="Container failed to exit within 10s of >
Aug 19 19:00:15 debug-device balenad[8744]: time="2022-08-19T19:00:15.074224040Z" level=info msg="shim reaped" id=e593ab6439fee4f5003e68e>
Aug 19 19:00:15 debug-device balenad[8721]: time="2022-08-19T19:00:15.077066553Z" level=info msg="ignoring event" container=e593ab6439fee>
Aug 19 19:00:16 debug-device balenad[8721]: time="2022-08-19T19:00:16.870509180Z" level=warning msg="Configured runtime \"runc\" is depre>
Aug 19 19:00:16 debug-device balenad[8744]: time="2022-08-19T19:00:16.911032847Z" level=warning msg="runtime v1 is deprecated since conta>
Aug 19 19:00:16 debug-device balenad[8744]: time="2022-08-19T19:00:16.916761744Z" level=info msg="shim balena-engine-containerd-shim star>
```

However, doing so has also had another side-effect. Because the Supervisor is
itself comprised of a container, restarting balenaEngine has _also_ stopped
and restarted the Supervisor. This is another good reason why balenaEngine
should only be stopped/restarted if absolutely necessary.

So, when is absolutely necessary? There are some issues which occasionally
occur that might require this.

One example could be due to corruption in the `/var/lib/docker` directory, usually related to:

- Memory exhaustion and investigation
- Container start/stop conflicts

Many examples of these are documented in the support knowledge base, so we will
not delve into them here.