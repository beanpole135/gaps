# GAPS (Go Application Package System)

This is a cross-platform, static-binary package system based around the [Go](https://golang.org/) programming language.

## Utilities
* gaps
   * Primary tool. Used to manage packages on a system
* gaps-pkg
   * Package creation utility. Used to build packages for all platforms.
* gaps-repo
   * Repository service daemon. Used to create/run a local repository of packages.

Currently supported Platforms

* Windows (64-bit)
* Mac (64-bit)
* Linux (64-bit)
* FreeBSD (64-bit)
* OpenBSD (64-bit)
* NetBSD (64-bit)
* DragonflyBSD (64-bit)
* Solaris (64-bit)
* Plan9 (64-bit)

## Build instructions
Simply run "make" or "make build" to compile any/all tools for your current platform.

Run "make package" to build any/all tools for all platforms.
