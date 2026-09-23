# m03l05-04: Tables make edge cases visible

You can express many input cases in one table test and get a useful name for every failure.

The value of a table is the cases it invites you to write down. Include empty input, a boundary value, and a normal value beside the obvious success case. Store expected errors or status values in the row so the policy is visible before the loop runs. Name each situation in plain language. When a new edge case appears, add a row first; only add branching to the test when the assertion genuinely differs.

- Put empty and boundary input in rows
- Keep expected errors in the table
- Use names that explain the situation
- Add a row before adding branching code

Notes or exercise panel; no executable listing.
