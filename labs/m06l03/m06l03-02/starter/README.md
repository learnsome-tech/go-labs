# m06l03-02: Waiting for SIGTERM

You can stop accepting traffic on SIGTERM, cancel active work, and let a Go service finish cleanly.

The signal channel receives SIGTERM from the operating system. Main announces ready, then blocks until the signal arrives. The stopping line is the start of shutdown work, such as calling server Shutdown and cancelling a root context. A buffered channel prevents the signal sender from waiting while main is between setup and receive.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
