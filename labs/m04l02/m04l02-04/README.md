# m04l02-04: Ranging over a closed channel

You can send values through channels, close them deliberately, and use channels as ownership boundaries.

This buffered channel receives two values before the sender closes it. The range loop reads both queued values, then ends when the channel is empty and closed. The final line proves that close marks completion without losing data already in the buffer. Keep close on the producer side, and let consumers range when a stream of work has a natural end.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
