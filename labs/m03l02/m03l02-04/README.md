# m03l02-04: Wrapping is part of your API

You can wrap an underlying error with context and test its identity with errors Is.

Wrapping is not only a formatting choice; it is part of the error contract between packages. Use the wrapping verb when callers should be able to inspect the cause. Use ordinary formatting when the lower level identity should remain private. If a sentinel is public, document the condition it represents and keep its meaning stable. The final message should still read naturally in logs, because an operator will usually see the text before anyone opens a debugger.

- Use percent w to preserve a cause
- Use percent v when identity is not needed
- Document sentinel errors callers may inspect
- Keep messages readable for operators

Notes or exercise panel; no executable listing.
