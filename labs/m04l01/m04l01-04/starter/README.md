# m04l01-04: Workers receive independent values

You can start goroutines, wait for their work, and recognise the ownership questions concurrency introduces.

Each loop value is passed as an argument to the goroutine, so the worker owns its own copy of the target name. The wait group counts both jobs and main waits for both to finish. Their scheduling order is not a contract, even though this small run may print the same order repeatedly. When output order matters, collect results through a coordination mechanism instead of relying on timing.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
