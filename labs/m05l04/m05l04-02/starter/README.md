# m05l04-02: Decoding into a struct

You can decode API responses into structs, encode reports, and reject malformed input clearly.

Reply describes the fields the tool needs, and tags map them to the JSON keys. Unmarshal fills the struct through its address and returns a syntax or type error when decoding fails. The print shows the two typed values and a nil error. Once data is in a struct, the rest of the checker can use ordinary Go fields instead of repeated string lookups.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
