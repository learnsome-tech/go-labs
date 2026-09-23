# m02l04-05: A tiny interface makes testing easy

You can define small interfaces, satisfy them without declarations, and use that flexibility to test and compose tooling code.

This loader has a boundary that is deliberately tiny. It needs a Reader with one method, and Fixture supplies a fixed value for the example. A test can pass the same kind of small fake without opening a file or making a network call. Production code can supply a different type with the same Read method. The loader does not know or care which one arrived, and that is the point of an implicit interface: replaceable behaviour without an inheritance hierarchy.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
