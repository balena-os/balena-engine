# balenaEngine Benchmarking

This directory contains some scripts we use to benchmark balenaEngine. They are
not super stable and ready for public consumption, but they are good enough for
informing decisions when working in improvements.

Currently, there's actually just one script here.

## `delta-benchmarks.sh`

This script collects some metrics on the generation of deltas. Namely, it
measures how long it takes, how much memory it uses, and how large are the
resulting deltas.

The script does this for a list of references (branches, tags, commits) defined
in the `branches` variable at its start, and for each of the test cases defined
in `testCases`. You can customize these two variables as you need.

You need to run this as the superuser ()`root`), from the root of the
balena-engine repository. Something like this should work:

```sh
sudo ./balena-benchmarking/delta-benchmarks.sh
```

All required images will be pulled into
`./balena-benchmarking/balenad-data-root`. If you already have pulled all
images, you can run a bit faster by using this:

```sh
sudo SKIP_PULL=y ./balena-benchmarking/delta-benchmarks.sh
```

Results are written to `./balena-benchmarking/delta.csv`.
