# m02l02-07: Try it yourself

You can choose arrays, slices, and maps for predictable collections and handle their zero values safely.

Build a small inventory tool. Start with a nil slice of service names and append three values. Create a map from each service name to its port, then print the ordered names and look up one service that is missing. Use the second result from the lookup so a missing entry cannot be confused with a real zero port. Add one more service after the first print and observe which collection preserves order and which one answers the lookup.

- Collect service names in a nil slice and append three values.
- Build a map from service name to port with make.
- Look up a missing service and print a clear absent message.

Notes or exercise panel; no executable listing.
