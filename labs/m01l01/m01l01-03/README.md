# m01l01-03: The build step, and what it produces

You can explain why platform and DevOps teams write their tooling in Go, compile and run a program with the go command, and describe exactly what the compiler hands you at the end of a build.

Now the same program from the shell. First ask the toolchain which version is installed, which also tells you the operating system and processor it will build for by default. Then build with an output name of your choosing. That command prints nothing at all, which in Go tooling means it worked. What it leaves behind is a file: machine code for this platform, with the Go runtime and every package you imported already inside it. Run the binary directly and you get the same two lines. Notice what you did not do: there was no interpreter, no virtual machine and no list of dependencies to install on the machine that runs it.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
