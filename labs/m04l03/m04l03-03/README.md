# m04l03-03: Default is a deliberate non wait

You can wait on several channel operations and keep a concurrent worker responsive.

A default case runs immediately when no channel operation is ready. That makes a poll, but a loop around a default can consume a core if it has no pause or useful work. Use it for a bounded inspection, or prefer a blocking select when waiting is the correct behaviour. Concurrency code should spend time doing work or waiting for a reason, not repeatedly asking an empty channel the same question.

- A default case runs immediately
- Use it for a bounded poll
- Avoid spinning without a pause
- Prefer blocking when work can wait

Notes or exercise panel; no executable listing.
