# m05l01-03: A small file in one call

You can read arguments and environment, read and write files safely, tell a missing file from a broken one, and set the exit code your pipeline branches on.

For a config file, a lock file or a small report you do not need a stream at all. WriteFile takes a name, a slice of bytes and a permission mode, and creates or truncates it in a single call. That mode is worth a moment: octal six hundred means readable and writable by the owner and by nobody else, which is what you want for anything holding a token. Permissions only apply when the file is created, and the process umask can still clear bits, so check rather than assume. ReadFile is the mirror image and hands you the whole file as bytes. Converting to a string is free enough for configuration, and printing it back shows the round trip worked.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
