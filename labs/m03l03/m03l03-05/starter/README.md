# m03l03-05: Deferred cleanup still runs

You can distinguish a programmer invariant from an expected operational error and use recover only at a deliberate boundary.

A panic unwinds through deferred calls in last in first out order. The recovery function runs first and prints recovered. The cleanup defer runs next, so resources registered with defer still get their chance to close. Control returns to the caller, which prints done. This ordering is useful at a worker boundary, but it should not encourage casual panics. The normal path should still return errors for conditions the caller can reasonably handle.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
