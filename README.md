# Rspscale

https://scale.ropsoft.cloud

Private WireGuard® networks made easy

## Overview

This repository contains the majority of Rspscale's open source code.
Notably, it includes the `rspscaled` daemon and
the `rspscale` CLI tool. The `rspscaled` daemon runs on Linux, Windows,
[macOS](https://scale.ropsoft.cloud/kb/1065/macos-variants/), and to varying degrees
on FreeBSD and OpenBSD. The Rspscale iOS and Android apps use this repo's
code, but this repo doesn't contain the mobile GUI code.

Other [Rspscale repos](https://github.com/orgs/rspscale/repositories) of note:

* the Android app is at https://github.com/ropsoft7/rspscale-android
* the Synology package is at https://github.com/ropsoft7/rspscale-synology
* the QNAP package is at https://github.com/ropsoft7/rspscale-qpkg
* the Chocolatey packaging is at https://github.com/ropsoft7/rspscale-chocolatey

For background on which parts of Rspscale are open source and why,
see [https://scale.ropsoft.cloud/opensource/](https://scale.ropsoft.cloud/opensource/).

## Using

We serve packages for a variety of distros and platforms at
[https://pkgs.scale.ropsoft.cloud](https://pkgs.scale.ropsoft.cloud/).

## Other clients

The [macOS, iOS, and Windows clients](https://scale.ropsoft.cloud/download)
use the code in this repository but additionally include small GUI
wrappers. The GUI wrappers on non-open source platforms are themselves
not open source.

## Building

We always require the latest Go release, currently Go 1.23. (While we build
releases with our [Go fork](https://github.com/tailscale/go/), its use is not
required.)

```
go install scale.ropsoft.cloud/cmd/rspscale{,d}
```

If you're packaging Rspscale for distribution, use `build_dist.sh`
instead, to burn commit IDs and version info into the binaries:

```
./build_dist.sh scale.ropsoft.cloud/cmd/rspscale
./build_dist.sh scale.ropsoft.cloud/cmd/rspscaled
```

If your distro has conventions that preclude the use of
`build_dist.sh`, please do the equivalent of what it does in your
distro's way, so that bug reports contain useful version information.

## Bugs

Please file any issues about this code or the hosted service on
[the issue tracker](https://github.com/ropsoft7/rspscale/issues).

## Contributing

PRs welcome! But please file bugs. Commit messages should [reference
bugs](https://docs.github.com/en/github/writing-on-github/autolinked-references-and-urls).

We require [Developer Certificate of
Origin](https://en.wikipedia.org/wiki/Developer_Certificate_of_Origin)
`Signed-off-by` lines in commits.

See `git log` for our commit message style. It's basically the same as
[Go's style](https://github.com/golang/go/wiki/CommitMessage).

## About Us

[Rspscale](https://scale.ropsoft.cloud/) is primarily developed by the
people at https://github.com/orgs/rspscale/people. For other contributors,
see:

* https://github.com/ropsoft7/rspscale/graphs/contributors
* https://github.com/ropsoft7/rspscale-android/graphs/contributors

## Legal

WireGuard is a registered trademark of Jason A. Donenfeld.
# rspscale
