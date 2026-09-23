# m05l05-03: Status and body are separate signals

You can issue an HTTP request with a context, check the response, and close the body on every path.

An HTTP response with a server error status is still a successful transport exchange, so the request call may return a response and a nil error. Check the status code explicitly before decoding the body as the success schema. Once a response exists, close its body when you are done. Keeping transport, status, and decoding checks separate makes an API failure easy to explain in a health report.

- A response can arrive with an error status
- Check the status code explicitly
- Close the body after a response
- Decode only an accepted body

Notes or exercise panel; no executable listing.
