# m02l04-02: A type satisfies an interface silently

You can define small interfaces, satisfy them without declarations, and use that flexibility to test and compose tooling code.

Speaker asks for one behaviour: Speak returns a string. Robot provides that method, and there is no implements line anywhere in the type declaration. The compiler checks the relationship when Robot is passed to announce, whose parameter accepts the interface. This is why implicit interfaces feel light in Go. The consumer states what it needs, and any type that already has that behaviour can be used. The concrete type stays independent of the consumer's package.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
