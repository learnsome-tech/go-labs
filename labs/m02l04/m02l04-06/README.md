# m02l04-06: The interface habits

You can define small interfaces, satisfy them without declarations, and use that flexibility to test and compose tooling code.

An interface is a method contract, and Go checks satisfaction implicitly when a value crosses that contract. Define the interface where the consumer uses it, keep it small, and accept it at a boundary where implementations can vary. Concrete types remain ordinary structs with ordinary methods. This gives tooling code a clean seam for tests, wrappers, and alternate transports without asking every type to register itself with a central hierarchy.

- Interfaces list methods, not data fields
- Satisfaction is implicit
- Define the smallest contract the caller needs
- Use boundaries to make tools replaceable

Notes or exercise panel; no executable listing.
