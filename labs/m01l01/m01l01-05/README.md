# m01l01-05: The compiler refuses an unused import

You can explain why platform and DevOps teams write their tooling in Go, compile and run a program with the go command, and describe exactly what the compiler hands you at the end of a build.

One early surprise, and it says a lot about the language. This program imports two packages but only uses one. In most languages that is a warning you learn to ignore. Try to build that in Go and the compiler refuses outright: the import of o s is there and unused, so the build fails with the file, the line and the column. The reasoning is deliberate. An unused import is either a leftover from code you deleted or a mistake, and in both cases it misleads the next person reading the file. The language would rather stop you now than let dead references pile up in a tool your whole team depends on.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
