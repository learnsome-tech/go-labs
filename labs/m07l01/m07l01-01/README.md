# m07l01-01: Start with a configuration contract

You can start the health checker with typed configuration, clear defaults, and a small target model.

The health checker starts with a configuration contract. Flags describe the operator input, defaults make a local run easy, and a target struct gives each check a typed shape. Validate that configuration before starting goroutines or opening network connections. Early failure is easier to explain than a worker discovering a bad host halfway through a run.

- Flags describe operator input
- Defaults make local runs easy
- A target is a typed value
- Validation fails before work starts

Notes or exercise panel; no executable listing.
