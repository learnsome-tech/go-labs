# m02l01-03: Passing a pointer to a function

You can read pointer syntax, mutate shared state deliberately, and recognise a nil pointer before it becomes a crash.

A function receives a copy of each argument, including a pointer. That copy still points at the same data, so the function can change the caller's value without returning a replacement. The parameter says this plainly: addTag accepts a pointer to a string. Inside, it follows the pointer, adds a suffix, and stores the result. The caller passes its address, then prints the changed status. This pattern is useful for a configuration object or a counter that several operations must update, while keeping the shared state visible in the function signature.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
