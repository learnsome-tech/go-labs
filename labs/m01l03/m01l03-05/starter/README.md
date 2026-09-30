# m01l03-05: Tidy and vendor, and what they are for

You can start a module, split a tool into packages with import paths that follow from the module path, keep dependencies tidy and vendored, and use a workspace file when a tool and a shared library sit in one checkout.

Two dependency commands earn their place in a pipeline. The first adds what you import and removes what you no longer import, rewriting the requirements to match the code as it is now rather than as it was. Run it before you commit; a diff on that file is a dependency review. The second copies every dependency into a directory inside the repository, so the build needs no network at all. That matters when your build agent has no route to the internet, or when you want an audit trail of exactly what went into a release. Our example imports only the standard library, so there is nothing to copy, and the build carries on regardless.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
