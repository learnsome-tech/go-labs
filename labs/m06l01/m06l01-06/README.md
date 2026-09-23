# m06l01-06: Exit codes and which stream to write to

You can build a command line tool with the standard flag package that takes typed options, reads positional arguments, falls back to environment variables, and fails with a usage message and the right exit code.

Two conventions that cost nothing and save a great deal of grief. First, exit codes. Zero means the work succeeded. Anything else means it did not, and the shell, the pipeline and the orchestrator all read that number. By convention a misused command line exits with two, which is what the flag package already does for you, and a genuine failure exits with one. Pick your codes, write them into the help text, and do not change them casually, because somebody has a script branching on them. Second, streams. Results go to standard output so they can be piped onward. Usage, warnings and errors go to standard error so they stay visible even when the output is redirected into a file.

- Exit zero only when the work really succeeded
- Two for a misused command line, one for a failed check
- Results to stdout, usage and diagnostics to stderr
- os.Exit skips deferred calls: leave from main, not deeper

Notes or exercise panel; no executable listing.
