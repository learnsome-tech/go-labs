# m01l04-03: A function that returns a value and an error

You can choose correctly between running, building and installing a program, write functions that return a value alongside an error, format output with the verbs you will use daily, and set an exit code that a pipeline can act on.

This is the shape of almost every function you will write in Go. Two packages imported, then a function taking two integers and returning two results: the answer, and an error. There is no exception to throw, so the second result carries the bad news, and a nil error is how the function says nothing went wrong. Look at the parameter list: two parameters of the same type share one type name. Then call it twice, once with sensible numbers and once with a total of zero. The first call gives a percentage and a nil error, which prints inside angle brackets. The second gives the zero value and a message. Printing both together is a teaching shortcut; real code checks the error and branches.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
