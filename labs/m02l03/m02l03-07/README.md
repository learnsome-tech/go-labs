# m02l03-07: Try it yourself

You can model tool data with structs, attach methods, and choose value or pointer receivers intentionally.

Define a Check struct with a name and a healthy boolean. Add a read only method that returns the name and status as text, then add a pointer method that marks the check healthy. Construct the zero value, call the mutating method, and print the report from the read only method. Notice how the receiver tells the reader whether a call can change the check, without making the caller inspect the method body.

- Define a Check struct with a name and a healthy boolean.
- Add a method that returns the name and status as text.
- Add a pointer method that marks the check healthy.

Notes or exercise panel; no executable listing.
