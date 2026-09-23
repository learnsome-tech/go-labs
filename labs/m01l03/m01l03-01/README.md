# m01l03-01: A module is the unit of dependency

You can start a module, split a tool into packages with import paths that follow from the module path, keep dependencies tidy and vendored, and use a workspace file when a tool and a shared library sit in one checkout.

Three words get used loosely and mean different things. A package is one directory of source files that are compiled together. A module is a tree of packages released and versioned as a unit, with a file at its root recording its path, the language version it needs, and the other modules it depends on. And a workspace is a way to tell the toolchain that several modules on your disk should be built together. Coming from Python, the relief here is that there is nothing global: no shared install directory, no environment to activate, no chance that two tools on the same machine disagree about a library version. The dependency set belongs to the module, not to the machine.

- A module is a tree of packages with one version and one path
- The go mod file records the path, the language version, the needs
- A package is a directory; the directory name is the import name
- Capitalised names are exported; lower case stays inside
- Nothing is global: there is no site packages to pollute

Notes or exercise panel; no executable listing.
