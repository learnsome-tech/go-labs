# m02l02-06: The collection decisions

You can choose arrays, slices, and maps for predictable collections and handle their zero values safely.

Arrays carry a fixed length in their type. Slices carry ordered data and grow through append, including from a nil zero value. Maps answer key based questions, but a missing key can look like a zero unless you keep the comma ok result. A nil map can be read but cannot be written, so make it before the first assignment. These decisions keep collection behaviour obvious, which is exactly what you want in a command that will run unattended.

- Use arrays when the length is part of the type
- Use slices for ordered, growing data
- Use maps for key based lookup
- Check comma ok for presence
- Make a map before writing

Notes or exercise panel; no executable listing.
