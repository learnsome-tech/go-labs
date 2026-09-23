# m02l05-04: Use any at the edge, not everywhere

You can inspect an interface value safely, use any for truly mixed data, and avoid hiding useful types behind it.

Any is a useful escape hatch, not a default design. It accepts every type, which makes it convenient for decoded data, logging fields, or a boundary that truly has mixed values. Inside the core of a tool, typed structs and small interfaces preserve the checks that catch mistakes before runtime. Each assertion is a decision point where an unexpected type can appear, so handle that possibility and return a useful error. Let wide values enter at the edge and become specific quickly.

- any accepts every type
- Assertions move back to a concrete type
- Typed structs preserve useful compiler checks
- A wide value is harder to reason about

Notes or exercise panel; no executable listing.
