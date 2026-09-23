# m02l03-04: Choose a pointer receiver for mutation

You can model tool data with structs, attach methods, and choose value or pointer receivers intentionally.

Receiver choice follows the same question as function arguments. A value receiver reads a copy, while a pointer receiver can change the original struct. Use a pointer receiver when the method updates state, when copying would be expensive, or when the type already has pointer based behaviour. Keep the choice consistent across the methods on one type so callers do not have to remember two mental models. For an addressable value, Go will take its address for the method call when the receiver requires it.

- Value receivers read a copy
- Pointer receivers can change the original
- Keep receiver choice consistent across a type
- Go automatically takes an address for addressable values

Notes or exercise panel; no executable listing.
