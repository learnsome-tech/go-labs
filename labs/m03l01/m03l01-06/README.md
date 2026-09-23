# m03l01-06: The error habits

You can return errors, inspect them, and keep failure handling explicit in a command line tool.

Errors are values, so a function returns its useful result and an error together. Check the error before using the result, return early when the failure path is complete, and include context that helps an operator act. Keep retry decisions and exit codes at the edge of the program. These simple habits make a tool's failure behaviour visible in code, easy to test, and safer to run from a pipeline.

- Return errors beside useful results
- Check before using the result
- Add context at the boundary
- Keep process policy out of libraries

Notes or exercise panel; no executable listing.
