# m04l02-02: Sending and receiving a value

You can send values through channels, close them deliberately, and use channels as ownership boundaries.

The channel carries strings. The worker sends ready, and the receiver waits on the arrow expression until that value arrives. Because the channel is unbuffered, the send and receive meet at one synchronization point. This is often clearer than sharing a variable protected by a lock: ownership of the value moves from the sender to the receiver.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
