# m02l04-07: Try it yourself

You can define small interfaces, satisfy them without declarations, and use that flexibility to test and compose tooling code.

Define a Probe interface with one Check method that returns a string. Create two concrete probe types, one that represents a real check and one that returns a fixed result for a test. Give both the same method, then write one function that accepts the interface and prints both results. Keep the interface beside that function and notice how no concrete type needs to mention the contract by name.

- Define a Probe interface with a Check method that returns a string.
- Create a real looking probe and a fixed probe that both satisfy it.
- Write one function that accepts the interface and prints both results.

Notes or exercise panel; no executable listing.
