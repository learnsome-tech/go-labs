# m02l01-04: Nil means there is no address

You can read pointer syntax, mutate shared state deliberately, and recognise a nil pointer before it becomes a crash.

Every pointer has a zero value, and that value is nil. Nil means there is no address to follow. Dereferencing nil causes a panic, so optional pointers need a check before use. A useful habit is to ask whether a pointer buys you something. If an empty string, a false value, or an empty slice already expresses the absence you need, use that ordinary zero value. Keep pointers for shared mutation, optional identity, or APIs where nil has a deliberate meaning.

- The zero value of a pointer is nil
- Dereferencing nil stops the program
- Check before following optional data
- Prefer a useful zero value when possible

Notes or exercise panel; no executable listing.
