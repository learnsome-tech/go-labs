# m05l03-05: The time habits

You can measure elapsed work, schedule ticks, and format timestamps without confusing a duration with a wall clock time.

Time names moments, while Duration names intervals. Measure latency with Since, format shared timestamps in UTC, and stop timers or tickers when their work ends. A context deadline usually makes request lifetime explicit. These choices keep logs comparable across hosts and prevent a tool from waiting longer than its operational contract.

- Use Duration for elapsed work
- Use UTC for shared timestamps
- Stop timers and tickers
- Prefer context for request lifetime

Notes or exercise panel; no executable listing.
