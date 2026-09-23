# m03l03-02: An ordinary error keeps control

You can distinguish a programmer invariant from an expected operational error and use recover only at a deliberate boundary.

The lookup checks its input and returns an error when the index is outside the list. The caller stays in ordinary control flow, prints a useful result, and can choose what to do next. This is the right shape for input that may be wrong in normal operation. A command line flag, a remote response, or a stale configuration can all reach this branch without meaning the program itself is broken.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
