# m03l01-05: A caller chooses the policy

You can return errors, inspect them, and keep failure handling explicit in a command line tool.

The check function knows how to identify an unknown host, but it does not know what a process should do next. Main owns that policy. It adds the operation name to the message and can later choose an exit code or a retry. The short declaration inside the if keeps the error scoped to the branch. When the check succeeds, the function returns nil and the success message remains clear. This is the boundary between reusable checking logic and the command that runs it.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
