# m06l05-02: A scratch Dockerfile

You can copy a static binary into a minimal image and choose the runtime files a service actually needs.

The first stage contains the compiler and source. The second stage starts from scratch and copies only the resulting executable. That keeps source code, package managers, and compiler tools out of the runtime surface. A real service that calls HTTPS would also copy a certificate bundle, and one that needs a user lookup would choose a base with that data.



Run: `sh run.sh`.

Verification: noVerify (Docker daemon is unavailable on this machine)
