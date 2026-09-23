# m05l04-01: JSON is a boundary format

You can decode API responses into structs, encode reports, and reject malformed input clearly.

JSON usually sits at the boundary between a tool and an API. Decode turns bytes into a typed struct, while Encode turns a typed result into a report. Struct tags map Go field names to the keys the service actually sends. The decoder validates syntax and returns an error, but your application still needs a policy for missing fields and unknown data. Keep the boundary typed as soon as possible.

- Struct tags map fields to keys
- Decode validates syntax
- Encode produces a report
- Unknown fields need a policy

Notes or exercise panel; no executable listing.
