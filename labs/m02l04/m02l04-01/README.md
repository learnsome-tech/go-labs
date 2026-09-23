# m02l04-01: An interface describes behaviour

You can define small interfaces, satisfy them without declarations, and use that flexibility to test and compose tooling code.

An interface describes behaviour through a set of methods. It does not describe the concrete data behind those methods. Go interfaces are implicit: a type satisfies an interface by having every method the interface asks for. The type never declares that it implements anything. This keeps the contract with the caller, which can define the smallest behaviour it needs. Small interfaces are easier to fake in a test and easier to compose in a command line tool.

- Methods are the contract
- The implementer does not name the interface
- Small contracts compose well
- Callers depend on what they use

Notes or exercise panel; no executable listing.
