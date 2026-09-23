# m02l05-01: An interface value has two parts

You can inspect an interface value safely, use any for truly mixed data, and avoid hiding useful types behind it.

An interface value carries two things: a dynamic type and a dynamic value. The static type tells the compiler which methods you may call directly, while the dynamic pair records what concrete value arrived. An interface with no value is nil, and it has neither part. This model explains why an interface can hold a string, a struct, or a pointer, yet still expose only the methods in its contract. When you need the concrete value, use a type assertion or a type switch.

- It carries a dynamic type
- It carries a dynamic value
- A nil interface has neither
- The static interface type limits direct operations

Notes or exercise panel; no executable listing.
