# m02l02-03: Maps need a comma ok check

You can choose arrays, slices, and maps for predictable collections and handle their zero values safely.

A map returns a value and a boolean when you ask for the comma ok form. That second result matters because a missing integer key would otherwise look like a real zero. Here the lookup asks for ssh, sees that the key is absent, and prints a useful message before returning. This is the same habit you use with environment settings and filesystem errors: preserve the difference between missing and present with an empty value. Use a map for lookup, and keep that presence check at the edge.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
