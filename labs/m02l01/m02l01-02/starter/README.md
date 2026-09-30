# m02l01-02: Taking an address and following it

You can read pointer syntax, mutate shared state deliberately, and recognise a nil pointer before it becomes a crash.

Here is the complete round trip. We start with a string named name, take its address into p, and then print the value reached through p. The star reads the value at the address. When we assign through that same star, we change the original variable, because both names point to one storage location. This is the part to remember for reviews: the pointer is not a second copy of the string. It is a route to the existing string. Run the example and the two lines prove the mutation happened in place.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
