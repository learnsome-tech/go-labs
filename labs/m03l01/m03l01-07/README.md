# m03l01-07: Try it yourself

You can return errors, inspect them, and keep failure handling explicit in a command line tool.

Write a function that accepts a service name and returns a short description plus an error. Reject an empty name with a useful message. Call the function once with a valid name and once with an empty name. In each case check the error first, then print either the description or a contextual failure line. Keep the function independent of exit codes so another caller could choose a different policy.

- Write a function that parses a non empty service name and returns an error otherwise.
- Call it with a valid and an empty name.
- Print the value on success and a contextual message on failure.

Notes or exercise panel; no executable listing.
