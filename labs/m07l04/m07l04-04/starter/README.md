# m07l04-04: Indented output for an operator

You can encode health results as stable JSON for pipelines, dashboards, and later automation.

An encoder can write directly to standard output and indent the document for a human operator. The report remains valid JSON, but its lines are easier to inspect in a terminal or saved artifact. For a pipeline where bytes matter, leave indentation off and keep diagnostics on standard error.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
