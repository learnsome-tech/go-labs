# m04l05-05: The cancellation rules

You can bound a concurrent operation with a timeout and return promptly when an orchestrator or caller cancels it.

Timeouts and cancellation give concurrent operations a finite lifetime. Derive a context at the boundary, pass it through every layer, and make blocked workers select on Done. Return the context error so the caller can distinguish a deadline from other failures. Before a process exits, cancel its work and join the workers that must clean up.

- Bound every remote operation
- Pass the caller context downward
- Return ctx Err on cancellation
- Join workers before exit

Notes or exercise panel; no executable listing.
