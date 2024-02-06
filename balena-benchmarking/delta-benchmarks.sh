#!/bin/bash

# Environment variables affecting the script behavior:
#
# - SKIP_PULL: if set to 'y', images will not be pulled before running the
#   benchmarks. Useful to spare some time if you are sure you already have all
#   needed images locally.

set -e

# Branches to benchmark.
branches=(
    "master"
    "lmb/librsync-memory"
)

# Test cases to benchmark (test case name, basis image, target image).
testCases=(
    "busybox-1.36.0-1.36.1          busybox:1.36.0                                   busybox:1.36.1"
    "busybox-1.25.0-1.36.1          busybox:1.25.0                                   busybox:1.36.1"
    "debian-10.0-11.7               debian:10.0                                      debian:11.7"
    "debian-11.6-11.7               debian:11.6                                      debian:11.7"
    "debian-slim-11.7               debian:11.7-slim                                 debian:11.7"
    "debian-11.7-slim               debian:11.7                                      debian:11.7-slim"
    "ubuntu-18.04-23.04             ubuntu:18.04                                     ubuntu:23.04"
    "alpine-3.7-3.18                alpine:3.7                                       alpine:3.18"
    "audio-aarch64-0.5.5-0.5.6      bh.cr/balenalabs/audio-aarch64/0.5.5             bh.cr/balenalabs/audio-aarch64/0.5.6"
    "audio-amd64-0.5.5-0.5.6        bh.cr/balenalabs/audio-amd64/0.5.5               bh.cr/balenalabs/audio-amd64/0.5.6"
    "browser-aarch64-2.3.7-2.4.7    bh.cr/balenalabs/browser-aarch64/2.3.7           bh.cr/balenalabs/browser-aarch64/2.4.7"
    "browser-amd64-2.3.7-2.4.7      bh.cr/balenalabs/browser-amd64/2.3.7             bh.cr/balenalabs/browser-amd64/2.4.7"
    "nodered-aarch64-2.4.0-2.4.1    bh.cr/balenalabs/balena-node-red-aarch64/2.4.0   bh.cr/balenalabs/balena-node-red-aarch64/2.4.1"
    "ca-priv-amd64-0.0.12-0.0.13    bh.cr/balena/ca-private-amd64/0.0.12             bh.cr/balena/ca-private-amd64/0.0.13"
    "ca-priv-amd64-0.0.13-0.0.12    bh.cr/balena/ca-private-amd64/0.0.13             bh.cr/balena/ca-private-amd64/0.0.12"
    "unzoner-armv7hf-1.2.0-1.2.23   bh.cr/belodetek/unzoner-armv7hf/1.2.0            bh.cr/belodetek/unzoner-armv7hf/1.2.23"
    "unzoner-armv7hf-1.2.23-1.2.0   bh.cr/belodetek/unzoner-armv7hf/1.2.23           bh.cr/belodetek/unzoner-armv7hf/1.2.0"
    # TODO maybe: https://gitlab.com/nvidia/container-images/l4t-base
)

balenadDataRoot="./balena-benchmarking/balenad-data-root"
balenadPIDFile="/var/run/balena-engine.pid"
deltaTag="balena-engine-delta-benchmark-image"

function assertRunningFromRepoRoot() {
    if [ ! -f "Makefile" ]; then
        echo "Please run from the root of the balena-engine repository."
        exit 1
    fi
}

# Build balenaEngine from branch $1.
function buildBalenaEngine() {
    echo
    echo "BUILDING BALENA ENGINE FROM BRANCH $1"
    echo
    git checkout "$1"
    make dynbinary
}

function startBalenad() {
    echo "Starting balenad..."
    mkdir -p "$balenadDataRoot"

    if [ -f "$balenadPIDFile" ]; then
        killBalenad
    fi

    balenad --data-root "$balenadDataRoot" --pidfile $balenadPIDFile &> /dev/null &
    echo -n "Waiting for balenad to start... "

    while [ ! -f "$balenadPIDFile" ]; do
        sleep 1
    done

    while [ ! balena-engine info &> /dev/null ]; do
        sleep 1
    done
    echo " done! (PID = $(cat "$balenadPIDFile"))"
}

function killBalenad() {
    echo "Killing balenad..."

    if [ ! -f "$balenadPIDFile" ]; then
        return
    fi
    kill $(cat $balenadPIDFile)
    sleep 5
    if [ -f "$balenadPIDFile" ]; then
        echo "balenaEngine still running, killing with -KILL"
        kill -KILL $(cat $balenadPIDFile)
    fi
}

function pullAllImages() {
    if [ "$SKIP_PULL" == "y" ]; then
        return
    fi

    echo
    echo "PULLING ALL IMAGES"
    echo

    buildBalenaEngine "master"

    startBalenad

    for testCase in "${testCases[@]}"; do
        tcBasis=$(echo $testCase | awk '{print $2}')
        tcTarget=$(echo $testCase | awk '{print $3}')

        balena-engine pull "$tcBasis"
        balena-engine pull "$tcTarget"
    done

    killBalenad
}

function balenadMaxMemory() {
    if [ ! -f "$balenadPIDFile" ]; then
        echo "balenad not running!"
        exit 1
    fi

    # Read the high water mark (VmHWM) of the balenad process.
    cat /proc/$(cat "$balenadPIDFile")/status | grep VmHWM | awk '{print $2}'
}

export PATH="$(pwd)/bundles/dynbinary-daemon:$PATH"

assertRunningFromRepoRoot

# Remember the current branch so we can switch back to it later.
originalBranch=$(git rev-parse --abbrev-ref HEAD)
echo "Running from this branch: $originalBranch"

pullAllImages

# The CSV file where results will be stored.
csvResults="./balena-benchmarking/delta.csv"
tmpResults="./balena-benchmarking/delta.tmp"

# Initialize the CSV files with headers.
echo "Case,Branch,BasisSize,DeltaSize,DeltaTime,DeltaMem" > "$csvResults"
rm -f "$tmpResults"

for branch in "${branches[@]}"; do
    echo "Running benchmarks for branch $branch"

    buildBalenaEngine "$branch"

    for testCase in "${testCases[@]}"; do
        tcName=$(echo $testCase | awk '{print $1}')
        tcBasis=$(echo $testCase | awk '{print $2}')
        tcTarget=$(echo $testCase | awk '{print $3}')

        echo "Running benchmark for $branch / $tcName"

        startBalenad

        # baselineMemInKB=$(balenadMaxMemory)
        deltaTimeInSecs=$(\time -f%e balena-engine image delta "$tcBasis" "$tcTarget" --tag "$deltaTag" 2>&1 | tail -n 1)
        usedMemInKB=$(balenadMaxMemory)
        usedMemInBytes=$((usedMemInKB * 1024))
        basisSizeInBytes=$(balena-engine inspect "$tcBasis" --format "{{.Size}}")
        deltaSizeInBytes=$(balena-engine inspect "$deltaTag" --format "{{.Size}}")

        # Collect data.
        echo "$tcName,$branch,$basisSizeInBytes,$deltaSizeInBytes,$deltaTimeInSecs,$usedMemInBytes" >> "$tmpResults"

        # Thanks Engine, you may go now.
        killBalenad
    done

    echo "Done with branch $branch"
done

echo "Preparing final results..."
sort "$tmpResults" >> "$csvResults"
rm -f "$tmpResults"

# Switch back to the original branch.
echo "Restoring original branch $originalBranch..."
git checkout "$originalBranch"

echo "Done with everything!"
