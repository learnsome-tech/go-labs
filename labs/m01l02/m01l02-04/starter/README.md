# m01l02-04: Formatting is settled, not argued about

You can install the Go toolchain the way that suits a laptop, a build agent or a container, and use the go command's own subcommands for formatting, checking and documentation instead of reaching for extra tools.

Go ships one formatter and no configuration for it. Tabs for indentation, one true brace style, and no discussion. The first command lists which files disagree with the canonical form, and here there is one. The second rewrites it in place. Run the listing again and there is nothing left to say, which is how this command reports success. For a team this removes a whole category of review comment, and in a pipeline the listing form is the check you want: if it prints any file name, somebody committed unformatted code, so fail the build and move on. Most editors run this on save, and you should let yours.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
