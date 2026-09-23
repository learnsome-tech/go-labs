# m06l02-05: The server habits

You can expose a small HTTP handler, return useful status codes, and configure a server with a timeout.

Handlers translate requests into status and body responses. ServeMux keeps routing explicit, and a configured server makes timeouts visible. Use separate liveness and readiness meanings, then test handler behaviour with a recorder before you involve a real listener. The next lesson gives that listener a safe lifetime under an orchestrator.

- Keep handlers small
- Route with ServeMux
- Set explicit service limits
- Test handlers without a listener

Notes or exercise panel; no executable listing.
