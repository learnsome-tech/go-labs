# m05l01-08: Try it yourself

You can read arguments and environment, read and write files safely, tell a missing file from a broken one, and set the exit code your pipeline branches on.

Your turn, and this one is a tool you can keep. Take a path from the arguments, and when nothing is passed fall back to an environment variable so a container can configure it. If the file is there, print its size. If it is not there, say so plainly and exit with the status three, so a pipeline can tell that case apart from any other failure. Then write a one line summary to a file that only the owner can read, read it back, and print the mode to prove the permission bits are what you asked for. Try it once with a directory instead of a file, and read the error you get.

- Write a tool that reads a path from os.Args and a default from an env var when no path is given.
- Print the file size when it exists, print absent and exit 3 when it does not.
- Write a summary line to a file with mode 0o600 and read it back to prove the mode stuck.

Notes or exercise panel; no executable listing.
