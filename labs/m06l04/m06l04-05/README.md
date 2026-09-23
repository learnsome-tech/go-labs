# m06l04-05: The binary build habits

You can build a portable Go binary with reproducible flags and inspect the result before shipping it.

Build a self contained binary with explicit target settings, disable cgo when your dependency graph permits it, and keep reproducible flags in the pipeline. Inspect the artifact, record its checksum and Go version, and run it in the target image before release. The next lesson puts that binary in minimal container images.

- Build with explicit target settings
- Disable cgo when possible
- Inspect and checksum the artifact
- Test inside the target image

Notes or exercise panel; no executable listing.
