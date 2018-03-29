#!/bin/sh

set -e

function cleanRebuild {
    rm -f bin/xplex-streamer
    mkdir -p bin/
}

function buildServerRelease {
    echo "Building xplex-streamer server release"
    CGO_ENABLED=0 GOOS=linux go build -o bin/xplex-streamer -a -ldflags '-extldflags "-static"' .
    echo "Compiled to './bin/xplex-streamer'"
}

function buildServerDev {
    echo "Building xplex-streamer server dev"
    go build -o bin/xplex-streamer .
    echo "Compiled to './bin/xplex-streamer'"
}

function buildDev {
    cleanRebuild
    buildServerDev
}

function buildRelease {
    cleanRebuild
    buildServerRelease
}

case "$1" in
    "release")
    buildRelease
    ;;
    "dev")
    buildDev
    ;;
    *)
    echo "Usage: ./build.sh [dev|release|help]"
    echo "Binaries are put in './bin/'"
    ;;
esac
