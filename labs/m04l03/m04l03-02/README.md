# m04l03-02: Selecting a result or a timeout

You can wait on several channel operations and keep a concurrent worker responsive.

The select waits for either a result or a timer. The worker sends quickly, so the result case wins and prints ready. If the worker took longer than the timer, the timeout case would run instead. This is the basic responsive shape for a network check: wait for useful work, but always keep an escape path for a caller that cannot wait forever.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
