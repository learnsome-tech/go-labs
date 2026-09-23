# m05l01-01: The four conversations every tool has

You can read arguments and environment, read and write files safely, tell a missing file from a broken one, and set the exit code your pipeline branches on.

Every command line tool you write has the same four conversations with the machine around it. It reads its arguments. It reads its environment. It touches files. And it leaves behind an exit status that some pipeline is going to branch on. In Python you would reach for sys, for pathlib and for open. In Go all four live in one package called o s, and the shapes are close enough that you will feel at home. One difference matters from the first line: nothing here raises. Every call that can fail hands you a value and an error side by side, and your tool decides what happens next.

- os.Args and os.Getenv: how a job hands you its input
- os.ReadFile and os.WriteFile: a small file in one call
- os.Open plus defer Close: anything you would rather stream
- os.Exit: the number your pipeline actually branches on

Notes or exercise panel; no executable listing.
