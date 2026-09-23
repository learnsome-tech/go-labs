# m07l03-02: Classifying a timeout

You can give every check a deadline, classify failures, and keep one slow target from blocking the report.

The cancelled context becomes an unhealthy result with a reason attached. A real request would use a deadline and return deadline exceeded when the endpoint takes too long. The checker keeps the error text for the report, while the coordinator can still classify the result as a timeout or cancellation.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
