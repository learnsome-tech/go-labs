# m07l03-05: The failure policy

You can give every check a deadline, classify failures, and keep one slow target from blocking the report.

Give every check a deadline, keep the target beside its error, and continue collecting results when one endpoint fails. Distinguish transport, timeout, and status failures in the report. Only after collection should the command choose whether any unhealthy target makes the process exit non zero.

- Give each check a deadline
- Wrap errors with target context
- Collect healthy and failed results
- Choose exit status after collection

Notes or exercise panel; no executable listing.
