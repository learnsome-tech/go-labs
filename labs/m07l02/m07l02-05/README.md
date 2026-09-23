# m07l02-05: The concurrent checker

You can run checks concurrently, collect one result per target, and preserve a useful report order.

Run independent checks concurrently, send typed results through a channel, and wait for every worker before closing that channel. Keep target identity in each result. If the report needs configuration order, collect and sort explicitly instead of depending on scheduling.

- Start one worker per target
- Send typed results
- Wait before closing the channel
- Sort when report order matters

Notes or exercise panel; no executable listing.
