# m06l01-09: Try it yourself

You can build a command line tool with the standard flag package that takes typed options, reads positional arguments, falls back to environment variables, and fails with a usage message and the right exit code.

Your turn. Take the probe and give it an option that switches the output to a single JSON object, so a script can consume it without any text wrangling. Then add a deadline option that falls back to an environment variable, using the helper pattern from earlier in the lesson. Finally, make the tool refuse to run when no hosts were given: a message on standard error and exit two, not a panic and not a quiet success. When that works, run it with an option that does not exist and check that the usage block still lists everything you added.

- Add a -json flag that prints the parsed configuration as one JSON object instead of lines.
- Add a -timeout duration flag that falls back to the PROBE_TIMEOUT environment variable.
- Make the tool exit two with your usage block when no host argument is given.

Notes or exercise panel; no executable listing.
