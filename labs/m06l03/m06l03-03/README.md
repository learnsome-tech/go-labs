# m06l03-03: Shutdown needs a deadline

You can stop accepting traffic on SIGTERM, cancel active work, and let a Go service finish cleanly.

The server Shutdown method stops new connections and waits for active handlers. Give it a context with a deadline so a stuck dependency cannot hold the process forever. After shutdown returns, close any remaining listeners and report whether the deadline expired. The deadline is part of the deployment contract and should fit inside the orchestrator's termination grace period.

- Shutdown stops new connections
- A context bounds the wait
- Close listeners after active work
- Report timeout as an operational error

Notes or exercise panel; no executable listing.
