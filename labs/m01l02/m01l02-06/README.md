# m01l02-06: Documentation without leaving the shell

You can install the Go toolchain the way that suits a laptop, a build agent or a container, and use the go command's own subcommands for formatting, checking and documentation instead of reaching for extra tools.

The documentation is in the toolchain too, generated from the comments in the source. Ask about one function and you get its package, its signature and its prose. Notice the signature in the first answer: Println takes any number of arguments and returns two values, a count and an error, which is a pattern you will see everywhere in this language. The second lookup shows a different shape: one argument in, one string out, and no error at all, because a missing environment variable is empty. This is the documentation for the package you have installed, not for whatever version a search engine decided to show you, and that distinction has saved me hours.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
