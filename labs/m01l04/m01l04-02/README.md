# m01l04-02: Watching the three verbs behave

You can choose correctly between running, building and installing a program, write functions that return a value alongside an error, format output with the verbs you will use daily, and set an exit code that a pipeline can act on.

Same program, three ways. Compile and throw away first: you get the output and nothing is left behind. Then keep the result with an output name, which prints nothing and leaves an executable. Run that executable and the output is identical, but this time no toolchain is involved at all, which is exactly what happens on the host you deploy to. The last one is a trap worth meeting now: extra words after the package are passed to your program, not to the compiler, so the compiler will not complain about them and your program has to decide what they mean. We will read them properly when we get to the flag package.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
