# m02l01-06: What to carry forward

You can read pointer syntax, mutate shared state deliberately, and recognise a nil pointer before it becomes a crash.

Pointers make sharing explicit. Take an address when a function needs to work on the caller's storage, follow it with the star operator, and remember that the pointer itself is only a route. Nil is the pointer zero value, so check it before you follow an optional route. Prefer ordinary zero values when they already model absence. These few rules cover most pointer code you will meet in platform tools, and they leave ownership visible to the next engineer reading the function signature.

- Pointers name shared storage
- Ampersand takes an address, star follows it
- A pointer parameter can mutate caller state
- Nil must be checked before dereference

Notes or exercise panel; no executable listing.
