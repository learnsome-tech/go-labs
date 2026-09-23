# m01l01-06: What to carry forward

You can explain why platform and DevOps teams write their tooling in Go, compile and run a program with the go command, and describe exactly what the compiler hands you at the end of a build.

So far: a Go program is a package named main containing a function named main. The build command turns it into one executable for one platform, with everything it needs already inside, and prints nothing when it succeeds. The compiler is strict in ways that feel abrupt at first and pay for themselves later, and an unused import failing the build is your first taste of that. Next lesson we install the toolchain properly and look at what the go command can do besides build, because almost everything you need in daily work is one subcommand rather than a separate tool you have to choose.

- Package main plus a main function makes an executable
- Build prints nothing when it succeeds
- One binary, one target platform, no runtime to install
- Unused imports are errors, not warnings

Notes or exercise panel; no executable listing.
