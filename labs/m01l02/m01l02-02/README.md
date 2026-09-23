# m01l02-02: Checking what you installed

You can install the Go toolchain the way that suits a laptop, a build agent or a container, and use the go command's own subcommands for formatting, checking and documentation instead of reaching for extra tools.

Two commands to confirm the install. The first tells you which toolchain and which target platform you have: the version, the operating system and the processor architecture. If that prints, the install is done and there is nothing else to configure; there is no path variable for libraries and no virtual environment to create. The second is the habit worth forming now. Every subcommand documents itself, and the first line is the usage line, which is usually the answer you wanted. When you cannot remember whether the output flag is on build or on install, this is faster than a web search and it matches the version you actually have.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
