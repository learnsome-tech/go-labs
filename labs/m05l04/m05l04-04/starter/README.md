# m05l04-04: Encoding a small report

You can decode API responses into structs, encode reports, and reject malformed input clearly.

Marshal turns the typed report into compact JSON. The tags define the keys that a pipeline or dashboard will consume, and the boolean remains a JSON boolean rather than a quoted string. The error is nil because the struct contains values JSON can encode. Keep this final conversion at the output boundary so internal names can evolve without changing the report accidentally.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
