# m07l04-01: The report is another API

You can encode health results as stable JSON for pipelines, dashboards, and later automation.

The health report is an API consumed by pipelines and dashboards. Give its fields stable names with JSON tags, collect every result before encoding, and write one document to standard output. Keep operational messages on standard error so a caller can parse the report without filtering human text.

- Use exported fields and tags
- Keep status names stable
- Encode once after collection
- Write JSON to standard output

Notes or exercise panel; no executable listing.
