# m02l04-04: Interfaces are strongest at boundaries

You can define small interfaces, satisfy them without declarations, and use that flexibility to test and compose tooling code.

Use an interface at a boundary where several implementations may enter the function. Accepting an interface gives the caller room to provide a real client, a fake client, or a wrapper that records calls. Returning a concrete type often gives the caller a clearer shape to work with. Keep the interface close to the consumer that needs it, and resist a large contract that bundles unrelated operations. A small boundary is easier to replace when a service moves from a local check to a remote one.

- Accept interfaces where callers vary
- Return concrete types when construction matters
- Keep interfaces close to their consumer
- Avoid large all purpose contracts

Notes or exercise panel; no executable listing.
