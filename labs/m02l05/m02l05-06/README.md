# m02l05-06: The safe path through an interface

You can inspect an interface value safely, use any for truly mixed data, and avoid hiding useful types behind it.

An interface value records a dynamic type and value, while its static type controls what you can call directly. Use a comma ok assertion when one concrete type is expected, and a type switch when a small known set is valid. Any makes a boundary flexible, but every assertion is a place where bad input can arrive. Narrow that input quickly into a typed struct or a small interface, then let the compiler help with the rest of the tool.

- Use comma ok assertions when a type may be wrong
- Use type switches for a small known set
- Prefer typed structs in core logic
- Keep any at genuine data boundaries

Notes or exercise panel; no executable listing.
