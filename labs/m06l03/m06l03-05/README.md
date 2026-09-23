# m06l03-05: The graceful shutdown sequence

You can stop accepting traffic on SIGTERM, cancel active work, and let a Go service finish cleanly.

Graceful shutdown is a sequence: receive SIGTERM, stop new traffic, cancel or finish active work, and call Shutdown with a deadline. Return from main only after cleanup has completed or the deadline has made the remaining failure explicit. That sequence is what makes a Go service behave predictably inside an orchestrator.

- Receive SIGTERM
- Stop new traffic
- Shutdown with a deadline
- Return after cleanup

Notes or exercise panel; no executable listing.
