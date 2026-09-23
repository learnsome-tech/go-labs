# m03l05-06: The table test recipe

You can express many input cases in one table test and get a useful name for every failure.

Table driven tests turn many scenarios into a small data set. Give every row a name, keep inputs and expected results together, and run one assertion per row. Subtests make failures searchable in the test output. Empty values and boundaries belong in the table because they are part of the contract. This style keeps the test readable as a tool grows and makes a new case a simple data change.

- Define named rows with inputs and wants
- Loop over rows with one assertion
- Use subtests for searchable failures
- Add edge cases as data

Notes or exercise panel; no executable listing.
