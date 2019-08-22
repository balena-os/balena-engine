# Patches on top of moby

## TODO

<details>
<summary>Commits</summary>

```
commit 80d443d400dcc85e87322f72866593b46bafb157
Author:     Resin CI <34882892+balena-ci@users.noreply.github.com>
AuthorDate: Thu Aug 22 15:17:52 2019 +0300
Commit:     Resin CI <34882892+balena-ci@users.noreply.github.com>
CommitDate: Thu Aug 22 15:17:52 2019 +0300

    v18.9.8

 CHANGELOG.md | 5 +++++
 VERSION      | 2 +-
 2 files changed, 6 insertions(+), 1 deletion(-)

---

commit 87ef8e0dca4281eac019704c4fd2f1a95c06824a
Author:     Resin CI <34882892+balena-ci@users.noreply.github.com>
AuthorDate: Wed Jun 26 11:49:23 2019 +0300
Commit:     Resin CI <34882892+balena-ci@users.noreply.github.com>
CommitDate: Wed Jun 26 11:49:23 2019 +0300

    v18.9.7

 CHANGELOG.md | 7 +++++++
 VERSION      | 2 +-
 2 files changed, 8 insertions(+), 1 deletion(-)

---

commit 178650445602475aeee1eab8076c5a83f9b6ed16
Author:     Robert Günzler <r@gnzler.io>
AuthorDate: Tue Jan 8 16:47:35 2019 +0100
Commit:     Robert Günzler <r@gnzler.io>
CommitDate: Tue Jun 25 14:31:02 2019 +0200

    contrib/install.sh: Improve error output
    
    Check for required tools and root, with user notification
    and clean exit.
    
    Signed-off-by: Robert Günzler <robertg@balena.io>

 contrib/install.sh | 37 +++++++++++++++++++++++++++++++++----
 1 file changed, 33 insertions(+), 4 deletions(-)

---

commit 8eb3f4af0e0b1b70d77cc3e456458fb430027b76
Author:     Robert Günzler <r@gnzler.io>
AuthorDate: Tue Jan 8 13:41:07 2019 +0100
Commit:     Robert Günzler <r@gnzler.io>
CommitDate: Tue Jun 25 14:31:02 2019 +0200

    contrib/install.sh: Add details to the success message
    
    First warn the user that balena-engine-daemon needs to be started.
    Including instructions on how to make the system ready for that:
    - service files
    - balena-engine group
    - how to allow non-root users to run containers
    
    Connects-to: #55
    Connects-to: #51
    Change-type: patch
    Signed-off-by: Robert Günzler <robertg@balena.io>

 contrib/install.sh | 20 ++++++++++++++++++++
 1 file changed, 20 insertions(+)

---

commit 64cb87f1e17b195836a8722cdfd6360a65fd6a91
Author:     Robert Günzler <r@gnzler.io>
AuthorDate: Tue Jan 8 13:12:32 2019 +0100
Commit:     Robert Günzler <r@gnzler.io>
CommitDate: Tue Jun 25 14:31:02 2019 +0200

    contrib/install.sh: Fail on error
    
    The install script should not print the success message if it didn't
    actually succeed to install anything
    
    Connects-to: #54
    Change-type: patch
    Signed-off-by: Robert Günzler <robertg@balena.io>

 contrib/install.sh | 2 ++
 1 file changed, 2 insertions(+)

---

commit 95c7371304f9cef494efe93f0a8ffd53a75eac21
Author:     Resin CI <34882892+balena-ci@users.noreply.github.com>
AuthorDate: Mon Jun 3 18:22:01 2019 +0300
Commit:     Resin CI <34882892+balena-ci@users.noreply.github.com>
CommitDate: Mon Jun 3 18:22:01 2019 +0300

    v18.9.6

 CHANGELOG.md | 5 +++++
 VERSION      | 2 +-
 2 files changed, 6 insertions(+), 1 deletion(-)

---

commit fa1d0b6490f9ecd1d5525bd5208522ab8b6713a5
Author:     Robert Günzler <r@gnzler.io>
AuthorDate: Mon Jun 3 12:39:58 2019 +0200
Commit:     Robert Günzler <r@gnzler.io>
CommitDate: Mon Jun 3 14:49:33 2019 +0200

    Bump containerd/cgroups to dbea6f2bd41658b84b00417ceefa416b97
    
    Fixes issues with systemd version >=420 and non-existent cgroups.
    
    Change-type: patch
    Connects-to: https://github.com/containerd/cgroups/issues/76
    Connects-to: https://github.com/docker/for-linux/issues/545
    Signed-off-by: Robert Günzler <robertg@balena.io>

 vendor.conf                                       |  2 +-
 vendor/github.com/containerd/cgroups/README.md    | 14 +++-
 vendor/github.com/containerd/cgroups/blkio.go     | 55 ++++++++------
 vendor/github.com/containerd/cgroups/cgroup.go    | 92 +++++++++++++++++++++--
 vendor/github.com/containerd/cgroups/control.go   | 11 +++
 vendor/github.com/containerd/cgroups/cpuset.go    | 10 +--
 vendor/github.com/containerd/cgroups/devices.go   |  3 +
 vendor/github.com/containerd/cgroups/net_prio.go  |  2 +-
 vendor/github.com/containerd/cgroups/opts.go      | 61 +++++++++++++++
 vendor/github.com/containerd/cgroups/paths.go     |  5 +-
 vendor/github.com/containerd/cgroups/subsystem.go |  2 +-
 vendor/github.com/containerd/cgroups/systemd.go   | 33 +++++++-
 vendor/github.com/containerd/cgroups/utils.go     | 29 ++++++-
 13 files changed, 278 insertions(+), 41 deletions(-)

---

commit 44824fbca7f7d95f2716367bf63515dae0cd62e5
Author:     Robert Günzler <r@gnzler.io>
AuthorDate: Thu Apr 18 12:17:28 2019 +0200
Commit:     Giovanni Garufi <nazrhom@gmail.com>
CommitDate: Fri Apr 26 14:13:53 2019 +0200

    v18.9.3

 VERSION | 2 +-
 1 file changed, 1 insertion(+), 1 deletion(-)

---

commit 5bf4b8087e14af7d7af83402ebfa7b00fa416648
Author:     Robert Günzler <r@gnzler.io>
AuthorDate: Thu Apr 18 15:18:05 2019 +0200
Commit:     Robert Günzler <r@gnzler.io>
CommitDate: Thu Apr 18 15:19:39 2019 +0200

    dockerfile: Rename docker-init to balena-engine-init
    
    Signed-off-by: Robert Günzler <robertg@balena.io>

 hack/dockerfile/install/tini.installer | 2 +-
 1 file changed, 1 insertion(+), 1 deletion(-)

---

commit 49810dc78e6865f2305969616245160e4df9460d
Author:     Robert Günzler <r@gnzler.io>
AuthorDate: Tue Apr 9 15:35:11 2019 +0200
Commit:     Robert Günzler <r@gnzler.io>
CommitDate: Tue Apr 9 18:10:58 2019 +0200

    Fix double locking in the event handling code of OOM events
    
    Change-type: patch
    Signed-off-by: Robert Günzler <robertg@balena.io>

 daemon/monitor.go | 2 --
 1 file changed, 2 deletions(-)

---

commit 9eca1531dae212ac85daa658c19fa652ed085364
Author:     Robert Günzler <r@gnzler.io>
AuthorDate: Mon Apr 8 19:21:58 2019 +0200
Commit:     Robert Günzler <r@gnzler.io>
CommitDate: Mon Apr 8 19:30:49 2019 +0200

    integration-tests: Add test for containers with memory,cpu constraints
    
    The only test from integration/ that covers any resource constrained
    container scenarios is the OomKilled check in integration/container/kill_test.go
    
    This adds two addional checks that try to create, startk, stop and
    inspect a busybox container with:
    a) a memory constraint like: balena-engine run -m 32m ..
    b) a memory constraint like: balena-engine run -cpus ".5" ..
    
    Change-type: patch
    Signed-off-by: Robert Günzler <robertg@balena.io>

 integration/container/resource_test.go | 104 +++++++++++++++++++++++++++++++++
 1 file changed, 104 insertions(+)

---

commit e284726bae2fe269d721f67e6a74171289ad551b
Author:     Resin CI <34882892+resin-ci@users.noreply.github.com>
AuthorDate: Mon Mar 18 20:07:50 2019 +0200
Commit:     Resin CI <34882892+resin-ci@users.noreply.github.com>
CommitDate: Mon Mar 18 20:07:50 2019 +0200

    v17.13.4

 CHANGELOG.md | 5 +++++
 VERSION      | 2 +-
 2 files changed, 6 insertions(+), 1 deletion(-)

---

commit fe78e2c9a69313007c53c83fff4b5525fbc2ba45
Author:     Resin CI <34882892+resin-ci@users.noreply.github.com>
AuthorDate: Mon Feb 25 15:11:03 2019 +0100
Commit:     Resin CI <34882892+resin-ci@users.noreply.github.com>
CommitDate: Mon Feb 25 15:11:03 2019 +0100

    v17.13.3

 CHANGELOG.md | 5 +++++
 VERSION      | 2 +-
 2 files changed, 6 insertions(+), 1 deletion(-)

---

commit ad1ae964378edc3d61be6488ff5f973d7228edb4
Author:     Robert Günzler <r@gnzler.io>
AuthorDate: Mon Feb 25 12:19:04 2019 +0100
Commit:     Robert Günzler <r@gnzler.io>
CommitDate: Mon Feb 25 12:25:27 2019 +0100

    vendor: Update runc to include fix for opencontainers/runc#1766
    
    Change-type: patch
    Signed-off-by: Robert Günzler <robertg@balena.io>

 vendor.conf                                   | 2 +-
 vendor/github.com/opencontainers/runc/kill.go | 8 +-------
 2 files changed, 2 insertions(+), 8 deletions(-)

---

commit 88330c9aac5556d0abc7a5afcb4d906604a07fa2
Author:     Sebastiaan van Stijn <github@gone.nl>
AuthorDate: Thu Feb 14 03:29:08 2019 +0100
Commit:     Sebastiaan van Stijn <github@gone.nl>
CommitDate: Sat Feb 23 01:49:12 2019 +0100

    Revert "Merge pull request #240 from seemethere/bundle_me_up_1809"
    
    This reverts commit eb137ff1765faeb29c2d99025bfd8ed41836dd06, reversing
    changes made to a79fabbfe84117696a19671f4aa88b82d0f64fc1.

    Signed-off-by: Sebastiaan van Stijn <github@gone.nl>

 Dockerfile                             |   1 -
 git-bundles/CVE-2019-5736.bundle       | Bin 4038 -> 0 bytes
 hack/dockerfile/install/runc.installer |  17 +----------------
 3 files changed, 1 insertion(+), 17 deletions(-)

---

commit a7b3bbea5f4f775818641ec3ba522ca912f641c2
Author:     Resin CI <34882892+resin-ci@users.noreply.github.com>
AuthorDate: Thu Feb 21 15:33:03 2019 +0100
Commit:     Resin CI <34882892+resin-ci@users.noreply.github.com>
CommitDate: Thu Feb 21 15:33:03 2019 +0100

    v17.13.2

 CHANGELOG.md | 5 +++++
 VERSION      | 2 +-
 2 files changed, 6 insertions(+), 1 deletion(-)

---

commit 7486436688c818c256c6584baac9055bf3178bb1
Author:     Robert Günzler <r@gnzler.io>
AuthorDate: Fri Feb 15 16:34:16 2019 +0100
Commit:     Robert Günzler <r@gnzler.io>
CommitDate: Thu Feb 21 15:01:57 2019 +0100

    travis: Only run builds for PRs and master/version branches

    Previously we only filtered out gh-pages and versionist branches.
    Travis was building PRs from this repo twice since they always create
    a branch as well.

    This replaces the branch rule with one that allows builds for anything
    that is not a push (like PRs), the master branch and branches that fit
    the naming scheme: 18.09-balena
    
    Signed-off-by: Robert Günzler <robertg@balena.io>

 .travis.yml | 7 +++----
 1 file changed, 3 insertions(+), 4 deletions(-)

---

commit fa432f08380ff5921ac75346586ec4d51bde9fb2
Author:     Robert Günzler <r@gnzler.io>
AuthorDate: Fri Feb 15 14:47:30 2019 +0100
Commit:     Robert Günzler <r@gnzler.io>
CommitDate: Thu Feb 21 15:01:57 2019 +0100

    travis: Build for armv7 and aarch64 as well

    Makes use of build stages to parallelize jobs.
    The `travis_wait` command is used to prevent timeouts of emulated builds
    See https://docs.travis-ci.com/user/common-build-problems/#build-times-out-because-no-output-was-received
    
    Signed-off-by: Robert Günzler <robertg@balena.io>

 .travis.yml | 43 +++++++++++++++++++++++++++++++++++--------
 1 file changed, 35 insertions(+), 8 deletions(-)

---

commit be3cfa585b98a425b053d3d3475fbb7a4793d26d
Author:     Robert Günzler <r@gnzler.io>
AuthorDate: Thu Feb 7 18:51:28 2019 +0100
Commit:     Robert Günzler <r@gnzler.io>
CommitDate: Thu Feb 21 15:01:57 2019 +0100

    travis: Use the minimal machine
    
    Since we build in docker anyway we can save the time it usually takes to
    set up the Go environment.
    See https://docs.travis-ci.com/user/languages/minimal-and-generic/
    
    Change-type: patch
    Signed-off-by: Robert Günzler <robertg@balena.io>

 .travis.yml | 5 +----
 1 file changed, 1 insertion(+), 4 deletions(-)

---

commit a010e2e2bcd935046d6bcf20aa69e1c93de1c8fc
Author:     Resin CI <34882892+resin-ci@users.noreply.github.com>
AuthorDate: Thu Feb 21 10:21:55 2019 +0100
Commit:     Resin CI <34882892+resin-ci@users.noreply.github.com>
CommitDate: Thu Feb 21 10:21:55 2019 +0100

    v17.13.1

 CHANGELOG.md | 5 +++++
 VERSION      | 2 +-
 2 files changed, 6 insertions(+), 1 deletion(-)

---

commit d6fc6068c2e5d02e0d6ddcfb5ee5b8a73ddf23f8
Author:     Robert Günzler <r@gnzler.io>
AuthorDate: Mon Feb 11 17:11:26 2019 +0100
Commit:     Robert Günzler <r@gnzler.io>
CommitDate: Wed Feb 20 18:27:58 2019 +0100

    vendor: Update runc to include fix for CVE-2019-5736
    
    Change-type: patch
    Signed-off-by: Robert Günzler <robertg@balena.io>

 vendor.conf                                        |   2 +-
 .../runc/libcontainer/nsenter/cloned_binary.c      | 268 +++++++++++++++++++++
 .../runc/libcontainer/nsenter/nsexec.c             |  11 +
 3 files changed, 280 insertions(+), 1 deletion(-)

---

commit 325f6ee47a8edaf093ea9f829c26962310c83759
Author:     Sebastiaan van Stijn <github@gone.nl>
AuthorDate: Wed Jan 23 23:47:10 2019 +0100
Commit:     Sebastiaan van Stijn <github@gone.nl>
CommitDate: Sat Feb 9 11:05:52 2019 +0100

    [18.09] Bump Golang 1.10.8 (CVE-2019-6486)
    
    See the milestone for details;
    https://github.com/golang/go/issues?q=milestone%3AGo1.10.8+label%3ACherryPickApproved
    
    Signed-off-by: Sebastiaan van Stijn <github@gone.nl>

 Dockerfile         | 4 ++--
 Dockerfile.e2e     | 2 +-
 Dockerfile.simple  | 2 +-
 Dockerfile.windows | 2 +-
 4 files changed, 5 insertions(+), 5 deletions(-)

---

commit 03dfb0ba53cc5f64b746a25aa5ed8a48763ea223
Author:     Eli Uriegas <eli.uriegas@docker.com>
AuthorDate: Wed Feb 6 00:25:54 2019 +0000
Commit:     Eli Uriegas <eli.uriegas@docker.com>
CommitDate: Wed Feb 6 00:25:54 2019 +0000

    Apply git bundles for CVE-2019-5736
    
    A git bundle allows us keep the same SHA, giving us the ability to
    validate our patch against a known entity and allowing us to push
    directly from our private forks to public forks without having to
    re-apply any patches.
    
    Signed-off-by: Eli Uriegas <eli.uriegas@docker.com>

 Dockerfile                             |   1 +
 git-bundles/CVE-2019-5736.bundle       | Bin 0 -> 4038 bytes
 hack/dockerfile/install/runc.installer |  17 ++++++++++++++++-
 3 files changed, 17 insertions(+), 1 deletion(-)

---

commit 9e5d8ad10bae6db6737944dbdcdf2b2486a84593
Author:     Resin CI <34882892+resin-ci@users.noreply.github.com>
AuthorDate: Thu Jan 31 14:35:39 2019 +0100
Commit:     Resin CI <34882892+resin-ci@users.noreply.github.com>
CommitDate: Thu Jan 31 14:35:39 2019 +0100

    v17.13.0

 CHANGELOG.md | 19 +++++++++++++++----
 VERSION      |  2 +-
 2 files changed, 16 insertions(+), 5 deletions(-)

---

commit 744c524d03a34ac36cb94163db89ca93e6e7d503
Author:     Robert Günzler <r@gnzler.io>
AuthorDate: Thu Jan 3 16:39:59 2019 +0100
Commit:     Robert Günzler <r@gnzler.io>
CommitDate: Mon Jan 14 13:00:29 2019 +0100

    Add cli for tagging delta images
    
    Update vendor.conf and vendor/ to include https://github.com/balena-os/balena-engine-cli/pull/7
    
    Change-type: patch
    Signed-off-by: Robert Günzler <robertg@balena.io>

 vendor.conf                                             |  2 +-
 vendor/github.com/docker/cli/cli/command/image/delta.go | 14 ++++++++++++--
 2 files changed, 13 insertions(+), 3 deletions(-)

---

commit 594b651faad03a87ce7f95b0608f2f2fc2c38af5
Author:     Robert Günzler <r@gnzler.io>
AuthorDate: Wed Dec 19 13:07:11 2018 +0100
Commit:     Robert Günzler <r@gnzler.io>
CommitDate: Mon Jan 14 13:00:29 2019 +0100

    Add delta integration test
    
    Signed-off-by: Robert Günzler <robertg@balena.io>

 integration/image/delta_test.go | 72 +++++++++++++++++++++++++++++++++++++++++
 1 file changed, 72 insertions(+)

---

commit bd66b913256fa6dfee063cd4e8653d6b8fc63a4f
Author:     Robert Günzler <r@gnzler.io>
AuthorDate: Thu Dec 13 20:01:58 2018 +0100
Commit:     Robert Günzler <r@gnzler.io>
CommitDate: Mon Jan 14 13:00:29 2019 +0100

    Allow tagging of image deltas on creation
    
    Similar to how the build command allows tagging of images this allows
    specifying a repo:tag indentifier to tag the delta with
    
    Requires: https://github.com/balena-os/balena-engine-cli/pull/7
    Change-type: minor
    Signed-off-by: Robert Günzler <robertg@balena.io>

 api/server/router/image/backend.go      |  2 +-
 api/server/router/image/image_routes.go |  6 +++++-
 api/types/client.go                     |  5 +++++
 client/image_delta.go                   | 15 +++++++++++++--
 client/interface.go                     |  2 +-
 daemon/create.go                        | 27 +++++++++++++++++++++++++--
 6 files changed, 50 insertions(+), 7 deletions(-)

---

commit dceb2fc48071b78a8a828e0468a15a479515385f
Author:     Robert Günzler <r@gnzler.io>
AuthorDate: Fri Dec 7 18:05:17 2018 +0100
Commit:     Robert Günzler <r@gnzler.io>
CommitDate: Mon Jan 14 13:00:29 2019 +0100

    Backport fixes for hanging on low-entropy situations
    
    Update vendor.conf and vendor/ to include
    https://github.com/balena-os/balena-containerd/pull/2
    
    Connects-to: https://github.com/balena-os/balena-engine/issues/105
    Connects-to: https://github.com/containerd/containerd/issues/2451
    Signed-off-by: Robert Günzler <robertg@balena.io>

 vendor.conf                                        |  2 +-
 .../containerd/containerd/cmd/containerd/main.go   |  3 ++
 .../containerd/containerd/cmd/ctr/main.go          |  3 ++
 .../containerd/containerd/diff/walking/differ.go   |  2 +-
 .../containerd/containerd/pkg/seed/seed.go         | 38 ++++++++++++++++++++++
 .../containerd/containerd/pkg/seed/seed_linux.go   | 24 ++++++++++++++
 .../containerd/containerd/pkg/seed/seed_other.go   | 28 ++++++++++++++++
 .../containerd/containerd/rootfs/apply.go          |  2 +-
 .../containerd/services/leases/service.go          |  2 +-
 .../github.com/cyphar/filepath-securejoin/VERSION  |  1 -
 vendor/github.com/docker/cli/VERSION               |  1 -
 vendor/github.com/opencontainers/runc/VERSION      |  1 -
 vendor/github.com/prometheus/client_golang/VERSION |  1 -
 .../theupdateframework/notary/NOTARY_VERSION       |  1 -
 14 files changed, 100 insertions(+), 9 deletions(-)

---

commit 377d55202abb397dd7b71034bfc87e2df2f5d414
Author:     Resin CI <34882892+resin-ci@users.noreply.github.com>
AuthorDate: Mon Jan 14 13:00:22 2019 +0100
Commit:     Resin CI <34882892+resin-ci@users.noreply.github.com>
CommitDate: Mon Jan 14 13:00:22 2019 +0100

    v17.12.1

 CHANGELOG.md | 6 ++++++
 VERSION      | 1 +
 2 files changed, 7 insertions(+)

---

commit 44bda6ccd8ca4988e9ba8dedeb2df16a0c0d9cf2
Author:     Robert Günzler <r@gnzler.io>
AuthorDate: Tue Jan 8 12:24:46 2019 +0100
Commit:     Robert Günzler <r@gnzler.io>
CommitDate: Mon Jan 14 12:26:53 2019 +0100

    Enable travis-ci
    
    This skips the legacy integration test suite to not exeed the job time
    limit imposed by travis
    
    Signed-off-by: Robert Günzler <robertg@balena.io>

 .travis.yml | 26 ++++++++++++++++++++++++++
 1 file changed, 26 insertions(+)

---

commit acf966f17d0c263adf48f44b0a663b974b8ecb81
Author:     Robert Günzler <r@gnzler.io>
AuthorDate: Wed Jan 9 16:05:24 2019 +0100
Commit:     Robert Günzler <r@gnzler.io>
CommitDate: Mon Jan 14 12:26:53 2019 +0100

    Add repo.yml
    
    Connects-to: #126
    Signed-off-by: Robert Günzler <robertg@balena.io>

 .resinci.yml | 3 +++
 repo.yml     | 2 ++
 2 files changed, 5 insertions(+)

---

commit 24f71e39980e8a4c6eabcea16e0a9efce1660bbe
Author:     Madhu Venugopal <madhu@docker.com>
AuthorDate: Fri Dec 28 09:21:57 2018 -0800
Commit:     Madhu Venugopal <madhu@docker.com>
CommitDate: Fri Dec 28 09:40:26 2018 -0800

    Revert "[18.09 backport] API: fix status code on conflicting service names"
    
    Signed-off-by: Madhu Venugopal <madhu@docker.com>

 integration/internal/swarm/service.go              | 13 ++++------
 integration/service/create_test.go                 | 30 ----------------------
 vendor.conf                                        |  2 +-
 .../docker/swarmkit/manager/controlapi/service.go  | 18 +++++--------
 4 files changed, 12 insertions(+), 51 deletions(-)

---

commit a9ae6c7547466f754da01a53c6be455c555e6102
Author:     Andrew Hsu <andrewhsu@docker.com>
AuthorDate: Mon Dec 17 12:06:35 2018 +0000
Commit:     Andrew Hsu <andrewhsu@docker.com>
CommitDate: Mon Dec 17 12:06:35 2018 +0000

    Revert "Propagate context to exec delete"
    
    This reverts commit b6430ba41388f0300ceea95c10738cbe1a9a7b10.
    
    Signed-off-by: Andrew Hsu <andrewhsu@docker.com>

 libcontainerd/client_daemon.go | 2 +-
 1 file changed, 1 insertion(+), 1 deletion(-)

---

commit 8afe9f422dc0183ce48e1db09189ccbde634080a
Author:     Sebastiaan van Stijn <github@gone.nl>
AuthorDate: Fri Dec 14 00:44:49 2018 +0100
Commit:     Sebastiaan van Stijn <github@gone.nl>
CommitDate: Fri Dec 14 00:44:49 2018 +0100

    Bump Golang 1.10.6 (CVE-2018-16875)
    
    go1.10.6 (released 2018/12/14)
    
    - crypto/x509: CPU denial of service in chain validation golang/go#29233
    - cmd/go: directory traversal in "go get" via curly braces in import paths golang/go#29231
    - cmd/go: remote command execution during "go get -u" golang/go#29230
    
    See the Go 1.10.6 milestone on the issue tracker for details:
    https://github.com/golang/go/issues?q=milestone%3AGo1.10.6
    
    Signed-off-by: Sebastiaan van Stijn <github@gone.nl>

 Dockerfile         | 4 ++--
 Dockerfile.e2e     | 2 +-
 Dockerfile.simple  | 2 +-
 Dockerfile.windows | 2 +-
 4 files changed, 5 insertions(+), 5 deletions(-)

---

commit ad7105260f3c2ff32a375ff78dce9a96e01d87cb
Author:     Sebastiaan van Stijn <github@gone.nl>
AuthorDate: Mon Dec 10 12:18:32 2018 +0100
Commit:     Sebastiaan van Stijn <github@gone.nl>
CommitDate: Mon Dec 10 12:18:32 2018 +0100

    Update swarmkit to return correct error-codes on conflicting names
    
    This updates the swarmkit vendoring to the latest version in the bump_v18.09
    branch
    
    Signed-off-by: Sebastiaan van Stijn <github@gone.nl>

 vendor.conf                                            |  2 +-
 .../docker/swarmkit/manager/controlapi/service.go      | 18 ++++++++++++------
 2 files changed, 13 insertions(+), 7 deletions(-)

---

commit 00ad8e7c5730f3c50ae2e548b47d1340202f72b2
Author:     Sebastiaan van Stijn <github@gone.nl>
AuthorDate: Fri Nov 30 20:43:05 2018 +0100
Commit:     Sebastiaan van Stijn <github@gone.nl>
CommitDate: Fri Nov 30 20:43:05 2018 +0100

    Bump Go to 1.10.5
    
    go1.10.5 (released 2018/11/02) includes fixes to the go command, linker,
    runtime and the database/sql package. See the milestone on the issue
    tracker for details:
    
    List of changes; https://github.com/golang/go/issues?q=milestone%3AGo1.10.5
    
    Signed-off-by: Sebastiaan van Stijn <github@gone.nl>

 Dockerfile         | 4 ++--
 Dockerfile.e2e     | 2 +-
 Dockerfile.simple  | 2 +-
 Dockerfile.windows | 2 +-
 4 files changed, 5 insertions(+), 5 deletions(-)

---

commit 45654ed0126aadaf6c3293b0a32ca8cf15021626
Author:     Tonis Tiigi <tonistiigi@gmail.com>
AuthorDate: Tue Nov 6 10:45:02 2018 -0800
Commit:     Tonis Tiigi <tonistiigi@gmail.com>
CommitDate: Tue Nov 6 10:52:34 2018 -0800

    builder: update copy to 0.1.9
    
    Signed-off-by: Tonis Tiigi <tonistiigi@gmail.com>

 builder/builder-next/builder.go | 2 +-
 1 file changed, 1 insertion(+), 1 deletion(-)

---

commit e1783a72d1b84bc3e32470c468d14445e5fba8db
Author:     Sebastiaan van Stijn <github@gone.nl>
AuthorDate: Tue Nov 6 12:33:50 2018 +0100
Commit:     Sebastiaan van Stijn <github@gone.nl>
CommitDate: Tue Nov 6 12:39:04 2018 +0100

    [18.09 backport] update libnetwork to fix iptables compatibility on debian
    
    Fixes a compatibility issue on recent debian versions, where iptables now uses
    nft by default.
    
    Signed-off-by: Sebastiaan van Stijn <github@gone.nl>

 hack/dockerfile/install/proxy.installer                  | 2 +-
 vendor.conf                                              | 2 +-
 vendor/github.com/docker/libnetwork/iptables/iptables.go | 9 +++++++--
 3 files changed, 9 insertions(+), 4 deletions(-)

---

commit 46dfcd83bf1bb820840df91629c04c47b32d1e21
Author:     Anshul Pundir <anshul.pundir@docker.com>
AuthorDate: Wed Oct 31 16:03:15 2018 -0700
Commit:     Anshul Pundir <anshul.pundir@docker.com>
CommitDate: Wed Oct 31 16:04:51 2018 -0700

    [18.09] Vendor swarmkit to 6186e40fb04a7681e25a9101dbc7418c37ef0c8b
    
    Signed-off-by: Anshul Pundir <anshul.pundir@docker.com>

 vendor.conf                                                    |  2 +-
 .../docker/swarmkit/manager/state/raft/transport/transport.go  | 10 ++++++++++
 2 files changed, 11 insertions(+), 1 deletion(-)

---

commit 7a8d0d21c9e047852f81cac8f6eeacfe565fa00c
Author:     Paulo Castro <paulo@balena.io>
AuthorDate: Tue Oct 23 23:26:41 2018 +0100
Commit:     Paulo Castro <paulo@balena.io>
CommitDate: Tue Oct 23 23:26:41 2018 +0100

    docs: Fix Docker capitalisation in balenaEngine docs
    
    Change-type: patch
    Signed-off-by: Paulo Castro <paulo@balena.io>

 docs/getting-started.md | 4 ++--
 1 file changed, 2 insertions(+), 2 deletions(-)

---

commit 2fb1f862d65a82617ce42290f76c460fe079ea14
Author:     Paulo Castro <paulo@balena.io>
AuthorDate: Tue Oct 23 20:23:18 2018 +0100
Commit:     Paulo Castro <paulo@balena.io>
CommitDate: Tue Oct 23 20:57:22 2018 +0100

    Update balenaEngine logo in README.md
    
    Change-type: patch
    Signed-off-by: Paulo Castro <paulo@balena.io>

 README.md                               |  2 +-
 docs/static_files/balena-engine.svg     | 17 ++++++++++
 docs/static_files/balena-logo-black.svg | 56 ---------------------------------
 3 files changed, 18 insertions(+), 57 deletions(-)

---

commit 6ee7d86a12fe83953eff0efd4de5878b4ff6814d
Author:     Sebastiaan van Stijn <github@gone.nl>
AuthorDate: Tue Oct 23 13:37:15 2018 +0200
Commit:     Sebastiaan van Stijn <github@gone.nl>
CommitDate: Tue Oct 23 13:37:15 2018 +0200

    Add note that we use the bump_v18.09 branch for SwarmKit
    
    Signed-off-by: Sebastiaan van Stijn <github@gone.nl>

 vendor.conf | 2 +-
 1 file changed, 1 insertion(+), 1 deletion(-)

---

commit 1222a7081ac9ebb0830a6c8008142258c49800b5
Author:     Drew Erny <drew.erny@docker.com>
AuthorDate: Mon Oct 22 13:24:55 2018 -0500
Commit:     Drew Erny <drew.erny@docker.com>
CommitDate: Mon Oct 22 15:10:20 2018 -0500

    Bump swarmkit
    
    Signed-off-by: Drew Erny <drew.erny@docker.com>

 vendor.conf                                                |  2 +-
 vendor/github.com/docker/swarmkit/agent/agent.go           |  2 +-
 vendor/github.com/docker/swarmkit/agent/errors.go          |  7 +------
 vendor/github.com/docker/swarmkit/agent/exec/controller.go |  2 +-
 .../docker/swarmkit/agent/exec/controller_stub.go          |  5 +++--
 vendor/github.com/docker/swarmkit/agent/exec/executor.go   |  3 ++-
 vendor/github.com/docker/swarmkit/agent/helpers.go         |  2 +-
 vendor/github.com/docker/swarmkit/agent/reporter.go        |  2 +-
 vendor/github.com/docker/swarmkit/agent/resource.go        |  3 ++-
 vendor/github.com/docker/swarmkit/agent/session.go         |  7 +++----
 vendor/github.com/docker/swarmkit/agent/task.go            |  2 +-
 vendor/github.com/docker/swarmkit/agent/worker.go          |  2 +-
 .../swarmkit/api/genericresource/resource_management.go    |  1 +
 .../docker/swarmkit/api/genericresource/validate.go        |  1 +
 vendor/github.com/docker/swarmkit/ca/auth.go               |  2 +-
 vendor/github.com/docker/swarmkit/ca/certificates.go       |  2 +-
 vendor/github.com/docker/swarmkit/ca/config.go             |  8 +++-----
 vendor/github.com/docker/swarmkit/ca/external.go           |  2 +-
 vendor/github.com/docker/swarmkit/ca/forward.go            |  3 ++-
 vendor/github.com/docker/swarmkit/ca/renewer.go            |  2 +-
 vendor/github.com/docker/swarmkit/ca/server.go             |  2 +-
 vendor/github.com/docker/swarmkit/ca/transport.go          |  8 +-------
 vendor/github.com/docker/swarmkit/log/context.go           |  2 +-
 vendor/github.com/docker/swarmkit/log/grpc.go              |  3 ++-
 .../docker/swarmkit/manager/allocator/allocator.go         |  2 +-
 .../manager/allocator/cnmallocator/networkallocator.go     |  5 ++---
 .../manager/allocator/cnmallocator/portallocator.go        |  4 ++--
 .../docker/swarmkit/manager/allocator/network.go           |  6 +++++-
 .../docker/swarmkit/manager/constraint/constraint.go       |  4 ++--
 .../docker/swarmkit/manager/controlapi/cluster.go          |  2 +-
 .../docker/swarmkit/manager/controlapi/config.go           |  2 +-
 .../docker/swarmkit/manager/controlapi/network.go          |  2 +-
 .../github.com/docker/swarmkit/manager/controlapi/node.go  |  2 +-
 .../docker/swarmkit/manager/controlapi/secret.go           |  2 +-
 .../docker/swarmkit/manager/controlapi/service.go          |  8 ++++----
 .../github.com/docker/swarmkit/manager/controlapi/task.go  |  3 ++-
 .../docker/swarmkit/manager/dispatcher/dispatcher.go       | 14 +++++---------
 .../github.com/docker/swarmkit/manager/dispatcher/nodes.go |  2 +-
 vendor/github.com/docker/swarmkit/manager/health/health.go |  2 +-
 .../docker/swarmkit/manager/keymanager/keymanager.go       |  2 +-
 .../github.com/docker/swarmkit/manager/logbroker/broker.go |  2 +-
 .../docker/swarmkit/manager/logbroker/subscription.go      |  2 +-
 vendor/github.com/docker/swarmkit/manager/manager.go       |  6 ++----
 .../docker/swarmkit/manager/metrics/collector.go           |  3 ---
 .../docker/swarmkit/manager/orchestrator/global/global.go  | 11 ++---------
 .../swarmkit/manager/orchestrator/replicated/replicated.go |  3 ++-
 .../swarmkit/manager/orchestrator/replicated/services.go   |  2 +-
 .../swarmkit/manager/orchestrator/replicated/slot.go       |  3 ++-
 .../swarmkit/manager/orchestrator/replicated/tasks.go      |  3 ++-
 .../swarmkit/manager/orchestrator/restart/restart.go       |  2 +-
 .../docker/swarmkit/manager/orchestrator/service.go        |  3 ++-
 .../docker/swarmkit/manager/orchestrator/taskinit/init.go  |  4 ++--
 .../manager/orchestrator/taskreaper/task_reaper.go         |  2 +-
 .../docker/swarmkit/manager/orchestrator/update/updater.go |  3 +--
 .../docker/swarmkit/manager/raftselector/raftselector.go   |  3 +--
 .../docker/swarmkit/manager/resourceapi/allocator.go       |  2 +-
 vendor/github.com/docker/swarmkit/manager/role_manager.go  |  2 +-
 .../docker/swarmkit/manager/scheduler/nodeinfo.go          |  2 +-
 .../docker/swarmkit/manager/scheduler/scheduler.go         |  2 +-
 .../github.com/docker/swarmkit/manager/state/proposer.go   |  3 ++-
 .../github.com/docker/swarmkit/manager/state/raft/raft.go  | 12 +++---------
 .../docker/swarmkit/manager/state/raft/storage.go          |  2 +-
 .../docker/swarmkit/manager/state/raft/storage/storage.go  |  3 +--
 .../docker/swarmkit/manager/state/raft/storage/walwrap.go  |  2 +-
 .../docker/swarmkit/manager/state/raft/transport/peer.go   |  3 +--
 .../swarmkit/manager/state/raft/transport/transport.go     |  3 +--
 .../github.com/docker/swarmkit/manager/state/raft/util.go  |  3 +--
 .../docker/swarmkit/manager/state/store/memory.go          |  4 ++--
 .../github.com/docker/swarmkit/manager/watchapi/server.go  |  2 +-
 vendor/github.com/docker/swarmkit/node/node.go             |  7 ++-----
 70 files changed, 106 insertions(+), 135 deletions(-)

---

commit 8a1313357a3c41a33a72c97fa5cfa180918e4ba6
Author:     Paulo Castro <paulo@balena.io>
AuthorDate: Thu Oct 18 23:06:31 2018 +0200
Commit:     Paulo Castro <paulo@balena.io>
CommitDate: Fri Oct 19 15:41:50 2018 +0200

    Update CHANGELOG.md
    
    Signed-off-by: Paulo Castro <paulo@balena.io>

 CHANGELOG.md | 67 ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 1 file changed, 67 insertions(+)

---

commit fd1fe0b702571865cc77d66937e4ca570b5b9cc3
Author:     Chris Telfer <ctelfer@docker.com>
AuthorDate: Thu Oct 18 10:52:57 2018 -0400
Commit:     Chris Telfer <ctelfer@docker.com>
CommitDate: Thu Oct 18 10:52:57 2018 -0400

    Bump libnetwork to 6da50d19 for DSR changes
    
    Bump libnetwork to 6da50d1978302f04c3e2089e29112ea24812f05b which
    is the current tip of libnetwork's bump_18.09 branch to get the DSR load
    balancing mode option changes for the 18.09 branch of Docker CE.
    
    Signed-off-by: Chris Telfer <ctelfer@docker.com>

 hack/dockerfile/install/proxy.installer            |  2 +-
 vendor.conf                                        |  2 +-
 vendor/github.com/docker/libnetwork/controller.go  | 37 +++++++---
 .../github.com/docker/libnetwork/ipvs/constants.go | 20 ++++++
 vendor/github.com/docker/libnetwork/network.go     | 81 +++++++++++++---------
 .../docker/libnetwork/osl/namespace_linux.go       | 30 ++++++++
 vendor/github.com/docker/libnetwork/osl/sandbox.go |  4 ++
 vendor/github.com/docker/libnetwork/sandbox.go     | 22 ++++++
 .../github.com/docker/libnetwork/service_linux.go  | 19 +++--
 9 files changed, 166 insertions(+), 51 deletions(-)

---

commit 63f30e90f9bd709821dd37412a68220e3dcf495b
Author:     Paulo Castro <paulo@balena.io>
AuthorDate: Wed Oct 17 16:59:26 2018 +0200
Commit:     Paulo Castro <paulo@balena.io>
CommitDate: Wed Oct 17 16:59:26 2018 +0200

    Update install.sh script for v17.12.0 release
    
    Signed-off-by: Paulo Castro <paulo@balena.io>

 contrib/install.sh | 2 +-
 1 file changed, 1 insertion(+), 1 deletion(-)

---

commit 25755b07fe92e331870307de4ab3641c88abe774
Author:     Paulo Castro <pefcastro@gmail.com>
AuthorDate: Mon Oct 15 20:47:15 2018 +0200
Commit:     Paulo Castro <paulo@balena.io>
CommitDate: Wed Oct 17 12:32:20 2018 +0200

    Use Balena's fork of golang.org/x/sys/unix (ARM SyncFileRange syscall)
    
    Signed-off-by: Paulo Castro <paulo@balena.io>

 Dockerfile                                         |    2 +-
 Dockerfile.build.aarch64                           |    2 +-
 Dockerfile.build.arm                               |    2 +-
 Dockerfile.build.i386                              |    2 +-
 Dockerfile.build.x86_64                            |    2 +-
 pkg/ioutils/eagerwriter.go                         |    7 +-
 vendor.conf                                        |    2 +-
 vendor/golang.org/x/sys/unix/affinity_linux.go     |  124 ++
 vendor/golang.org/x/sys/unix/aliases.go            |   14 +
 vendor/golang.org/x/sys/unix/asm_dragonfly_amd64.s |   10 +-
 vendor/golang.org/x/sys/unix/asm_linux_386.s       |   36 +-
 vendor/golang.org/x/sys/unix/asm_linux_amd64.s     |   30 +-
 vendor/golang.org/x/sys/unix/asm_linux_arm.s       |   35 +-
 vendor/golang.org/x/sys/unix/asm_linux_arm64.s     |   30 +-
 vendor/golang.org/x/sys/unix/asm_linux_mips64x.s   |   36 +-
 vendor/golang.org/x/sys/unix/asm_linux_mipsx.s     |   33 +-
 vendor/golang.org/x/sys/unix/asm_linux_ppc64x.s    |   30 +-
 vendor/golang.org/x/sys/unix/asm_linux_s390x.s     |   28 +
 vendor/golang.org/x/sys/unix/cap_freebsd.go        |   30 +-
 vendor/golang.org/x/sys/unix/constants.go          |    2 +-
 vendor/golang.org/x/sys/unix/dev_aix_ppc.go        |   27 +
 vendor/golang.org/x/sys/unix/dev_aix_ppc64.go      |   29 +
 vendor/golang.org/x/sys/unix/dirent.go             |   91 +-
 vendor/golang.org/x/sys/unix/env_unix.go           |    6 +-
 vendor/golang.org/x/sys/unix/env_unset.go          |   14 -
 .../golang.org/x/sys/unix/{flock.go => fcntl.go}   |   10 +
 .../{flock_linux_32bit.go => fcntl_linux_32bit.go} |    0
 vendor/golang.org/x/sys/unix/file_unix.go          |   27 -
 vendor/golang.org/x/sys/unix/gccgo.go              |   16 +
 vendor/golang.org/x/sys/unix/gccgo_c.c             |   10 +-
 vendor/golang.org/x/sys/unix/ioctl.go              |   30 +
 vendor/golang.org/x/sys/unix/openbsd_pledge.go     |   77 +-
 vendor/golang.org/x/sys/unix/pagesize_unix.go      |    2 +-
 vendor/golang.org/x/sys/unix/race0.go              |    2 +-
 vendor/golang.org/x/sys/unix/sockcmsg_unix.go      |    4 +-
 vendor/golang.org/x/sys/unix/str.go                |    2 +-
 vendor/golang.org/x/sys/unix/syscall.go            |   13 +-
 vendor/golang.org/x/sys/unix/syscall_aix.go        |  564 ++++++
 vendor/golang.org/x/sys/unix/syscall_aix_ppc.go    |   34 +
 ...scall_openbsd_amd64.go => syscall_aix_ppc64.go} |   19 +-
 vendor/golang.org/x/sys/unix/syscall_bsd.go        |   68 +-
 vendor/golang.org/x/sys/unix/syscall_darwin.go     |  217 ++-
 vendor/golang.org/x/sys/unix/syscall_darwin_arm.go |    4 +
 vendor/golang.org/x/sys/unix/syscall_dragonfly.go  |  146 +-
 vendor/golang.org/x/sys/unix/syscall_freebsd.go    |  340 +---
 vendor/golang.org/x/sys/unix/syscall_linux.go      |  376 ++--
 vendor/golang.org/x/sys/unix/syscall_linux_386.go  |   17 +-
 .../golang.org/x/sys/unix/syscall_linux_amd64.go   |   36 +-
 vendor/golang.org/x/sys/unix/syscall_linux_arm.go  |   16 +-
 .../golang.org/x/sys/unix/syscall_linux_arm64.go   |   62 +-
 vendor/golang.org/x/sys/unix/syscall_linux_gc.go   |   14 +
 .../golang.org/x/sys/unix/syscall_linux_gc_386.go  |   16 +
 .../x/sys/unix/syscall_linux_gccgo_386.go          |   30 +
 .../x/sys/unix/syscall_linux_gccgo_arm.go          |   20 +
 .../golang.org/x/sys/unix/syscall_linux_mips64x.go |   13 +-
 .../golang.org/x/sys/unix/syscall_linux_mipsx.go   |   21 +-
 .../golang.org/x/sys/unix/syscall_linux_ppc64x.go  |   31 +-
 ...all_linux_arm64.go => syscall_linux_riscv64.go} |   66 +-
 .../golang.org/x/sys/unix/syscall_linux_s390x.go   |   18 +
 .../golang.org/x/sys/unix/syscall_linux_sparc64.go |    4 +
 vendor/golang.org/x/sys/unix/syscall_netbsd.go     |  150 +-
 vendor/golang.org/x/sys/unix/syscall_no_getwd.go   |   11 -
 vendor/golang.org/x/sys/unix/syscall_openbsd.go    |  162 +-
 .../golang.org/x/sys/unix/syscall_openbsd_amd64.go |    4 +
 vendor/golang.org/x/sys/unix/syscall_solaris.go    |   66 +-
 .../golang.org/x/sys/unix/syscall_solaris_amd64.go |    5 -
 vendor/golang.org/x/sys/unix/syscall_unix.go       |  115 +-
 vendor/golang.org/x/sys/unix/timestruct.go         |   22 +-
 vendor/golang.org/x/sys/unix/xattr_bsd.go          |  231 +++
 vendor/golang.org/x/sys/unix/zerrors_aix_ppc.go    | 1372 ++++++++++++++
 vendor/golang.org/x/sys/unix/zerrors_aix_ppc64.go  | 1373 ++++++++++++++
 vendor/golang.org/x/sys/unix/zerrors_darwin_386.go |  388 ++--
 .../golang.org/x/sys/unix/zerrors_darwin_amd64.go  |  388 ++--
 vendor/golang.org/x/sys/unix/zerrors_darwin_arm.go |  388 ++--
 .../golang.org/x/sys/unix/zerrors_darwin_arm64.go  |  388 ++--
 .../x/sys/unix/zerrors_dragonfly_amd64.go          |  354 ++--
 .../golang.org/x/sys/unix/zerrors_freebsd_386.go   |  306 +--
 .../golang.org/x/sys/unix/zerrors_freebsd_amd64.go |  306 +--
 .../golang.org/x/sys/unix/zerrors_freebsd_arm.go   |  306 +--
 vendor/golang.org/x/sys/unix/zerrors_linux_386.go  |  787 ++++++--
 .../golang.org/x/sys/unix/zerrors_linux_amd64.go   |  786 ++++++--
 vendor/golang.org/x/sys/unix/zerrors_linux_arm.go  |  788 ++++++--
 .../golang.org/x/sys/unix/zerrors_linux_arm64.go   |  787 ++++++--
 vendor/golang.org/x/sys/unix/zerrors_linux_mips.go |  790 ++++++--
 .../golang.org/x/sys/unix/zerrors_linux_mips64.go  |  790 ++++++--
 .../x/sys/unix/zerrors_linux_mips64le.go           |  790 ++++++--
 .../golang.org/x/sys/unix/zerrors_linux_mipsle.go  |  790 ++++++--
 .../golang.org/x/sys/unix/zerrors_linux_ppc64.go   |  786 ++++++--
 .../golang.org/x/sys/unix/zerrors_linux_ppc64le.go |  786 ++++++--
 ...ors_linux_amd64.go => zerrors_linux_riscv64.go} |  803 ++++++--
 .../golang.org/x/sys/unix/zerrors_linux_s390x.go   |  785 ++++++--
 .../golang.org/x/sys/unix/zerrors_linux_sparc64.go |  350 ++--
 vendor/golang.org/x/sys/unix/zerrors_netbsd_386.go |  322 ++--
 .../golang.org/x/sys/unix/zerrors_netbsd_amd64.go  |  322 ++--
 vendor/golang.org/x/sys/unix/zerrors_netbsd_arm.go |  322 ++--
 .../golang.org/x/sys/unix/zerrors_openbsd_386.go   |  322 ++--
 .../golang.org/x/sys/unix/zerrors_openbsd_amd64.go |  572 ++++--
 .../golang.org/x/sys/unix/zerrors_openbsd_arm.go   |  322 ++--
 .../golang.org/x/sys/unix/zerrors_solaris_amd64.go |  371 ++--
 vendor/golang.org/x/sys/unix/zsyscall_aix_ppc.go   | 1519 +++++++++++++++
 vendor/golang.org/x/sys/unix/zsyscall_aix_ppc64.go | 1519 +++++++++++++++
 .../golang.org/x/sys/unix/zsyscall_darwin_386.go   |  163 +-
 .../golang.org/x/sys/unix/zsyscall_darwin_amd64.go |  163 +-
 .../golang.org/x/sys/unix/zsyscall_darwin_arm.go   |  167 +-
 .../golang.org/x/sys/unix/zsyscall_darwin_arm64.go |  163 +-
 .../x/sys/unix/zsyscall_dragonfly_amd64.go         |   57 +
 .../golang.org/x/sys/unix/zsyscall_freebsd_386.go  |   55 +-
 .../x/sys/unix/zsyscall_freebsd_amd64.go           |   55 +-
 .../golang.org/x/sys/unix/zsyscall_freebsd_arm.go  |   55 +-
 vendor/golang.org/x/sys/unix/zsyscall_linux_386.go |  339 +++-
 .../golang.org/x/sys/unix/zsyscall_linux_amd64.go  |  350 +++-
 vendor/golang.org/x/sys/unix/zsyscall_linux_arm.go |  343 +++-
 .../golang.org/x/sys/unix/zsyscall_linux_arm64.go  |  253 ++-
 .../golang.org/x/sys/unix/zsyscall_linux_mips.go   |  369 +++-
 .../golang.org/x/sys/unix/zsyscall_linux_mips64.go |  319 +++-
 .../x/sys/unix/zsyscall_linux_mips64le.go          |  319 +++-
 .../golang.org/x/sys/unix/zsyscall_linux_mipsle.go |  369 +++-
 .../golang.org/x/sys/unix/zsyscall_linux_ppc64.go  |  356 +++-
 .../x/sys/unix/zsyscall_linux_ppc64le.go           |  356 +++-
 ...ll_linux_amd64.go => zsyscall_linux_riscv64.go} |  385 ++--
 .../golang.org/x/sys/unix/zsyscall_linux_s390x.go  |  324 +++-
 .../x/sys/unix/zsyscall_linux_sparc64.go           |  461 ++++-
 .../golang.org/x/sys/unix/zsyscall_netbsd_386.go   |  296 +++
 .../golang.org/x/sys/unix/zsyscall_netbsd_amd64.go |  296 +++
 .../golang.org/x/sys/unix/zsyscall_netbsd_arm.go   |  296 +++
 .../golang.org/x/sys/unix/zsyscall_openbsd_386.go  |   93 +
 .../x/sys/unix/zsyscall_openbsd_amd64.go           |   93 +
 .../golang.org/x/sys/unix/zsyscall_openbsd_arm.go  |   93 +
 .../x/sys/unix/zsyscall_solaris_amd64.go           |  311 +++
 .../golang.org/x/sys/unix/zsysctl_openbsd_386.go   |    2 +-
 .../golang.org/x/sys/unix/zsysctl_openbsd_amd64.go |   40 +-
 .../golang.org/x/sys/unix/zsysctl_openbsd_arm.go   |    2 +-
 vendor/golang.org/x/sys/unix/zsysnum_darwin_386.go |   60 +-
 .../golang.org/x/sys/unix/zsysnum_darwin_amd64.go  |   60 +-
 vendor/golang.org/x/sys/unix/zsysnum_darwin_arm.go |   14 +-
 .../golang.org/x/sys/unix/zsysnum_darwin_arm64.go  |   14 +-
 .../golang.org/x/sys/unix/zsysnum_freebsd_386.go   |  736 ++++----
 .../golang.org/x/sys/unix/zsysnum_freebsd_amd64.go |  736 ++++----
 .../golang.org/x/sys/unix/zsysnum_freebsd_arm.go   |  736 ++++----
 vendor/golang.org/x/sys/unix/zsysnum_linux_386.go  |    2 +
 .../golang.org/x/sys/unix/zsysnum_linux_amd64.go   |    2 +
 vendor/golang.org/x/sys/unix/zsysnum_linux_arm.go  |    1 +
 .../golang.org/x/sys/unix/zsysnum_linux_arm64.go   |    1 +
 vendor/golang.org/x/sys/unix/zsysnum_linux_mips.go |    2 +
 .../golang.org/x/sys/unix/zsysnum_linux_mips64.go  |    2 +
 .../x/sys/unix/zsysnum_linux_mips64le.go           |    2 +
 .../golang.org/x/sys/unix/zsysnum_linux_mipsle.go  |    2 +
 .../golang.org/x/sys/unix/zsysnum_linux_ppc64.go   |    5 +
 .../golang.org/x/sys/unix/zsysnum_linux_ppc64le.go |    5 +
 ...num_linux_arm64.go => zsysnum_linux_riscv64.go} |    6 +-
 .../golang.org/x/sys/unix/zsysnum_linux_s390x.go   |   48 +-
 .../golang.org/x/sys/unix/zsysnum_linux_sparc64.go |    2 +-
 vendor/golang.org/x/sys/unix/zsysnum_netbsd_386.go |    2 +-
 .../golang.org/x/sys/unix/zsysnum_netbsd_amd64.go  |    2 +-
 vendor/golang.org/x/sys/unix/zsysnum_netbsd_arm.go |    2 +-
 .../golang.org/x/sys/unix/zsysnum_openbsd_386.go   |    2 +-
 .../golang.org/x/sys/unix/zsysnum_openbsd_amd64.go |   25 +-
 .../golang.org/x/sys/unix/zsysnum_openbsd_arm.go   |    2 +-
 .../golang.org/x/sys/unix/zsysnum_solaris_amd64.go |   13 -
 vendor/golang.org/x/sys/unix/ztypes_aix_ppc.go     |  345 ++++
 vendor/golang.org/x/sys/unix/ztypes_aix_ppc64.go   |  354 ++++
 vendor/golang.org/x/sys/unix/ztypes_darwin_386.go  |  130 +-
 .../golang.org/x/sys/unix/ztypes_darwin_amd64.go   |  176 +-
 vendor/golang.org/x/sys/unix/ztypes_darwin_arm.go  |  130 +-
 .../golang.org/x/sys/unix/ztypes_darwin_arm64.go   |  176 +-
 .../x/sys/unix/ztypes_dragonfly_amd64.go           |  146 +-
 vendor/golang.org/x/sys/unix/ztypes_freebsd_386.go |   39 +-
 .../golang.org/x/sys/unix/ztypes_freebsd_amd64.go  |   39 +-
 vendor/golang.org/x/sys/unix/ztypes_freebsd_arm.go |   39 +-
 vendor/golang.org/x/sys/unix/ztypes_linux_386.go   | 1492 +++++++++++++--
 vendor/golang.org/x/sys/unix/ztypes_linux_amd64.go | 1498 +++++++++++++--
 vendor/golang.org/x/sys/unix/ztypes_linux_arm.go   | 1511 +++++++++++++--
 vendor/golang.org/x/sys/unix/ztypes_linux_arm64.go | 1500 +++++++++++++--
 vendor/golang.org/x/sys/unix/ztypes_linux_mips.go  | 1471 +++++++++++++--
 .../golang.org/x/sys/unix/ztypes_linux_mips64.go   | 1496 +++++++++++++--
 .../golang.org/x/sys/unix/ztypes_linux_mips64le.go | 1496 +++++++++++++--
 .../golang.org/x/sys/unix/ztypes_linux_mipsle.go   | 1471 +++++++++++++--
 vendor/golang.org/x/sys/unix/ztypes_linux_ppc64.go | 1506 +++++++++++++--
 .../golang.org/x/sys/unix/ztypes_linux_ppc64le.go  | 1506 +++++++++++++--
 .../golang.org/x/sys/unix/ztypes_linux_riscv64.go  | 1995 ++++++++++++++++++++
 vendor/golang.org/x/sys/unix/ztypes_linux_s390x.go | 1406 ++++++++++++--
 .../golang.org/x/sys/unix/ztypes_linux_sparc64.go  |  218 ++-
 vendor/golang.org/x/sys/unix/ztypes_netbsd_386.go  |   48 +-
 .../golang.org/x/sys/unix/ztypes_netbsd_amd64.go   |   48 +-
 vendor/golang.org/x/sys/unix/ztypes_netbsd_arm.go  |   48 +-
 vendor/golang.org/x/sys/unix/ztypes_openbsd_386.go |  137 +-
 .../golang.org/x/sys/unix/ztypes_openbsd_amd64.go  |  228 ++-
 vendor/golang.org/x/sys/unix/ztypes_openbsd_arm.go |  137 +-
 .../golang.org/x/sys/unix/ztypes_solaris_amd64.go  |  201 +-
 vendor/golang.org/x/sys/windows/aliases.go         |   13 +
 vendor/golang.org/x/sys/windows/asm_windows_386.s  |    4 +-
 .../golang.org/x/sys/windows/asm_windows_amd64.s   |    2 +-
 vendor/golang.org/x/sys/windows/asm_windows_arm.s  |   11 +
 vendor/golang.org/x/sys/windows/env_unset.go       |   15 -
 vendor/golang.org/x/sys/windows/env_windows.go     |    4 +
 vendor/golang.org/x/sys/windows/registry/key.go    |   10 +-
 .../x/sys/windows/registry/zsyscall_windows.go     |    2 +-
 .../golang.org/x/sys/windows/security_windows.go   |   45 +-
 vendor/golang.org/x/sys/windows/service.go         |   19 +
 .../golang.org/x/sys/windows/svc/debug/service.go  |    2 +-
 vendor/golang.org/x/sys/windows/svc/mgr/config.go  |   40 +-
 .../golang.org/x/sys/windows/svc/mgr/recovery.go   |  135 ++
 vendor/golang.org/x/sys/windows/svc/service.go     |    4 +-
 vendor/golang.org/x/sys/windows/svc/sys_amd64.s    |    2 +-
 vendor/golang.org/x/sys/windows/svc/sys_arm.s      |   38 +
 vendor/golang.org/x/sys/windows/syscall.go         |    3 +
 vendor/golang.org/x/sys/windows/syscall_windows.go |  191 +-
 vendor/golang.org/x/sys/windows/types_windows.go   |  237 ++-
 .../golang.org/x/sys/windows/types_windows_arm.go  |   22 +
 .../golang.org/x/sys/windows/zsyscall_windows.go   |  274 ++-
 210 files changed, 46172 insertions(+), 10008 deletions(-)

---

commit 3f8ae5574074f1e9cd6687e085c32cd77718c929
Author:     Paulo Castro <paulo@balena.io>
AuthorDate: Tue Oct 16 17:23:41 2018 +0200
Commit:     Paulo Castro <paulo@balena.io>
CommitDate: Tue Oct 16 17:45:26 2018 +0200

    Have the balena-engine binary accept being called as balena, balenad...
    
    Addresses resin-os/meta-resin#1215
    
    Signed-off-by: Paulo Castro <paulo@balena.io>

 cmd/balena-engine/main.go | 14 +++++++-------
 cmd/mobynit/main.go       |  2 +-
 2 files changed, 8 insertions(+), 8 deletions(-)

---

commit a9c160256fee050d7021a9631dd3fe11399a8620
Author:     Paulo Castro <pefcastro@gmail.com>
AuthorDate: Fri Oct 12 13:28:07 2018 +0100
Commit:     Paulo Castro <pefcastro@gmail.com>
CommitDate: Mon Oct 15 23:38:26 2018 +0200

    Run vndr tool for balena-cli and balena-libnetwork (balenaEngine rename)
    
    Signed-off-by: Paulo Castro <paulo@resin.io>

 vendor.conf                                                   | 8 ++++----
 vendor/github.com/containerd/containerd/plugin/plugin_go18.go | 2 +-
 2 files changed, 5 insertions(+), 5 deletions(-)

---

commit 40c33e33d23882936ec402b1563fca7456af001d
Author:     Paulo Castro <pefcastro@gmail.com>
AuthorDate: Thu Oct 11 13:27:57 2018 +0100
Commit:     Paulo Castro <pefcastro@gmail.com>
CommitDate: Thu Oct 11 18:33:37 2018 +0100

    Fix daemon/cluster/executor/container/ unit tests
    
    Deleted unsupported SetContainerDependencyStore method (swarm feature)
    from the daemon.cluster.executor.Backend interface.
    
    Signed-off-by: Paulo Castro <paulo@resin.io>

 Makefile                                     | 2 +-
 daemon/cluster/executor/backend.go           | 2 --
 daemon/cluster/executor/container/adapter.go | 4 ----
 3 files changed, 1 insertion(+), 7 deletions(-)

---

commit b40c26d2aec12f21f46aeb96ae8c170e96809f6a
Author:     Paulo Castro <pefcastro@gmail.com>
AuthorDate: Wed Oct 3 21:08:53 2018 +0100
Commit:     Paulo Castro <pefcastro@gmail.com>
CommitDate: Thu Oct 11 09:18:22 2018 +0100

    Rename balena to balena-engine (executables) or balenaEngine (project)
    
    Signed-off-by: Paulo Castro <paulo@resin.io>

 .github/ISSUE_TEMPLATE.md                          |  6 ++--
 .github/PULL_REQUEST_TEMPLATE.md                   |  4 +--
 CHANGELOG.md                                       | 26 ++++++++--------
 Dockerfile                                         |  6 ++--
 Dockerfile.build.aarch64                           |  4 +--
 Dockerfile.build.arm                               |  4 +--
 Dockerfile.build.i386                              |  4 +--
 Dockerfile.build.x86_64                            |  4 +--
 FAQ.md                                             |  4 +--
 README.md                                          | 18 +++++------
 build-allarch.sh                                   | 20 ++++++------
 build.sh                                           | 18 +++++------
 builder/dockerfile/parser/testfile-line/Dockerfile |  2 +-
 .../testfiles/brimstone-consuldock/Dockerfile      |  2 +-
 cli/config/configdir.go                            |  2 +-
 client/client_unix.go                              |  2 +-
 client/errors.go                                   |  4 +--
 client/request_test.go                             |  4 +--
 cmd/{balena => balena-engine}/main.go              | 14 ++++-----
 cmd/dockerd/config.go                              | 10 +++---
 cmd/dockerd/config_common_unix.go                  |  8 ++---
 cmd/dockerd/config_unix.go                         |  4 +--
 cmd/dockerd/daemon_unix.go                         |  4 +--
 cmd/dockerd/docker.go                              |  4 +--
 cmd/dockerd/hack/malformed_host_override_test.go   | 16 +++++-----
 cmd/mobynit/main.go                                |  2 +-
 container/container.go                             |  2 +-
 contrib/apparmor/template.go                       |  2 +-
 .../openrc/{balena.confd => balena-engine.confd}   | 12 ++++----
 .../openrc/{balena.initd => balena-engine.initd}   |  4 +--
 .../{balena.service => balena-engine.service}      |  4 +--
 ...alena.service.rpm => balena-engine.service.rpm} |  2 +-
 .../{balena.socket => balena-engine.socket}        |  6 ++--
 .../init/sysvinit-debian/{balena => balena-engine} | 12 ++++----
 .../{balena.default => balena-engine.default}      |  0
 .../init/sysvinit-redhat/{balena => balena-engine} | 10 +++---
 .../{balena.sysconfig => balena-engine.sysconfig}  |  2 +-
 .../upstart/{balena.conf => balena-engine.conf}    |  8 ++---
 contrib/install.sh                                 |  2 +-
 .../docker-engine-selinux/docker.if                |  4 +--
 .../docker-engine-selinux/docker.if                |  4 +--
 daemon/config/config.go                            |  2 +-
 daemon/config/config_common_unix_test.go           |  2 +-
 daemon/config/config_test.go                       |  6 ++--
 daemon/daemon_linux.go                             |  2 +-
 daemon/daemon_unix.go                              |  6 ++--
 daemon/graphdriver/devmapper/deviceset.go          |  4 +--
 daemon/info.go                                     |  2 +-
 daemon/listeners/group_unix.go                     |  2 +-
 daemon/metrics.go                                  |  2 +-
 daemon/volumes_unix.go                             |  2 +-
 docs/api/v1.18.md                                  |  2 +-
 docs/api/v1.19.md                                  |  2 +-
 docs/api/v1.20.md                                  |  2 +-
 docs/api/v1.21.md                                  |  2 +-
 docs/api/v1.22.md                                  |  2 +-
 docs/api/v1.23.md                                  |  2 +-
 docs/api/v1.24.md                                  |  2 +-
 docs/getting-started.md                            | 30 +++++++++---------
 hack/dockerfile/install-binaries.sh                |  2 +-
 hack/integration-cli-on-swarm/README.md            |  2 +-
 hack/integration-cli-on-swarm/host/compose.go      |  2 +-
 hack/make/.binary                                  |  2 +-
 hack/make/.binary-setup                            | 16 +++++-----
 hack/make/.detect-daemon-osarch                    |  6 ++--
 hack/make/.ensure-emptyfs                          |  4 +--
 hack/make/.integration-daemon-start                | 20 ++++++------
 hack/make/.integration-daemon-stop                 |  2 +-
 hack/make/.integration-test-helpers                |  2 +-
 hack/make/binary-balena                            |  4 +--
 hack/make/dynbinary-balena                         |  9 ++++--
 hack/make/run                                      |  8 ++---
 integration-cli/check_test.go                      |  4 +--
 integration-cli/daemon/daemon.go                   |  8 ++---
 integration-cli/docker_api_exec_test.go            |  2 +-
 integration-cli/docker_cli_daemon_test.go          | 36 +++++++++++-----------
 integration-cli/docker_cli_exec_test.go            |  4 +--
 integration-cli/docker_cli_search_test.go          |  2 +-
 integration-cli/docker_cli_volume_test.go          |  2 +-
 integration/container/restart_test.go              |  2 +-
 integration/plugin/authz/main_test.go              |  2 +-
 integration/service/main_test.go                   |  2 +-
 landr.conf.js                                      |  6 ++--
 libcontainerd/remote_daemon.go                     |  4 +--
 libcontainerd/remote_daemon_linux.go               |  4 +--
 libcontainerd/remote_daemon_windows.go             |  4 +--
 opts/hosts.go                                      |  2 +-
 opts/hosts_test.go                                 | 12 ++++----
 pkg/chrootarchive/cgroup_unix.go                   |  2 +-
 registry/config_unix.go                            |  2 +-
 vendor.conf                                        |  4 +--
 .../containerd/containerd/plugin/plugin_go18.go    |  2 +-
 vendor/github.com/docker/cli/cli/config/config.go  |  2 +-
 vendor/github.com/docker/cli/cmd/docker/docker.go  | 12 ++++----
 vendor/github.com/docker/cli/opts/hosts.go         |  2 +-
 .../docker/libnetwork/portmapper/proxy.go          |  2 +-
 96 files changed, 276 insertions(+), 271 deletions(-)

---

commit aa6df9901d7d439f2c980ef5b25e24b815d2ee44
Author:     Paulo Castro <pefcastro@gmail.com>
AuthorDate: Mon Oct 8 09:47:54 2018 +0100
Commit:     Paulo Castro <pefcastro@gmail.com>
CommitDate: Mon Oct 8 10:10:16 2018 +0100

    Disable incompatible integration tests
    
    Change-type: patch
    Signed-off-by: Paulo Castro <paulo@resin.io>

 Dockerfile                                         |   2 +-
 integration-cli/docker_cli_build_test.go           |  15 ++
 integration-cli/docker_cli_by_digest_test.go       |   6 +
 integration-cli/docker_cli_create_test.go          |   9 +
 integration-cli/docker_cli_daemon_test.go          | 291 +++++++++++++++++++++
 integration-cli/docker_cli_events_unix_test.go     |  12 +
 integration-cli/docker_cli_exec_test.go            |   3 +
 .../docker_cli_external_volume_driver_unix_test.go |  42 +++
 integration-cli/docker_cli_info_test.go            |   9 +
 integration-cli/docker_cli_logout_test.go          |   2 +
 integration-cli/docker_cli_network_unix_test.go    |  51 ++++
 integration-cli/docker_cli_oom_killed_test.go      |   6 +
 integration-cli/docker_cli_proxy_test.go           |   3 +
 integration-cli/docker_cli_prune_unix_test.go      |   6 +
 integration-cli/docker_cli_pull_test.go            |   3 +
 integration-cli/docker_cli_pull_trusted_test.go    |  15 ++
 integration-cli/docker_cli_push_test.go            |   6 +
 .../docker_cli_registry_user_agent_test.go         |   3 +
 integration-cli/docker_cli_run_test.go             |  15 ++
 integration-cli/docker_cli_run_unix_test.go        |  21 ++
 integration-cli/docker_cli_userns_test.go          |   3 +
 integration-cli/docker_cli_v2_only_test.go         |   6 +
 integration-cli/docker_hub_pull_suite_test.go      |   3 +
 integration/container/restart_test.go              |   3 +
 24 files changed, 534 insertions(+), 1 deletion(-)

---

commit 3e2973d26934bd22c46f07764afb1ed8b11bf6a1
Author:     Zubair Lutfullah Kakakhel <zubair@resin.io>
AuthorDate: Thu Aug 23 12:12:50 2018 +0000
Commit:     Zubair Lutfullah Kakakhel <zubair@resin.io>
CommitDate: Wed Sep 26 17:58:10 2018 +0000

    mobynit: Add support to mount rootfs from a custom location
    
    This patch adds support to pass an argument to mobynit which allows
    mobynit to mount a rootfs from a custom path.
    
    e.g. ./mobynit -sysroot /mnt/sysroot/inactive
    will mount the rootfs partition from /mnt/sysroot/inactive and return
    the destination path in stdout.
    
    Signed-off-by: Zubair Lutfullah Kakakhel <zubair@resin.io>

 cmd/mobynit/main.go | 100 ++++++++++++++++++++++++++++++++++------------------
 1 file changed, 65 insertions(+), 35 deletions(-)

---

commit cce1763d57b5c8fc446b0863517bb5313e7e53be
Author:     Tibor Vass <tibor@docker.com>
AuthorDate: Sat Sep 22 00:14:43 2018 +0000
Commit:     Tibor Vass <tibor@docker.com>
CommitDate: Sat Sep 22 01:24:11 2018 +0000

    vendor: remove boltdb dependency which is superseded by bbolt
    
    This also brings in these PRs from swarmkit:
    - https://github.com/docker/swarmkit/pull/2691
    - https://github.com/docker/swarmkit/pull/2744
    - https://github.com/docker/swarmkit/pull/2732
    - https://github.com/docker/swarmkit/pull/2729
    - https://github.com/docker/swarmkit/pull/2748
    
    Signed-off-by: Tibor Vass <tibor@docker.com>

 daemon/cluster/noderunner.go                       |   20 +-
 vendor.conf                                        |   12 +-
 vendor/github.com/boltdb/bolt/LICENSE              |   20 -
 vendor/github.com/boltdb/bolt/README.md            |  857 ----------------
 vendor/github.com/boltdb/bolt/bolt_386.go          |   10 -
 vendor/github.com/boltdb/bolt/bolt_amd64.go        |   10 -
 vendor/github.com/boltdb/bolt/bolt_arm.go          |   28 -
 vendor/github.com/boltdb/bolt/bolt_arm64.go        |   12 -
 vendor/github.com/boltdb/bolt/bolt_linux.go        |   10 -
 vendor/github.com/boltdb/bolt/bolt_openbsd.go      |   27 -
 vendor/github.com/boltdb/bolt/bolt_ppc.go          |    9 -
 vendor/github.com/boltdb/bolt/bolt_ppc64.go        |    9 -
 vendor/github.com/boltdb/bolt/bolt_ppc64le.go      |   12 -
 vendor/github.com/boltdb/bolt/bolt_s390x.go        |   12 -
 vendor/github.com/boltdb/bolt/bolt_unix.go         |   89 --
 vendor/github.com/boltdb/bolt/bolt_unix_solaris.go |   90 --
 vendor/github.com/boltdb/bolt/bolt_windows.go      |  144 ---
 vendor/github.com/boltdb/bolt/boltsync_unix.go     |    8 -
 vendor/github.com/boltdb/bolt/bucket.go            |  778 ---------------
 vendor/github.com/boltdb/bolt/cursor.go            |  400 --------
 vendor/github.com/boltdb/bolt/db.go                | 1036 --------------------
 vendor/github.com/boltdb/bolt/doc.go               |   44 -
 vendor/github.com/boltdb/bolt/errors.go            |   71 --
 vendor/github.com/boltdb/bolt/freelist.go          |  248 -----
 vendor/github.com/boltdb/bolt/node.go              |  604 ------------
 vendor/github.com/boltdb/bolt/page.go              |  178 ----
 vendor/github.com/boltdb/bolt/tx.go                |  682 -------------
 vendor/github.com/containerd/containerd/README.md  |    3 +-
 .../api/services/content/v1/content.pb.go          |    4 +-
 .../api/services/content/v1/content.proto          |    2 +-
 .../containerd/api/services/events/v1/events.pb.go |    4 +-
 .../containerd/api/services/events/v1/events.proto |    2 +-
 .../containerd/archive/compression/compression.go  |   88 +-
 vendor/github.com/containerd/containerd/cio/io.go  |    9 +
 .../containerd/containerd/container_opts_unix.go   |   58 --
 .../containerd/containerd/content/helpers.go       |    6 +-
 .../containerd/containerd/content/local/store.go   |    4 +-
 .../containerd/containerd/content/local/writer.go  |    8 +-
 .../containerd/content/proxy/content_writer.go     |    6 +-
 .../containerd/contrib/seccomp/seccomp.go          |    2 +-
 .../containerd/events/exchange/exchange.go         |    2 +-
 .../containerd/containerd/{import.go => export.go} |   64 +-
 .../containerd/images/archive/importer.go          |  254 +++++
 .../containerd/images/archive/reference.go         |   86 ++
 .../containerd/containerd/images/importexport.go   |    2 +-
 vendor/github.com/containerd/containerd/import.go  |  160 +--
 vendor/github.com/containerd/containerd/install.go |   41 +-
 .../containerd/containerd/install_opts.go          |    9 +
 .../containerd/containerd/metadata/bolt.go         |    2 +-
 .../containerd/metadata/boltutil/helpers.go        |    2 +-
 .../containerd/containerd/metadata/buckets.go      |    8 +-
 .../containerd/containerd/metadata/containers.go   |    2 +-
 .../containerd/containerd/metadata/content.go      |   17 +-
 .../containerd/containerd/metadata/db.go           |    4 +-
 .../containerd/containerd/metadata/gc.go           |    2 +-
 .../containerd/containerd/metadata/images.go       |    2 +-
 .../containerd/containerd/metadata/leases.go       |    2 +-
 .../containerd/containerd/metadata/migrations.go   |   14 +-
 .../containerd/containerd/metadata/namespaces.go   |    2 +-
 .../containerd/containerd/metadata/snapshot.go     |    2 +-
 .../containerd/containerd/mount/mount_windows.go   |    4 +
 .../github.com/containerd/containerd/oci/spec.go   |  219 ++++-
 .../containerd/containerd/oci/spec_opts.go         |  914 ++++++++++++++++-
 .../containerd/containerd/oci/spec_opts_unix.go    |  733 --------------
 .../containerd/containerd/oci/spec_opts_windows.go |   89 --
 .../containerd/containerd/oci/spec_unix.go         |  188 ----
 .../containerd/containerd/platforms/defaults.go    |    5 -
 .../spec_windows.go => platforms/defaults_unix.go} |   32 +-
 .../defaults_windows.go}                           |   20 +-
 .../containerd/remotes/docker/fetcher.go           |    2 +-
 .../containerd/remotes/docker/httpreadseeker.go    |    2 +-
 .../containerd/remotes/docker/schema1/converter.go |   55 +-
 .../containerd/runtime/v1/linux/proc/exec.go       |    2 +-
 .../containerd/runtime/v1/linux/proc/exec_state.go |    8 +-
 .../containerd/runtime/v1/linux/proc/init.go       |   13 +-
 .../containerd/runtime/v1/linux/proc/io.go         |    1 -
 .../containerd/runtime/v1/linux/runtime.go         |   16 +-
 .../containerd/containerd/runtime/v1/linux/task.go |    5 +-
 .../containerd/runtime/v1/shim/service.go          |   47 +-
 .../containerd/services/server/server.go           |    2 +-
 .../containerd/containerd/sys/socket_unix.go       |    2 +-
 vendor/github.com/containerd/containerd/task.go    |    5 +-
 .../github.com/containerd/containerd/task_opts.go  |   62 ++
 .../{task_opts_linux.go => task_opts_unix.go}      |   29 +-
 .../github.com/containerd/containerd/vendor.conf   |   41 +-
 vendor/github.com/containerd/continuity/context.go |  657 +++++++++++++
 vendor/github.com/containerd/continuity/digests.go |   88 ++
 .../containerd/continuity/driver/driver_unix.go    |   13 -
 .../containerd/continuity/driver/lchmod_linux.go   |   19 +
 .../containerd/continuity/driver/lchmod_unix.go    |   14 +
 vendor/github.com/containerd/continuity/fs/du.go   |    4 +-
 .../github.com/containerd/continuity/fs/du_unix.go |    8 +-
 .../containerd/continuity/fs/du_windows.go         |    8 +-
 .../containerd/continuity/groups_unix.go           |  113 +++
 .../github.com/containerd/continuity/hardlinks.go  |   57 ++
 .../containerd/continuity/hardlinks_unix.go        |   36 +
 .../containerd/continuity/hardlinks_windows.go     |   12 +
 vendor/github.com/containerd/continuity/ioutils.go |   47 +
 .../github.com/containerd/continuity/manifest.go   |  144 +++
 .../github.com/containerd/continuity/proto/gen.go  |    3 +
 .../containerd/continuity/proto/manifest.pb.go     |  181 ++++
 .../containerd/continuity/proto/manifest.proto     |   97 ++
 .../github.com/containerd/continuity/resource.go   |  574 +++++++++++
 .../containerd/continuity/resource_unix.go         |   37 +
 .../containerd/continuity/resource_windows.go      |   12 +
 .../containerd/continuity/sysx/README.md           |    3 +
 vendor/github.com/containerd/continuity/sysx/asm.s |   10 -
 .../containerd/continuity/sysx/chmod_darwin.go     |   18 -
 .../containerd/continuity/sysx/chmod_darwin_386.go |   25 -
 .../continuity/sysx/chmod_darwin_amd64.go          |   25 -
 .../containerd/continuity/sysx/chmod_freebsd.go    |   17 -
 .../continuity/sysx/chmod_freebsd_amd64.go         |   25 -
 .../containerd/continuity/sysx/chmod_linux.go      |   12 -
 .../containerd/continuity/sysx/chmod_solaris.go    |   11 -
 .../github.com/containerd/continuity/sysx/sys.go   |   37 -
 .../github.com/containerd/continuity/sysx/xattr.go |   48 +-
 .../containerd/continuity/sysx/xattr_darwin.go     |   71 --
 .../containerd/continuity/sysx/xattr_darwin_386.go |  111 ---
 .../continuity/sysx/xattr_darwin_amd64.go          |  111 ---
 .../containerd/continuity/sysx/xattr_freebsd.go    |   12 -
 .../containerd/continuity/sysx/xattr_linux.go      |   44 -
 .../containerd/continuity/sysx/xattr_openbsd.go    |    7 -
 .../containerd/continuity/sysx/xattr_solaris.go    |   12 -
 .../continuity/sysx/xattr_unsupported.go           |    9 +-
 .../github.com/containerd/continuity/vendor.conf   |    2 +-
 vendor/github.com/containerd/cri/LICENSE           |  201 ++++
 vendor/github.com/containerd/cri/README.md         |  176 ++++
 .../containerd/cri/pkg/util/deep_copy.go           |   42 +
 vendor/github.com/containerd/cri/pkg/util/id.go    |   29 +
 vendor/github.com/containerd/cri/pkg/util/image.go |   50 +
 .../github.com/containerd/cri/pkg/util/strings.go  |   59 ++
 .../containerd/{containerd => cri}/vendor.conf     |  120 ++-
 vendor/github.com/containerd/go-runc/console.go    |    2 +
 vendor/github.com/containerd/go-runc/io.go         |   99 +-
 vendor/github.com/containerd/go-runc/io_unix.go    |   76 ++
 vendor/github.com/containerd/go-runc/io_windows.go |   62 ++
 vendor/github.com/containerd/go-runc/runc.go       |    3 +-
 vendor/github.com/docker/libkv/README.md           |    6 +-
 .../github.com/docker/libkv/store/boltdb/boltdb.go |    2 +-
 vendor/github.com/docker/libkv/store/etcd/etcd.go  |  174 ++--
 vendor/github.com/docker/swarmkit/agent/config.go  |    2 +-
 vendor/github.com/docker/swarmkit/agent/session.go |    2 +-
 vendor/github.com/docker/swarmkit/agent/storage.go |    2 +-
 vendor/github.com/docker/swarmkit/agent/worker.go  |    2 +-
 vendor/github.com/docker/swarmkit/api/specs.pb.go  |  449 ++++++---
 vendor/github.com/docker/swarmkit/api/specs.proto  |   17 +-
 vendor/github.com/docker/swarmkit/api/types.pb.go  |  652 ++++++------
 vendor/github.com/docker/swarmkit/api/types.proto  |    1 +
 .../docker/swarmkit/manager/allocator/allocator.go |   21 +-
 .../manager/allocator/cnmallocator/drivers_ipam.go |   15 +-
 .../allocator/cnmallocator/networkallocator.go     |   14 +-
 .../docker/swarmkit/manager/allocator/network.go   |   10 +-
 .../docker/swarmkit/manager/controlapi/cluster.go  |   13 +
 .../swarmkit/manager/dispatcher/dispatcher.go      |   40 +-
 .../github.com/docker/swarmkit/manager/manager.go  |   35 +-
 vendor/github.com/docker/swarmkit/node/node.go     |   14 +-
 vendor/github.com/docker/swarmkit/vendor.conf      |    2 +-
 157 files changed, 5799 insertions(+), 8116 deletions(-)

---

commit b501aa82d5151b8af73d6670772cc4e8ba94765f
Author:     Tonis Tiigi <tonistiigi@gmail.com>
AuthorDate: Fri Sep 21 17:02:32 2018 -0700
Commit:     Tonis Tiigi <tonistiigi@gmail.com>
CommitDate: Fri Sep 21 17:06:25 2018 -0700

    vendor: update bolt to bbolt
    
    Signed-off-by: Tonis Tiigi <tonistiigi@gmail.com>

 builder/builder-next/adapters/snapshot/layer.go    | 2 +-
 builder/builder-next/adapters/snapshot/snapshot.go | 2 +-
 builder/fscache/fscache.go                         | 2 +-
 volume/service/db.go                               | 2 +-
 volume/service/db_test.go                          | 2 +-
 volume/service/restore.go                          | 2 +-
 volume/service/store.go                            | 2 +-
 7 files changed, 7 insertions(+), 7 deletions(-)

---

commit 46a703bb3bfe75e99de2cc457dc0d568a1976a6b
Author:     Tonis Tiigi <tonistiigi@gmail.com>
AuthorDate: Fri Sep 21 17:06:01 2018 -0700
Commit:     Tonis Tiigi <tonistiigi@gmail.com>
CommitDate: Fri Sep 21 17:06:25 2018 -0700

    vendor: add bbolt v1.3.1-etcd.8
    
    Signed-off-by: Tonis Tiigi <tonistiigi@gmail.com>

 vendor.conf                                  |    1 +
 vendor/go.etcd.io/bbolt/LICENSE              |   20 +
 vendor/go.etcd.io/bbolt/README.md            |  953 +++++++++++++++++++++
 vendor/go.etcd.io/bbolt/bolt_386.go          |   10 +
 vendor/go.etcd.io/bbolt/bolt_amd64.go        |   10 +
 vendor/go.etcd.io/bbolt/bolt_arm.go          |   28 +
 vendor/go.etcd.io/bbolt/bolt_arm64.go        |   12 +
 vendor/go.etcd.io/bbolt/bolt_linux.go        |   10 +
 vendor/go.etcd.io/bbolt/bolt_mips64x.go      |   12 +
 vendor/go.etcd.io/bbolt/bolt_mipsx.go        |   12 +
 vendor/go.etcd.io/bbolt/bolt_openbsd.go      |   27 +
 vendor/go.etcd.io/bbolt/bolt_ppc.go          |   12 +
 vendor/go.etcd.io/bbolt/bolt_ppc64.go        |   12 +
 vendor/go.etcd.io/bbolt/bolt_ppc64le.go      |   12 +
 vendor/go.etcd.io/bbolt/bolt_s390x.go        |   12 +
 vendor/go.etcd.io/bbolt/bolt_unix.go         |   93 +++
 vendor/go.etcd.io/bbolt/bolt_unix_solaris.go |   88 ++
 vendor/go.etcd.io/bbolt/bolt_windows.go      |  141 ++++
 vendor/go.etcd.io/bbolt/boltsync_unix.go     |    8 +
 vendor/go.etcd.io/bbolt/bucket.go            |  775 ++++++++++++++++++
 vendor/go.etcd.io/bbolt/cursor.go            |  396 +++++++++
 vendor/go.etcd.io/bbolt/db.go                | 1138 ++++++++++++++++++++++++++
 vendor/go.etcd.io/bbolt/doc.go               |   44 +
 vendor/go.etcd.io/bbolt/errors.go            |   71 ++
 vendor/go.etcd.io/bbolt/freelist.go          |  333 ++++++++
 vendor/go.etcd.io/bbolt/node.go              |  604 ++++++++++++++
 vendor/go.etcd.io/bbolt/page.go              |  197 +++++
 vendor/go.etcd.io/bbolt/tx.go                |  707 ++++++++++++++++
 28 files changed, 5738 insertions(+)

---

commit 9f4cd6a7ea39ec0c1ad62a44b98f3f02b70efa78
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Thu Sep 20 17:28:35 2018 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Thu Sep 20 17:28:35 2018 -0700

    update containerd/console to fix race: lock Cond before Signal
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>
    Connects-To: https://github.com/moby/moby/issues/35865
    Connects-To: https://github.com/containerd/console/pull/27

 vendor.conf                                        |   2 +-
 vendor/github.com/containerd/console/console.go    |  16 +++
 .../github.com/containerd/console/console_linux.go |  36 ++++--
 .../github.com/containerd/console/console_unix.go  |  18 ++-
 .../containerd/console/console_windows.go          | 124 +++++++++------------
 vendor/github.com/containerd/console/tc_darwin.go  |  16 +++
 vendor/github.com/containerd/console/tc_freebsd.go |  16 +++
 vendor/github.com/containerd/console/tc_linux.go   |  34 ++++--
 .../containerd/console/tc_openbsd_cgo.go           |  51 +++++++++
 .../containerd/console/tc_openbsd_nocgo.go         |  47 ++++++++
 .../containerd/console/tc_solaris_cgo.go           |  16 +++
 .../containerd/console/tc_solaris_nocgo.go         |  16 +++
 vendor/github.com/containerd/console/tc_unix.go    |  18 ++-
 13 files changed, 316 insertions(+), 94 deletions(-)

---

commit deba4bbc54f980b92f5aeab674688265486aa3b1
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Wed Sep 5 16:00:14 2018 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Wed Sep 5 16:00:14 2018 -0700

    delta: use chain ids to decide whether to skip a layer
    
    With the previous code if two layers had identical contents (and diff
    ids) but different parents (and different chain ids) they would be
    considered common and skipped.
    
    However the pulling code uses the chain id to decide whether it should
    pull a layer so a delta created under these conditions was leading to
    the pulling code trying to apply a non-existent delta.
    
    This is fixed by using the same logic with chain ids in the generation
    phase.
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 daemon/create.go | 17 +++++++++++++----
 1 file changed, 13 insertions(+), 4 deletions(-)

---

commit 4032b6778df39f53fda0e6e54f0256c9a3b1d618
Author:     Kir Kolyshkin <kolyshkin@gmail.com>
AuthorDate: Thu Aug 30 15:32:14 2018 -0700
Commit:     Kir Kolyshkin <kolyshkin@gmail.com>
CommitDate: Thu Aug 30 17:34:59 2018 -0700

    Fix relabeling local volume source dir
    
    In case a volume is specified via Mounts API, and SELinux is enabled,
    the following error happens on container start:
    
    > $ docker volume create testvol
    > $ docker run --rm --mount source=testvol,target=/tmp busybox true
    > docker: Error response from daemon: error setting label on mount
    > source '': no such file or directory.
    
    The functionality to relabel the source of a local mount specified via
    Mounts API was introduced in commit 5bbf5cc and later broken by commit
    e4b6adc, which removed setting mp.Source field.
    
    With the current data structures, the host dir is already available in
    v.Mountpoint, so let's just use it.
    
    Fixes: e4b6adc
    Signed-off-by: Kir Kolyshkin <kolyshkin@gmail.com>

 daemon/volumes.go | 2 ++
 1 file changed, 2 insertions(+)

---

commit c87589c33b9974a1eeceede2b9606fbbddf3a8f5
Author:     Gergely Imreh <gergely@resin.io>
AuthorDate: Tue Aug 28 11:38:40 2018 +0100
Commit:     Gergely Imreh <gergely@resin.io>
CommitDate: Tue Aug 28 11:38:40 2018 +0100

    version: Fix balena CLI version string
    
    The balena CLI version, git commit, and build time variables are
    set at build time using the appropriate flags.
    
    Signed-off-by: Gergely Imreh <gergely@resin.io>

 hack/make.sh | 7 ++++++-
 1 file changed, 6 insertions(+), 1 deletion(-)

---

commit 9d1d910e5d293dddd6dfb96b9f50316d03697850
Author:     Gergely Imreh <gergely@resin.io>
AuthorDate: Tue Aug 28 11:35:16 2018 +0100
Commit:     Gergely Imreh <gergely@resin.io>
CommitDate: Tue Aug 28 11:35:16 2018 +0100

    version: Fix balena server version string
    
    The balena server version info is loaded from the `VERSION`
    environment variable in the build script, but that variable
    wasn't passed.
    
    Signed-off-by: Gergely Imreh <gergely@resin.io>

 build.sh | 2 +-
 1 file changed, 1 insertion(+), 1 deletion(-)

---

commit 3685c83503cb7ebdfc0fff3f900123cec7913c73
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Mon Aug 27 09:49:14 2018 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Aug 27 09:49:14 2018 -0700

    pkg/chrootarchive: disable memory cgroups until pending issues are fixed
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 pkg/chrootarchive/archive_unix.go | 1 -
 pkg/chrootarchive/diff_unix.go    | 1 -
 2 files changed, 2 deletions(-)

---

commit 85b036bd3a57a0cece7e09bc947cbafd7ba4fa4e
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Fri Aug 24 17:50:25 2018 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Aug 27 09:40:30 2018 -0700

    vendor: update libnetwork to include stale default bridge fix
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 vendor.conf                                                  | 2 +-
 vendor/github.com/docker/libnetwork/drivers/bridge/bridge.go | 6 ++----
 2 files changed, 3 insertions(+), 5 deletions(-)

---

commit b706f5daf673cde20571a611a33ae62d9fba26cb
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Tue Jul 17 10:46:18 2018 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Tue Jul 17 10:46:18 2018 -0700

    pkg/ioutils: implement eager writer
    
    EagerFileWriter will schedule immediate writeback of dirty pages. The
    writes will only block if the device's write queue is full, which
    provides throttling without incuring the latency cost of fsync.
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 pkg/ioutils/eagerwriter.go | 54 ++++++++++++++++++++++++++++++++++++++++++++++
 1 file changed, 54 insertions(+)

---

commit 08b01efe225263dd8e7f7f82fd0fddc403f267b9
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Wed Jul 4 15:24:38 2018 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Thu Jul 5 20:33:22 2018 -0700

    Revert "vendor: update golang/x/sys to support fadvise for arm64"
    
    This reverts commit 5ead292ff82763acdc68efd8175dc7442682a9f0.

 vendor.conf                                          |  2 +-
 vendor/golang.org/x/sys/unix/syscall_linux_arm64.go  |  1 -
 vendor/golang.org/x/sys/unix/zsyscall_linux_arm64.go | 10 ----------
 3 files changed, 1 insertion(+), 12 deletions(-)

---

commit 60f2a21c95a6ee96b5a03367cc3c2e463be3787c
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Wed Jul 4 15:14:48 2018 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Thu Jul 5 20:33:21 2018 -0700

    pull: rely on memory cgroups to avoid page cache thrashing
    
    This improves the previous attempt at reducing page cache thrashing
    introduced in 8c0ceab ("pkg/archive: use fadvise to prevent pagecache
    thrashing"). While the approached worked, it introduced severe delays
    during pull especially on high latency hard drives since it had to wait
    for every file.
    
    A much better approach implemented in this commit is to run the
    unpacking process in a memory constrained cgroup to limit the maximum
    page cache usage while at the same time hinting the kernel to start the
    writeback as soon as possible to reduce the amount of dirty pages as
    much as possible.
    
    This makes `balena pull` as fast as `docker pull && sync` while having
    minimal impact on page cache usage.
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 pkg/archive/archive.go            | 14 +++-----------
 pkg/chrootarchive/archive_unix.go |  1 +
 pkg/chrootarchive/cgroup_unix.go  | 29 +++++++++++++++++++++++++++++
 pkg/chrootarchive/diff_unix.go    |  1 +
 4 files changed, 34 insertions(+), 11 deletions(-)

---

commit 38b223b0013390cb47276e1f4f5ddf6f44cb5db3
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Wed Jun 27 14:29:45 2018 +0300
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Thu Jun 28 12:18:45 2018 +0300

    pkg/stringid: don't bother seeding math/random with crypto grade seed
    
    Seeding the math/random RNG with a crypto grade seed slows down startup
    on embedded devices that don't have enough entropy during early boot.
    
    Furthermore math/rand is seeded by a lot of other packages with
    non-crypto grade seeds which makes it a last write wins situation.
    
    Since math/rand is only used in the GenerateNonCryptoID function, we
    make it always use the current time as a seed to avoid startup delays.
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 pkg/stringid/stringid.go | 14 +-------------
 1 file changed, 1 insertion(+), 13 deletions(-)

---

commit f08057baa72ce97af7de809d9f944bbb7f5275fe
Author:     Gergely Imreh <imrehg@gmail.com>
AuthorDate: Thu Jun 7 22:17:02 2018 +0100
Commit:     Gergely Imreh <imrehg@gmail.com>
CommitDate: Thu Jun 7 22:17:02 2018 +0100

    vendor: update btrfs dependency
    
    Provides updates for licensing and a build problem.
    
    Signed-off-by: Gergely Imreh <imrehg@gmail.com>

 vendor.conf                                   |  2 +-
 vendor/github.com/containerd/btrfs/README.md  | 20 +++++++++++++++
 vendor/github.com/containerd/btrfs/btrfs.go   | 36 +++++++++++++++++++--------
 vendor/github.com/containerd/btrfs/doc.go     | 17 +++++++++++++
 vendor/github.com/containerd/btrfs/helpers.go | 15 +++++++++++
 vendor/github.com/containerd/btrfs/info.go    | 15 +++++++++++
 vendor/github.com/containerd/btrfs/ioctl.go   | 15 +++++++++++
 7 files changed, 108 insertions(+), 12 deletions(-)

---

commit 519ed006c67b4637f5782dc737f78b33d53248e0
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Tue Mar 13 08:54:58 2018 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon May 7 20:18:00 2018 -0700

    container: remove extraneous lock leading to deadlocks
    
    defers are called in the opposite way that they were declared. The
    container is already being locked in the function
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 daemon/rename.go | 2 --
 1 file changed, 2 deletions(-)

---

commit 2e2f9df86df5cf35b47f2fbaa186f09a942f710b
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Mon May 7 20:14:50 2018 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon May 7 20:18:00 2018 -0700

    tests: more integration test fixes
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 hack/dockerfile/binaries-commits        | 2 +-
 hack/make/.binary                       | 2 +-
 integration-cli/docker_api_info_test.go | 1 -
 3 files changed, 2 insertions(+), 3 deletions(-)

---

commit 276ee9d99df2a6ff8b1f1f3c4c8d1b46dad71bd9
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Mon May 7 20:14:16 2018 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon May 7 20:18:00 2018 -0700

    cmd/mobynit: adapt to new internal API
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 cmd/mobynit/main.go | 11 ++++++-----
 1 file changed, 6 insertions(+), 5 deletions(-)

---

commit 8e47b094929f8a7e47a96e51c883e9ff072ca01e
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Mon May 7 20:13:51 2018 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon May 7 20:18:00 2018 -0700

    build: switch the default build to be the dynamically linked binary
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 Makefile | 2 +-
 1 file changed, 1 insertion(+), 1 deletion(-)

---

commit 137b066c421a4a8390b0232648b992801356f7fc
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Mon May 7 20:12:52 2018 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon May 7 20:18:00 2018 -0700

    tests: remove plugin support in tests
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 integration/plugin/authz/main_test.go | 2 ++
 internal/test/environment/clean.go    | 3 ---
 internal/test/environment/protect.go  | 3 ---
 3 files changed, 2 insertions(+), 6 deletions(-)

---

commit 64f52ee1e3ed120fc617aab21d3047983d252b24
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Mon May 7 20:12:18 2018 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon May 7 20:18:00 2018 -0700

    tests: skip swarm tests
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 daemon/config/config_test.go        | 14 --------------
 daemon/config/opts.go               |  5 ++++-
 integration/service/create_test.go  |  2 ++
 integration/service/inspect_test.go |  2 ++
 4 files changed, 8 insertions(+), 15 deletions(-)

---

commit e0e5db31fc8c0388f22dd1f664150b0b0ba738cf
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Mon May 7 20:11:23 2018 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon May 7 20:18:00 2018 -0700

    fix regression of DockerSuite.TestAPINetworkCreateCheckDuplicate
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 api/server/router/network/network_routes.go | 8 ++++++++
 1 file changed, 8 insertions(+)

---

commit 5955d382f3c9829a960dbfb9136bfb137477e5d6
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Mon May 7 20:09:48 2018 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon May 7 20:18:00 2018 -0700

    build: do not install embedded binaries separately
    
    skip runc, containerd, proxy, dockercli
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 Dockerfile | 2 +-
 1 file changed, 1 insertion(+), 1 deletion(-)

---

commit a466c05736c506467667623f51134a7839c20abd
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Mon May 7 20:07:17 2018 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon May 7 20:18:00 2018 -0700

    cmd/balena: exit with non-zero code if called with unknown command
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 cmd/balena/main.go | 9 ++++++---
 1 file changed, 6 insertions(+), 3 deletions(-)

---

commit 3a1be7aa960835cf3f2e34f92532c4549c19cc66
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Mon May 7 20:05:41 2018 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon May 7 20:18:00 2018 -0700

    a lot of balena rename fixes for integration tests
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 daemon/daemon_unix.go                     |  4 +-
 hack/dockerfile/install-binaries.sh       |  2 +-
 hack/make/.integration-test-helpers       |  2 +-
 hack/make/dynbinary                       |  6 +-
 integration-cli/check_test.go             |  4 +-
 integration-cli/daemon/daemon.go          |  2 +-
 integration-cli/docker_api_exec_test.go   |  2 +-
 integration-cli/docker_cli_daemon_test.go | 92 +++++++++++++++----------------
 integration-cli/docker_cli_exec_test.go   |  4 +-
 integration-cli/docker_cli_run_test.go    |  2 +-
 integration-cli/docker_cli_search_test.go |  2 +-
 integration-cli/docker_cli_volume_test.go |  2 +-
 integration/container/restart_test.go     |  2 +-
 integration/plugin/authz/main_test.go     |  2 +-
 integration/service/main_test.go          |  2 +-
 libcontainerd/remote_daemon.go            |  4 +-
 libcontainerd/remote_daemon_linux.go      |  4 +-
 libcontainerd/remote_daemon_windows.go    |  4 +-
 18 files changed, 71 insertions(+), 71 deletions(-)

---

commit f3b6b8a1aea4186d0d3a0e49241dc3750f97481a
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Tue Feb 13 13:48:29 2018 -0800
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon May 7 20:18:00 2018 -0700

    vendor: update containerd
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 vendor.conf                                        |  2 +-
 .../containerd/cmd/containerd-shim/main_unix.go    | 24 +++++++++++++---------
 .../containerd/linux/shim/client/client.go         |  1 +
 3 files changed, 16 insertions(+), 11 deletions(-)

---

commit b64eefe3a2fe7311d27f8a1a6a24565ba005e6f2
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Mon Feb 12 19:59:46 2018 -0800
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon May 7 20:18:00 2018 -0700

    build: switch to statically linked builds
    
    It should fix issue #52
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 build.sh | 4 ++--
 1 file changed, 2 insertions(+), 2 deletions(-)

---

commit 9ed4298d6637ee54f92a7b45ce4edd23bd3d171c
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Mon Feb 12 19:59:27 2018 -0800
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 20:00:44 2018 -0800

    build: let the go compiler do the stripping
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 build.sh | 3 +--
 1 file changed, 1 insertion(+), 2 deletions(-)

---

commit bd23724524ac213d3cb807048f3cdeb8f1bed203
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Mon Feb 12 19:58:49 2018 -0800
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 20:00:44 2018 -0800

    build: limit max go procs to avoid qemu hangs
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 build.sh | 2 +-
 1 file changed, 1 insertion(+), 1 deletion(-)

---

commit 5ead292ff82763acdc68efd8175dc7442682a9f0
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Tue Aug 29 16:17:17 2017 +0300
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 20:00:44 2018 -0800

    vendor: update golang/x/sys to support fadvise for arm64
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 vendor.conf                                          |  2 +-
 vendor/golang.org/x/sys/unix/syscall_linux_arm64.go  |  1 +
 vendor/golang.org/x/sys/unix/zsyscall_linux_arm64.go | 10 ++++++++++
 3 files changed, 12 insertions(+), 1 deletion(-)

---

commit 0386158271fa492c85e6aacd38c683519c7d64a1
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Mon Feb 12 18:47:16 2018 -0800
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 20:00:44 2018 -0800

    build: add libudev dependency
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 Dockerfile.build.aarch64 | 4 +++-
 Dockerfile.build.arm     | 3 ++-
 Dockerfile.build.i386    | 4 +++-
 Dockerfile.build.x86_64  | 4 +++-
 4 files changed, 11 insertions(+), 4 deletions(-)

---

commit fd78fe46cc62efec5f8fa45d9e217f69ae7072fb
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Mon Feb 12 18:42:48 2018 -0800
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 20:00:43 2018 -0800

    vendor: update containerd to non-plugin version
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 vendor.conf                                        |  2 +-
 .../containerd/containerd/plugin/plugin_go18.go    | 40 +---------------------
 2 files changed, 2 insertions(+), 40 deletions(-)

---

commit a1191cbdff1bd13a0cc3575ef7f63c7efc75dcef
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Mon Feb 12 15:29:59 2018 -0800
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 20:00:43 2018 -0800

    daemon/config: remove swarm support
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 daemon/config/opts.go | 16 +++-------------
 1 file changed, 3 insertions(+), 13 deletions(-)

---

commit ddaa8c17ff345a4ea0aa52fb59b4167e1f9a240b
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Wed Jan 17 17:26:52 2018 -0800
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 20:00:43 2018 -0800

    daemon: add appropriate container locks to avoid races
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 daemon/daemon.go  | 2 ++
 daemon/monitor.go | 2 ++
 daemon/rename.go  | 2 ++
 daemon/start.go   | 3 +++
 4 files changed, 9 insertions(+)

---

commit c24bda923c9347977b0e0a65ff224c7c026d2baa
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Mon Dec 11 09:02:24 2017 -0800
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 20:00:43 2018 -0800

    healthcheck: fix docker segfaulting
    
    We can't just convert struct pointer that is `nil` to an anonymous field
    type. Anonymous fields behave in the same way like fields and will
    segfault if accessed without checking.
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 container/container.go | 8 +++++++-
 daemon/monitor.go      | 7 ++++++-
 2 files changed, 13 insertions(+), 2 deletions(-)

---

commit 97505a483051f77a65f62f1800cf8dd2508365e8
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Mon Feb 12 18:55:00 2018 -0800
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 20:00:42 2018 -0800

    vendor: update vendor.conf with all required dependencies
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 vendor.conf | 36 +++++++++++++++++++++++++-----------
 1 file changed, 25 insertions(+), 11 deletions(-)

---

commit 8c124159175f3ad53482eccbafc3dfdbb931832c
Author:     Yossi Eliaz <yossi@resin.io>
AuthorDate: Sun Dec 3 08:57:04 2017 -0600
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 20:00:42 2018 -0800

    restartmanager: fixed the unit test
    
    Signed-off-by: Yossi Eliaz <yossi@resin.io>

 restartmanager/restartmanager_test.go | 7 +++++--
 1 file changed, 5 insertions(+), 2 deletions(-)

---

commit 8af842e3ac86b58574c702676554ccb987afc210
Author:     Yossi Eliaz <yossi@resin.io>
AuthorDate: Sun Dec 3 07:50:03 2017 -0600
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 20:00:42 2018 -0800

    tests: renamed runc to balena-runc
    
    Signed-off-by: Yossi Eliaz <yossi@resin.io>

 integration-cli/docker_cli_daemon_test.go | 8 ++++----
 1 file changed, 4 insertions(+), 4 deletions(-)

---

commit 55f437987735d6e4ccf356c9e5d4cfe13ada1211
Author:     Yossi Eliaz <yossi@resin.io>
AuthorDate: Tue Nov 14 21:06:11 2017 -0600
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 20:00:42 2018 -0800

    fixed balena version error
    
    Signed-off-by: Yossi Eliaz <yossi@resin.io>

 hack/make/.detect-daemon-osarch | 4 ++--
 1 file changed, 2 insertions(+), 2 deletions(-)

---

commit 24b643b1b57d606a603b3bfd1f7f8831dff3e2d7
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Sun Nov 26 16:12:25 2017 -0800
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 20:00:42 2018 -0800

    daemon: experimental: restart container when they become unhealthy
    
    Start treating an unhealthy container as a failure condition. Restarts
    are only initiated if a restart policy is set.
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 container/container.go           |  2 +-
 daemon/health.go                 | 19 +++++++++++++++++++
 daemon/monitor.go                |  2 +-
 restartmanager/restartmanager.go |  7 ++++---
 4 files changed, 25 insertions(+), 5 deletions(-)

---

commit b430038f36453d8f9006d380cc9f3030617076f4
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Wed Nov 22 02:29:04 2017 -0800
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 20:00:42 2018 -0800

    daemon: only attempt to prune local networks since swarm is disabled
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 daemon/prune.go | 3 ---
 1 file changed, 3 deletions(-)

---

commit eac6aa078447edb7fcda36975a5aa594ff8c5408
Author:     Akis Kesoglou <akiskesoglou@gmail.com>
AuthorDate: Mon Nov 6 15:34:23 2017 +0200
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 20:00:41 2018 -0800

    Updated init scripts for Balena
    
    This changes all mentions of "docker" to "balena" in init scripts, including file names. Does not touch "Docker" when it's inferred (from context) that it's used to refer to the Moby project.

 contrib/init/openrc/{docker.confd => balena.confd}           | 10 +++++-----
 contrib/init/openrc/{docker.initd => balena.initd}           |  4 ++--
 contrib/init/systemd/{docker.service => balena.service}      |  4 ++--
 .../init/systemd/{docker.service.rpm => balena.service.rpm}  |  2 +-
 contrib/init/systemd/{docker.socket => balena.socket}        |  4 ++--
 contrib/init/sysvinit-debian/{docker => balena}              | 12 ++++++------
 .../init/sysvinit-debian/{docker.default => balena.default}  |  0
 contrib/init/sysvinit-redhat/{docker => balena}              | 10 +++++-----
 .../sysvinit-redhat/{docker.sysconfig => balena.sysconfig}   |  2 +-
 contrib/init/upstart/{docker.conf => balena.conf}            |  4 ++--
 10 files changed, 26 insertions(+), 26 deletions(-)

---

commit 062cf0e404d063811bb1d35313bf085b71f94468
Author:     Yossi Eliaz <yossi@resin.io>
AuthorDate: Mon Oct 30 10:17:12 2017 -0500
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 20:00:41 2018 -0800

    Updated github hooks for balena
    
    Signed-off-by: Yossi Eliaz <yossi@resin.io>

 .github/ISSUE_TEMPLATE.md        | 7 ++-----
 .github/PULL_REQUEST_TEMPLATE.md | 9 ++-------
 2 files changed, 4 insertions(+), 12 deletions(-)

---

commit 07e8c0a0f9e261343ea4680c6c047d1ae16783cf
Author:     craig-mulligan <craig@resin.io>
AuthorDate: Fri Oct 13 11:29:41 2017 +0100
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 20:00:41 2018 -0800

    Update website copy

 landr.conf.js | 4 ++--
 1 file changed, 2 insertions(+), 2 deletions(-)

---

commit 5d81d5a215f5e8842ef8e999a30bc10c70988f8e
Author:     Edward Vielmetti <edward.vielmetti@gmail.com>
AuthorDate: Fri Oct 13 08:44:42 2017 -0400
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 20:00:41 2018 -0800

    Issue template should refer to balena throughout

 .github/ISSUE_TEMPLATE.md | 4 ++--
 1 file changed, 2 insertions(+), 2 deletions(-)

---

commit a8846e22160da0ce2f9367769d92d04a0acfe1c2
Author:     Yossi Eliaz <yossi@resin.io>
AuthorDate: Wed Oct 18 01:11:46 2017 -0500
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 20:00:41 2018 -0800

    updated the mock of xfer to pass unit test
    
    Signed-off-by: Yossi Eliaz <yossi@resin.io>

 distribution/xfer/download_test.go | 5 +++++
 1 file changed, 5 insertions(+)

---

commit 8f898bb3add828e82ec16f8514bd6f7f020326d7
Author:     Yossi Eliaz <yossi@resin.io>
AuthorDate: Fri Oct 13 01:26:00 2017 -0500
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 20:00:41 2018 -0800

    fixed integration with balena
    
    Signed-off-by: Yossi Eliaz <yossi@resin.io>

 Dockerfile                                           |  6 +++---
 builder/dockerfile/parser/testfile-line/Dockerfile   |  2 +-
 .../parser/testfiles/brimstone-consuldock/Dockerfile |  2 +-
 client/request_test.go                               |  4 ++--
 cmd/dockerd/hack/malformed_host_override_test.go     | 16 ++++++++--------
 contrib/apparmor/template.go                         |  2 +-
 contrib/init/openrc/docker.confd                     |  2 +-
 contrib/init/systemd/docker.socket                   |  2 +-
 contrib/init/upstart/docker.conf                     |  4 ++--
 .../docker-engine-selinux/docker.if                  |  4 ++--
 .../docker-engine-selinux/docker.if                  |  4 ++--
 daemon/config/config_test.go                         |  6 +++---
 daemon/volumes_unix.go                               |  2 +-
 docs/api/v1.18.md                                    |  2 +-
 docs/api/v1.19.md                                    |  2 +-
 docs/api/v1.20.md                                    |  2 +-
 docs/api/v1.21.md                                    |  2 +-
 docs/api/v1.22.md                                    |  2 +-
 docs/api/v1.23.md                                    |  2 +-
 docs/api/v1.24.md                                    |  2 +-
 hack/integration-cli-on-swarm/README.md              |  2 +-
 hack/integration-cli-on-swarm/host/compose.go        |  2 +-
 hack/make/.binary                                    |  2 +-
 hack/make/.binary-setup                              | 14 +++++++-------
 hack/make/.detect-daemon-osarch                      |  2 +-
 hack/make/.ensure-emptyfs                            |  4 ++--
 hack/make/.integration-daemon-start                  | 20 ++++++++++----------
 hack/make/.integration-daemon-stop                   |  2 +-
 hack/make/run                                        |  8 ++++----
 integration-cli/daemon/daemon.go                     |  6 +++---
 integration-cli/docker_cli_daemon_test.go            |  4 ++--
 opts/hosts_test.go                                   | 12 ++++++------
 32 files changed, 74 insertions(+), 74 deletions(-)

---

commit 60cb5cb2427f14ea100b77e450902a0c6115d321
Author:     Yossi Eliaz <yossi@resin.io>
AuthorDate: Fri Oct 13 00:07:12 2017 -0500
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 20:00:40 2018 -0800

    Renaming target to support balena
    
    Signed-off-by: Yossi Eliaz <yossi@resin.io>

 Makefile | 16 ++++++++--------
 1 file changed, 8 insertions(+), 8 deletions(-)

---

commit 5d3045460459918738a55c181a69a1c9019fbf53
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Sun Oct 15 14:06:24 2017 +0200
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 20:00:40 2018 -0800

    fix addidental mention of balaena name instead of balena
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 docs/getting-started.md | 2 +-
 1 file changed, 1 insertion(+), 1 deletion(-)

---

commit 189482e3725ced0d066afc44cbceb6b49ccbbaf8
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Fri Oct 13 02:48:52 2017 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 20:00:40 2018 -0800

    build: temporary switch to other base images
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 Dockerfile.build.arm  | 2 +-
 Dockerfile.build.i386 | 2 +-
 contrib/install.sh    | 4 ++--
 3 files changed, 4 insertions(+), 4 deletions(-)

---

commit b196586536c568f0f8e977a076c90c2da10a9328
Author:     craig-mulligan <craig@resin.io>
AuthorDate: Fri Oct 13 08:46:05 2017 +0100
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:13:26 2018 -0800

    add analytics

 landr.conf.js | 10 +++++-----
 1 file changed, 5 insertions(+), 5 deletions(-)

---

commit 2e78618717feb287c182240541ead3faf9b9bd9f
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Thu Oct 12 22:41:34 2017 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:13:26 2018 -0800

    docs: write a getting started guide
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 docs/getting-started.md | 67 +++++++++++++++++++++++++++++++++++++++++++------
 1 file changed, 59 insertions(+), 8 deletions(-)

---

commit 1003600d7200c13f1b48624ecb4d57b13dfae0ff
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Thu Oct 12 21:22:30 2017 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:13:26 2018 -0800

    build: use resin base images

 Dockerfile.build.arm    | 2 +-
 Dockerfile.build.i386   | 2 +-
 Dockerfile.build.x86_64 | 2 +-
 3 files changed, 3 insertions(+), 3 deletions(-)

---

commit b4a4651f0069902d77f05d92d82a9a4081c1a06e
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Thu Oct 12 21:04:24 2017 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:13:26 2018 -0800

    readme: Improve docker-ce section

    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 README.md | 13 +++++++++----
 1 file changed, 9 insertions(+), 4 deletions(-)

---

commit 99f6d0c5405d1d8a15f1edb02f347cb804d34e0e
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Thu Oct 12 16:44:09 2017 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:13:25 2018 -0800

    balena build scripts

    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 Dockerfile.build.aarch64 | 17 ++++++++++++++++
 Dockerfile.build.arm     | 18 ++++++++++++++++
 Dockerfile.build.i386    | 13 ++++++++++++
 Dockerfile.build.x86_64  | 13 ++++++++++++
 build-allarch.sh         | 25 +++++++++++++++++++++++
 build.sh                 | 40 ++++++++++++++++++++++++++++++++++++
 contrib/install.sh       | 53 ++++++++++++++++++++++++++++++++++++++++++++++++
 7 files changed, 179 insertions(+)

---

commit ec7e8dcafe3962da51692a2c98e02d3e1c963867
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Wed Oct 4 07:58:47 2017 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:13:23 2018 -0800

    daemon: return engine name as part of the version information

    This can be used by clients that know how to talk to the extended API
    balaena provides, e.g creating deltas

    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 api/types/types.go | 1 +
 daemon/info.go     | 1 +
 2 files changed, 2 insertions(+)

---

commit 2585aa19338217724fe2dac6d34741c8bb672ee1
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Sun Oct 1 10:02:11 2017 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:13:23 2018 -0800

    vendor: whitelist VERSION files

    Since we're consolidating binaries we need the VERSION of individual
    components to properly implement integration tests

    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 hack/vendor.sh | 2 +-
 1 file changed, 1 insertion(+), 1 deletion(-)

---

commit dd56c7ab7921f14a66f35f7d39d00c3f172c202d
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Tue Mar 28 20:00:01 2017 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:13:23 2018 -0800

    distribution: separate layer and image config for v1 pushes

    When content addressablity was introduced in #17924, a compatibility
    layer for registry v1 pushes was added. When the engine is asked to
    push an image to a v1 registry it needs to create v1 IDs for the images.

    The strategy so far has been to use the full imageID for the first v1
    layer and the ChainID for all other layers, effectively creating as many
    v1 layers as there are in the image. Only the top most layer contained
    the image configuration and the other layers had a dummy json containing
    only a parent reference.

    This becomes problematic when the first layer of the image is big.
    Consinder the following two Dockerfiles:

    FROM busybox
    RUN create_very_big_file
    CMD /foo

    FROM busybox
    RUN create_very_big_file
    CMD /bar

    Both of these images will have the exact same layers, with the layer
    created by `RUN create_very_big_file` being the topmost one, but their
    imageIDs will differ since they have a different CMD and therefore
    different image configs.

    When pushing to a v1 registry, the `RUN create_very_big_file` layer will
    be pushed twice, once with the v1 ID set to foo's imageID and once with
    the v1 ID set to bar's imageID. Also, any clients wanting to pull those
    images won't realise it's the same layer and will proceed to download it
    twice.

    This commit solves this problem by separating the layers from the image
    configuration information when pushing to a v1 registry. To do this, all
    layers of an image are pushed with their ChainIDs and a synthetic top
    level layer is created with its contents set to the EmptyLayer, it's
    config set to the image config, and its v1 ID set to the imageID. This
    will have the side-effect of adding one layer.

    To prevent new layers being piled on top of each other forever, the code
    checks if the topmost layer is already an empty layer and in that case
    it uses that for the image configuration.

    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 distribution/push_v1.go | 10 ++++++++--
 1 file changed, 8 insertions(+), 2 deletions(-)

---

commit b37c7dfd27eb9139efd32b957b87f427f2fa7cba
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Wed Sep 28 01:00:07 2016 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:13:23 2018 -0800

    build: support bind mounts during docker build

    This can be used to implement various caching schemes that are finer
    grained than layer invalidation.

    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 api/server/router/build/build_routes.go |  9 +++++++++
 api/types/client.go                     |  1 +
 builder/dockerfile/internals.go         | 10 ++++++++++
 client/image_build.go                   |  7 +++++++
 4 files changed, 27 insertions(+)

---

commit 23a7f4da7ce39a3442dc8bcfdf47d82a610f7412
Author:     Yossi Eliaz <yossi@resin.io>
AuthorDate: Thu Sep 14 17:23:33 2017 -0500
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:13:22 2018 -0800

    rce-runc: changed the build to work only with resin-os/runc

    Changed the install binaries and build to clone and work
    only with the resin-os/runc repo.

    Signed-off-by: Yossi Eliaz <yossi@resin.io>

 hack/dockerfile/binaries-commits    | 4 ++--
 hack/dockerfile/install-binaries.sh | 2 +-
 2 files changed, 3 insertions(+), 3 deletions(-)

---

commit 35e9dadbdc864be2d21aae070d4e3217bf7ac45d
Author:     Yossi Eliaz <yossi@resin.io>
AuthorDate: Thu Sep 14 17:00:26 2017 -0500
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:13:22 2018 -0800

    rce-runc: porpogated the git commit and version

    Propogated properly the git commit id
    and the version to the runc main.go file from
    the Makefile.

    Signed-off-by: Yossi Eliaz <yossi@resin.io>

 hack/make/.binary | 5 +++++
 1 file changed, 5 insertions(+)

---

commit 29ece15118234d3be7ba6c56be504afd5ffa105d
Author:     Yossi Eliaz <yossi@resin.io>
AuthorDate: Tue Aug 22 00:08:04 2017 -0500
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:13:21 2018 -0800

    distribution: Added a warning when download failes

    Adds a warning when the fs download fails
    and logs the error before retrying again.

    Signed-off-by: Yossi Eliaz <yossi@resin.io>

 distribution/pull_v2.go | 1 +
 1 file changed, 1 insertion(+)

---
```

</details>

## Project maintenance

### 18.09 upstream update

<details>
<summary>Commits</summary>

```
commit beb7f70265dd873c1e674ae65fefcf9bb3e372b4
Author:     Robert Günzler <r@gnzler.io>
AuthorDate: Thu Apr 18 12:47:37 2019 +0200
Commit:     Robert Günzler <r@gnzler.io>
CommitDate: Thu Apr 18 15:19:39 2019 +0200

    integration-tests: Skip tests relying on swarm,plugin support

    18.09 moved many tests that were previously part of the
    legacy test suite under the new integration tests.
    Because of that the filtering that was done at one point in
    the past did not catch tests that make use of features no
    supported by balenaEngine.

    Some previously skipped tests that now get run by default
    require further investigation, those are marked with "TODO"

    Signed-off-by: Robert Günzler <robertg@balena.io>

 integration/config/config_test.go                | 15 +++++++++++++++
 integration/network/inspect_test.go              |  3 +++
 integration/network/ipvlan/ipvlan_test.go        | 21 +++++++++++++++++++++
 integration/network/macvlan/macvlan_test.go      |  4 ++++
 integration/network/service_test.go              |  9 +++++++++
 integration/plugin/common/plugin_test.go         |  3 +++
 integration/plugin/graphdriver/external_test.go  |  4 ++++
 integration/plugin/logging/logging_linux_test.go |  2 ++
 integration/plugin/logging/validation_test.go    |  2 ++
 integration/plugin/volumes/mounts_test.go        |  2 ++
 integration/secret/secret_test.go                | 11 +++++++++++
 integration/service/create_test.go               | 14 ++++++++++++++
 integration/service/network_test.go              |  3 +++
 integration/service/plugin_test.go               |  3 +++
 14 files changed, 96 insertions(+)

---

commit 2c29eccd53f91349fc8896d0521f5a6e9712a181
Author:     Robert Günzler <r@gnzler.io>
AuthorDate: Mon Mar 4 19:41:40 2019 +0100
Commit:     Robert Günzler <r@gnzler.io>
CommitDate: Mon Apr 8 19:29:22 2019 +0200

    Update Dockerfiles used for build to Go 1.10.8
    
    Change-type: patch
    Signed-off-by: Robert Günzler <robertg@balena.io>

 Dockerfile.build.aarch64 | 3 +--
 Dockerfile.build.arm     | 3 +--
 Dockerfile.build.i386    | 3 +--
 Dockerfile.build.x86_64  | 3 +--
 4 files changed, 4 insertions(+), 8 deletions(-)

---

commit 2189e871a666ac5263425f831b495ec932c14ac0
Author:     Robert Günzler <r@gnzler.io>
AuthorDate: Tue Feb 5 18:01:30 2019 +0100
Commit:     Robert Günzler <r@gnzler.io>
CommitDate: Mon Apr 8 19:29:22 2019 +0200

    delta: Move implementation under ImageService
    
    Signed-off-by: Robert Günzler <robertg@balena.io>

 daemon/create.go             | 242 -----------------------------------------
 daemon/daemon.go             |  22 ++--
 daemon/images/image_delta.go | 253 +++++++++++++++++++++++++++++++++++++++++++
 daemon/images/image_pull.go  |   9 +-
 daemon/images/service.go     |  11 ++
 5 files changed, 274 insertions(+), 263 deletions(-)

---

commit b43966d4a90590a78cddaf24326be0adadd1e1e9
Author:     Robert Günzler <r@gnzler.io>
AuthorDate: Tue Feb 5 16:32:28 2019 +0100
Commit:     Robert Günzler <r@gnzler.io>
CommitDate: Mon Apr 8 19:29:22 2019 +0200

    builder-next: Implement xfer.DownloadDescriptor for layerDescriptor
    
    Signed-off-by: Robert Günzler <robertg@balena.io>

 builder/builder-next/adapters/containerimage/pull.go | 8 ++++++++
 builder/builder-next/worker/worker.go                | 8 ++++++++
 2 files changed, 16 insertions(+)

---

commit 0601a3917f796ab28cb5fd34d721e291f4a51713
Author:     Robert Günzler <r@gnzler.io>
AuthorDate: Thu Mar 7 13:29:27 2019 +0100
Commit:     Robert Günzler <r@gnzler.io>
CommitDate: Mon Apr 8 19:29:22 2019 +0200

    Add 18.09.3 changelog from upstream
    
    Signed-off-by: Robert Günzler <robertg@balena.io>

 CHANGELOG.md | 103 +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 1 file changed, 103 insertions(+)

---

commit 3a3129812cf193a4bd167a035fecad2892330a00
Author:     Robert Günzler <r@gnzler.io>
AuthorDate: Mon Mar 4 15:33:08 2019 +0100
Commit:     Robert Günzler <r@gnzler.io>
CommitDate: Mon Apr 8 19:29:22 2019 +0200

    Fixes after update to 18.09
    
    Signed-off-by: Robert Günzler <robertg@balena.io>

 Dockerfile                                  | 23 -----------------------
 api/server/router/network/network_routes.go |  7 +++++++
 api/swagger.yaml                            |  8 ++++----
 builder/builder-next/executor_unix.go       |  2 +-
 cmd/dockerd/daemon.go                       |  1 +
 cmd/dockerd/docker_unix.go                  |  2 +-
 daemon/daemon_unix.go                       |  4 ++++
 daemon/kill.go                              |  1 -
 distribution/pull_v2.go                     |  2 +-
 dockerversion/useragent.go                  |  2 +-
 hack/dind                                   |  2 +-
 hack/make/.binary                           |  2 +-
 hack/make/.integration-daemon-start         |  2 +-
 hack/make/binary-balena                     |  1 +
 hack/make/dynbinary-balena                  |  1 +
 image/store.go                              |  4 ++--
 integration/container/restart_test.go       |  3 ++-
 integration/network/service_test.go         |  7 +++----
 integration/service/inspect_test.go         |  1 -
 internal/test/daemon/daemon.go              |  2 +-
 20 files changed, 33 insertions(+), 44 deletions(-)

---

commit 93ba024772c0fa3f7b58966d19130907e706b162
Author:     Robert Günzler <r@gnzler.io>
AuthorDate: Mon Mar 4 16:17:30 2019 +0100
Commit:     Robert Günzler <r@gnzler.io>
CommitDate: Mon Apr 8 19:28:07 2019 +0200

    vendor: Update dependencies to 18.09
    
    Signed-off-by: Robert Günzler <robertg@balena.io>

 distribution/xfer/download.go                      |    2 +-
 hack/make/.binary                                  |    4 +-
 vendor.conf                                        |   26 +-
 vendor/archive/tar/LICENSE                         |   27 -
 vendor/archive/tar/README.md                       |   27 -
 vendor/archive/tar/common.go                       |  720 ---
 vendor/archive/tar/format.go                       |  303 --
 vendor/archive/tar/reader.go                       |  855 ---
 vendor/archive/tar/stat_actime1.go                 |   20 -
 vendor/archive/tar/stat_actime2.go                 |   20 -
 vendor/archive/tar/stat_unix.go                    |   76 -
 vendor/archive/tar/strconv.go                      |  326 --
 vendor/archive/tar/writer.go                       |  644 ---
 .../Microsoft/hcsshim/internal/runhcs/container.go |   71 +
 .../Microsoft/hcsshim/internal/runhcs/util.go      |   16 +
 .../Microsoft/hcsshim/internal/runhcs/vm.go        |   43 +
 .../Microsoft/hcsshim/pkg/go-runhcs/LICENSE        |  201 +
 .../Microsoft/hcsshim/pkg/go-runhcs/NOTICE         |   22 +
 .../Microsoft/hcsshim/pkg/go-runhcs/runhcs.go      |  134 +
 .../hcsshim/pkg/go-runhcs/runhcs_create-scratch.go |   10 +
 .../hcsshim/pkg/go-runhcs/runhcs_create.go         |  101 +
 .../hcsshim/pkg/go-runhcs/runhcs_delete.go         |   33 +
 .../Microsoft/hcsshim/pkg/go-runhcs/runhcs_exec.go |   88 +
 .../Microsoft/hcsshim/pkg/go-runhcs/runhcs_kill.go |   11 +
 .../Microsoft/hcsshim/pkg/go-runhcs/runhcs_list.go |   28 +
 .../hcsshim/pkg/go-runhcs/runhcs_pause.go          |   10 +
 .../Microsoft/hcsshim/pkg/go-runhcs/runhcs_ps.go   |   20 +
 .../hcsshim/pkg/go-runhcs/runhcs_resize-tty.go     |   33 +
 .../hcsshim/pkg/go-runhcs/runhcs_resume.go         |   10 +
 .../hcsshim/pkg/go-runhcs/runhcs_start.go          |   10 +
 .../hcsshim/pkg/go-runhcs/runhcs_state.go          |   20 +
 .../{resin-os => balena-os}/circbuf/LICENSE        |    0
 .../{resin-os => balena-os}/circbuf/README.md      |    0
 .../{resin-os => balena-os}/circbuf/circbuf.go     |    0
 .../{resin-os => balena-os}/librsync-go/LICENSE    |    0
 .../{resin-os => balena-os}/librsync-go/README.md  |    4 +-
 .../{resin-os => balena-os}/librsync-go/delta.go   |    2 +-
 .../{resin-os => balena-os}/librsync-go/match.go   |    0
 .../{resin-os => balena-os}/librsync-go/op.go      |    0
 .../{resin-os => balena-os}/librsync-go/patch.go   |    0
 .../{resin-os => balena-os}/librsync-go/rollsum.go |    0
 .../librsync-go/signature.go                       |    0
 .../github.com/checkpoint-restore/go-criu/LICENSE  |  201 +
 .../checkpoint-restore/go-criu/README.md           |   58 +
 .../go-criu/rpc/rpc.pb.go}                         |  321 +-
 vendor/github.com/containerd/aufs/LICENSE          |  201 +
 vendor/github.com/containerd/aufs/README.md        |   23 +
 .../snapshots/overlay/overlay.go => aufs/aufs.go}  |  169 +-
 vendor/github.com/containerd/containerd/README.md  |   17 +-
 .../containerd/containerd/archive/tar.go           |    4 +-
 vendor/github.com/containerd/containerd/cio/io.go  |    4 +
 .../containerd/containerd/cio/io_unix.go           |    4 -
 .../containerd/containerd/cio/io_windows.go        |   20 +
 vendor/github.com/containerd/containerd/client.go  |   79 +-
 .../containerd/cmd/containerd-shim/main_unix.go    |   78 +-
 .../containerd/cmd/containerd-shim/shim_darwin.go  |   22 +-
 .../containerd/cmd/containerd-shim/shim_freebsd.go |   46 +
 .../containerd/cmd/containerd-shim/shim_linux.go   |   32 +-
 .../containerd/cmd/containerd-shim/shim_unix.go    |   27 -
 .../containerd/cmd/containerd/builtins.go          |   20 +-
 .../cmd/containerd/builtins_btrfs_linux.go         |   18 +-
 .../cmd/containerd/builtins_cri_linux.go           |   21 +
 .../containerd/cmd/containerd/builtins_linux.go    |   22 +-
 .../containerd/cmd/containerd/builtins_unix.go     |   18 +-
 .../containerd/cmd/containerd/builtins_windows.go  |   22 +
 .../containerd/cmd/containerd/command/config.go    |   71 +
 .../cmd/containerd/command/config_linux.go         |   34 +
 .../cmd/containerd/command/config_unsupported.go   |   38 +
 .../cmd/containerd/command/config_windows.go       |   34 +
 .../cmd/containerd/{ => command}/main.go           |  147 +-
 .../cmd/containerd/{ => command}/main_unix.go      |   46 +-
 .../cmd/containerd/command/main_windows.go         |   91 +
 .../containerd/cmd/containerd/command/oci-hook.go  |  164 +
 .../cmd/containerd/{ => command}/publish.go        |   25 +-
 .../cmd/containerd/command/service_unsupported.go  |   44 +
 .../cmd/containerd/command/service_windows.go      |  467 ++
 .../containerd/containerd/cmd/containerd/config.go |   22 -
 .../containerd/cmd/containerd/config_linux.go      |   21 -
 .../cmd/containerd/config_unsupported.go           |   22 -
 .../containerd/cmd/containerd/config_windows.go    |   20 -
 .../containerd/containerd/cmd/containerd/main.go   |  199 +-
 .../containerd/cmd/containerd/main_windows.go      |   42 -
 .../containerd/cmd/ctr/{ => app}/main.go           |   37 +-
 .../containerd/containerd/cmd/ctr/app/main_unix.go |   25 +
 .../containerd/cmd/ctr/builtins_cri_linux.go       |   25 +
 .../containerd/cmd/ctr/commands/client.go          |   20 +-
 .../containerd/cmd/ctr/commands/commands.go        |  102 +
 .../containerd/cmd/ctr/commands/commands_unix.go   |   33 +
 .../cmd/ctr/commands/commands_windows.go           |   30 +
 .../cmd/ctr/commands/containers/containers.go      |  236 +-
 .../containerd/cmd/ctr/commands/content/content.go |   44 +-
 .../containerd/cmd/ctr/commands/content/fetch.go   |   99 +-
 .../containerd/cmd/ctr/commands/events/events.go   |   69 +-
 .../containerd/cmd/ctr/commands/images/export.go   |   34 +-
 .../containerd/cmd/ctr/commands/images/images.go   |   20 +-
 .../containerd/cmd/ctr/commands/images/import.go   |   88 +-
 .../containerd/cmd/ctr/commands/images/pull.go     |   88 +-
 .../containerd/cmd/ctr/commands/images/push.go     |   21 +-
 .../containerd/cmd/ctr/commands/install/install.go |   68 +
 .../containerd/cmd/ctr/commands/leases/leases.go   |  202 +
 .../cmd/ctr/commands/namespaces/namespaces.go      |   16 +
 .../containerd/cmd/ctr/commands/plugins/plugins.go |   25 +
 .../containerd/cmd/ctr/commands/pprof/pprof.go     |   28 +-
 .../cmd/ctr/commands/pprof/pprof_unix.go           |   16 +
 .../cmd/ctr/commands/pprof/pprof_windows.go        |   16 +
 .../containerd/cmd/ctr/commands/resolver.go        |   29 +-
 .../containerd/cmd/ctr/commands/run/run.go         |  172 +-
 .../containerd/cmd/ctr/commands/run/run_unix.go    |  196 +-
 .../containerd/cmd/ctr/commands/run/run_windows.go |  168 +-
 .../containerd/cmd/ctr/commands/shim/io_unix.go    |   27 +-
 .../containerd/cmd/ctr/commands/shim/shim.go       |   61 +-
 .../cmd/ctr/commands/signal_map_linux.go           |   44 -
 .../containerd/cmd/ctr/commands/signal_map_unix.go |   42 -
 .../cmd/ctr/commands/signal_map_windows.go         |   23 -
 .../containerd/cmd/ctr/commands/signals.go         |   39 +-
 .../cmd/ctr/commands/snapshots/snapshots.go        |  222 +-
 .../containerd/cmd/ctr/commands/tasks/attach.go    |   20 +-
 .../cmd/ctr/commands/tasks/checkpoint.go           |   81 +-
 .../containerd/cmd/ctr/commands/tasks/delete.go    |   64 +-
 .../containerd/cmd/ctr/commands/tasks/exec.go      |   65 +-
 .../containerd/cmd/ctr/commands/tasks/kill.go      |   33 +-
 .../containerd/cmd/ctr/commands/tasks/list.go      |   18 +-
 .../containerd/cmd/ctr/commands/tasks/metrics.go   |  117 +
 .../containerd/cmd/ctr/commands/tasks/pause.go     |   16 +
 .../containerd/cmd/ctr/commands/tasks/ps.go        |   16 +
 .../containerd/cmd/ctr/commands/tasks/resume.go    |   16 +
 .../containerd/cmd/ctr/commands/tasks/start.go     |   75 +-
 .../containerd/cmd/ctr/commands/tasks/tasks.go     |   16 +
 .../cmd/ctr/commands/tasks/tasks_unix.go           |   62 +-
 .../cmd/ctr/commands/tasks/tasks_windows.go        |   36 +-
 .../containerd/cmd/ctr/commands/version/version.go |   16 +
 .../containerd/containerd/cmd/ctr/main.go          |  102 +-
 .../containerd/containerd/cmd/ctr/main_unix.go     |    9 -
 .../github.com/containerd/containerd/container.go  |   76 +
 .../containerd/container_checkpoint_opts.go        |  155 +
 .../containerd/containerd/container_opts_unix.go   |   69 -
 .../containerd/container_restore_opts.go           |  150 +
 .../{seccomp/seccomp.go => apparmor/apparmor.go}   |   55 +-
 .../containerd/contrib/apparmor/template.go        |  209 +
 .../containerd/containerd/contrib/nvidia/nvidia.go |  207 +
 .../containerd/contrib/seccomp/seccomp.go          |    4 +-
 .../containerd/containerd/diff/apply/apply.go      |  128 +
 .../containerd/containerd/diff/lcow/lcow.go        |  197 +
 .../containerd/containerd/diff/walking/differ.go   |  323 +-
 .../containerd/diff/walking/plugin/plugin.go       |   55 +
 .../containerd/containerd/diff/windows/windows.go  |  192 +
 .../containerd/containerd/errdefs/grpc.go          |    2 +-
 .../containerd/events/exchange/exchange.go         |    8 +-
 vendor/github.com/containerd/containerd/export.go  |   25 +-
 .../containerd/containerd/filters/parser.go        |    2 +-
 .../containerd/gc/scheduler/scheduler.go           |   53 +-
 .../containerd/containerd/identifiers/validate.go  |    2 +-
 vendor/github.com/containerd/containerd/image.go   |   28 +-
 .../containerd/containerd/images/mediatypes.go     |   13 +-
 .../containerd/containerd/images/oci/exporter.go   |   65 +-
 .../containerd/containerd/images/oci/importer.go   |  188 -
 vendor/github.com/containerd/containerd/install.go |    3 +-
 .../containerd/containerd/metadata/containers.go   |    2 +-
 .../containerd/containerd/metadata/content.go      |   41 +-
 .../containerd/containerd/metadata/db.go           |   26 +-
 .../containerd/containerd/metadata/images.go       |    2 +-
 .../containerd/containerd/metadata/leases.go       |    2 +-
 .../containerd/containerd/metrics/cgroups/blkio.go |   16 +
 .../containerd/metrics/cgroups/cgroups.go          |   51 +-
 .../containerd/containerd/metrics/cgroups/cpu.go   |   16 +
 .../containerd/metrics/cgroups/hugetlb.go          |   16 +
 .../containerd/metrics/cgroups/memory.go           |   16 +
 .../containerd/metrics/cgroups/metric.go           |   16 +
 .../containerd/metrics/cgroups/metrics.go          |   81 +-
 .../containerd/containerd/metrics/cgroups/oom.go   |   21 +-
 .../containerd/containerd/metrics/cgroups/pids.go  |   16 +
 .../containerd/containerd/mount/mountinfo_linux.go |    8 +-
 .../containerd/containerd/oci/spec_opts.go         |   57 +-
 .../containerd/containerd/oci/spec_opts_windows.go |   41 +
 .../containerd/{ => pkg}/progress/bar.go           |   16 +
 .../containerd/containerd/pkg/progress/doc.go      |   18 +
 .../containerd/containerd/pkg/progress/escape.go   |   24 +
 .../containerd/{ => pkg}/progress/humaans.go       |   16 +
 .../containerd/{ => pkg}/progress/writer.go        |   57 +-
 .../containerd/containerd/plugin/plugin_go18.go    |   40 +-
 .../containerd/containerd/progress/doc.go          |    2 -
 .../containerd/containerd/progress/escape.go       |    8 -
 .../containerd/remotes/docker/authorizer.go        |    8 +-
 .../containerd/remotes/docker/converter.go         |   88 +
 .../containerd/remotes/docker/resolver.go          |   29 +-
 .../containerd/runtime/linux/runctypes/runc.pb.go  |  238 +-
 .../containerd/runtime/linux/runctypes/runc.proto  |    4 +
 .../containerd/runtime/restart/monitor/change.go   |   75 +
 .../containerd/runtime/restart/monitor/monitor.go  |  223 +
 .../containerd/runtime/restart/restart.go          |   78 +
 .../containerd/runtime/v1/linux/bundle.go          |   17 +-
 .../containerd/runtime/v1/linux/proc/init.go       |   11 +-
 .../containerd/runtime/v1/linux/proc/types.go      |    1 +
 .../containerd/runtime/v1/linux/runtime.go         |   34 +-
 .../containerd/containerd/runtime/v1/shim.go       |   38 +
 .../containerd/runtime/v1/shim/client/client.go    |   35 +-
 .../containerd/runtime/v1/shim/service.go          |    7 +
 .../containerd/containerd/runtime/v2/README.md     |  174 +
 .../containerd/containerd/runtime/v2/binary.go     |  167 +
 .../containerd/containerd/runtime/v2/bundle.go     |  144 +
 .../containerd/containerd/runtime/v2/manager.go    |  254 +
 .../containerd/runtime/v2/manager_unix.go          |   28 +
 .../containerd/runtime/v2/manager_windows.go       |   34 +
 .../containerd/containerd/runtime/v2/process.go    |  160 +
 .../containerd/runtime/v2/runc/options/doc.go      |   17 +
 .../runc.pb.go => v2/runc/options/oci.pb.go}       |  853 ++-
 .../containerd/runtime/v2/runc/options/oci.proto   |   58 +
 .../containerd/runtime/v2/runhcs/options/doc.go    |   17 +
 .../runtime/v2/runhcs/options/runhcs.pb.go         |  571 ++
 .../runtime/v2/runhcs/options/runhcs.proto         |   43 +
 .../containerd/containerd/runtime/v2/shim.go       |  409 ++
 .../containerd/runtime/v2/shim/reaper_unix.go      |  109 +
 .../containerd/containerd/runtime/v2/shim/shim.go  |  288 +
 .../containerd/runtime/v2/shim/shim_darwin.go      |   29 +
 .../containerd/runtime/v2/shim/shim_freebsd.go     |   29 +
 .../containerd/runtime/v2/shim/shim_linux.go       |   30 +
 .../containerd/runtime/v2/shim/shim_unix.go        |  130 +
 .../containerd/runtime/v2/shim/shim_windows.go     |  302 ++
 .../containerd/containerd/runtime/v2/shim/util.go  |  142 +
 .../containerd/runtime/v2/shim/util_unix.go        |   70 +
 .../containerd/runtime/v2/shim/util_windows.go     |   90 +
 .../containerd/containerd/runtime/v2/shim_unix.go  |   32 +
 .../containerd/runtime/v2/shim_windows.go          |   87 +
 .../containerd/containerd/runtime/v2/task/doc.go   |   17 +
 .../containerd/runtime/v2/task/shim.pb.go          | 5693 ++++++++++++++++++++
 .../containerd/runtime/v2/task/shim.proto          |  184 +
 .../containerd/services/containers/helpers.go      |   16 +
 .../containerd/services/containers/local.go        |  245 +
 .../containerd/services/containers/service.go      |  188 +-
 .../containerd/services/content/service.go         |   86 +-
 .../containerd/services/content/store.go           |   71 +
 .../services/diff/{service.go => local.go}         |   65 +-
 .../containerd/containerd/services/diff/service.go |  155 +-
 .../containerd/services/diff/service_unix.go       |   23 +
 .../containerd/services/diff/service_windows.go    |   23 +
 .../containerd/services/events/service.go          |   19 +-
 .../containerd/services/healthcheck/service.go     |   16 +
 .../containerd/services/images/helpers.go          |   16 +
 .../containerd/containerd/services/images/local.go |  182 +
 .../containerd/services/images/service.go          |  185 +-
 .../containerd/services/introspection/service.go   |   27 +-
 .../containerd/containerd/services/leases/local.go |  109 +
 .../containerd/services/leases/service.go          |  121 +-
 .../services/namespaces/{service.go => local.go}   |   86 +-
 .../containerd/services/namespaces/service.go      |  201 +-
 .../containerd/services/opt/path_unix.go           |   21 +
 .../containerd/services/opt/path_windows.go        |   25 +
 .../containerd/containerd/services/opt/service.go  |   68 +
 .../containerd/services/server/config/config.go    |   38 +
 .../containerd/services/server/server.go           |  394 ++
 .../containerd/services/server/server_linux.go     |   55 +
 .../containerd/services/server/server_solaris.go   |   27 +
 .../services/server/server_unsupported.go          |   29 +
 .../containerd/services/server/server_windows.go   |   29 +
 .../containerd/containerd/services/services.go     |   36 +
 .../containerd/services/snapshots/service.go       |   67 +-
 .../containerd/services/snapshots/snapshotters.go  |   98 +
 .../services/tasks/{service.go => local.go}        |  371 +-
 .../containerd/services/tasks/local_unix.go        |   54 +
 .../containerd/services/tasks/service.go           |  592 +-
 .../containerd/services/version/service.go         |   19 +-
 .../containerd/containerd/snapshots/btrfs/btrfs.go |   25 +-
 .../containerd/containerd/snapshots/lcow/lcow.go   |  438 ++
 .../snapshots/{naive/naive.go => native/native.go} |   33 +-
 .../containerd/snapshots/overlay/check.go          |   88 +
 .../containerd/snapshots/overlay/overlay.go        |  223 +-
 .../containerd/snapshots/storage/bolt.go           |   38 +-
 .../containerd/snapshots/storage/metastore.go      |   18 +-
 .../containerd/snapshots/windows/windows.go        |  331 +-
 vendor/github.com/containerd/containerd/task.go    |   34 +
 .../github.com/containerd/containerd/task_opts.go  |   61 +
 .../github.com/containerd/containerd/vendor.conf   |   32 +-
 .../containerd/containerd/version/version.go       |   18 +-
 .../containerd/containerd/windows/hcsshim.go       |   96 +-
 .../github.com/containerd/containerd/windows/io.go |   18 +-
 .../containerd/containerd/windows/meta.go          |   54 -
 .../containerd/containerd/windows/pid_pool.go      |   16 +
 .../containerd/containerd/windows/process.go       |   36 +-
 .../containerd/containerd/windows/runtime.go       |  256 +-
 .../containerd/containerd/windows/task.go          |  117 +-
 vendor/github.com/containerd/cri/cli/cli.go        |   78 +
 vendor/github.com/containerd/cri/cri.go            |  195 +
 .../containerd/cri/pkg/annotations/annotations.go  |   38 +
 .../github.com/containerd/cri/pkg/api/v1/api.pb.go |  597 ++
 .../github.com/containerd/cri/pkg/api/v1/api.proto |   30 +
 .../containerd/cri/pkg/atomic/atomic_boolean.go    |   54 +
 .../github.com/containerd/cri/pkg/client/client.go |   45 +
 .../github.com/containerd/cri/pkg/config/config.go |  212 +
 .../containerd/cri/pkg/constants/constants.go      |   26 +
 .../cri/pkg/containerd/importer/importer.go        |  356 ++
 .../cri/pkg/containerd/opts/container.go           |  118 +
 .../containerd/cri/pkg/containerd/opts/spec.go     |   51 +
 .../containerd/cri/pkg/containerd/opts/task.go     |   38 +
 .../containerd/cri/pkg/containerd/util/util.go     |   46 +
 .../containerd/cri/pkg/ioutil/read_closer.go       |   57 +
 .../containerd/cri/pkg/ioutil/write_closer.go      |  102 +
 .../containerd/cri/pkg/ioutil/writer_group.go      |  105 +
 vendor/github.com/containerd/cri/pkg/log/log.go    |   29 +
 .../github.com/containerd/cri/pkg/netns/netns.go   |  220 +
 vendor/github.com/containerd/cri/pkg/os/os.go      |  142 +
 .../containerd/cri/pkg/registrar/registrar.go      |  102 +
 .../containerd/cri/pkg/server/container_attach.go  |   82 +
 .../containerd/cri/pkg/server/container_create.go  | 1044 ++++
 .../containerd/cri/pkg/server/container_exec.go    |   36 +
 .../cri/pkg/server/container_execsync.go           |  195 +
 .../containerd/cri/pkg/server/container_list.go    |  101 +
 .../cri/pkg/server/container_log_reopen.go         |   51 +
 .../containerd/cri/pkg/server/container_remove.go  |  120 +
 .../containerd/cri/pkg/server/container_start.go   |  182 +
 .../containerd/cri/pkg/server/container_stats.go   |   47 +
 .../cri/pkg/server/container_stats_list.go         |  185 +
 .../containerd/cri/pkg/server/container_status.go  |  164 +
 .../containerd/cri/pkg/server/container_stop.go    |  143 +
 .../cri/pkg/server/container_update_resources.go   |  161 +
 .../github.com/containerd/cri/pkg/server/events.go |  424 ++
 .../containerd/cri/pkg/server/helpers.go           |  498 ++
 .../containerd/cri/pkg/server/image_list.go        |   38 +
 .../containerd/cri/pkg/server/image_load.go        |   56 +
 .../containerd/cri/pkg/server/image_pull.go        |  285 +
 .../containerd/cri/pkg/server/image_remove.go      |   65 +
 .../containerd/cri/pkg/server/image_status.go      |  105 +
 .../containerd/cri/pkg/server/imagefs_info.go      |   51 +
 .../cri/pkg/server/instrumented_service.go         |  479 ++
 .../containerd/cri/pkg/server/io/container_io.go   |  234 +
 .../containerd/cri/pkg/server/io/exec_io.go        |  146 +
 .../containerd/cri/pkg/server/io/helpers.go        |  141 +
 .../containerd/cri/pkg/server/io/logger.go         |  136 +
 .../containerd/cri/pkg/server/restart.go           |  464 ++
 .../containerd/cri/pkg/server/sandbox_list.go      |  100 +
 .../cri/pkg/server/sandbox_portforward.go          |  125 +
 .../containerd/cri/pkg/server/sandbox_remove.go    |  111 +
 .../containerd/cri/pkg/server/sandbox_run.go       |  641 +++
 .../containerd/cri/pkg/server/sandbox_status.go    |  193 +
 .../containerd/cri/pkg/server/sandbox_stop.go      |  139 +
 .../containerd/cri/pkg/server/service.go           |  269 +
 .../containerd/cri/pkg/server/snapshots.go         |  120 +
 .../github.com/containerd/cri/pkg/server/status.go |   77 +
 .../containerd/cri/pkg/server/streaming.go         |  226 +
 .../cri/pkg/server/update_runtime_config.go        |   78 +
 .../containerd/cri/pkg/server/version.go           |   42 +
 .../cri/pkg/store/container/container.go           |  169 +
 .../cri/pkg/store/container/fake_status.go         |   62 +
 .../containerd/cri/pkg/store/container/metadata.go |   87 +
 .../containerd/cri/pkg/store/container/status.go   |  206 +
 .../github.com/containerd/cri/pkg/store/errors.go  |   27 +
 .../containerd/cri/pkg/store/image/fake_image.go   |   34 +
 .../containerd/cri/pkg/store/image/image.go        |  259 +
 .../containerd/cri/pkg/store/sandbox/metadata.go   |   82 +
 .../containerd/cri/pkg/store/sandbox/sandbox.go    |  146 +
 .../containerd/cri/pkg/store/sandbox/status.go     |  100 +
 .../containerd/cri/pkg/store/snapshot/snapshot.go  |   87 +
 vendor/github.com/containerd/cri/pkg/store/util.go |   42 +
 vendor/github.com/docker/cli/cli/cobra.go          |   18 +-
 .../docker/cli/cli/command/builder/cmd.go          |   22 +
 .../cli/cli/command/{image => builder}/prune.go    |   63 +-
 vendor/github.com/docker/cli/cli/command/cli.go    |  175 +-
 .../docker/cli/cli/command/commands/commands.go    |   17 +-
 .../docker/cli/cli/command/container/attach.go     |   42 +-
 .../docker/cli/cli/command/container/commit.go     |    2 +-
 .../docker/cli/cli/command/container/cp.go         |    2 +-
 .../docker/cli/cli/command/container/create.go     |   21 +-
 .../docker/cli/cli/command/container/diff.go       |    3 +-
 .../docker/cli/cli/command/container/exec.go       |    2 +-
 .../docker/cli/cli/command/container/export.go     |    2 +-
 .../docker/cli/cli/command/container/hijack.go     |    2 +-
 .../docker/cli/cli/command/container/inspect.go    |    3 +-
 .../docker/cli/cli/command/container/kill.go       |    2 +-
 .../docker/cli/cli/command/container/list.go       |    2 +-
 .../docker/cli/cli/command/container/logs.go       |    2 +-
 .../docker/cli/cli/command/container/opts.go       |   40 +-
 .../docker/cli/cli/command/container/pause.go      |    2 +-
 .../docker/cli/cli/command/container/port.go       |    2 +-
 .../docker/cli/cli/command/container/prune.go      |    4 +-
 .../docker/cli/cli/command/container/rename.go     |    2 +-
 .../docker/cli/cli/command/container/restart.go    |    2 +-
 .../docker/cli/cli/command/container/rm.go         |    2 +-
 .../docker/cli/cli/command/container/run.go        |   17 +-
 .../docker/cli/cli/command/container/start.go      |    5 +-
 .../docker/cli/cli/command/container/stats.go      |    2 +-
 .../cli/cli/command/container/stats_helpers.go     |    2 +-
 .../docker/cli/cli/command/container/stop.go       |    2 +-
 .../docker/cli/cli/command/container/top.go        |    2 +-
 .../docker/cli/cli/command/container/tty.go        |    2 +-
 .../docker/cli/cli/command/container/unpause.go    |    2 +-
 .../docker/cli/cli/command/container/update.go     |    2 +-
 .../docker/cli/cli/command/container/utils.go      |    9 +-
 .../docker/cli/cli/command/container/wait.go       |    2 +-
 .../docker/cli/cli/command/engine/activate_unix.go |   13 +
 .../cli/cli/command/engine/activate_windows.go     |    9 +
 .../docker/cli/cli/command/engine/auth.go          |   34 +
 .../docker/cli/cli/command/engine/check.go         |  125 +
 .../docker/cli/cli/command/engine/cmd.go           |   22 +
 .../docker/cli/cli/command/engine/init.go          |   10 +
 .../docker/cli/cli/command/engine/update.go        |   55 +
 .../docker/cli/cli/command/formatter/buildcache.go |  179 +
 .../docker/cli/cli/command/formatter/container.go  |   26 +-
 .../docker/cli/cli/command/formatter/disk_usage.go |  155 +-
 .../docker/cli/cli/command/formatter/node.go       |    8 +-
 .../docker/cli/cli/command/formatter/service.go    |  127 +-
 .../docker/cli/cli/command/formatter/stack.go      |   36 +-
 .../docker/cli/cli/command/formatter/trust.go      |   15 -
 .../docker/cli/cli/command/formatter/updates.go    |   73 +
 .../docker/cli/cli/command/image/build.go          |  161 +-
 .../docker/cli/cli/command/image/build/context.go  |   83 +-
 .../docker/cli/cli/command/image/build_buildkit.go |  441 ++
 .../docker/cli/cli/command/image/build_session.go  |   16 +-
 .../docker/cli/cli/command/image/delta.go          |    9 +-
 .../docker/cli/cli/command/image/history.go        |    2 +-
 .../docker/cli/cli/command/image/import.go         |    9 +-
 .../docker/cli/cli/command/image/inspect.go        |    2 +-
 .../docker/cli/cli/command/image/list.go           |    3 +-
 .../docker/cli/cli/command/image/load.go           |    3 +-
 .../docker/cli/cli/command/image/prune.go          |   36 +-
 .../docker/cli/cli/command/image/pull.go           |   25 +-
 .../docker/cli/cli/command/image/push.go           |   21 +-
 .../docker/cli/cli/command/image/remove.go         |    5 +-
 .../docker/cli/cli/command/image/save.go           |    7 +-
 .../github.com/docker/cli/cli/command/image/tag.go |    2 +-
 .../docker/cli/cli/command/image/trust.go          |   29 +-
 .../docker/cli/cli/command/manifest/annotate.go    |   97 +
 .../docker/cli/cli/command/manifest/cmd.go         |   45 +
 .../docker/cli/cli/command/manifest/create_list.go |   82 +
 .../docker/cli/cli/command/manifest/inspect.go     |  148 +
 .../docker/cli/cli/command/manifest/push.go        |  281 +
 .../docker/cli/cli/command/manifest/util.go        |   81 +
 .../docker/cli/cli/command/network/connect.go      |    3 +-
 .../docker/cli/cli/command/network/create.go       |    2 +-
 .../docker/cli/cli/command/network/disconnect.go   |    2 +-
 .../docker/cli/cli/command/network/inspect.go      |    2 +-
 .../docker/cli/cli/command/network/list.go         |   13 +-
 .../docker/cli/cli/command/network/prune.go        |    4 +-
 .../docker/cli/cli/command/network/remove.go       |    3 +-
 .../docker/cli/cli/command/orchestrator.go         |   72 +-
 .../github.com/docker/cli/cli/command/registry.go  |   58 +-
 .../docker/cli/cli/command/registry/login.go       |  107 +-
 .../docker/cli/cli/command/registry/logout.go      |    3 +-
 .../docker/cli/cli/command/registry/search.go      |   17 +-
 .../docker/cli/cli/command/system/cmd.go           |    1 +
 .../github.com/docker/cli/cli/command/system/df.go |   19 +-
 .../docker/cli/cli/command/system/dial_stdio.go    |  107 +
 .../docker/cli/cli/command/system/events.go        |    2 +-
 .../docker/cli/cli/command/system/info.go          |  134 +-
 .../docker/cli/cli/command/system/inspect.go       |   38 +-
 .../docker/cli/cli/command/system/prune.go         |   42 +-
 .../docker/cli/cli/command/system/version.go       |   54 +-
 vendor/github.com/docker/cli/cli/command/trust.go  |   40 +-
 .../github.com/docker/cli/cli/command/trust/cmd.go |   10 +-
 .../docker/cli/cli/command/trust/common.go         |   45 +-
 .../docker/cli/cli/command/trust/inspect.go        |   46 +-
 .../command/trust/{view.go => inspect_pretty.go}   |   36 +-
 .../github.com/docker/cli/cli/command/trust/key.go |    2 +-
 .../docker/cli/cli/command/trust/revoke.go         |    2 +-
 .../docker/cli/cli/command/trust/sign.go           |    7 +-
 .../docker/cli/cli/command/trust/signer.go         |    2 +-
 .../docker/cli/cli/command/trust/signer_add.go     |    2 +-
 .../docker/cli/cli/command/trust/signer_remove.go  |   39 +-
 vendor/github.com/docker/cli/cli/command/utils.go  |    4 +-
 .../docker/cli/cli/command/volume/create.go        |    4 +-
 .../docker/cli/cli/command/volume/inspect.go       |    3 +-
 .../docker/cli/cli/command/volume/list.go          |   16 +-
 .../docker/cli/cli/command/volume/prune.go         |    8 +-
 .../docker/cli/cli/command/volume/remove.go        |    2 +-
 .../cli/cli/compose/interpolation/interpolation.go |    7 +-
 .../docker/cli/cli/compose/loader/interpolate.go   |    9 +-
 .../docker/cli/cli/compose/loader/loader.go        |  182 +-
 .../docker/cli/cli/compose/loader/merge.go         |  233 +
 .../docker/cli/cli/compose/schema/bindata.go       |  726 ++-
 .../docker/cli/cli/compose/schema/schema.go        |    4 +-
 .../docker/cli/cli/compose/template/template.go    |  216 +-
 .../docker/cli/cli/compose/types/types.go          |  342 +-
 vendor/github.com/docker/cli/cli/config/config.go  |   10 +-
 .../docker/cli/cli/config/configfile/file.go       |   74 +-
 .../cli/config/credentials/default_store_linux.go  |    4 +-
 .../credentials/default_store_unsupported.go       |    4 +-
 .../cli/cli/config/credentials/file_store.go       |    9 +
 .../docker/cli/cli/connhelper/connhelper.go        |  302 ++
 .../docker/cli/cli/connhelper/connhelper_linux.go  |   12 +
 .../cli/cli/connhelper/connhelper_nolinux.go       |   10 +
 .../docker/cli/cli/connhelper/ssh/ssh.go           |   70 +
 vendor/github.com/docker/cli/cli/flags/client.go   |    1 -
 vendor/github.com/docker/cli/cli/flags/common.go   |   21 +-
 .../docker/cli/cli/manifest/store/store.go         |  180 +
 .../docker/cli/cli/manifest/types/types.go         |  114 +
 .../docker/cli/cli/registry/client/client.go       |  211 +
 .../docker/cli/cli/registry/client/endpoint.go     |  133 +
 .../docker/cli/cli/registry/client/fetcher.go      |  308 ++
 vendor/github.com/docker/cli/cli/trust/trust.go    |   19 +-
 vendor/github.com/docker/cli/cli/version.go        |    2 +-
 vendor/github.com/docker/cli/cmd/docker/docker.go  |  133 +-
 .../cli/internal/containerizedengine/containerd.go |   78 +
 .../cli/internal/containerizedengine/progress.go   |  215 +
 .../cli/internal/containerizedengine/types.go      |   49 +
 .../cli/internal/containerizedengine/update.go     |  183 +
 .../docker/cli/internal/versions/versions.go       |  127 +
 vendor/github.com/docker/cli/opts/envfile.go       |   61 +-
 .../docker/cli/opts/{envfile.go => file.go}        |   54 +-
 vendor/github.com/docker/cli/opts/hosts.go         |    2 +
 vendor/github.com/docker/cli/opts/opts.go          |   29 +-
 vendor/github.com/docker/cli/opts/parse.go         |   22 +-
 vendor/github.com/docker/cli/opts/port.go          |   27 +-
 vendor/github.com/docker/cli/types/types.go        |   88 +
 vendor/github.com/docker/cli/vendor.conf           |  158 +-
 .../docker-credential-helpers/pass/pass_linux.go   |  208 -
 vendor/github.com/docker/libkv/store/etcd/etcd.go  |  604 ---
 .../github.com/docker/libnetwork/cmd/proxy/main.go |    7 +-
 .../docker/libnetwork/cmd/proxy/proxy.go           |    4 +
 .../docker/libnetwork/cmd/proxy/sctp_proxy.go      |   93 +
 .../libnetwork/drivers/macvlan/macvlan_endpoint.go |   96 -
 .../drivers/macvlan/macvlan_joinleave.go           |  144 -
 .../libnetwork/drivers/macvlan/macvlan_network.go  |  260 -
 .../libnetwork/drivers/macvlan/macvlan_setup.go    |  209 -
 .../libnetwork/drivers/macvlan/macvlan_store.go    |  351 --
 .../libnetwork/drivers/overlay/encryption.go       |  639 ---
 .../docker/libnetwork/drivers/overlay/joinleave.go |  232 -
 .../libnetwork/drivers/overlay/ostweaks_linux.go   |   23 -
 .../libnetwork/drivers/overlay/ov_endpoint.go      |  252 -
 .../libnetwork/drivers/overlay/ov_network.go       | 1155 ----
 .../docker/libnetwork/drivers/overlay/overlay.go   |  392 --
 .../libnetwork/drivers/overlay/overlay.pb.go       |  455 --
 .../docker/libnetwork/drivers/overlay/peerdb.go    |  526 --
 .../docker/libnetwork/portmapper/proxy.go          |    2 +-
 vendor/github.com/docker/libnetwork/store.go       |    6 -
 vendor/github.com/hashicorp/go-version/LICENSE     |  354 ++
 vendor/github.com/hashicorp/go-version/README.md   |   65 +
 .../github.com/hashicorp/go-version/constraint.go  |  204 +
 vendor/github.com/hashicorp/go-version/version.go  |  326 ++
 .../hashicorp/go-version/version_collection.go     |   17 +
 .../session/auth/authprovider/authprovider.go      |   44 +
 .../session/secrets/secretsprovider/file.go        |   54 +
 .../secrets/secretsprovider/secretsprovider.go     |   60 +
 .../sshforward/sshprovider/agentprovider.go        |  198 +
 .../moby/buildkit/util/appcontext/appcontext.go    |   41 +
 .../buildkit/util/appcontext/appcontext_unix.go    |   11 +
 .../buildkit/util/appcontext/appcontext_windows.go |    7 +
 .../buildkit/util/progress/progressui/display.go   |  432 ++
 .../buildkit/util/progress/progressui/printer.go   |  248 +
 vendor/github.com/morikuni/aec/LICENSE             |   21 +
 vendor/github.com/morikuni/aec/README.md           |  178 +
 vendor/github.com/morikuni/aec/aec.go              |  137 +
 vendor/github.com/morikuni/aec/ansi.go             |   59 +
 vendor/github.com/morikuni/aec/builder.go          |  388 ++
 vendor/github.com/morikuni/aec/sgr.go              |  202 +
 vendor/github.com/opencontainers/runc/README.md    |    5 +-
 .../github.com/opencontainers/runc/checkpoint.go   |   19 +-
 vendor/github.com/opencontainers/runc/delete.go    |    3 +-
 vendor/github.com/opencontainers/runc/events.go    |   38 +-
 vendor/github.com/opencontainers/runc/exec.go      |    6 +
 vendor/github.com/opencontainers/runc/kill.go      |    5 +-
 .../runc/libcontainer/capabilities_linux.go        |    3 +-
 .../runc/libcontainer/cgroups/fs/apply_raw.go      |   64 +-
 .../runc/libcontainer/cgroups/fs/cpu.go            |   12 +-
 .../runc/libcontainer/cgroups/fs/cpuset.go         |    8 +-
 .../runc/libcontainer/cgroups/fs/kmem.go           |   62 +
 .../runc/libcontainer/cgroups/fs/kmem_disabled.go  |   15 +
 .../runc/libcontainer/cgroups/fs/memory.go         |   47 +-
 .../libcontainer/cgroups/systemd/apply_systemd.go  |  141 +-
 .../runc/libcontainer/cgroups/utils.go             |   38 +-
 .../runc/libcontainer/configs/validate/rootless.go |   48 +-
 .../libcontainer/configs/validate/validator.go     |   45 +-
 .../opencontainers/runc/libcontainer/container.go  |    7 +
 .../runc/libcontainer/container_linux.go           |  447 +-
 .../runc/libcontainer/container_solaris.go         |   20 -
 .../runc/libcontainer/container_windows.go         |   20 -
 .../runc/libcontainer/criu_opts_linux.go           |    2 +-
 .../runc/libcontainer/criu_opts_windows.go         |    6 -
 .../runc/libcontainer/criurpc/criurpc.proto        |  209 -
 .../runc/libcontainer/factory_linux.go             |   52 +-
 .../opencontainers/runc/libcontainer/init_linux.go |   63 +-
 .../runc/libcontainer/intelrdt/intelrdt.go         |  454 +-
 .../runc/libcontainer/intelrdt/stats.go            |   16 +
 .../runc/libcontainer/keys/keyctl.go               |   10 +-
 .../runc/libcontainer/message_linux.go             |   26 +-
 .../runc/libcontainer/mount/mount_freebsd.go       |   41 -
 .../runc/libcontainer/mount/mount_unsupported.go   |   12 -
 .../runc/libcontainer/network_linux.go             |  157 -
 .../runc/libcontainer/nsenter/cloned_binary.c      |  384 +-
 .../runc/libcontainer/nsenter/nsexec.c             |   11 +
 .../opencontainers/runc/libcontainer/process.go    |    3 +
 .../runc/libcontainer/process_linux.go             |  120 +-
 .../runc/libcontainer/rootfs_linux.go              |  177 +-
 .../runc/libcontainer/setns_init_linux.go          |   22 +-
 .../runc/libcontainer/specconv/example.go          |    6 +-
 .../runc/libcontainer/specconv/spec_linux.go       |  116 +-
 .../runc/libcontainer/standard_init_linux.go       |   68 +-
 .../runc/libcontainer/state_linux.go               |    8 +-
 .../runc/libcontainer/stats_freebsd.go             |    5 -
 .../runc/libcontainer/stats_solaris.go             |    7 -
 .../runc/libcontainer/stats_windows.go             |    5 -
 .../opencontainers/runc/libcontainer/sync.go       |    5 +-
 .../runc/libcontainer/utils/utils.go               |   15 -
 vendor/github.com/opencontainers/runc/main.go      |   19 +-
 .../opencontainers/runc/notify_socket.go           |   12 +-
 vendor/github.com/opencontainers/runc/pause.go     |   31 +-
 vendor/github.com/opencontainers/runc/ps.go        |   10 +-
 vendor/github.com/opencontainers/runc/restore.go   |   10 +-
 .../opencontainers/runc/rootless_linux.go          |   58 +
 vendor/github.com/opencontainers/runc/signals.go   |    5 +-
 vendor/github.com/opencontainers/runc/spec.go      |   11 +-
 vendor/github.com/opencontainers/runc/tty.go       |    9 +-
 vendor/github.com/opencontainers/runc/update.go    |   23 +-
 vendor/github.com/opencontainers/runc/utils.go     |   11 +
 .../github.com/opencontainers/runc/utils_linux.go  |   35 +-
 vendor/github.com/opencontainers/runc/vendor.conf  |    5 +-
 .../opencontainers/runtime-spec/specs-go/config.go |   12 +-
 .../selinux/go-selinux/label/label_selinux.go      |    7 +-
 .../go-selinux/{selinux_linux.go => selinux.go}    |  278 +-
 .../selinux/go-selinux/selinux_stub.go             |  188 -
 .../opencontainers/selinux/go-selinux/xattrs.go    |    2 +-
 vendor/github.com/tonistiigi/units/LICENSE         |   21 +
 vendor/github.com/tonistiigi/units/bytes.go        |  125 +
 vendor/github.com/tonistiigi/units/readme.md       |   29 +
 vendor/golang.org/x/crypto/blake2b/blake2b.go      |   84 +-
 .../x/crypto/blake2b/blake2bAVX2_amd64.go          |   26 +-
 .../x/crypto/blake2b/blake2bAVX2_amd64.s           |   12 -
 .../golang.org/x/crypto/blake2b/blake2b_amd64.go   |    7 +-
 vendor/golang.org/x/crypto/blake2b/blake2b_amd64.s |    9 -
 .../x/crypto/internal/chacha20/chacha_generic.go   |  264 +
 .../x/crypto/internal/chacha20/chacha_noasm.go     |   16 +
 .../x/crypto/internal/chacha20/chacha_s390x.go     |   30 +
 .../x/crypto/internal/chacha20/chacha_s390x.s      |  283 +
 .../golang.org/x/crypto/internal/chacha20/xor.go   |   43 +
 vendor/golang.org/x/crypto/ssh/agent/client.go     |  683 +++
 vendor/golang.org/x/crypto/ssh/agent/forward.go    |  103 +
 vendor/golang.org/x/crypto/ssh/agent/keyring.go    |  215 +
 vendor/golang.org/x/crypto/ssh/agent/server.go     |  523 ++
 vendor/golang.org/x/crypto/ssh/buffer.go           |   97 +
 vendor/golang.org/x/crypto/ssh/certs.go            |  521 ++
 vendor/golang.org/x/crypto/ssh/channel.go          |  633 +++
 vendor/golang.org/x/crypto/ssh/cipher.go           |  770 +++
 vendor/golang.org/x/crypto/ssh/client.go           |  278 +
 vendor/golang.org/x/crypto/ssh/client_auth.go      |  525 ++
 vendor/golang.org/x/crypto/ssh/common.go           |  383 ++
 vendor/golang.org/x/crypto/ssh/connection.go       |  143 +
 vendor/golang.org/x/crypto/ssh/doc.go              |   21 +
 vendor/golang.org/x/crypto/ssh/handshake.go        |  646 +++
 vendor/golang.org/x/crypto/ssh/kex.go              |  540 ++
 vendor/golang.org/x/crypto/ssh/keys.go             | 1035 ++++
 vendor/golang.org/x/crypto/ssh/mac.go              |   61 +
 vendor/golang.org/x/crypto/ssh/messages.go         |  766 +++
 vendor/golang.org/x/crypto/ssh/mux.go              |  330 ++
 vendor/golang.org/x/crypto/ssh/server.go           |  593 ++
 vendor/golang.org/x/crypto/ssh/session.go          |  647 +++
 vendor/golang.org/x/crypto/ssh/streamlocal.go      |  116 +
 vendor/golang.org/x/crypto/ssh/tcpip.go            |  474 ++
 vendor/golang.org/x/crypto/ssh/transport.go        |  353 ++
 vendor/golang.org/x/sys/cpu/cpu.go                 |   70 +
 vendor/golang.org/x/sys/cpu/cpu_arm.go             |    9 +
 vendor/golang.org/x/sys/cpu/cpu_arm64.go           |   67 +
 vendor/golang.org/x/sys/cpu/cpu_gc_x86.go          |   16 +
 vendor/golang.org/x/sys/cpu/cpu_gccgo.c            |   43 +
 vendor/golang.org/x/sys/cpu/cpu_gccgo.go           |   26 +
 vendor/golang.org/x/sys/cpu/cpu_linux.go           |   55 +
 vendor/golang.org/x/sys/cpu/cpu_mips64x.go         |   11 +
 vendor/golang.org/x/sys/cpu/cpu_mipsx.go           |   11 +
 vendor/golang.org/x/sys/cpu/cpu_ppc64x.go          |   11 +
 vendor/golang.org/x/sys/cpu/cpu_s390x.go           |    9 +
 vendor/golang.org/x/sys/cpu/cpu_x86.go             |   55 +
 vendor/golang.org/x/sys/cpu/cpu_x86.s              |   27 +
 vendor/golang.org/x/sys/unix/asm_aix_ppc64.s       |   17 +
 vendor/golang.org/x/sys/unix/asm_linux_ppc64x.s    |   12 -
 vendor/golang.org/x/sys/unix/dirent.go             |    2 +-
 vendor/golang.org/x/sys/unix/env_unix.go           |    2 +-
 vendor/golang.org/x/sys/unix/gccgo.go              |    1 +
 vendor/golang.org/x/sys/unix/gccgo_c.c             |    1 +
 vendor/golang.org/x/sys/unix/openbsd_pledge.go     |  152 +-
 vendor/golang.org/x/sys/unix/openbsd_unveil.go     |   44 +
 vendor/golang.org/x/sys/unix/sockcmsg_unix.go      |   25 +-
 vendor/golang.org/x/sys/unix/syscall.go            |    2 +-
 vendor/golang.org/x/sys/unix/syscall_aix.go        |   31 +-
 vendor/golang.org/x/sys/unix/syscall_darwin.go     |   36 +-
 vendor/golang.org/x/sys/unix/syscall_dragonfly.go  |   12 +-
 vendor/golang.org/x/sys/unix/syscall_freebsd.go    |  557 +-
 vendor/golang.org/x/sys/unix/syscall_linux.go      |  228 +-
 .../golang.org/x/sys/unix/syscall_linux_amd64.go   |   31 +-
 vendor/golang.org/x/sys/unix/syscall_linux_arm.go  |    8 +
 .../golang.org/x/sys/unix/syscall_linux_arm64.go   |    9 +-
 .../golang.org/x/sys/unix/syscall_linux_mipsx.go   |    7 +-
 .../golang.org/x/sys/unix/syscall_linux_ppc64x.go  |   22 +-
 .../golang.org/x/sys/unix/syscall_linux_riscv64.go |    9 +-
 .../golang.org/x/sys/unix/syscall_linux_s390x.go   |   13 +
 vendor/golang.org/x/sys/unix/syscall_netbsd.go     |   51 +-
 vendor/golang.org/x/sys/unix/syscall_openbsd.go    |   52 +-
 .../golang.org/x/sys/unix/syscall_openbsd_386.go   |    4 +
 .../golang.org/x/sys/unix/syscall_openbsd_arm.go   |    4 +
 vendor/golang.org/x/sys/unix/syscall_solaris.go    |    4 +-
 vendor/golang.org/x/sys/unix/syscall_unix.go       |   10 +-
 vendor/golang.org/x/sys/unix/syscall_unix_gc.go    |    2 +-
 ...yscall_unix_gc.go => syscall_unix_gc_ppc64x.go} |   21 +-
 vendor/golang.org/x/sys/unix/timestruct.go         |    2 +-
 vendor/golang.org/x/sys/unix/xattr_bsd.go          |   15 +-
 .../x/sys/unix/zerrors_dragonfly_amd64.go          |   66 +-
 .../golang.org/x/sys/unix/zerrors_freebsd_386.go   |   29 +
 .../golang.org/x/sys/unix/zerrors_freebsd_amd64.go |   29 +
 .../golang.org/x/sys/unix/zerrors_freebsd_arm.go   |   29 +
 vendor/golang.org/x/sys/unix/zerrors_linux_386.go  |  146 +-
 .../golang.org/x/sys/unix/zerrors_linux_amd64.go   |  146 +-
 vendor/golang.org/x/sys/unix/zerrors_linux_arm.go  |  146 +-
 .../golang.org/x/sys/unix/zerrors_linux_arm64.go   |  146 +-
 vendor/golang.org/x/sys/unix/zerrors_linux_mips.go |  145 +-
 .../golang.org/x/sys/unix/zerrors_linux_mips64.go  |  145 +-
 .../x/sys/unix/zerrors_linux_mips64le.go           |  145 +-
 .../golang.org/x/sys/unix/zerrors_linux_mipsle.go  |  145 +-
 .../golang.org/x/sys/unix/zerrors_linux_ppc64.go   |  145 +-
 .../golang.org/x/sys/unix/zerrors_linux_ppc64le.go |  145 +-
 .../golang.org/x/sys/unix/zerrors_linux_riscv64.go |   49 +-
 .../golang.org/x/sys/unix/zerrors_linux_s390x.go   |  146 +-
 .../golang.org/x/sys/unix/zerrors_linux_sparc64.go |  350 +-
 vendor/golang.org/x/sys/unix/zerrors_netbsd_386.go |   44 +
 .../golang.org/x/sys/unix/zerrors_netbsd_amd64.go  |   44 +
 vendor/golang.org/x/sys/unix/zerrors_netbsd_arm.go |   44 +
 .../golang.org/x/sys/unix/zerrors_openbsd_386.go   |   54 +
 .../golang.org/x/sys/unix/zerrors_openbsd_amd64.go |   59 +
 .../golang.org/x/sys/unix/zerrors_openbsd_arm.go   |   54 +
 .../golang.org/x/sys/unix/zerrors_solaris_amd64.go |   35 +
 vendor/golang.org/x/sys/unix/zsyscall_aix_ppc.go   |   97 +-
 vendor/golang.org/x/sys/unix/zsyscall_aix_ppc64.go | 1073 ++--
 .../golang.org/x/sys/unix/zsyscall_aix_ppc64_gc.go | 1162 ++++
 .../x/sys/unix/zsyscall_aix_ppc64_gccgo.go         | 1042 ++++
 .../golang.org/x/sys/unix/zsyscall_darwin_386.go   |   59 +-
 .../golang.org/x/sys/unix/zsyscall_darwin_amd64.go |   59 +-
 .../golang.org/x/sys/unix/zsyscall_darwin_arm.go   |   59 +-
 .../golang.org/x/sys/unix/zsyscall_darwin_arm64.go |   59 +-
 .../x/sys/unix/zsyscall_dragonfly_amd64.go         |  133 +-
 .../golang.org/x/sys/unix/zsyscall_freebsd_386.go  |  110 +-
 .../x/sys/unix/zsyscall_freebsd_amd64.go           |  110 +-
 .../golang.org/x/sys/unix/zsyscall_freebsd_arm.go  |  110 +-
 vendor/golang.org/x/sys/unix/zsyscall_linux_386.go |  174 +-
 .../golang.org/x/sys/unix/zsyscall_linux_amd64.go  |  206 +-
 vendor/golang.org/x/sys/unix/zsyscall_linux_arm.go |  184 +-
 .../golang.org/x/sys/unix/zsyscall_linux_arm64.go  |  174 +-
 .../golang.org/x/sys/unix/zsyscall_linux_mips.go   |  186 +-
 .../golang.org/x/sys/unix/zsyscall_linux_mips64.go |  174 +-
 .../x/sys/unix/zsyscall_linux_mips64le.go          |  174 +-
 .../golang.org/x/sys/unix/zsyscall_linux_mipsle.go |  186 +-
 .../golang.org/x/sys/unix/zsyscall_linux_ppc64.go  |  209 +-
 .../x/sys/unix/zsyscall_linux_ppc64le.go           |  209 +-
 .../x/sys/unix/zsyscall_linux_riscv64.go           |   53 +-
 .../golang.org/x/sys/unix/zsyscall_linux_s390x.go  |  189 +-
 .../x/sys/unix/zsyscall_linux_sparc64.go           |  204 +-
 .../golang.org/x/sys/unix/zsyscall_netbsd_386.go   |  389 +-
 .../golang.org/x/sys/unix/zsyscall_netbsd_amd64.go |  389 +-
 .../golang.org/x/sys/unix/zsyscall_netbsd_arm.go   |  389 +-
 .../golang.org/x/sys/unix/zsyscall_openbsd_386.go  |  186 +-
 .../x/sys/unix/zsyscall_openbsd_amd64.go           |  186 +-
 .../golang.org/x/sys/unix/zsyscall_openbsd_arm.go  |  186 +-
 .../x/sys/unix/zsyscall_solaris_amd64.go           |  256 +
 .../golang.org/x/sys/unix/zsysctl_openbsd_amd64.go |   13 +
 vendor/golang.org/x/sys/unix/zsysnum_linux_386.go  |    2 +-
 .../golang.org/x/sys/unix/zsysnum_linux_amd64.go   |    2 +-
 vendor/golang.org/x/sys/unix/zsysnum_linux_arm.go  |    3 +-
 .../golang.org/x/sys/unix/zsysnum_linux_arm64.go   |    3 +-
 vendor/golang.org/x/sys/unix/zsysnum_linux_mips.go |    2 +-
 .../golang.org/x/sys/unix/zsysnum_linux_mips64.go  |    2 +-
 .../x/sys/unix/zsysnum_linux_mips64le.go           |    2 +-
 .../golang.org/x/sys/unix/zsysnum_linux_mipsle.go  |    2 +-
 .../golang.org/x/sys/unix/zsysnum_linux_ppc64.go   |    4 +-
 .../golang.org/x/sys/unix/zsysnum_linux_ppc64le.go |    4 +-
 .../golang.org/x/sys/unix/zsysnum_linux_riscv64.go |    3 +-
 .../golang.org/x/sys/unix/zsysnum_linux_s390x.go   |    4 +-
 .../golang.org/x/sys/unix/zsysnum_openbsd_386.go   |   25 +-
 .../golang.org/x/sys/unix/zsysnum_openbsd_amd64.go |    1 +
 .../golang.org/x/sys/unix/zsysnum_openbsd_arm.go   |   13 +-
 vendor/golang.org/x/sys/unix/ztypes_darwin_386.go  |   10 +-
 .../golang.org/x/sys/unix/ztypes_darwin_amd64.go   |   10 +-
 vendor/golang.org/x/sys/unix/ztypes_darwin_arm.go  |   10 +-
 .../golang.org/x/sys/unix/ztypes_darwin_arm64.go   |   10 +-
 .../x/sys/unix/ztypes_dragonfly_amd64.go           |   27 +-
 vendor/golang.org/x/sys/unix/ztypes_freebsd_386.go |  276 +-
 .../golang.org/x/sys/unix/ztypes_freebsd_amd64.go  |  294 +-
 vendor/golang.org/x/sys/unix/ztypes_freebsd_arm.go |  298 +-
 vendor/golang.org/x/sys/unix/ztypes_linux_386.go   |  149 +-
 vendor/golang.org/x/sys/unix/ztypes_linux_amd64.go |  151 +-
 vendor/golang.org/x/sys/unix/ztypes_linux_arm.go   |  150 +-
 vendor/golang.org/x/sys/unix/ztypes_linux_arm64.go |  151 +-
 vendor/golang.org/x/sys/unix/ztypes_linux_mips.go  |  150 +-
 .../golang.org/x/sys/unix/ztypes_linux_mips64.go   |  151 +-
 .../golang.org/x/sys/unix/ztypes_linux_mips64le.go |  151 +-
 .../golang.org/x/sys/unix/ztypes_linux_mipsle.go   |  150 +-
 vendor/golang.org/x/sys/unix/ztypes_linux_ppc64.go |  151 +-
 .../golang.org/x/sys/unix/ztypes_linux_ppc64le.go  |  151 +-
 .../golang.org/x/sys/unix/ztypes_linux_riscv64.go  |   26 +-
 vendor/golang.org/x/sys/unix/ztypes_linux_s390x.go |  151 +-
 .../golang.org/x/sys/unix/ztypes_linux_sparc64.go  |   10 +-
 vendor/golang.org/x/sys/unix/ztypes_netbsd_386.go  |   27 +-
 .../golang.org/x/sys/unix/ztypes_netbsd_amd64.go   |   27 +-
 vendor/golang.org/x/sys/unix/ztypes_netbsd_arm.go  |   27 +-
 vendor/golang.org/x/sys/unix/ztypes_openbsd_386.go |  120 +-
 .../golang.org/x/sys/unix/ztypes_openbsd_amd64.go  |  120 +-
 vendor/golang.org/x/sys/unix/ztypes_openbsd_arm.go |  204 +-
 .../golang.org/x/sys/unix/ztypes_solaris_amd64.go  |   27 +-
 vendor/golang.org/x/sys/windows/syscall_windows.go |   60 +-
 vendor/golang.org/x/sys/windows/types_windows.go   |  115 +-
 vendor/golang.org/x/text/width/kind_string.go      |    2 +-
 .../x/text/width/{tables.go => tables10.0.0.go}    |  322 +-
 .../x/text/width/{tables.go => tables9.0.0.go}     |    4 +-
 vendor/golang.org/x/text/width/trieval.go          |    2 +-
 vendor/vbom.ml/util/LICENSE                        |   17 +
 vendor/vbom.ml/util/README.md                      |    5 +
 vendor/vbom.ml/util/sortorder/README.md            |    5 +
 vendor/vbom.ml/util/sortorder/doc.go               |    5 +
 vendor/vbom.ml/util/sortorder/natsort.go           |   76 +
 801 files changed, 71381 insertions(+), 18586 deletions(-)

```

</details>

## Allow configuring max download/upload attempts for pull/push

<details>
<summary>Commits</summary>

```
commit 7cab3339d1c1e4ad39b5baa49f4a38d5c1eb1ad5
Author:     Resin CI <34882892+balena-ci@users.noreply.github.com>
AuthorDate: Mon Apr 29 20:02:29 2019 +0300
Commit:     Resin CI <34882892+balena-ci@users.noreply.github.com>
CommitDate: Mon Apr 29 20:02:29 2019 +0300

    v18.9.5

 CHANGELOG.md | 5 +++++
 VERSION      | 2 +-
 2 files changed, 6 insertions(+), 1 deletion(-)

---

commit c88f6f82b644c6c0b0b7932d980a86542be411b1
Author:     Robert Günzler <r@gnzler.io>
AuthorDate: Wed Apr 24 14:15:56 2019 +0200
Commit:     Robert Günzler <r@gnzler.io>
CommitDate: Fri Apr 26 19:19:29 2019 +0200

    Add daemon flags to configure max download/upload attempts during pull/push
    
    The defaults remain the same (dl=5, ul=5), but are moved from distribution/xfer to
    daemon/config.
    
    Connects-to: https://github.com/balena-os/balena-engine/issues/160
    Change-type: patch
    Signed-off-by: Robert Günzler <robertg@balena.io>

 cmd/dockerd/config.go              |  5 +++++
 daemon/config/config.go            | 24 ++++++++++++++++++++++++
 daemon/daemon.go                   |  2 ++
 daemon/images/service.go           |  8 ++++++--
 daemon/reload.go                   | 31 +++++++++++++++++++++++++++++++
 distribution/xfer/download.go      |  8 ++++----
 distribution/xfer/download_test.go |  9 ++++++---
 distribution/xfer/upload.go        |  8 ++++----
 distribution/xfer/upload_test.go   |  9 ++++++---
 plugin/backend_linux.go            |  2 +-
 10 files changed, 89 insertions(+), 17 deletions(-)

---

```
</details>

## Include marketing website sources

<details>
<summary>Commits</summary>

```
commit 5c46120af9340bb71c47befa8e39a027e627a8e2
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Fri Oct 13 03:04:55 2017 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 20:00:40 2018 -0800

    landr: add correct feature descriptions
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 landr.conf.js | 4 ++--
 1 file changed, 2 insertions(+), 2 deletions(-)

---

commit 090ec746e0d59c4ee38fb1f4720494c727b11233
Author:     craig-mulligan <craig@resin.io>
AuthorDate: Fri Oct 13 09:40:05 2017 +0100
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:13:27 2018 -0800

    3 -> 3.5

 README.md     | 2 +-
 landr.conf.js | 2 +-
 2 files changed, 2 insertions(+), 2 deletions(-)

---

commit 3e2c3baab164afe9c82680b301d7ee9e7e25dbd9
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Thu Oct 12 21:10:13 2017 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:13:26 2018 -0800

    landr: use install.sh
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 landr.conf.js | 4 ++--
 1 file changed, 2 insertions(+), 2 deletions(-)

---

commit 392319c984f0d66b8c965d82f5a1d22fab80565e
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Thu Oct 12 21:00:08 2017 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:13:26 2018 -0800

    landr: fix architecture extraction from assets
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 landr.conf.js | 12 +++---------
 1 file changed, 3 insertions(+), 9 deletions(-)

---

commit f5922f9f421566d99d356c2a660b0f70fac995bb
Author:     craig-mulligan <craig@resin.io>
AuthorDate: Fri Oct 13 03:31:50 2017 +0100
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:13:26 2018 -0800

    add post-build hook

 landr.conf.js     |   7 +++
 package-lock.json | 143 ------------------------------------------------------
 2 files changed, 7 insertions(+), 143 deletions(-)

---

commit e13fca33295c06d3f94a6e69e702250dd9c83a89
Author:     craig-mulligan <craig@resin.io>
AuthorDate: Thu Oct 12 23:05:57 2017 +0100
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:13:25 2018 -0800

    add website data

 FAQ.md                                    |  24 +++++
 docs/getting-started.md                   |  28 ++++++
 landr.conf.js                             |  87 ++++++++++++++++++
 package-lock.json                         | 143 ++++++++++++++++++++++++++++++
 www/static/balena.svg                     |  65 ++++++++++++++
 www/static/favicon.ico                    | Bin 0 -> 15086 bytes
 www/static/features/bandwidth.svg         |  22 +++++
 www/static/features/failure-resistant.svg |  29 ++++++
 www/static/features/footprint.svg         |  28 ++++++
 www/static/features/multiple.svg          |  24 +++++
 www/static/features/storage.svg           |  25 ++++++
 www/static/features/undisturbed.svg       |  15 ++++
 12 files changed, 490 insertions(+)

---
```

</details>

## Rename to balena-engine

<details>
<summary>Commits</summary>

```
commit 44ace080a685a4c04107de18fcd430d2ada7949e
Author:     Robert Günzler <r@gnzler.io>
AuthorDate: Tue Jan 8 13:37:52 2019 +0100
Commit:     Robert Günzler <r@gnzler.io>
CommitDate: Tue Jun 25 14:31:02 2019 +0200

    contrib/install.sh: Rename balena to balenaEngine in ASCII art output

    Change-type: patch
    Signed-off-by: Robert Günzler <robertg@balena.io>

 contrib/install.sh | 14 +++++++-------
 1 file changed, 7 insertions(+), 7 deletions(-)

---

commit d254fc040d5cec0b169078378241b0f6383797c5
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Thu Oct 12 15:08:02 2017 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:13:25 2018 -0800

    readme: describe what balena is
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 README.md                               | 76 ++++++++++++++-------------------
 docs/static_files/balena-logo-black.svg | 56 ++++++++++++++++++++++++
 2 files changed, 87 insertions(+), 45 deletions(-)

---

commit 385bc86e9096a6aa62575e1b6f612ee842a7492b
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Thu Oct 5 22:34:13 2017 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:13:25 2018 -0800

    rename container engine to balena
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 cli/config/configdir.go                              |  2 +-
 client/client_unix.go                                |  2 +-
 client/errors.go                                     |  4 ++--
 cmd/{rce/rce.go => balena/main.go}                   | 14 +++++++-------
 cmd/dockerd/config.go                                | 10 +++++-----
 cmd/dockerd/config_common_unix.go                    |  8 ++++----
 cmd/dockerd/config_unix.go                           |  4 ++--
 cmd/dockerd/daemon_unix.go                           |  4 ++--
 cmd/dockerd/docker.go                                |  4 ++--
 cmd/mobynit/main.go                                  |  2 +-
 container/container.go                               |  2 +-
 daemon/config/config.go                              |  2 +-
 daemon/config/config_common_unix_test.go             |  2 +-
 daemon/daemon_linux.go                               |  2 +-
 daemon/daemon_unix.go                                |  2 +-
 daemon/graphdriver/devmapper/device_setup.go         | 14 +++++++-------
 daemon/graphdriver/devmapper/deviceset.go            | 10 +++++-----
 daemon/listeners/group_unix.go                       |  2 +-
 daemon/metrics.go                                    |  2 +-
 hack/make.sh                                         |  4 ++--
 hack/make/.binary-setup                              |  2 +-
 hack/make/binary-balena                              | 15 +++++++++++++++
 hack/make/binary-rce-docker                          | 15 ---------------
 hack/make/{dynbinary-rce-docker => dynbinary-balena} |  8 ++++----
 opts/hosts.go                                        |  2 +-
 registry/config_unix.go                              |  2 +-
 26 files changed, 70 insertions(+), 70 deletions(-)

---
commit d22c96a33573d42551838c75a9472ce3e6b5ea7d
Author:     Yossi Eliaz <yossi@resin.io>
AuthorDate: Tue Aug 22 00:13:05 2017 -0500
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:13:21 2018 -0800

    hack/make/cross: Disabled the Windows target
    
    Signed-off-by: Yossi Eliaz <yossi@resin.io>

 hack/make/cross | 2 +-
 1 file changed, 1 insertion(+), 1 deletion(-)

---

commit 429ac7e5bb07f3adacb5ea61db0a7fff918e7dcc
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Fri Aug 4 18:00:35 2017 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:06:54 2018 -0800

    hack: revert binary stripping by default
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 hack/make/binary-rce-docker    | 2 +-
 hack/make/dynbinary-rce-docker | 1 -
 2 files changed, 1 insertion(+), 2 deletions(-)

---

commit bc9b3093c8707afef3a183e3be18f709404ec486
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Fri Aug 4 17:13:07 2017 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:06:54 2018 -0800

    hack: create both dynamic and static flavours of rce
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 hack/make/binary-rce-docker                           | 2 --
 hack/make/{binary-rce-docker => dynbinary-rce-docker} | 1 +
 2 files changed, 1 insertion(+), 2 deletions(-)

---

commit 55dcad3cacb620c605c4752b53b0e5156e988c81
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Thu Jul 27 22:31:37 2017 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:06:51 2018 -0800

    hack: create all appropriate symlinks after building rce
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 hack/make/binary-rce-docker | 2 +-
 1 file changed, 1 insertion(+), 1 deletion(-)

---

commit ae5453e659a263f830615ff738651212980ff0f4
Author:     Yossi Eliaz <yossi@resin.io>
AuthorDate: Fri Jun 2 08:19:56 2017 -0500
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:06:50 2018 -0800

    Renamed the target and binary name
    
    In order to avoid conflict with previous code.
    We renamed the rce into rce-docker.

 Makefile                                    | 4 ++--
 hack/make.sh                                | 2 +-
 hack/make/.binary-setup                     | 2 +-
 hack/make/{binary-rce => binary-rce-docker} | 0
 4 files changed, 4 insertions(+), 4 deletions(-)

---

commit c46694c87c2665c64885858b08bc1255e71f3bab
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Tue Jul 25 23:51:23 2017 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:06:50 2018 -0800

    rce: create stripped binary
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 hack/make/binary-rce | 1 +
 1 file changed, 1 insertion(+)

---

commit 7cd827f8eb2923f079cb6dd2c429ba3ad1b9dfec
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Tue Jul 25 23:06:19 2017 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:06:50 2018 -0800

    hack: allow variables to be set by the environment
    
    We need this to propage the values from hack/make/binary-rce
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 hack/make/.binary | 4 ++--
 1 file changed, 2 insertions(+), 2 deletions(-)

---
```

</details>

## Delta layers feature

<details>
<summary>Commits</summary>

```
commit db7076c18a833998c0413900013854629b458028
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Thu Oct 12 15:51:45 2017 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:13:25 2018 -0800

    add changelog for 17.06+rev1
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 CHANGELOG.md | 3577 +---------------------------------------------------------
 1 file changed, 15 insertions(+), 3562 deletions(-)

---

commit 49ddb26883b698cce2242d2b0e71febbc69079e9
Author:     Akis Kesoglou <akiskesoglou@gmail.com>
AuthorDate: Thu Oct 12 15:08:51 2017 +0300
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:13:25 2018 -0800

    Fix typo in identifier name

 daemon/create.go | 8 ++++----
 1 file changed, 4 insertions(+), 4 deletions(-)

---

commit 3a8b3ac46396fb9af64bac83474aa435776af652
Author:     Akis Kesoglou <akiskesoglou@gmail.com>
AuthorDate: Thu Oct 12 15:08:25 2017 +0300
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:13:24 2018 -0800

    Fix invalid API URL reference in swagger.yml

 api/swagger.yaml | 2 +-
 1 file changed, 1 insertion(+), 1 deletion(-)

---

commit a33a1d80f721bf03b74ff09d91c1075ffb3548c9
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Wed Oct 11 18:21:56 2017 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:13:24 2018 -0800

    daemon: compute and print summary of delta efficiency
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 daemon/create.go | 21 ++++++++++++++++++++-
 1 file changed, 20 insertions(+), 1 deletion(-)

---

commit 6a0936bba24099ee20f7f1e084bb82fa4c4018df
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Fri Oct 6 16:38:05 2017 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:13:24 2018 -0800

    delta: revert automatic tag generation for delta image
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 daemon/create.go | 11 -----------
 1 file changed, 11 deletions(-)

---

commit 697bf0500a7cdecdcd30683a08af220ac04e060a
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Fri Oct 6 14:10:06 2017 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:13:24 2018 -0800

    api: move delta creation under POST images/delta
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 api/server/router/delta/backend.go      | 11 -----------
 api/server/router/delta/delta.go        | 30 ------------------------------
 api/server/router/delta/delta_routes.go | 32 --------------------------------
 api/server/router/image/backend.go      |  1 +
 api/server/router/image/image.go        |  1 +
 api/server/router/image/image_routes.go | 22 ++++++++++++++++++++++
 api/swagger.yaml                        |  2 +-
 client/image_delta.go                   | 22 ++++++++++++++++++++++
 cmd/dockerd/daemon.go                   |  2 --
 9 files changed, 47 insertions(+), 76 deletions(-)

---

commit b67a57f6b44a4ea25ca6233a8ae482c1dac235bd
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Fri Oct 6 13:36:57 2017 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:13:24 2018 -0800

    image/delta: client support with progress reporting
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 api/server/router/delta/backend.go      |  6 ++-
 api/server/router/delta/delta_routes.go | 21 +++++++----
 client/interface.go                     |  1 +
 daemon/create.go                        | 66 +++++++++++++++++++++++++--------
 4 files changed, 70 insertions(+), 24 deletions(-)

---

commit 43f44b0b1e793cd8c42776c38a226426ffca2891
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Thu Sep 21 17:51:18 2017 +0300
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:13:22 2018 -0800

    delta: remove false assumption about src image
    
    Previously the code assumed the source image had at least as many layers
    as the destination image.
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 daemon/create.go | 2 +-
 1 file changed, 1 insertion(+), 1 deletion(-)

---

commit e0537fc5b06b53668fc321f5b32d6d77d6577270
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Wed Jul 19 18:38:07 2017 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:13:20 2018 -0800

    daemon: allow a secondary daemon store to be loaded
    
    This commit allows the daemon to load a secondary daemon store to be
    used as source for deltas.
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 cmd/dockerd/config.go   |  3 +++
 daemon/config/config.go |  3 +++
 daemon/daemon.go        | 37 +++++++++++++++++++++++++++++++++++++
 daemon/image_pull.go    |  8 +++++++-
 daemon/image_push.go    |  2 +-
 distribution/config.go  | 10 ++++++++--
 6 files changed, 59 insertions(+), 4 deletions(-)

---

commit 2d5b662827ffcc2c4dc700b79d3d0908246b2426
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Wed Jul 12 22:55:20 2017 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:13:18 2018 -0800

    distribution: implement delta pull
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 distribution/pull_v1.go       |  4 ++++
 distribution/pull_v2.go       | 43 +++++++++++++++++++++++++++++++++++++++++++
 distribution/xfer/download.go | 35 +++++++++++++++++++++++++++++++++--
 3 files changed, 80 insertions(+), 2 deletions(-)

---

commit 0713bcd1cb014d6fde2ecd00c0f169f8c18972b4
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Wed Jul 12 20:19:32 2017 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:12:43 2018 -0800

    distribution: serialise pulling the config and layers
    
    Pull the config before starting anything else. We need the config before
    starting the layer download to prepare for a delta pull if this config
    happens to be a delta config.
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 distribution/pull_v2.go | 103 +++++++++++-------------------------------------
 1 file changed, 23 insertions(+), 80 deletions(-)

---

commit 0b6d9b43a9cbd62845787ec927d273d72dafa1ef
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Mon Jul 10 22:45:52 2017 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:08:00 2018 -0800

    deltas: implement image diff based on librsync
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 daemon/create.go | 164 ++++++++++++++++++++++++++++++++++++++++++++++++++++++-
 1 file changed, 163 insertions(+), 1 deletion(-)

---

commit b3816803a038a44c411a95cba11db166307c5dc3
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Sat Jul 8 20:42:10 2017 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:06:55 2018 -0800

    api: skeleton for POST /deltas/create method
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 api/server/router/delta/backend.go      |  7 +++++++
 api/server/router/delta/delta.go        | 30 ++++++++++++++++++++++++++++++
 api/server/router/delta/delta_routes.go | 27 +++++++++++++++++++++++++++
 api/swagger.yaml                        | 31 +++++++++++++++++++++++++++++++
 cmd/dockerd/daemon.go                   |  2 ++
 daemon/create.go                        |  6 ++++++
 6 files changed, 103 insertions(+)

---

commit 826321faf8a1eb42385362a017d68bbc807f1a34
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Thu Jul 6 20:21:37 2017 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:06:55 2018 -0800

    make GetTarSeekStream part of the ImageConfigStore interface
    
    This will allow distribution code to request a seekable tar stream in
    order to compute a binary delta on top of it.
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 distribution/config.go  |  6 ++++++
 plugin/backend_linux.go | 10 ++++++++++
 plugin/blobstore.go     |  6 ++++++
 3 files changed, 22 insertions(+)

---

commit 4d8f686b16efdf239e53ae28b49139002b8ad7ac
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Thu Jul 6 20:07:50 2017 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:06:55 2018 -0800

    image: implement seekable tar stream for a whole image
    
    This method creates a seekable stream that is the concatenation of the
    tar seekable streams of the layers the image is composed of. This method
    is intended to be used as a basis for delta based updates where bits of
    the previous image can be reused to reconstruct a layer of a future
    image.
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 image/store.go | 43 +++++++++++++++++++++++++++++++++++++++++++
 1 file changed, 43 insertions(+)

---

commit 98d01b27a5c82ae1afc003185b3dc3457e803ec0
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Thu Jul 6 18:54:08 2017 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:06:55 2018 -0800

    layer: implement seekable tar stream
    
    The TarSeekStream is added to the Layer interface to allow any user of
    layers to request a ReadSeeker of the tar archive that would have been
    produced in a normal TarStream().
    
    This allows reading parts of the resulting archive on the fly, without
    having to buffer the file on disk and then seek on top of it.
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 layer/empty.go       |  9 +++++++++
 layer/layer.go       |  5 +++++
 layer/layer_store.go | 36 ++++++++++++++++++++++++++++++++++++
 layer/ro_layer.go    |  7 +++++++
 4 files changed, 57 insertions(+)

---

commit c8643e4d0d3c5dd1c4b8c9e03044857aea8d7478
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Thu Jul 27 20:00:04 2017 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:06:55 2018 -0800

    pkg/tarsplitutils: optimise reads using binary search
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 pkg/tarsplitutils/tarstream.go | 51 +++++++++++++++++-------------------------
 1 file changed, 20 insertions(+), 31 deletions(-)

---

commit c2b1d1bbe130d292553d31f1dbd070d683987412
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Thu Jul 6 18:40:10 2017 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:06:55 2018 -0800

    pkg/tarsplitutils: implement random access seeker stream
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 pkg/tarsplitutils/tarstream.go | 138 +++++++++++++++++++++++++++++++++++++++++
 1 file changed, 138 insertions(+)

---

commit c6133b90f337ea8a7c4ebbb73ee62df779bf1c9d
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Thu Jul 13 16:30:08 2017 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:06:55 2018 -0800

    ioutils: implement ReadSeekCloser concat stream
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 pkg/ioutils/concat.go | 123 ++++++++++++++++++++++++++++++++++++++++++++++++++
 1 file changed, 123 insertions(+)

---

commit 34769d6eb18be01b51225ad44e7410178b4e3520
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Thu Jul 6 18:04:28 2017 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:06:54 2018 -0800

    ioutils: implement ReadSeekCloser interface and wrapper
    
    This is a useful abstraction that exists in a few places in docker's
    codebase
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 pkg/ioutils/readers.go | 23 +++++++++++++++++++++++
 1 file changed, 23 insertions(+)

---

commit caa9b94f726aa25812044a8cf58c52648d3d5c89
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Thu Aug 17 07:25:00 2017 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:06:53 2018 -0800

    distribution: check for nil before closing the download

 distribution/pull_v2.go | 4 +++-
 1 file changed, 3 insertions(+), 1 deletion(-)

---
```

</details>

## Report total progress on pulling

<details>
<summary>Commits</summary>

```
commit 08de939889c3046ff734f948e6f285f5554b523d
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Fri Oct 6 13:30:02 2017 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:13:24 2018 -0800

    pkg/ioutils: export SeekerSize utility

    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 pkg/ioutils/concat.go | 6 +++---
 1 file changed, 3 insertions(+), 3 deletions(-)

---

commit 374fa017bcd6eb5ee60db134afb62e96fd117ea9
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Tue Oct 3 16:13:43 2017 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:13:24 2018 -0800

    distribution: calculate combined progress when pulling

    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 distribution/pull_v1.go       |  4 ++++
 distribution/pull_v2.go       |  4 ++++
 distribution/xfer/download.go | 12 ++++++++----
 3 files changed, 16 insertions(+), 4 deletions(-)

---

commit 128b42eb29f8f7378dc011494b1b09554e4b5f53
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Tue Oct 3 16:11:31 2017 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:13:23 2018 -0800

    pkg/progress: add progress sink

    This Sink can have multiple writers writing to it so that it calculates
    a progress for all of them. For example, if you are downloading 5 files,
    1, 2, 3, 4, 5 MB each, you can create a ProgressSink with its size set
    to 15MB and use io.Tee() with each of those files' Readers to calculate
    a combined progress.

    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 pkg/progress/progresssink.go | 53 ++++++++++++++++++++++++++++++++++++++++++++
 1 file changed, 53 insertions(+)

```
</details>

## Prevent pagecache thrashing during unpacking (on pull)

<details>
<summary>Commits</summary>

```
commit fcf3865fddb99e7a45dd7a1004f644b2e5d0a93b
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Sun Oct 8 15:26:23 2017 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:13:27 2018 -0800

    pkg/archive: sync files before issuing the fadvise syscall
    
    Linux ignores fadvise for dirty pages so we need to make sure we sync to
    have them marked as clean before discarding
    
    Ref: https://github.com/torvalds/linux/blob/v4.13/mm/truncate.c#L489-L493
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 pkg/archive/archive.go | 24 +++++++++++-------------
 1 file changed, 11 insertions(+), 13 deletions(-)

---

commit 8c0ceaba294daa7c0b1c1a7c722c77070e8b8786
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Thu Aug 3 19:55:56 2017 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:06:52 2018 -0800

    pkg/archive: use fadvise to prevent pagecache thrashing
    
    During a docker pull a very large amount of files are touched during
    unpacking. This causes linux to fill up the page cache with entries that
    most probably won't be used.
    
    There are two issues with this. The first one is that by putting a lot
    of pressure on the page cache memory fragmentation occurs. This can
    cause filesystem corruption on some platforms that can't handle memory
    allocation failures correctly.
    
    The second issue is that by not hinting our intentions to the kernel, we
    might evict useful pages from the cache that could be in use by some
    running container, and therefore affecting the performance of the
    application.
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 pkg/archive/archive.go | 15 +++++++++++++++
 1 file changed, 15 insertions(+)

---
```

</details>


## Fix upstream bugs

### invalid container config

https://github.com/moby/moby/issues/33018

<details>
<summary>Commits</summary>

```
commit 10ae0733da29c67488126159370e87000bc018b0
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Sat Jul 1 18:39:13 2017 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:06:52 2018 -0800

    container: make sure config on disk has a valid Config
    
    We've seen cases where container config on disk is a valid JSON but
    misses the Config object. Due to the way docker loads it in memory, it
    can cause a nul pointer exception.
    
    This patch checks for that and doesn't load the container in this case
    
    Upstream-Status: Investigating https://github.com/moby/moby/issues/33018
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 container/container.go | 4 ++++
 1 file changed, 4 insertions(+)

---
```

</details>

## Make pulls more resilient against power cuts

<details>
<summary>Commits</summary>

```
commit 175199490819176eec7c78cdeb95466828e3eb28
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Thu Nov 3 17:25:12 2016 +0000
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:06:52 2018 -0800

    pkg/ioutils: sync parent directory too
    
    After renaming the file to the target path the parent directory needs to
    be synced to make sure the rename hits the disk.
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 pkg/ioutils/fswriters.go | 23 +++++++++++++++++++----
 1 file changed, 19 insertions(+), 4 deletions(-)

---

commit 684d8ba6109c853b355bf11ca3733c4099f14b92
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Thu Nov 3 00:38:14 2016 +0000
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:06:52 2018 -0800

    aufs,overlay: durably write layer on disk before returning
    
    This patch makes sure the layer contents are synced to disk before
    reporting the ApplyDiff operation as successful. This prevents
    /var/lib/docker corruption but the method used here is not the most
    efficient since it will sync all the currently mounted filesystems.
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 daemon/graphdriver/aufs/aufs.go        | 5 +++++
 daemon/graphdriver/overlay2/overlay.go | 5 +++++
 2 files changed, 10 insertions(+)

---
```

</details>

### Allow turning off resilience measures

Improve pull performance at the cost of increased risk of filesystem corruption

`--storage.opt={overlay2,aufs}.sync_diffs=false`

<details>
<summary>Commits</summary>

```
commit cf02022c246852e4503e5f63f30cea96350feaff
Author:     Resin CI <34882892+resin-ci@users.noreply.github.com>
AuthorDate: Fri Apr 26 14:17:29 2019 +0300
Commit:     Giovanni Garufi <nazrhom@gmail.com>
CommitDate: Fri Apr 26 14:16:30 2019 +0200

    v18.9.4

 CHANGELOG.md | 11 ++++++++++-
 VERSION      |  2 +-
 2 files changed, 11 insertions(+), 2 deletions(-)

---

commit 5e472e9eb556ac6aeb4e0346551d12b652c62a7f
Author:     Robert Günzler <r@gnzler.io>
AuthorDate: Thu Apr 25 18:42:02 2019 +0200
Commit:     Giovanni Garufi <nazrhom@gmail.com>
CommitDate: Fri Apr 26 14:13:59 2019 +0200

    integration-tests: Skip aufs test, doesn't work with dind

    Signed-off-by: Robert Günzler <robertg@balena.io>

 integration/image/pull_test.go | 1 +
 1 file changed, 1 insertion(+)

---

commit 4c3cafd8c56102bcf83773dcb50a597e944d382a
Author:     Robert Günzler <r@gnzler.io>
AuthorDate: Fri Apr 19 16:37:07 2019 +0200
Commit:     Giovanni Garufi <nazrhom@gmail.com>
CommitDate: Fri Apr 26 14:13:59 2019 +0200

    integration-tests: Add image pull tests

    Download an image on aufs/overlay2 once with syncDiffs enabled and
    disabled, comparing the speeds and checking that without syncing is
    faster.

    Signed-off-by: Robert Günzler <robertg@balena.io>

 integration/image/pull_test.go | 56 ++++++++++++++++++++++++++++++++++++++++++
 1 file changed, 56 insertions(+)

---

commit 0ef2d9dc7c6456b97e3b019f5110b6a67603c76c
Author:     Robert Günzler <r@gnzler.io>
AuthorDate: Fri Apr 19 15:38:34 2019 +0200
Commit:     Giovanni Garufi <nazrhom@gmail.com>
CommitDate: Fri Apr 26 14:13:59 2019 +0200

    aufs,overlay2: Add driver opts for disk sync

    This patch adds a driver option to enalble/disable the to disk syncing introduced in
    684d8ba6109c853b355bf11ca3733c4099f14b92.

    The default is still to sync all currently mounted filesystems before
    reporting an ApplyDiff as successful.

    Connects-to: https://github.com/balena-os/balena-engine/issues/133
    Change-type: patch
    Signed-off-by: Robert Günzler <robertg@balena.io>

 daemon/graphdriver/aufs/aufs.go        | 39 +++++++++++++++++++++++++++++++---
 daemon/graphdriver/overlay2/overlay.go | 18 +++++++++++-----
 2 files changed, 49 insertions(+), 8 deletions(-)
```

</details>

## Booting into a container filesystem - `mobynit`

<details>
<summary>Commits</summary>

```
commit 0c56d80d4982d41bd05e64aa05bbefebeae82eba
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Tue Oct 3 19:49:51 2017 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:13:23 2018 -0800

    cmd/mobynit: fix storage driver path
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 cmd/mobynit/main.go | 2 +-
 1 file changed, 1 insertion(+), 1 deletion(-)

---

commit 19da842517cd7fe6fe1b21297cf4f14d2307b4de
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Tue Oct 3 14:53:08 2017 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:13:23 2018 -0800

    cmd/mobynit: get graphdriver parameter from disk
    
    Since this is PID 1 it's easier to parameterise it through a file on the
    disk rather than having to rely on the bootloader to pass the right
    arguments
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 cmd/mobynit/main.go | 24 +++++++++++-------------
 1 file changed, 11 insertions(+), 13 deletions(-)

---

commit 2496793a9d26e0a1015d27a3a4b51227c965f9ce
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Mon Aug 7 17:06:18 2017 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:13:20 2018 -0800

    cmd/mobynit: permanently remount root as rw
    
    We need the root to be read/write in order to run any docker daemon on
    top of it. The pivot root remains read-only however.
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 cmd/mobynit/main.go | 9 ++++-----
 1 file changed, 4 insertions(+), 5 deletions(-)

---

commit 4192b6d153d88d47dabfee13026cc7863e906ac2
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Sun Aug 6 19:15:41 2017 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:13:20 2018 -0800

    cmd/mobynit: switch the pivot path to /mnt/sysroot/active
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 cmd/mobynit/main.go | 2 +-
 1 file changed, 1 insertion(+), 1 deletion(-)

---

commit 0dcd39eeba55bca5f112fbe4e92f02a27523bb61
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Thu Aug 3 18:25:10 2017 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:06:52 2018 -0800

    cmd/mobynit: propagate initrd mounts to chroot
    
    If there was an initrd before mobynit that initialised the filesystem
    transfer those mounts to the new root.
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 cmd/mobynit/main.go | 13 +++++++++++++
 1 file changed, 13 insertions(+)

---

commit 471e34a0182ad8e93763d831500a0606fd68dd71
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Thu Aug 3 18:24:21 2017 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:06:52 2018 -0800

    mobynit: read the containerID from /current symlink
    
    Make /current be a symlink so that it can be atomically switched from
    one host app to another.
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 cmd/mobynit/main.go | 6 ++----
 1 file changed, 2 insertions(+), 4 deletions(-)

---

commit fb65dbf59a338352b4d7e010b5e91ef697f8ee1e
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Thu Aug 3 00:58:22 2017 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:06:52 2018 -0800

    daemon: revert short circuit of volume setup for bare containers
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 daemon/create_unix.go | 3 ---
 1 file changed, 3 deletions(-)

---

commit 9d7480d8b6a0aaf6f35eb7de2620e23e00dbbba8
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Tue Aug 1 17:15:01 2017 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:06:51 2018 -0800

    cmd/mobynit: accept a flag for the graph driver
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 cmd/mobynit/main.go | 12 +++++++++++-
 1 file changed, 11 insertions(+), 1 deletion(-)

---

commit 7191674c2c95069ef8918a063d4a6f33b283543f
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Thu Jul 27 13:38:23 2017 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:06:51 2018 -0800

    daemon: skip initLayer for bare runtime containers
    
    Containers that are meant to be booted from do not need the initLayer.
    
    The default init layer shadows /etc/resolv.conf and other files from the
    filesystem but this can cause problems if we're bootstrapping the
    system.
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 daemon/create.go      | 11 ++++++++---
 daemon/create_unix.go |  3 +++
 daemon/daemon_unix.go |  1 +
 3 files changed, 12 insertions(+), 3 deletions(-)

---

commit 60649c7a18a69dd39ba32216fe6c2a6b9f5f4f24
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Tue Jul 25 13:13:30 2017 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:06:51 2018 -0800

    cmd: add mobynit for host app booting
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 cmd/mobynit/main.go | 91 +++++++++++++++++++++++++++++++++++++++++++++++++++++
 1 file changed, 91 insertions(+)

---
```

</details>

## Stream data directly into layer store on pull

<details>
<summary>Commits</summary>

```
commit ad509615ba39b00d14c00f667f7444dd2619d961
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Thu Jul 27 21:24:50 2017 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:06:51 2018 -0800

    distribution: resume streaming download in case of failure
    
    Attempt to resume a download if it stops midway for up to
    maxDownloadAttempts times.
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 distribution/pull_v2.go | 49 +++++++++++++++++++++++++++++++++++++++++--------
 1 file changed, 41 insertions(+), 8 deletions(-)

---

commit bf0b86825bc72b60ef6d74091f08089779864b7f
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Thu Jul 27 21:13:19 2017 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:06:51 2018 -0800

    ioutils: implement TeeReadCloser
    
    Taken from https://github.com/Azure/go-autorest/blob/v8.1.1/autorest/utility.go#L52-L70
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 pkg/ioutils/readers.go | 20 ++++++++++++++++++++
 1 file changed, 20 insertions(+)

---

commit fd1683789adbfb83e1e21c0b6191b95092efd5d2
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Mon Jul 24 15:36:46 2017 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:06:51 2018 -0800

    distribution: stream download directly to the layer store
    
    start extracting as soon as the data starts coming from the network.
    This reduces the free space required for a docker pull to happen since
    there are no temp files.
    
    Digest verification happens when the last byte of the layer is read.
    If verification fails we return an error that will cascade to the layer
    store and cancel the pull.
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 distribution/pull_v2.go | 171 ++++++------------------------------------------
 1 file changed, 19 insertions(+), 152 deletions(-)

---
```

</details>


## Remove "cloud-focused" features

* service discovery backends
    - consul
    - etcd
    - zookeeper
* logdrivers
    - awslogs 
    - etwlogs 
    - fluentd 
    - gcplogs 
    - gelf 
    - logentries
    - splunk 
    - syslog 
    - local 
* docker-swarm related features

<details>
<summary>Commits</summary>

```
commit bce9bc78ee79cc80bfa47857dd756714dd8bdd54
Author:     Yossi Eliaz <yossi@resin.io>
AuthorDate: Tue Oct 10 21:15:37 2017 -0500
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 20:00:40 2018 -0800

    Fixed the runc version test
    
    Fixed TestInfoAPIRuncCommit and verified
    version is consistent
    
    Signed-off-by: Yossi Eliaz <yossi@resin.io>

 hack/make/.integration-daemon-start     | 4 ++++
 hack/make/.integration-test-helpers     | 1 +
 integration-cli/docker_api_info_test.go | 8 ++++++++
 3 files changed, 13 insertions(+)

---

commit add016d96df3477e3f6e5b86b96e60e1c5419099
Author:     Yossi Eliaz <yossi@resin.io>
AuthorDate: Tue Oct 10 21:26:17 2017 -0500
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 20:00:40 2018 -0800

    skip tests of unsopported components
    
    skip TestRunStoppedLoggingDriverNoLeak
    add a safe conversion of numerical values
    skip not supported log-drivers
    skip swarm related tests
    
    Signed-off-by: Yossi Eliaz <yossi@resin.io>

 integration-cli/docker_cli_daemon_test.go     | 14 +++++++++-----
 integration-cli/docker_cli_info_test.go       |  4 ++++
 integration-cli/docker_cli_prune_unix_test.go |  2 ++
 integration-cli/docker_cli_run_test.go        |  2 ++
 4 files changed, 17 insertions(+), 5 deletions(-)

---

commit b1d10730870cfe7d767e17f05c43917156263d93
Author:     Yossi Eliaz <yossi@resin.io>
AuthorDate: Fri Sep 8 10:05:55 2017 -0500
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:13:22 2018 -0800

    integration: Skip plugin tests
    
    rce-docker disabled plugin support, therefore we
    removed the tests.
    
    Signed-off-by: Yossi Eliaz <yossi@resin.io>

 integration-cli/docker_cli_daemon_plugins_test.go  | 22 ++++++++++++++++++++++
 integration-cli/docker_cli_events_test.go          |  2 ++
 integration-cli/docker_cli_inspect_test.go         |  2 ++
 integration-cli/docker_cli_network_unix_test.go    |  2 ++
 .../docker_cli_plugins_logdriver_test.go           |  4 ++++
 integration-cli/docker_cli_plugins_test.go         | 19 +++++++++++++++++++
 6 files changed, 51 insertions(+)

---

commit 63199860001969b8e552555f752a490d15a58605
Author:     Yossi Eliaz <yossi@resin.io>
AuthorDate: Thu Aug 24 19:38:56 2017 -0500
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:13:22 2018 -0800

    integration-cli: Disabled swarm tests
    
    Signed-off-by: Yossi Eliaz <yossi@resin.io>

 integration-cli/check_test.go | 12 ------------
 1 file changed, 12 deletions(-)

---

commit ea20d07becb11e6949c4fc001f0e9b9819fd9140
Author:     Yossi Eliaz <yossi@resin.io>
AuthorDate: Thu Aug 24 19:37:41 2017 -0500
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:13:22 2018 -0800

    Makefile: Revert cli-integration on swarm disable
    
    Re-enabled the cli integartion tests on swarm
    
    Signed-off-by: Yossi Eliaz <yossi@resin.io>

 Makefile | 6 ++++++
 1 file changed, 6 insertions(+)

---

commit 860bb33b77acd6278fdb44255b417784286159f1
Author:     Yossi Eliaz <yossi@resin.io>
AuthorDate: Wed Aug 23 18:30:15 2017 -0500
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:13:22 2018 -0800

    hack/make Modified docker-py to pass integration w/o swarm
    
    Updated test-docker-py to use zozo123/docker-py.git
    in order to skip swarm related tests.
    This also put a FIXME skip in tests that
    need to adjusted further.
    
    Signed-off-by: Yossi Eliaz <yossi@resin.io>

 hack/make/test-docker-py | 7 ++-----
 1 file changed, 2 insertions(+), 5 deletions(-)

---

commit e58efd2e9593d086af59ecb163cbb78f2e8d1a32
Author:     Yossi Eliaz <yossi@resin.io>
AuthorDate: Tue Aug 22 10:42:52 2017 -0500
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:13:21 2018 -0800

    hack/make: Adjusted integration tests for rce-docker
    
    - Switched the dynbinary to the rce-docker one
    - Removed swarm integration tests from the Makefiles.
    
    Signed-off-by: Yossi Eliaz <yossi@resin.io>

 Makefile                            | 18 ++++++------------
 hack/make.sh                        |  1 +
 hack/make/.integration-daemon-start |  2 +-
 3 files changed, 8 insertions(+), 13 deletions(-)

---

commit ea1aaacf90a90f61342643075483797d8f6b32ea
Author:     Yossi Eliaz <yossi@resin.io>
AuthorDate: Tue Aug 22 01:02:21 2017 -0500
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:13:21 2018 -0800

    daemon: Adjusted unix_test for the container struct
    
    Signed-off-by: Yossi Eliaz <yossi@resin.io>

 daemon/daemon_unix_test.go | 1 +
 1 file changed, 1 insertion(+)

---

commit 6b4eb9670fa8c867ff089fd39eac3bc4b06b6975
Author:     Yossi Eliaz <yossi@resin.io>
AuthorDate: Tue Aug 22 00:47:38 2017 -0500
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:13:21 2018 -0800

    hack/make: Removed cluster unit tests
    
    Removed cluster unit tests from the receipt
    
    Signed-off-by: Yossi Eliaz <yossi@resin.io>

 hack/make.ps1 | 1 +
 1 file changed, 1 insertion(+)

---

commit aab108d7020b5af19c5a3fefa9e12bf505165755
Author:     Yossi Eliaz <yossi@resin.io>
AuthorDate: Tue Aug 22 00:26:01 2017 -0500
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:13:21 2018 -0800

    migratev1_test: Fixed compilation errors
    
    Signed-off-by: Yossi Eliaz <yossi@resin.io>

 migrate/v1/migratev1_test.go | 6 ++++++
 1 file changed, 6 insertions(+)

---

commit ca4c0ad02675de3385ec19a48036b08e9c42837e
Author:     Yossi Eliaz <yossi@resin.io>
AuthorDate: Tue Aug 22 00:16:23 2017 -0500
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:13:21 2018 -0800

    xfer/download_test: Fixed compilation errors
    
    Signed-off-by: Yossi Eliaz <yossi@resin.io>

 distribution/xfer/download_test.go | 10 ++++++++++
 1 file changed, 10 insertions(+)

---

commit 3b2046fdc9e2d40ada9027b847e9e7898b72f503
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Thu Aug 17 19:58:17 2017 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:06:53 2018 -0800

    cmd/dockerd: remove support for docker plugins
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 cmd/dockerd/daemon.go | 2 --
 1 file changed, 2 deletions(-)

---

commit 0b51dbc886d0f534c4920dfb0ac418ff52e18317
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Thu Aug 17 19:48:18 2017 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:06:53 2018 -0800

    daemon/events: remove swarm related events
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 daemon/events.go | 200 -------------------------------------------------------
 1 file changed, 200 deletions(-)

---

commit 82b451966319a5fe9a697dff5a28551eadb8b1e8
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Thu Aug 17 19:45:39 2017 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:06:53 2018 -0800

    container_operations: remove swarm functionalities
    
    Removes support for docker Secrets, and docker Configs which are only
    usable within swarm mode
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 container/container.go              |   3 +-
 daemon/container_operations_unix.go | 101 +-----------------------------------
 daemon/dependency.go                |  17 ------
 3 files changed, 3 insertions(+), 118 deletions(-)

---

commit 77818f12ce92072b69f3097afaa63dbe8129eb03
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Thu Aug 17 19:40:20 2017 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:06:53 2018 -0800

    cmd/dockerd: drop support for swarm and checkpoint commands
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 cmd/dockerd/daemon.go | 46 ----------------------------------------------
 1 file changed, 46 deletions(-)

---

commit 98e2ea0b9506ac6211b25d765ae8ea851280c38c
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Thu Aug 17 19:36:23 2017 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:06:53 2018 -0800

    router/system: remove swarm dependency
    
    Makes the system commands only deal with the local daemon, ignoring
    swarm clusters.
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 api/server/router/system/system.go        | 5 +----
 api/server/router/system/system_routes.go | 3 ---
 cmd/dockerd/daemon.go                     | 2 +-
 3 files changed, 2 insertions(+), 8 deletions(-)

---

commit 35906f35bcbd8d969616061ec1420cb0136abbe5
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Thu Aug 17 19:34:34 2017 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:06:53 2018 -0800

    router/network: remove swarm dependency
    
    Network commands are built to query both the active cluster (if any) and
    the daemon. This commits makes them only deal with local networks.
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 api/server/router/network/network.go        |  5 +-
 api/server/router/network/network_routes.go | 81 +----------------------------
 cmd/dockerd/daemon.go                       |  2 +-
 3 files changed, 3 insertions(+), 85 deletions(-)

---

commit e752e6aa4c219da7ba74434cd4838fcedb5c229e
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Tue Jul 25 13:10:40 2017 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:06:50 2018 -0800

    pkg/discovery: remove consul,etcd,zookeeper backends
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 pkg/discovery/kv/kv.go | 12 ------------
 1 file changed, 12 deletions(-)

---

commit e468a65bedc93ce942c0bcc7f7c55889270d17e2
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Tue Jul 25 13:03:52 2017 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:06:50 2018 -0800

    daemon: only support journald and jsonfile log drivers
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 daemon/logdrivers_linux.go | 7 -------
 1 file changed, 7 deletions(-)

---
```

</details>

## Consolidate required binaries

This compiles all auxillary binaries required to run containers into one, that get called based on `argv[0]` (similar to how busybox works).

This includes:
* engine
* runc
* contained
* contained-ctr
* contained-shim
* proxy

<details>
<summary>Commits</summary>

```
commit 6e257f07dfd7f928251250f99adca929e0c6197b
Author:     Petros Angelatos <petrosagg@gmail.com>
AuthorDate: Tue Jul 25 17:09:09 2017 -0700
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:06:50 2018 -0800

    cmd/rce: adapt imports to new packages
    
    Signed-off-by: Petros Angelatos <petrosagg@gmail.com>

 cmd/rce/rce.go | 8 ++++----
 1 file changed, 4 insertions(+), 4 deletions(-)

---

commit 8b8b55c8a0229339848cfd11bd8967c4363f5e2c
Author:     Yossi Eliaz <yossi@resin.io>
AuthorDate: Fri May 5 16:56:03 2017 -0500
Commit:     Petros Angelatos <petrosagg@gmail.com>
CommitDate: Mon Feb 12 17:06:47 2018 -0800

    cmd/rce: Added the main of binary consolidation
    
    Added the rce.go main code which boxes all the binary into one suite.
    
    Signed-off-by: Yossi Eliaz <yossi@resin.io>

 Makefile                           |  4 ++--
 cmd/dockerd/config.go              |  2 +-
 cmd/dockerd/config_common_unix.go  |  2 +-
 cmd/dockerd/config_experimental.go |  2 +-
 cmd/dockerd/config_solaris.go      |  2 +-
 cmd/dockerd/config_unix.go         |  2 +-
 cmd/dockerd/config_unix_test.go    |  2 +-
 cmd/dockerd/config_windows.go      |  2 +-
 cmd/dockerd/daemon.go              |  2 +-
 cmd/dockerd/daemon_freebsd.go      |  2 +-
 cmd/dockerd/daemon_linux.go        |  2 +-
 cmd/dockerd/daemon_test.go         |  2 +-
 cmd/dockerd/daemon_unix.go         |  2 +-
 cmd/dockerd/daemon_unix_test.go    |  2 +-
 cmd/dockerd/daemon_windows.go      |  2 +-
 cmd/dockerd/docker.go              |  5 +++--
 cmd/dockerd/docker_windows.go      |  3 ++-
 cmd/dockerd/metrics.go             |  2 +-
 cmd/dockerd/options.go             |  2 +-
 cmd/dockerd/options_test.go        |  2 +-
 cmd/dockerd/service_unsupported.go |  2 +-
 cmd/dockerd/service_windows.go     |  2 +-
 cmd/rce/rce.go                     | 41 ++++++++++++++++++++++++++++++++++++++
 hack/make.sh                       |  1 +
 hack/make/.binary-setup            |  2 ++
 hack/make/binary-rce               | 16 +++++++++++++++
 26 files changed, 86 insertions(+), 24 deletions(-)

---
```

</details>

## Various small patches

<details>
<summary>Commits</summary>

```
```

</details>
