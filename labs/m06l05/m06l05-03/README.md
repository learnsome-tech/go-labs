# m06l05-03: Minimal images change debugging

You can copy a static binary into a minimal image and choose the runtime files a service actually needs.

Minimal images reduce attack surface but remove familiar debugging tools. There may be no shell to exec into, so the service must write useful logs to standard output and error. Put health checks in the orchestrator or a sidecar, and keep a separate debug image for investigation rather than adding a shell to every production release. The image contract should remain small and explicit.

- There may be no shell to exec
- Logs must go to standard streams
- Health checks belong outside the image
- Keep a debug image for investigation

Notes or exercise panel; no executable listing.
