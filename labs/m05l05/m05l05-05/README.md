# m05l05-05: The HTTP client habits

You can issue an HTTP request with a context, check the response, and close the body on every path.

Build HTTP requests with a context, send them through a client with a deliberate timeout, and separate transport errors from response status errors. Check the status before decoding and close every response body. A fake transport gives tests a deterministic boundary, while the production client can reuse connections and enforce the same cancellation contract.

- Attach context to requests
- Configure a timeout
- Check transport and status separately
- Close every response body

Notes or exercise panel; no executable listing.
