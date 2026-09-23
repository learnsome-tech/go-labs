# m07l03-04: A failed result stays attached

You can give every check a deadline, classify failures, and keep one slow target from blocking the report.

Result carries the target and its error together. The coordinator can print or encode this value without looking up a separate error map. A successful result can leave Err nil, while a timeout or transport failure keeps a readable cause. The project will use this one shape for every target.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
