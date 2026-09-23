# m06l03-01: Orchestrators ask processes to leave

You can stop accepting traffic on SIGTERM, cancel active work, and let a Go service finish cleanly.

An orchestrator usually sends SIGTERM before it removes a process. Treat that signal as a request to leave cleanly. Stop accepting new traffic, give active requests a bounded window to finish, cancel work that cannot finish, and exit after resources are closed. A process that ignores the signal is eventually killed, which can cut off responses and lose buffered output.

- SIGTERM is a request to stop
- Stop accepting new traffic
- Give active work a deadline
- Exit only after cleanup

Notes or exercise panel; no executable listing.
