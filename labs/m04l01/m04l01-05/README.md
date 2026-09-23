# m04l01-05: The goroutine rules

You can start goroutines, wait for their work, and recognise the ownership questions concurrency introduces.

A goroutine lets a function run concurrently, and a wait group lets its owner join that work before returning. Pass independent values into workers, protect any shared mutation, and treat scheduling order as unpredictable. These rules keep concurrency visible and make the lifetime of each job something a reviewer can reason about.

- Start work with go
- Join work with a wait group
- Pass values when sharing is unnecessary
- Never rely on scheduling order

Notes or exercise panel; no executable listing.
