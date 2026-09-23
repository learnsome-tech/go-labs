# m07l02-03: Concurrency does not define report order

You can run checks concurrently, collect one result per target, and preserve a useful report order.

Concurrent workers finish at different times. A channel preserves the order in which results arrive, not necessarily the order in which targets were configured. If a stable report matters, store an index in each target or sort the collected results before encoding them. The coordinator owns closing the result stream after every worker has finished.

- Workers finish at different times
- A channel preserves delivery, not input order
- Sort results when order matters
- Keep the coordinator as the owner

Notes or exercise panel; no executable listing.
