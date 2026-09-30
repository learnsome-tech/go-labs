# m06l03-04: A bounded shutdown call

You can stop accepting traffic on SIGTERM, cancel active work, and let a Go service finish cleanly.

The server has no listener in this small example, so Shutdown returns nil immediately. The important shape is the context with a finite deadline and the deferred cancel that releases its timer. In a running service, call this after SIGTERM, then wait for the shutdown result before returning from main.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
