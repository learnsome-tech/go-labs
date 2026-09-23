# m04l05-01: Every remote wait needs an escape

You can bound a concurrent operation with a timeout and return promptly when an orchestrator or caller cancels it.

A remote operation without an escape can hold a goroutine forever. A timeout gives the operation a maximum wait, while cancellation follows the caller when the request or process no longer needs the result. Derive one context for the whole check and pass it to every operation that can wait. The worker must observe that context so a timeout ends the work instead of leaving a sleeping goroutine behind.

- Timeouts bound an operation
- Cancellation follows the caller
- Use one context for the whole check
- Do not leak a sleeping worker

Notes or exercise panel; no executable listing.
