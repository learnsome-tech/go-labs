# m06l01-07: Where the standard library stops and Cobra starts

You can build a command line tool with the standard flag package that takes typed options, reads positional arguments, falls back to environment variables, and fails with a usage message and the right exit code.

So when do you reach for Cobra? When the tool grows subcommands, each with its own flags and its own help page, and you want shell completion generated for it. Cobra does that well, and it is what the Kubernetes command line tool and dozens of others are built on. Underneath it uses a package called pflag, which adds the double dash long option style you expect from GNU tools. But a probe that does one job needs none of that. The standard library version has no dependency to audit, no version to bump and nothing to vendor, and for an internal tool that ends up inside a container image those are real advantages. Start here, and move only when the shape of the tool demands it.

- One tool, one job: flag is the whole answer
- Subcommands, completion, nested help: reach for Cobra
- Cobra sits on pflag, which adds GNU style long options
- Every dependency is one more thing your build must trust

Notes or exercise panel; no executable listing.
