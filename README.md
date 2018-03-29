xplex streamer
==============

Media streaming server for xplex with RTMP support

What it does:

- Work as edge servers for accepting RTMP ingestions
- Fetch configurations from rig
- Relay RTMP stream to destinations that it obtains from rig

## Install

- [Install Golang 1.9+](https://golang.org/doc/install)
- Install [dep](https://golang.github.io/dep/docs/installation.html)

In project root, run:

```sh
$ dep ensure
```

Compile debug builds using `build.sh`:

```sh
$ ./build.sh dev
```

Compile static binaries for release profile:

```sh
$ ./build.sh release
```

Compiled binaries are put in `./bin/`.
