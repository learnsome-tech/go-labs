# m03l01-04: Errors should carry action and context

You can return errors, inspect them, and keep failure handling explicit in a command line tool.

A useful error tells the operator what failed and gives enough context to act. Include the operation and the resource when that helps, but keep the underlying cause available for code to inspect later. A boolean only says that something went wrong; an error can explain what happened. Let lower layers return facts, and let the command line edge choose whether to retry, print a message, or set the exit status. That separation keeps reusable packages free of process policy.

- Say which operation failed
- Include the resource when it helps
- Do not hide errors in a boolean
- Choose retry and exit policy at the edge

Notes or exercise panel; no executable listing.
