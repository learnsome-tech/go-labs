# m01l03-07: One workspace over two modules

You can start a module, split a tool into packages with import paths that follow from the module path, keep dependencies tidy and vendored, and use a workspace file when a tool and a shared library sit in one checkout.

Two module directories, side by side: a tool and a library it imports. Create the workspace naming both. Look at what it records: a language version and a list of directories in use, and no versions at all, because these are the copies on your disk. Now build the tool. It resolves the library to the directory next door, so an edit in the library is visible to the tool immediately with no publish, no tag and no version bump. Keep this file out of the repository: it describes your machine, not the project. And when you want to prove the module still builds the way a stranger would build it, set the workspace variable to off for one command.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
