# m05l03-04: Formatting a UTC timestamp

You can measure elapsed work, schedule ticks, and format timestamps without confusing a duration with a wall clock time.

Date constructs a known moment in UTC, and Format renders it with the standard RFC timestamp layout. The layout is a reference time written with words that look unusual at first, but the constant gives you the correct portable form. Keep UTC in logs and machine to machine data, then convert to a local zone only when a person needs that view.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
