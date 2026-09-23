# m05l01-07: What to carry forward

You can read arguments and environment, read and write files safely, tell a missing file from a broken one, and set the exit code your pipeline branches on.

So: arguments and environment come in through Args and LookupEnv, whole small files go in and out with ReadFile and WriteFile, and anything larger gets Open with a defer on the very next line. Errors are values here, so a missing file is something you test for with errors dot Is rather than something you catch. And the number you exit with is part of your tool's public interface, so choose it deliberately and write it in the readme. Next lesson we pick up the file handle we just opened and look at the two interfaces the whole standard library is built around, which is where the real leverage in Go tooling lives.

- Args and LookupEnv cover argument and environment input
- ReadFile and WriteFile for small files, Open for streams
- defer Close on the line after a successful Open
- errors.Is with os.ErrNotExist, never a string compare
- Pick your exit codes deliberately and document them

Notes or exercise panel; no executable listing.
