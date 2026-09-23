# m05l03-01: Time has two useful meanings

You can measure elapsed work, schedule ticks, and format timestamps without confusing a duration with a wall clock time.

The time package represents both a moment and an interval. Time names when something happened, while Duration names how long it took. Use a monotonic measurement such as Since when you care about elapsed work, because wall clocks can jump. Use UTC for timestamps that cross machines and regions, then format them at the display boundary for people.

- Time names a moment
- Duration names an interval
- Since measures elapsed work
- Use UTC for portable output

Notes or exercise panel; no executable listing.
