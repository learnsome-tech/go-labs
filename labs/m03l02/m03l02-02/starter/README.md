# m03l02-02: Wrapping a missing resource

You can wrap an underlying error with context and test its identity with errors Is.

The sentinel error names a condition that callers may need to recognise. Load wraps it with the operation read config and the percent w verb, which preserves the cause in the chain. Printing the error gives an operator useful context. The errors Is call still finds the original sentinel through that wrapper and prints true. A string comparison would be brittle because the human message can change while the condition remains the same.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
