# m03l01-03: Handling a failed operation

You can return errors, inspect them, and keep failure handling explicit in a command line tool.

This call takes the failure branch because the name is empty. The function returns an empty string and an error created with formatted text. The caller prints a useful context line and returns before touching the empty result. Returning early keeps the happy path at the bottom of the function and makes the failure policy easy to follow. For a command line program, that same branch could write to standard error and choose a non zero exit status.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
