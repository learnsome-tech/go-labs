# m04l05-03: Cancellation must reach every layer

You can bound a concurrent operation with a timeout and return promptly when an orchestrator or caller cancels it.

Cancellation only works when every layer passes the same context downward and chooses context aware operations. A child function should not replace a caller deadline with a background context, because that breaks the caller's guarantee. When a blocking operation ends with cancellation, return promptly and let the outer layer add useful context. Tests should cover the successful result and the cancelled path so both lifetimes stay intentional.

- Pass the same context downward
- Do not replace a caller deadline
- Return promptly from blocked work
- Test both success and cancellation

Notes or exercise panel; no executable listing.
