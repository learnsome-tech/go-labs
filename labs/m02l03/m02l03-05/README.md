# m02l03-05: A method that changes state

You can model tool data with structs, attach methods, and choose value or pointer receivers intentionally.

Counter has a Value field whose zero value is ready for use. Add has a pointer receiver because it must change that field on the original counter. The caller holds an addressable value, so Go lets it call Add directly and supplies the address behind the scenes. Two calls leave the counter at two, which proves that the receiver did not mutate a discarded copy. This is a small example, but the same shape fits retry counts, bytes processed, and health results collected during a run.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
