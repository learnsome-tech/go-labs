# m03l03-01: A panic is not a routine error

You can distinguish a programmer invariant from an expected operational error and use recover only at a deliberate boundary.

Use an error for an expected operational failure such as a missing file, a refused connection, or invalid input. A panic signals that the program reached an invariant it was not supposed to violate. It unwinds the current goroutine and runs deferred calls while it travels. Recover can turn that panic back into a value, but only at a deliberate boundary such as a request handler or a worker supervisor. Do not use panic as a shorter spelling of ordinary error handling.

- Return expected failures as errors
- Panic signals a broken invariant
- A panic unwinds the current goroutine
- Recover belongs at a clear boundary

Notes or exercise panel; no executable listing.
