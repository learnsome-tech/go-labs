# m01l02-05: Vet catches what the compiler allows

You can install the Go toolchain the way that suits a laptop, a build agent or a container, and use the go command's own subcommands for formatting, checking and documentation instead of reaching for extra tools.

This program compiles perfectly. It is also wrong: the format string asks for a decimal number and the argument is a string, so at runtime you get mangled output instead of a host name. The compiler has no opinion, because printing takes any arguments. Run vet over it and you get the file, the line, the verb and the type that does not match. Vet is a collection of checks for exactly this sort of thing: format strings that do not line up with their arguments, unreachable code, locks copied by value, struct tags that will not decode. It is part of the toolchain, it runs in about a second, and it belongs in your pipeline next to the tests.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
