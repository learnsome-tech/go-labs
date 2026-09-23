# m01l03-03: Splitting a tool into packages

You can start a module, split a tool into packages with import paths that follow from the module path, keep dependencies tidy and vendored, and use a workspace file when a tool and a shared library sit in one checkout.

Here is a second package in the same module. The package clause names it, and by convention the name matches its directory. Then a type worth sharing, describing one thing our tool watches. Then one exported function that renders a target for a human. Notice how visibility works: a capital letter means other packages can use this name, and a lower case name is private to the directory. There is no keyword for public or private, only that spelling rule. The directory called internal is special too: the toolchain refuses imports of it from outside this module, which is how you publish a tool without publishing its guts. When the main package calls it, we get one line back.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
