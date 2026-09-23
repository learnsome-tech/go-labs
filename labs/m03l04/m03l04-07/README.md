# m03l04-07: Try it yourself

You can write focused Go tests with the standard testing package and read failures as feedback about behaviour.

Write a small host validator that returns an error for an empty name. Add one test for a valid host and another for the empty case. Run the package tests, then deliberately change one expected result and read the failure message. Restore the expectation and run again. That loop is the basic feedback cycle you will use before every change to a tool.

- Write a function that returns an error for an empty host.
- Create a test for a valid host and a test for the empty case.
- Run go test and read the failure after deliberately changing one expectation.

Notes or exercise panel; no executable listing.
