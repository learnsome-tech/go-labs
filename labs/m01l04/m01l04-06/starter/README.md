# m01l04-06: Exit codes are the contract with your pipeline

You can choose correctly between running, building and installing a program, write functions that return a value alongside an error, format output with the verbs you will use daily, and set an exit code that a pipeline can act on.

Here is the part that makes a program a tool rather than a script somebody runs by hand. A pipeline does not read your output; it reads your exit status. Zero means success and anything else means failure, and that single number decides whether a deployment continues. So the summary for humans goes to standard output, the reason for failure goes to standard error, and the process exits with one. Run it and look at the two streams together. Two rules come with this: exit immediately terminates the process, so deferred cleanup does not run, and choose your codes deliberately, because somebody will eventually write a pipeline condition against them.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
