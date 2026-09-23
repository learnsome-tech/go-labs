# m07l03-03: Errors are report data

You can give every check a deadline, classify failures, and keep one slow target from blocking the report.

A failed target is report data, not necessarily a failed checker process. Keep the error beside the target and continue collecting other results. Separate a transport failure from a response status failure so the report can say whether a service was unreachable or merely unhealthy. Decide the overall process exit policy after all checks are complete.

- Do not stop the whole run for one target
- Keep the error beside the target
- Separate transport from status failure
- Choose the process exit policy later

Notes or exercise panel; no executable listing.
