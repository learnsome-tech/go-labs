# m03l02-06: The wrapping pattern

You can wrap an underlying error with context and test its identity with errors Is.

Wrap an underlying error when you add an operation, and preserve its cause with the wrapping verb. Use errors Is to inspect that chain and branch on a stable condition. Do not compare the printed message, because its wording exists for people and may gain useful context over time. Keep the final decision about defaults, retries, and exits at the outer boundary that understands the job being performed.

- Wrap with operation context
- Inspect identity with errors Is
- Never branch on error text
- Choose policy at the outer boundary

Notes or exercise panel; no executable listing.
