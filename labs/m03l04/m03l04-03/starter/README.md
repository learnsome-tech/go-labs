# m03l04-03: A test can report several checks

You can write focused Go tests with the standard testing package and read failures as feedback about behaviour.

Error records a failure and lets the current test continue, which is useful when several independent expectations should be reported together. This test checks both a successful response code and a server error. The helper keeps the policy in one place, and the test names the two promises it relies on. Prefer a small assertion message that tells the reader what the code should have done, rather than repeating the entire implementation.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
