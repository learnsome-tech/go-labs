# m03l02-07: Try it yourself

You can wrap an underlying error with context and test its identity with errors Is.

Create a sentinel error for an unavailable service. Write a function that wraps it with the service name and an operation. Print the resulting message, then use errors Is to decide whether the caller should retry. Change the wording of the wrapper and confirm that the branch still works. That experiment shows why error identity belongs to code while the message belongs to the operator.

- Create a sentinel error for an unavailable service.
- Wrap it in a function that adds the service name.
- Use errors Is to print a retry message without comparing text.

Notes or exercise panel; no executable listing.
