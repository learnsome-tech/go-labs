# m01l03-08: What to remember about layout

You can start a module, split a tool into packages with import paths that follow from the module path, keep dependencies tidy and vendored, and use a workspace file when a tool and a shared library sit in one checkout.

To summarise: one module file at the root, packages as directories beneath it, import paths that follow from the module path, and visibility decided by capital letters. Keep the guts in a directory called internal when you do not want strangers importing them. Tidy before committing so the requirements match the code. Vendor when the build agent has no network or when you want the dependency tree in the repository. And reach for a workspace when two modules on your disk have to be built together, remembering it is a local file describing local reality. Next lesson: what the difference between running and building really is, and how a tool reports success or failure to whatever called it.

- One go mod file at the root; packages are directories under it
- Capital letters export; internal hides a directory from the world
- Tidy before you commit; vendor when the agent has no network
- A workspace is for your disk, not for the repository

Notes or exercise panel; no executable listing.
