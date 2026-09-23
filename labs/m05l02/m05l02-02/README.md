# m05l02-02: Copying a stream

You can copy streams, read bounded input, and treat files and network bodies through common interfaces.

The source buffer is a Reader and the target buffer is a Writer. io Copy moves bytes between them and returns the count plus an error. The function does not know that both ends are memory buffers; the same call can connect a file to standard output or an HTTP body to a report file. The example prints ten bytes, a nil error, and the copied text.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
