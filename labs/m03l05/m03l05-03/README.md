# m03l05-03: Subtests give each row a name

You can express many input cases in one table test and get a useful name for every failure.

Subtests let the testing package report each row as its own named check. The map provides two cases, and the loop opens a subtest with the row name. A failing row appears in the test output with that name, which makes a large table easier to search. For tables with more fields, a slice of structs is usually clearer than a map because it preserves order and lets each row carry all of its expected values.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
