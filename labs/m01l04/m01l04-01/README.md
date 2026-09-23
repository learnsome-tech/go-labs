# m01l04-01: Run, build, install: one compiler, three verbs

You can choose correctly between running, building and installing a program, write functions that return a value alongside an error, format output with the verbs you will use daily, and set an exit code that a pipeline can act on.

Three verbs, one compiler underneath, and the only difference is where the executable ends up. Run compiles into a temporary directory, executes it and throws it away, which is what you want while you are changing something every few seconds. Build writes the binary where you ask and stops, which is what a pipeline does before it copies the result into an image. Install builds and then moves the result into your binary directory so you can type its name from anywhere, which is how you deploy small tools to your own machine. People sometimes assume run is an interpreter. It is not: every one of these compiles your code completely before a line of it executes.

- Run compiles to a temporary place and executes: for the loop
- Build leaves a binary where you asked: for shipping
- Install puts it on your path: for tools you use daily
- All three compile the same way; only the destination differs

Notes or exercise panel; no executable listing.
