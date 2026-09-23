# m06l05-05: The container habits

You can copy a static binary into a minimal image and choose the runtime files a service actually needs.

Multi stage Dockerfiles keep compilers out of production. Scratch gives the smallest surface but no runtime data, while distroless supplies a curated base and a nonroot user. Copy only what the binary needs, send logs to standard streams, and verify the image in the same environment that will run it. That closes the service delivery path from source to orchestrator.

- Compile in a builder stage
- Copy only runtime inputs
- Choose scratch or distroless deliberately
- Keep production images observable

Notes or exercise panel; no executable listing.
