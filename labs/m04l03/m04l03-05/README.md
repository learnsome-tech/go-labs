# m04l03-05: The select rules

You can wait on several channel operations and keep a concurrent worker responsive.

Select lets a goroutine wait for work and an escape path together. Add a timer when a result must arrive by a deadline, and use default only when an immediate poll is intentional. Ready cases have no priority based on their order. This keeps workers responsive without hiding an unbounded wait inside a channel receive.

- Select chooses a ready channel case
- A timer provides a deadline
- Default makes the operation non blocking
- Never rely on case order

Notes or exercise panel; no executable listing.
