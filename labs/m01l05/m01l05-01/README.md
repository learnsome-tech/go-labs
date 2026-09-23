# m01l05-01: Static types, written down only when useful

You can declare variables both ways and say which to use where, predict the zero value of any type and design configuration structs around it, define constant sets with iota, and convert between numeric types without guessing.

Go is statically typed, so every variable has exactly one type for its whole life, and the compiler knows it. What Go avoids is making you write that type down when it is already obvious. Inside a function the short form with a colon and an equals sign declares and assigns at once, taking the type from the value, and that is what you will write ninety nine times in a hundred. The longer form with the var keyword is for package level declarations, where the short form is not allowed, and for the times you want a specific type rather than the inferred one. And as with imports, an unused local variable is a compile error. The compiler will not let you leave litter behind.

- Every variable has one type, fixed when it is declared
- Colon equals infers the type from the value: use it in functions
- The var keyword is for package level and for an explicit type
- An unused local variable is a compile error, not a warning

Notes or exercise panel; no executable listing.
