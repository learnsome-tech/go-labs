# m06l02-02: A health handler

You can expose a small HTTP handler, return useful status codes, and configure a server with a timeout.

The health function receives the writer and request required by the handler contract. It sets a content type and writes a small response. Main registers the path on a new multiplexer and prints true to prove the route exists. Starting the listener is a separate concern, which keeps this example easy to test with an in memory request and recorder.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
