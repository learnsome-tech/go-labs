# m03l03-03: Recover at a worker boundary

You can distinguish a programmer invariant from an expected operational error and use recover only at a deliberate boundary.

This worker boundary uses a named error result and a deferred recovery function. If the worker panics, recover captures the value and changes the returned error. Main can report the failure while the supervisor remains alive. Recover is useful here because one isolated job should not tear down a process that can continue serving other jobs. The boundary is explicit, and the rest of the program still handles a normal error value.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
