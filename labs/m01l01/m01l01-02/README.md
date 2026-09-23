# m01l01-02: A first program, top to bottom

You can explain why platform and DevOps teams write their tooling in Go, compile and run a program with the go command, and describe exactly what the compiler hands you at the end of a build.

Here is the whole of a Go program. Every file starts by naming the package it belongs to, and the name main is special: it marks a program rather than a library, something that can be built into an executable. Then the imports, one per line or in a block, naming the packages this file uses. Here that is f m t, the formatting package, which is where printing lives. Last comes the function named main, which is where a Go program begins. Two calls to Println, each printing one line. Let us run it and see those two lines come back.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
