# m06l05-01: The runtime image only needs runtime files

You can copy a static binary into a minimal image and choose the runtime files a service actually needs.

A multi stage build compiles in a Go image, then copies only the binary and required runtime files into the final image. Scratch is nearly empty, so it has no shell, certificate store, or user database. Distroless adds a small curated runtime without a package manager. Choose the image from the files your service needs, and run as a non root user when the workload allows it.

- Multi stage builds keep compilers out
- Scratch has no shell or certificates
- Distroless adds selected runtime support
- Run as a non root user when possible

Notes or exercise panel; no executable listing.
