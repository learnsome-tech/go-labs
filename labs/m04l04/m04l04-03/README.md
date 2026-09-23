# m04l04-03: Cancellation is cooperative

You can pass context through a call chain and use its deadline and values at the right boundaries.

Cancellation is cooperative. Calling the cancel function closes the context Done signal, but a worker must select on that signal and return when it sees it. The context error explains whether the deadline expired or the caller cancelled. Create a derived context for a bounded operation, and always call its cancel function so timers and children can be released promptly.

- Cancel closes the Done signal
- Functions must observe Done
- Return ctx Err when stopping
- Always call the cancel function

Notes or exercise panel; no executable listing.
