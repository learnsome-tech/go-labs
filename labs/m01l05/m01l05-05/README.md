# m01l05-05: Constants and iota for a set of states

You can declare variables both ways and say which to use where, predict the zero value of any type and design configuration structs around it, define constant sets with iota, and convert between numeric types without guessing.

Go has no enumerated type, so this is the idiom that replaces one. First a named type based on a whole number, which gives the compiler something to check. Then the constant block, where the counter identifier counts from zero down the block, one per line, and the type and the assignment carry down with it. Notice that the first constant is the zero value of the type, which is why unknown deserves that slot rather than healthy: an unset state should not read as fine. Constants are compile time values, so there is no memory and no lookup, and a mistyped comparison is a build failure. Compare them and print one, and remember that printing shows the number until we teach the type to describe itself, which is a method, and methods are next module.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
