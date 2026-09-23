# m06l02-01: A handler is a function with a contract

You can expose a small HTTP handler, return useful status codes, and configure a server with a timeout.

The net HTTP server starts with a handler contract. A request carries the method and path, while a response writer sends the status and body. ServeMux routes paths to handlers, and an HTTP server groups that handler with timeouts and an address. Put limits on a real service rather than relying on defaults, so a slow client cannot hold resources forever.

- Request carries method and path
- ResponseWriter sends the status and body
- ServeMux routes paths
- Server fields make limits explicit

Notes or exercise panel; no executable listing.
