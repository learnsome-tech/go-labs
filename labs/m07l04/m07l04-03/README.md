# m07l04-03: Keep machine output clean

You can encode health results as stable JSON for pipelines, dashboards, and later automation.

Machine readable output works when standard output contains only the document. Put diagnostics on standard error, and use indentation when a person is reading a saved report. The process exit status can still summarize whether any target was unhealthy, while the JSON carries the detail needed to decide what happened.

- JSON belongs on standard output
- Diagnostics belong on standard error
- Use indentation only for people
- Exit status still carries the summary

Notes or exercise panel; no executable listing.
