# m03l03-04: Recover does not make bad state safe

You can distinguish a programmer invariant from an expected operational error and use recover only at a deliberate boundary.

Recover contains a panic; it does not repair the state that caused it. Use it when an isolated request or worker must be prevented from taking down its supervisor. Record enough context to investigate, and decide whether the process should stop after the recovery. Do not continue with state you know is corrupt. In many cases, validating input before the invariant is reached is clearer and safer than relying on a recovery path.

- Use it to contain an isolated failure
- Log enough context to investigate
- Do not continue with corrupt state
- Prefer validation before the invariant

Notes or exercise panel; no executable listing.
