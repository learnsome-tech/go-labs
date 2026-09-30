# m01l05-06: Conversions are explicit, always

You can declare variables both ways and say which to use where, predict the zero value of any type and design configuration structs around it, define constant sets with iota, and convert between numeric types without guessing.

Last piece of the module, and the one that catches everybody arriving from a dynamic language. Go will not mix numeric types for you. A whole number and a number with a decimal point cannot be multiplied together, even though the machine could do it perfectly well, and there is no silent widening. You convert, in writing, every time. Run those three lines and read the middle one closely: converting the decimal to a whole number truncates it towards zero, so a factor of one point two five becomes one, and the multiplication quietly does nothing. That is not a bug in Go, it is arithmetic you asked for, and it is exactly the kind of mistake that produces a capacity report nobody trusts. Convert deliberately, and convert late.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
