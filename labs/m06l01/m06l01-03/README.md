# m06l01-03: Getting it wrong: usage and exit code two

You can build a command line tool with the standard flag package that takes typed options, reads positional arguments, falls back to environment variables, and fails with a usage message and the right exit code.

Callers get flags wrong, and what happens then is part of your tool's interface. The package already does the sensible thing: an unknown flag stops parsing, prints a message, prints the usage block, and exits with status two. That usage block is generated from the declarations, which is why a decent help string next to each option pays for itself. You can replace the default heading by assigning your own function to Usage, and look at where it writes: standard error, not standard output, so a pipeline consuming your real output is never polluted by help text. Let us pass a flag that does not exist and read the whole thing, exit status and all.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
