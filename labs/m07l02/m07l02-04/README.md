# m07l02-04: A result carries identity

You can run checks concurrently, collect one result per target, and preserve a useful report order.

Result keeps the target identity beside the health decision. The final report can therefore explain which target passed and which target failed without consulting a separate map. The checker workers will create these values, and the coordinator will collect them for JSON output in a later lesson.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
