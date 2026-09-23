# m02l05-05: Checking a typed configuration value

You can inspect an interface value safely, use any for truly mixed data, and avoid hiding useful types behind it.

This boundary accepts any value, then narrows it immediately. The assertion requires an integer, and the second check rejects a value below one, so the function returns a boolean that tells the caller whether the setting is usable. A real configuration loader might do this after decoding a generic map, then pass the resulting integer through typed code. Notice that the invalid string never travels deeper into the program. Convert uncertainty at the edge, and keep the rest of the tool simple.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
