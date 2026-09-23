# m01l05-02: Both declaration forms, and what was inferred

You can declare variables both ways and say which to use where, predict the zero value of any type and design configuration structs around it, define constant sets with iota, and convert between numeric types without guessing.

One package level variable, outside every function, declared with the var keyword because that is the only form allowed there. Then inside, the short form four times over: a string, a whole number, a number with a decimal point. Notice what inference chose when we did not say: plain int for the whole number and the wider floating point type for the decimal. When we want a wider integer we have to say so, which is the line with the explicit type. The declaration with no value at all is the interesting one, and the next segment is about exactly that. Print the types and then the values, and read the first two lines carefully: those are the compiler's decisions, made visible.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
