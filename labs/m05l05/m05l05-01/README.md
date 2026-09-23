# m05l05-01: An HTTP client has a full lifecycle

You can issue an HTTP request with a context, check the response, and close the body on every path.

An HTTP client call has a lifecycle that should be visible in code. Build a request with the caller's context, send it through a client, check the response status before trusting the body, and close the body on every successful response. The default client is fine for a small tool, while a long running service should configure timeouts and transport reuse deliberately.

- Build a request with context
- Send it through a client
- Check status before decoding
- Close the response body

Notes or exercise panel; no executable listing.
