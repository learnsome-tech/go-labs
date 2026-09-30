# m06l01-02: Declaring typed options and parsing them

You can build a command line tool with the standard flag package that takes typed options, reads positional arguments, falls back to environment variables, and fails with a usage message and the right exit code.

Here is the shape of every tool built on this package. After the imports, each option is declared before parsing, and the declaration returns a pointer to the value that parsing will fill in later. We declare each option in turn: a string for the address, an integer for the retry count, a boolean for verbosity, and a duration for the deadline on each attempt. Notice the duration: the package turns two seconds or five hundred milliseconds into a real duration value for you, which is exactly the arithmetic you would hand roll in Bash. Then we call Parse, which walks the arguments and fills every pointer. Only after that are the values safe to read, so we print what we parsed. Let us run it with values for three options and let the fourth fall back to its default.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
