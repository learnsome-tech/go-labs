# m01l02-03: One command, a dozen verbs

You can install the Go toolchain the way that suits a laptop, a build agent or a container, and use the go command's own subcommands for formatting, checking and documentation instead of reaching for extra tools.

Almost everything you do with Go is a verb on one command. Compile and execute in one step, or compile and keep the binary. Run the tests. Check for suspicious code. Format the source. Manage dependencies. Read documentation. Print the environment the toolchain thinks it is in. That is deliberate, and for an operations audience it matters more than it sounds: a project needs no build file, no task runner and no plugin list before a stranger can compile it. If you can clone it and you have the toolchain, you can build it, and the commands are the same in every Go repository you will ever open.

- run and build: compile, with or without keeping the binary
- test, vet and fmt: the checks, all in the box
- mod and work: dependencies and multi module checkouts
- doc, env and clean: look things up, look around, tidy up

Notes or exercise panel; no executable listing.
