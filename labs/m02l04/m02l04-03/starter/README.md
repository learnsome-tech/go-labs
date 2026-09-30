# m02l04-03: Two concrete types, one function

You can define small interfaces, satisfy them without declarations, and use that flexibility to test and compose tooling code.

Here two unrelated concrete types share one small contract. Healthy and Failed each provide Report, so printReport can accept either value without a type switch or a base class. The function depends on the behaviour it uses and nothing more. In a health checker that lets a real network probe and a deterministic test probe share the same reporting path. The interface stays small because the caller only needs one method, and that makes every implementation easy to inspect.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
