# m01l04-05: Where install puts things, and why it matters

You can choose correctly between running, building and installing a program, write functions that return a value alongside an error, format output with the verbs you will use daily, and set an exit code that a pipeline can act on.

Install compiles and then places the binary in a bin directory under the go path, which defaults to a go directory in your home. Put that on your shell path once and any tool you install becomes a command. The version worth remembering is install with an at suffix naming a module and a version: it fetches that module, builds it and installs the binary, without touching the module you happen to be standing in. That is how a great many tools in this ecosystem are distributed, and it is why Go tooling instructions are usually one line long rather than a page about package managers. It is also a supply chain decision, so pin the version rather than taking latest.

- Install writes to the bin directory under the go path
- Put that directory on your path once and forget it
- Install with an at version suffix fetches, builds and installs
- That is how most Go tools you use were delivered to you

Notes or exercise panel; no executable listing.
