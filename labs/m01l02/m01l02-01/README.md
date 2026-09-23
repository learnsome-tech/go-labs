# m01l02-01: Getting the toolchain onto a machine

You can install the Go toolchain the way that suits a laptop, a build agent or a container, and use the go command's own subcommands for formatting, checking and documentation instead of reaching for extra tools.

There are three ways to get a Go toolchain, and which one you want depends on where the machine lives. On a laptop, a package manager is the least trouble. On a server or in a pipeline you usually want the official archive, because it is one download, one directory on the path, and no opinions from a distribution about which version you get. In continuous integration the honest answer is often neither: use the official Go container image and let your build agent stay empty. One detail worth knowing early is that a module can record the toolchain version it wants, and a modern go command will fetch and use exactly that version rather than whatever happens to be installed.

- Official archive: unpack it, put one directory on the path
- Package managers are fine for a laptop: brew, apt, winget
- In pipelines, use the official golang container image
- A module can pin the toolchain it wants to be built with

Notes or exercise panel; no executable listing.
