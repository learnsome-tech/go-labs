# m01l04-07: The four things to keep

You can choose correctly between running, building and installing a program, write functions that return a value alongside an error, format output with the verbs you will use daily, and set an exit code that a pipeline can act on.

Four things to keep from this lesson. Use run while you are iterating, build when something else will carry the binary, install for tools you want on your own path. Write functions that return their answer alongside an error, and read a nil error as nothing went wrong. Learn the printing verbs, especially the one that prints field names, because it will tell you what your data really contains. And treat the exit status as a contract: zero for success, non zero with a human readable reason on standard error. Next lesson, the last of this module, is about variables, types and the zero value, which is where Go quietly differs from the languages you already use.

- Run while iterating, build to ship, install for your own tools
- Return a value and an error; nil means nothing went wrong
- The plus v verb is your fastest debugging instrument
- Exit zero for success, non zero with a reason on standard error

Notes or exercise panel; no executable listing.
