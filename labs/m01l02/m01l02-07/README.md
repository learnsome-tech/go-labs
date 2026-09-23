# m01l02-07: The toolchain in four habits

You can install the Go toolchain the way that suits a laptop, a build agent or a container, and use the go command's own subcommands for formatting, checking and documentation instead of reaching for extra tools.

Four habits to leave with. Install the toolchain once and stop thinking about it, because there is no second layer to configure. Let your editor format on save, and have the pipeline fail when any file is unformatted. Run vet next to the tests, since it costs a second and catches a class of bug that reviews miss. And look things up with the documentation subcommand rather than the web, because it answers for the version in front of you. Next lesson we give our code a home: modules, import paths, and the workspace file that lets a tool and a shared library live in one checkout.

- Install once; there is nothing else to set up
- Format on save, and list unformatted files in the pipeline
- Vet alongside the tests, not once a quarter
- Look things up with doc, against the version you have

Notes or exercise panel; no executable listing.
