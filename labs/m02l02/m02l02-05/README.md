# m02l02-05: Building a result and a lookup table

You can choose arrays, slices, and maps for predictable collections and handle their zero values safely.

This is a common tool pattern. Start with a nil slice for results and append as you discover names. Then create a map with make before the loop stores membership flags. The range loop gives us an index and a value, and the blank identifier discards the index because it is not useful here. At the end the slice preserves order, while the map answers a membership question quickly. One collection carries the report, and the other carries the lookup table.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
