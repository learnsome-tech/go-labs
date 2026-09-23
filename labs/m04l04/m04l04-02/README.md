# m04l04-02: Passing a context through a function

You can pass context through a call chain and use its deadline and values at the right boundaries.

The context starts from a background root and carries one request scoped value into report. The function accepts context first, so a caller can later add cancellation or a deadline without changing the shape. Values should identify metadata that belongs to this request, such as a trace label. Keep ordinary configuration in explicit parameters or a typed struct instead of hiding it in context.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
