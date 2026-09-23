# m06l02-03: Status codes are part of the API

You can expose a small HTTP handler, return useful status codes, and configure a server with a timeout.

A response status is part of the service API. WriteHeader sends an explicit status, while the first body write defaults to success. Use a client error status for invalid input and a server error status for a dependency failure. Keep a liveness endpoint about process health and a readiness endpoint about whether this instance should receive traffic.

- WriteHeader sends the status
- Write defaults to success
- Return four hundred for bad input
- Keep health and readiness distinct

Notes or exercise panel; no executable listing.
