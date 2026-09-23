# m05l02-03: Bound what you read

You can copy streams, read bounded input, and treat files and network bodies through common interfaces.

ReadAll is convenient for a small response, but it keeps the whole stream in memory. For input from a network or a command line pipe, choose a bound. LimitReader exposes at most the number of bytes you permit, while ReadFull asks for an exact amount and reports short input. A bound is an operational decision as much as a memory decision, because it keeps an unexpected response from filling a process.

- ReadAll uses memory for the whole stream
- LimitReader caps the bytes consumed
- ReadFull requires an exact amount
- Choose a bound for untrusted input

Notes or exercise panel; no executable listing.
