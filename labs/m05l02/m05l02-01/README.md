# m05l02-01: Reader and Writer are the universal seams

You can copy streams, read bounded input, and treat files and network bodies through common interfaces.

The io package gives Go tooling two small interfaces. A Reader supplies bytes, and a Writer accepts bytes. Files, network bodies, buffers, and standard input can all satisfy those interfaces, so a function that copies a Reader to a Writer does not care where the bytes came from. This is one of the most useful seams in the standard library: test with a buffer, then run with a file or a response body.

- Reader supplies bytes
- Writer accepts bytes
- Copy connects any two streams
- Close belongs to the owner

Notes or exercise panel; no executable listing.
