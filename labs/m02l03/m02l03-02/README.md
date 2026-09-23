# m02l03-02: Defining and constructing a struct

You can model tool data with structs, attach methods, and choose value or pointer receivers intentionally.

The type declaration says a Target has a Name and a Port. Both fields are exported because their names begin with capital letters, so another package could read them. The literal uses field names instead of relying on order, which keeps a later field addition from silently changing every construction site. Then ordinary field selection reads the values. This is the shape you might pass from a flag parser into a checker: one value with a name that explains what the fields mean.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
