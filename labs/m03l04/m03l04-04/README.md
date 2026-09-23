# m03l04-04: Keep tests deterministic

You can write focused Go tests with the standard testing package and read failures as feedback about behaviour.

A unit test should be quick and repeatable. Avoid reaching a real network, waiting on wall clock time, or depending on files left by another test. Use a fake implementation at an interface boundary, a temporary directory for file work, and fixed input for randomness. When one behaviour fails, the output should point to one decision. Deterministic tests give a platform team confidence that a change is safe before it reaches a build pipeline.

- Avoid real networks in unit tests
- Control time and randomness
- Use temporary directories for files
- Make one behaviour fail at a time

Notes or exercise panel; no executable listing.
