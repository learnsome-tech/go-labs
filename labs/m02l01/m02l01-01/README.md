# m02l01-01: A pointer is an address you can use

You can read pointer syntax, mutate shared state deliberately, and recognise a nil pointer before it becomes a crash.

A pointer is a value that tells you where another value lives. That sounds low level, but the useful idea is small. A normal variable gives you the value itself. A pointer gives you a handle to that value, so two parts of a program can agree on one piece of state. The ampersand operator takes an address, and the star operator follows an address. Go makes both operations explicit, which means a reader can see when code shares state instead of guessing from a function call.

- A value lives somewhere in memory
- A pointer stores where that value lives
- An ampersand takes an address
- A star follows an address to the value

Notes or exercise panel; no executable listing.
