# m07l01-05: Project setup

You can start the health checker with typed configuration, clear defaults, and a small target model.

The project begins at the edge with flags, environment fallback, a typed target struct, and validation before work starts. Once configuration is valid, workers can focus on checking targets rather than interpreting operator input. That separation will make the concurrent checker easier to test and easier to shut down.

- Parse configuration at the edge
- Model targets with a struct
- Validate before concurrency
- Pass typed values to workers

Notes or exercise panel; no executable listing.
