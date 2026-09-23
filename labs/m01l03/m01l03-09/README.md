# m01l03-09: Try it yourself

You can start a module, split a tool into packages with import paths that follow from the module path, keep dependencies tidy and vendored, and use a workspace file when a tool and a shared library sit in one checkout.

Build the shape rather than reading about it. Start a module with a path of your own choosing, add a package in a directory under internal, export one function from it and call that function from your main package. Then make the mistake on purpose: create a second module in a sibling directory, try to import the first module's internal package, and read the refusal carefully. Knowing what that error looks like is worth a great deal when you meet it in somebody else's repository and have to decide whether the boundary is deliberate.

- Start a module, then add a package under a directory called internal.
- Export one function from it, call it from the main package, and build.
- Try importing that internal package from a second module and read the error.

Notes or exercise panel; no executable listing.
