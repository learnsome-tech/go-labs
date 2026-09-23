# m04l01-03: Concurrency needs an ownership story

You can start goroutines, wait for their work, and recognise the ownership questions concurrency introduces.

Concurrency is a design choice about ownership, not a decoration on a function call. Prefer passing values into a worker when that is enough. If workers share mutable state, protect access with a channel or a synchronization primitive. Join every goroutine you start, and test concurrent code with the race detector on a machine that supports it. A fast program with an unclear owner is harder to operate than a slower program with a clear handoff.

- Prefer passing values to sharing pointers
- Protect shared mutation
- Join every goroutine you start
- The race detector finds conflicting access

Notes or exercise panel; no executable listing.
