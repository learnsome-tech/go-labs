# m06l01-08: What to carry forward

You can build a command line tool with the standard flag package that takes typed options, reads positional arguments, falls back to environment variables, and fails with a usage message and the right exit code.

To carry forward: declare every option before parsing, read the values only after Parse has run, and take the positional arguments from Args. Let the package generate the usage block from your declarations, keep it on standard error, and honour the exit code conventions so whatever calls your tool can tell success from failure from misuse. Wire the environment in as the default value and the precedence order arranges itself. Next lesson the probe grows a server: the same standard library, a multiplexer for routing, and the four timeouts that decide whether your service survives contact with the open internet.

- Declare, then Parse, then read through the pointers
- Args holds the positional arguments Parse did not take
- Usage and errors to stderr, results to stdout
- Environment as the default, an explicit flag overrides

Notes or exercise panel; no executable listing.
