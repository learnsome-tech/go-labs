# m04l02-05: The channel rules

You can send values through channels, close them deliberately, and use channels as ownership boundaries.

Channels transfer typed values and can synchronize the handoff. Unbuffered channels wait for a receiver, while buffered channels provide a bounded queue. The sender owns closing, and the receiver can range until the close arrives. Treat the channel lifecycle as part of the design, because a send or close at the wrong time becomes a panic.

- Use channels to transfer ownership
- Buffer only a known amount
- The sender closes the stream
- Range stops after close

Notes or exercise panel; no executable listing.
