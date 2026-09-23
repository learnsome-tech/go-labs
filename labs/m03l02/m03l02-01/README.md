# m03l02-01: Context without losing the cause

You can wrap an underlying error with context and test its identity with errors Is.

As errors move up through packages, each layer should add the operation it was attempting without throwing away the original cause. Wrapping does that. The resulting message reads well to a person, while the error chain stays available to code. The errors Is function searches through that chain for a sentinel or a matching value. This lets a caller decide that a missing file is expected, while a permission failure still stops the run.

- Wrapping adds a human useful operation
- The original error stays in the chain
- errors Is checks the chain
- Sentinel errors support decisions

Notes or exercise panel; no executable listing.
