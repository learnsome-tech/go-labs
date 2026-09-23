# m01l05-04: Why this changes how you write configuration

You can declare variables both ways and say which to use where, predict the zero value of any type and design configuration structs around it, define constant sets with iota, and convert between numeric types without guessing.

This has a direct consequence for the kind of code an operations team writes. When you define a settings struct and someone creates one without filling it in, every field already holds a sensible starting point, so you can design your defaults to be the zero values and delete a page of initialisation. Strict mode off, retry count zero, no tags. There is a real cost, and you should know it before it bites: the zero value is indistinguishable from a deliberate zero, so you cannot tell whether somebody set the retry count to zero or never mentioned it. When that difference matters, and in a configuration file parser it usually does, you reach for a pointer field, which is nil when absent, or you carry a separate flag saying it was set.

- A fresh config struct is already valid, not empty and dangerous
- Choose defaults so the zero value means the safe behaviour
- The cost: you cannot tell absent from deliberately zero
- When that matters, use a pointer field or a second boolean

Notes or exercise panel; no executable listing.
