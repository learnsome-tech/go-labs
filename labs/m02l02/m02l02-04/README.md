# m02l02-04: Zero values make empty collections safe

You can choose arrays, slices, and maps for predictable collections and handle their zero values safely.

The zero values for collections are friendly in one direction and sharp in another. Appending to a nil slice works, so a function can build a result without a special initialization line. Reading from a nil map returns the value zero and says the key was absent. Writing to a nil map is different: it panics, because there is no table to receive the entry. Make a map with make before storing into it, and let nil slices represent an empty result when that is useful to your caller.

- Appending to a nil slice is safe
- Reading a nil map returns a zero value
- Writing to a nil map panics
- Make a map before storing into it

Notes or exercise panel; no executable listing.
