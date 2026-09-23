# m04l04-01: Context carries a request lifetime

You can pass context through a call chain and use its deadline and values at the right boundaries.

Context carries the lifetime of a request or operation through a call chain. Pass it as the first parameter, and let each layer observe cancellation or a deadline. A context can also carry request scoped metadata, but it should not become a general bag of configuration. The context package gives a worker a shared signal for stopping when the caller no longer needs the result.

- Pass context as the first parameter
- A context can be cancelled
- Deadlines describe a maximum wait
- Values are for request scoped metadata

Notes or exercise panel; no executable listing.
