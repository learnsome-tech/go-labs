# m01l03-06: Workspaces: the piece the roadmap skips

You can start a module, split a tool into packages with import paths that follow from the module path, keep dependencies tidy and vendored, and use a workspace file when a tool and a shared library sit in one checkout.

This next piece is authored for this course rather than taken from the language roadmap, which ships no topic for it at all, even though anyone who maintains more than one Go repository runs into it in the first week. The situation is this: you have a tool and a shared library, you are changing both, and the tool must build against the library on your disk rather than the last published version. The old answer was to add a replace directive to the requirements file and try to remember to remove it before committing, which people forgot, constantly. A workspace solves it outside the module: one file, listing the module directories you want built together.

- Authored for this course: no roadmap topic covers go work
- One checkout, several modules, built as you have them on disk
- Replaces editing a replace directive you must remember to undo
- The workspace file is yours: keep it out of the repository

Notes or exercise panel; no executable listing.
