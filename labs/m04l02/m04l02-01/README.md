# m04l02-01: A channel is a typed handoff

You can send values through channels, close them deliberately, and use channels as ownership boundaries.

A channel is a typed handoff between goroutines. An unbuffered send waits until another goroutine receives, which creates synchronization at the exchange. A buffered channel can hold a bounded number of values before a sender waits. Closing a channel says that no more values will arrive; it does not discard values already waiting. The receiver can use the comma ok result or a range loop to notice the close.

- Channels carry values between goroutines
- An unbuffered send waits for a receiver
- A buffered channel holds a bounded queue
- Close says no more values will arrive

Notes or exercise panel; no executable listing.
