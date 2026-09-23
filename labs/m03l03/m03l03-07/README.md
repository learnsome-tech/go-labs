# m03l03-07: Try it yourself

You can distinguish a programmer invariant from an expected operational error and use recover only at a deliberate boundary.

Build a tiny supervisor. Let the worker return an ordinary error for bad input, and add one separate invariant check that panics. Recover only in the supervisor boundary, turn the panic into a clear message, and let the ordinary error travel as an error. Run both paths and compare how the caller can distinguish an expected failure from a contained programming fault.

- Write a worker that returns an error for bad input.
- Add a separate invariant check that panics.
- Recover at the supervisor boundary and print whether the worker failed or was contained.

Notes or exercise panel; no executable listing.
