# m02l03-06: Structs become useful when behaviour joins data

You can model tool data with structs, attach methods, and choose value or pointer receivers intentionally.

Structs give related data a named shape, and methods give that shape a vocabulary. Use a value receiver when the operation only reads a small value. Use a pointer receiver when the operation changes state or copying would be costly, and keep receiver choices consistent across the type. Named fields make configuration readable at construction sites. Together, these habits turn a pile of flags and local variables into a small model that your tool can pass around safely.

- Struct fields describe a stable data shape
- Methods keep operations beside that shape
- Use value receivers for read only behaviour
- Use pointer receivers when state changes

Notes or exercise panel; no executable listing.
