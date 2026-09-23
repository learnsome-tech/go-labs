# m01l02-08: Try it yourself

You can install the Go toolchain the way that suits a laptop, a build agent or a container, and use the go command's own subcommands for formatting, checking and documentation instead of reaching for extra tools.

Three small things to do before the next lesson. Print your toolchain version and read the usage line for the test subcommand, so the help habit starts today. Then take a working file, wreck its indentation and spacing deliberately, confirm the formatter lists it, and let the formatter fix it rather than fixing it by hand. Finally write a print call whose format verb does not match the argument you gave it, compile it to prove the compiler is happy, and then run vet and read the complaint. That contrast between what compiles and what vet objects to is the thing worth internalising.

- Print the version, then read the usage line of the test subcommand.
- Break the formatting of a file on purpose and fix it with the formatter.
- Write a print call whose format verb does not match its argument, then run vet.

Notes or exercise panel; no executable listing.
