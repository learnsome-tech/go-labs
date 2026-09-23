# m07l05-03: Build and run the same artifact

You can test the finished checker, build a small runtime image, and connect its exit status to an orchestrator.

Build the release artifact after tests pass, then copy that same binary into the runtime image. Do not rebuild different source in a later packaging step. Verify the container command, its user, and the health probe in an environment that matches the orchestrator. This closes the gap between what the pipeline tested and what the cluster runs.

- Compile once for the target
- Run tests before packaging
- Copy the binary into the image
- Verify the container command and probe

Notes or exercise panel; no executable listing.
