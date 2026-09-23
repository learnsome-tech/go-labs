# m06l02-04: Testing a handler in memory

You can expose a small HTTP handler, return useful status codes, and configure a server with a timeout.

The recorder and request let us exercise the handler without opening a port. After the call, the recorder exposes the status code and body that a real client would receive. In a test this same setup gives a deterministic check for status, headers, and JSON. Keep network binding out of unit tests and reserve it for an integration check.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
