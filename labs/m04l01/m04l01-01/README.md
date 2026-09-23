# m04l01-01: A goroutine is a concurrent function

You can start goroutines, wait for their work, and recognise the ownership questions concurrency introduces.

A goroutine is a function running concurrently with its caller. The go statement starts it and the caller continues immediately. That makes parallel work cheap, but it also creates an ownership question: who knows when the work is finished, and who owns the data it changes? A process exits when main returns, even if a goroutine is still running, so every concurrent operation needs a way to join or signal completion.

- The go statement starts work
- The caller continues immediately
- A process exits when main returns
- Shared state needs coordination

Notes or exercise panel; no executable listing.
