# m04l04-04: Observing a cancelled context

You can pass context through a call chain and use its deadline and values at the right boundaries.

The derived context is cancelled immediately. Its Done channel is closed, so the select receives and the context error explains the reason. A real worker would put this case beside its work channel and return from the function. The caller owns cancel here, and calling it releases the resources associated with the derived context.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
