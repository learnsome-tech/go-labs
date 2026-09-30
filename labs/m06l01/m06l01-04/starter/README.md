# m06l01-04: Positional arguments are what is left over

You can build a command line tool with the standard flag package that takes typed options, reads positional arguments, falls back to environment variables, and fails with a usage message and the right exit code.

Flags are the options; the positional arguments are the things you are acting on. Call Args after parsing and you get everything that was not consumed as a flag, in order, as a slice of strings. Here the tool wants at least one host, so an empty slice is a misuse of the command line: a message on standard error and exit two, matching what the package itself does for a bad flag. One habit worth forming: parse first, validate second, work third. Validation failures are cheap and belong before you open a socket or a file, not halfway through a run. Let us run it with one flag and three hosts and watch the split.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
