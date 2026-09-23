# m04l03-01: Select waits on several cases

You can wait on several channel operations and keep a concurrent worker responsive.

Select waits on several channel operations and runs one case that is ready. If more than one is ready, the choice is deliberately unspecified, so do not use case order as a priority rule. A default case makes the select non blocking, which is useful for a quick poll but can also create a busy loop. In tooling code, select is the bridge between work, a timer, and a cancellation signal.

- Each case is a channel operation
- A ready case runs without priority
- A default case avoids waiting
- Use select for time and cancellation

Notes or exercise panel; no executable listing.
