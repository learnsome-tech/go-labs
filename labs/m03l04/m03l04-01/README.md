# m03l04-01: Tests are ordinary Go code

You can write focused Go tests with the standard testing package and read failures as feedback about behaviour.

The standard testing package keeps tests close to the code they protect. A file ending in underscore test is compiled for tests, and a function beginning with Test receives a pointer to testing information. The go test command builds a temporary test binary, runs it, and reports failures with file and line context. A test should describe behaviour a caller depends on, so a failure tells you which promise needs attention.

- Test files end with underscore test
- Test functions receive a testing pointer
- The go test command builds a test binary
- A failure names the behaviour to inspect

Notes or exercise panel; no executable listing.
