# m06l04-02: The build command

You can build a portable Go binary with reproducible flags and inspect the result before shipping it.

The build script targets Linux and a common server architecture, disables cgo, and removes local path details from the binary. The linker flags trim symbols that a production image does not need. The file command then gives a quick inspection point for the artifact. A real pipeline should also record the Go version and checksum beside the binary.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
