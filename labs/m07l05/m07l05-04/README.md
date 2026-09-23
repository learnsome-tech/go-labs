# m07l05-04: A release command sequence

You can test the finished checker, build a small runtime image, and connect its exit status to an orchestrator.

The release sequence runs tests first, builds a Linux binary with cgo disabled, and then builds the image from that artifact. Docker would execute the final step when a daemon is available. On a machine without Go or Docker, the sequence remains an accurate transcript of the release contract and is marked accordingly.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain and Docker are unavailable on this machine)
