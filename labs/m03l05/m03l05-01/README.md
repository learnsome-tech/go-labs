# m03l05-01: A table is a list of contracts

You can express many input cases in one table test and get a useful name for every failure.

A table driven test stores many cases as data and runs the same assertion over each row. Every row can name the situation, carry its input, and state the expected result. The loop removes repeated setup and keeps the test focused on the contract. When a row fails, its name tells you which input deserves attention. This shape works especially well for parsers, flag validation, and the status mapping in a health checker.

- Each row names one case
- Inputs and expected results stay together
- A loop removes repeated test code
- Names make failures searchable

Notes or exercise panel; no executable listing.
