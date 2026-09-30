# m07l01-04: Validating a target

You can start the health checker with typed configuration, clear defaults, and a small target model.

The validator checks the complete target before any request starts. A missing URL returns a clear error, and main prints it as the configuration failure. Later lessons can assume that every target has the fields required for a check, which keeps worker code focused on network behaviour.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
