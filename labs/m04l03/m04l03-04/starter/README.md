# m04l03-04: A non blocking poll

You can wait on several channel operations and keep a concurrent worker responsive.

The buffered channel has no value waiting, so the receive case is not ready. The default case runs immediately and reports no update. This is a useful one time poll, such as checking whether a status event has already arrived before continuing. For repeated polling, add a ticker or a sleep so the loop yields time to other work.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
