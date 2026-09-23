# m03l04-05: Testing an error result

You can write focused Go tests with the standard testing package and read failures as feedback about behaviour.

This example checks the error side of a result instead of only checking success. The empty input must fail, so the test asserts that the error is not nil. In your own code use a real package error value rather than inventing one in a test. The important shape is the same: give the function input at the boundary, then assert the observable result and error contract.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
