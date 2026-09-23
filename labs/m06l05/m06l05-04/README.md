# m06l05-04: A distroless runtime

You can copy a static binary into a minimal image and choose the runtime files a service actually needs.

This variant uses a distroless static image and selects the nonroot user supplied by that image. The compiler remains in the build stage, and the runtime receives only the application. Docker would build this file when the daemon and image access are available; on a restricted machine it remains a transcript that still shows the complete release shape.



Run: `sh run.sh`.

Verification: noVerify (Docker daemon is unavailable on this machine)
