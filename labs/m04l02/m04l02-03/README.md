# m04l02-03: Close the sending side

You can send values through channels, close them deliberately, and use channels as ownership boundaries.

Closing belongs to the side that knows production is complete, usually the sender. A receiver can range over the channel and stop naturally when it is closed. Sending after close panics, and closing twice panics, so make one component responsible for that lifecycle. A channel is a handoff with an owner, and the close is the owner's promise that the stream has ended.

- Only the sender should close
- Receivers can range until close
- A send after close panics
- Closing twice also panics

Notes or exercise panel; no executable listing.
