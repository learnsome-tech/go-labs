# m03l02-05: A wrapped error keeps its identity

You can wrap an underlying error with context and test its identity with errors Is.

Start returns a wrapped busy error. The first print gives the operation and its cause in one line. The second branch uses errors Is to recognise that busy is retryable and prints a different policy. This is a small but important distinction for automation: the message can be enriched at every layer, while the decision still rests on a stable error identity. Keep sentinel values focused on conditions that callers truly need to distinguish.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
