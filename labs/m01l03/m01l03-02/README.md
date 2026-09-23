# m01l03-02: Starting a module from an empty directory

You can start a module, split a tool into packages with import paths that follow from the module path, keep dependencies tidy and vendored, and use a workspace file when a tool and a shared library sit in one checkout.

Starting a project is two commands. Make a directory, then declare the module path. That path is the name other code will import, so it is conventionally where the code lives: a host, an owner and a repository name. For something private that will never be published, any unique path will do, and you will see me use a local one throughout this course. Look at what it wrote: the module path and the language version, and nothing else. Dependencies appear here as you add them, and not before. There is no scaffolding step, no template and no generator; that file plus one source file is a complete Go project.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
