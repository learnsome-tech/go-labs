# m02l03-03: Methods give the type a vocabulary

You can model tool data with structs, attach methods, and choose value or pointer receivers intentionally.

A method is a function with a receiver, and the receiver gives the type a vocabulary that callers can discover. Address reads a Target and returns the host and port in the form a dialer needs. This receiver is a value receiver, so the method gets a copy and cannot change the original. That is a good default for small data values and formatting helpers. Callers read t dot Address as naturally as they read a field, while the implementation keeps the formatting rule in one place.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
