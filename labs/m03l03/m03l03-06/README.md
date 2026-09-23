# m03l03-06: Error or panic

You can distinguish a programmer invariant from an expected operational error and use recover only at a deliberate boundary.

Return an error when an operation can fail in normal use. Reserve panic for a broken programmer invariant, and recover only where an isolated failure must be contained. Deferred calls run during the unwind, which helps cleanup, but recovery does not make corrupted state trustworthy. Validate inputs early and keep the ordinary error path easy to see in every tool you ship.

- Expected problems return errors
- Broken invariants may panic
- Recover only at an intentional boundary
- Never hide corrupt state

Notes or exercise panel; no executable listing.
