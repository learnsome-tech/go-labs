# m01l01-04: What the compiler actually hands you

You can explain why platform and DevOps teams write their tooling in Go, compile and run a program with the go command, and describe exactly what the compiler hands you at the end of a build.

It is worth being precise about what you get, because it shapes how you ship. The output is machine code for one operating system and one processor family, with the Go runtime, the garbage collector and every package you imported linked in already. There is nothing beside it to install. If you need a different target, you set two environment variables and build again, which is how a laptop produces a Linux binary for a container. The cost of all this is that the file is bigger than a script: a few megabytes rather than a few kilobytes. For a tool that has to run on a stripped down image, that is a trade almost every operations team takes happily.

- Machine code for one operating system and one processor
- The runtime and every imported package linked in already
- No interpreter, no virtual machine, no dependency install
- Cross compile for another target by setting two variables

Notes or exercise panel; no executable listing.
