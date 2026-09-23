# m04l04-05: The context rules

You can pass context through a call chain and use its deadline and values at the right boundaries.

Context gives a request a shared lifetime signal. Pass it first through the call chain, observe Done in every worker that can wait, and return its error when stopping. Use values only for request metadata. Whenever you create a child context with cancellation or a deadline, call the cancel function in the owner that created it.

- Pass context first
- Use it for lifetime and deadlines
- Observe Done in every worker
- Call cancel when you create a child

Notes or exercise panel; no executable listing.
