# m02l03-01: A struct names the data together

You can model tool data with structs, attach methods, and choose value or pointer receivers intentionally.

A struct groups related values under one named type. For a platform tool that might be a target with a name, a port, and a timeout. The field names document the shape wherever the value travels, and a composite literal makes construction visible at the call site. Structs have a zero value too, so you can often declare one and fill the fields you need. This gives configuration and report data a stable shape before you add any behaviour.

- Fields give related values one type
- Composite literals make construction visible
- Field names document the data shape
- A zero struct is often ready to use

Notes or exercise panel; no executable listing.
