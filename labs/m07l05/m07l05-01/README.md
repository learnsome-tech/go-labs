# m07l05-01: The finished tool has three contracts

You can test the finished checker, build a small runtime image, and connect its exit status to an orchestrator.

The finished health checker has three contracts. Tests protect its checking and reporting behaviour. The static binary is the release artifact, and the container image defines the runtime assumptions around it. Finally, the exit code and health endpoint give an orchestrator a simple signal while the JSON report preserves detail for a pipeline.

- Tests protect checker behaviour
- The binary is the release artifact
- The image defines runtime assumptions
- The probe and exit code serve orchestration

Notes or exercise panel; no executable listing.
