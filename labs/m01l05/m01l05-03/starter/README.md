# m01l05-03: Every type has a zero value, and it is useful

You can declare variables both ways and say which to use where, predict the zero value of any type and design configuration structs around it, define constant sets with iota, and convert between numeric types without guessing.

Declare something without a value in Go and it is not undefined, and it is not a null you have to guard against. It is the zero value for its type, and the zero values were chosen to be useful: zero for numbers, the empty string, false for booleans, and nil for slices, maps and pointers. Print them all and notice the second line especially. That slice is nil, and its length is still zero rather than an error, so you can ask a nil slice how long it is and even append to it. A nil map can be read from safely and only panics when you write to it. And a struct with no values is a struct whose every field is its own zero.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
