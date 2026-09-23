# m02l01-05: Checking an optional pointer

You can read pointer syntax, mutate shared state deliberately, and recognise a nil pointer before it becomes a crash.

This small helper makes the safe shape visible. The caller starts with a nil pointer, and describe checks it before trying to read anything. That path prints unset and returns. Then we create a real string and pass its address, so the second path can safely follow the pointer and print the value. In a service, this check is the difference between a missing optional setting and a process that panics during startup. Keep the guard close to the dereference so a future edit cannot accidentally move the safety boundary.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
