# m02l05-02: A checked type assertion

You can inspect an interface value safely, use any for truly mixed data, and avoid hiding useful types behind it.

The checked assertion asks whether an interface currently holds a string. The comma ok form returns the string and a boolean, so a value of another type follows the safe branch instead of panicking. The function uses any, which is the modern spelling of an empty interface and means any value may arrive. That freedom is useful at a serialization edge, but it also removes compile time detail. Keep the assertion close to that edge and convert to a useful concrete type before the rest of the program continues.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
