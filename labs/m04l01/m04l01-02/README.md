# m04l01-02: Starting work with go

You can start goroutines, wait for their work, and recognise the ownership questions concurrency introduces.

The anonymous function starts as a goroutine. A wait group records that one piece of work exists, and the deferred Done call tells the group when the worker returns. Main waits before printing its own line, so the process cannot exit early. The output shows the worker completing before main finishes. A wait group is useful when a caller starts a known number of independent jobs and needs to join them before returning a result.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
