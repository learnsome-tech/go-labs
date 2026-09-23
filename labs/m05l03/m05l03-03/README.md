# m05l03-03: Timers and tickers have lifetimes

You can measure elapsed work, schedule ticks, and format timestamps without confusing a duration with a wall clock time.

A timer sends one event after a duration, while a ticker sends repeated events. Both hold resources, so stop them when the operation ends, especially in a long running service. A ticker is useful for periodic refresh or sampling. When the lifetime belongs to a request, a context deadline often communicates the cancellation contract more clearly than a timer hidden inside a helper.

- A timer fires once
- A ticker sends repeated ticks
- Stop them when work ends
- Context is often clearer for cancellation

Notes or exercise panel; no executable listing.
