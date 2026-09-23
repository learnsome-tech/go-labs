# m05l02-04: Reading a bounded prefix

You can copy streams, read bounded input, and treat files and network bodies through common interfaces.

The limit reader wraps a longer source and exposes only its first four bytes. ReadAll is now safe for this deliberately bounded stream, and the output proves that the remaining input was not consumed. In a tool that accepts a response body, choose the bound from the protocol or an operational limit and return an error when the input is larger than your contract allows.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
