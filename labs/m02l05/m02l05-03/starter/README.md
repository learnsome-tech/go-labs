# m02l05-03: A type switch handles several cases

You can inspect an interface value safely, use any for truly mixed data, and avoid hiding useful types behind it.

A type switch is the readable form when several concrete types are expected. The switch tests the dynamic type and gives each case a value with the matching concrete type. Strings and integers get specific output, while the default branch keeps an unexpected value from crashing the tool. This is useful when decoding deliberately mixed input, such as a small set of command values. If every caller should provide one known behaviour, an ordinary interface is usually clearer than a switch over any.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
