# m03l01-01: An error is an ordinary result

You can return errors, inspect them, and keep failure handling explicit in a command line tool.

In Go, an error is an ordinary value returned by a function. There is no hidden exception path waiting to jump across your code. A call gives you its useful result and, when something went wrong, an error beside it. The caller decides whether to retry, report, or stop. That makes failure handling visible in a review and predictable in a tool that runs unattended. Build the habit of checking the error before you use the result it came with.

- Errors are values returned by functions
- The caller decides what failure means
- Handle an error before using the result
- Keep success and failure paths visible

Notes or exercise panel; no executable listing.
