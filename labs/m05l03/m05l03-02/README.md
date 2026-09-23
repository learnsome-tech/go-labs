# m05l03-02: Measuring elapsed work

You can measure elapsed work, schedule ticks, and format timestamps without confusing a duration with a wall clock time.

The program records a start time, waits briefly, and asks how much time has elapsed. Since returns a duration, and the comparison proves that the work took a positive amount of time. In production you would record the duration for a metric or a log line rather than print a boolean. Use elapsed time for latency and deadlines, not a formatted timestamp subtraction.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
