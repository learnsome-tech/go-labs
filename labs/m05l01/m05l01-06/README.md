# m05l01-06: Exit codes are the contract with CI

You can read arguments and environment, read and write files safely, tell a missing file from a broken one, and set the exit code your pipeline branches on.

Your exit status is an interface, even though it does not look like one. A pipeline step, a Kubernetes probe and a shell script with the dash e flag all read that number and nothing else. Name them at the top of the file rather than scattering bare numbers through the code. Zero means the check passed. One means the check itself failed, which is a result, not a crash. Two here means the caller used the tool wrongly, and keeping those apart lets a pipeline retry one and never the other. One warning: o s dot Exit stops the program immediately and skips every deferred call, so flush and close before you reach it. Run it and the shell reports it.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
