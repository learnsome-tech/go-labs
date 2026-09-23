# m01l03-04: Import paths follow from the module path

You can start a module, split a tool into packages with import paths that follow from the module path, keep dependencies tidy and vendored, and use a workspace file when a tool and a shared library sit in one checkout.

An import path is just the module path with the directory appended, so moving a directory changes its import path and nothing else has to be registered anywhere. Standard library packages are the ones with short bare names. There is one rule here that shapes designs more than anything else in this lesson: import cycles are a compile error. If your storage package imports your server package and the server imports storage, the build stops, and you have to decide which one owns the shared type. Teams new to Go find this annoying for a week and then find that their tools have a layering they can describe out loud, which is not the usual experience with a Python package tree.

- Import path is the module path plus the directory inside it
- Renaming a directory renames the import: no registry to update
- Standard library imports are bare names like os and net h t t p
- An import cycle is a compile error, which forces honest layering

Notes or exercise panel; no executable listing.
