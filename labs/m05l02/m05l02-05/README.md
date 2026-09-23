# m05l02-05: The io habits

You can copy streams, read bounded input, and treat files and network bodies through common interfaces.

Reader and Writer let one function work with files, buffers, standard streams, and network bodies. io Copy transfers bytes without knowing the concrete types. ReadAll is fine for a known small value; otherwise put a limit around untrusted input. The owner that opens a resource remains responsible for closing it after the last read or write.

- Design functions around Reader and Writer
- Use Copy for direct transfer
- Bound untrusted input
- Close resources at the owner

Notes or exercise panel; no executable listing.
