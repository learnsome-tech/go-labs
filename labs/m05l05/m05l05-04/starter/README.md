# m05l05-04: Testing a response with a fake round tripper

You can issue an HTTP request with a context, check the response, and close the body on every path.

A custom transport lets a test or a small example supply a deterministic response without a network call. The fake round tripper returns a successful status and a body that can be closed. The client still follows the normal lifecycle, so the calling code can test status handling and decoding without waiting on DNS or a remote service. In production the transport would be the real network implementation.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
