# m03l02-03: Choosing a branch with errors Is

You can wrap an underlying error with context and test its identity with errors Is.

This caller treats a missing settings file as a normal first run. The wrapped error still carries the sentinel, so errors Is selects the use defaults branch. Any other cause would fall through to stop. The policy lives at the edge because a library should report what happened rather than decide whether a deployment may continue. When you add context with wrapping, you get a message for people and a stable identity for code.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
