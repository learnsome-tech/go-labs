# m03l01-02: Returning a value and an error

You can return errors, inspect them, and keep failure handling explicit in a command line tool.

Divide returns two results. A successful calculation returns the quotient and a nil error, while a zero count returns a zero value and a descriptive error. The caller checks the error immediately. Only after that check does it print the result. This ordering prevents a bad result from travelling further into the program. In a health checker, the same shape lets a probe return a response and an error, leaving the caller in charge of whether a failed probe should be retried or reported.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
