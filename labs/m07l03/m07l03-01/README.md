# m07l03-01: A health check must finish

You can give every check a deadline, classify failures, and keep one slow target from blocking the report.

A health checker is useful only when it finishes within an operational window. Derive a context with a timeout for each request, return the context error as a result, and wrap transport failures with the target name. One slow endpoint should not block healthy results from the other targets. The coordinator can then report every target with a reason that an operator can act on.

- Derive a context per check
- Return timeout as a result
- Wrap transport errors with target context
- Let other targets continue

Notes or exercise panel; no executable listing.
