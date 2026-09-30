# m06l04-04: Inspecting build metadata

You can build a portable Go binary with reproducible flags and inspect the result before shipping it.

The runtime package exposes the operating system and architecture selected for this build. On the checking machine the example reports its local pair. A cross compiled release would report Linux and the target architecture when it runs there. Keep this metadata in diagnostics when operators need to confirm which artifact reached a host.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
