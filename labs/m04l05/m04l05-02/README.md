# m04l05-02: A timeout returns an error

You can bound a concurrent operation with a timeout and return promptly when an orchestrator or caller cancels it.

Work waits for either its simulated result or the context Done signal. The derived context has a very short timeout, so the deadline case wins and work returns the context error. Main defers cancel to release the timer even though the deadline will arrive soon. This same pattern wraps an HTTP request, a file operation, or a health probe so the caller always gets control back.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
