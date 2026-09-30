# m04l05-04: A caller cancels a worker

You can bound a concurrent operation with a timeout and return promptly when an orchestrator or caller cancels it.

The worker blocks on the context Done signal instead of sleeping forever. Main starts it, cancels the context, and waits for a separate done signal so the process knows the worker has returned. The worker prints its stop message before the program exits. In a service, the same handoff lets shutdown cancel active requests and then wait for their cleanup.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
