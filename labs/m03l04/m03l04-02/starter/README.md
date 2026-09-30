# m03l04-02: A first test function

You can write focused Go tests with the standard testing package and read failures as feedback about behaviour.

This test calls the function like any other Go code and checks the result with an if statement. Fatalf records a formatted failure and stops this test when the value is wrong. The test name says what behaviour it protects, while the message includes the value that arrived. Running the package with go test produces an ok result when the check passes. The same command will show the file and line when the expectation is broken.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
