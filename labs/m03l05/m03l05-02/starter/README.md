# m03l05-02: The first table test

You can express many input cases in one table test and get a useful name for every failure.

The cases slice carries a name, an input status code, and the expected health result. The loop feeds each code to the same helper and reports the case name when the result differs. One test function now covers success, redirect, and server error behaviour without three copies of the assertion. As the policy grows, add a row and keep the code that explains the rule in one place.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
