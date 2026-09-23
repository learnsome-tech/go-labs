# m06l01-01: The command line is an interface too

You can build a command line tool with the standard flag package that takes typed options, reads positional arguments, falls back to environment variables, and fails with a usage message and the right exit code.

Every tool you have been annoyed by got its command line wrong: it ignored an option, it printed a stack trace where a usage line belonged, or it exited zero after failing. Go ships a package called flag in the standard library, and for a single purpose tool it is the whole answer. You get typed options with defaults, a usage message generated from those declarations, and whatever positional arguments are left over. What you do not get is subcommands, shell completion and nested help, and we will look at where that line sits. Throughout this lesson the tool we are building is a reachability probe, the sort of thing a pipeline runs against a list of hosts.

- Typed options with defaults, not hand rolled string parsing
- A usage block generated from the declarations themselves
- Positional arguments for the things being acted on
- Exit codes and streams a calling script can rely on

Notes or exercise panel; no executable listing.
