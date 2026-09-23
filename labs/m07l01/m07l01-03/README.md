# m07l01-03: Configuration belongs at the edge

You can start the health checker with typed configuration, clear defaults, and a small target model.

Parse flags and environment values once in main, then pass typed configuration into the checker. Reject empty names and URLs before starting work. Keeping configuration at the edge prevents every worker from inventing its own defaults and makes the project predictable inside a container or pipeline.

- Parse once in main
- Pass typed values inward
- Reject empty names and URLs
- Keep environment fallback explicit

Notes or exercise panel; no executable listing.
