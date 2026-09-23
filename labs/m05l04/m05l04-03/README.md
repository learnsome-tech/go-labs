# m05l04-03: Encode only what you mean to publish

You can decode API responses into structs, encode reports, and reject malformed input clearly.

Encoding follows the same boundary in the other direction. Only exported fields are visible to the JSON package, and tags can rename or omit fields when an empty value carries no information. Marshal returns bytes and an error. Treat the resulting report as an API: choose stable names, avoid leaking internal fields, and test the shape that downstream automation will parse.

- Exported fields are visible to json
- Omit empty fields when useful
- Marshal returns bytes and an error
- A stable report is an API

Notes or exercise panel; no executable listing.
