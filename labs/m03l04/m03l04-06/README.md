# m03l04-06: The testing workflow

You can write focused Go tests with the standard testing package and read failures as feedback about behaviour.

The standard testing package is ordinary Go with a small reporting API. Put tests in underscore test files, name each behaviour with a Test function, and use fatal reporting when the rest of the test cannot continue. Keep independent checks together when that gives a clearer failure report. Avoid real networks and uncontrolled time so the same test means the same thing on a laptop and in a pipeline.

- Name behaviour in a Test function
- Use Fatalf for required setup
- Use Error for independent checks
- Keep unit tests deterministic

Notes or exercise panel; no executable listing.
