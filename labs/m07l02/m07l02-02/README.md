# m07l02-02: Collecting concurrent results

You can run checks concurrently, collect one result per target, and preserve a useful report order.

The buffered result channel can hold both worker results while the coordinator waits. Each worker receives its own target name, sends a report, and marks its work done. After the wait group reaches zero, the coordinator closes the channel and ranges over the complete stream. In a production report, sort or index results if output order must match configuration order.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
