# m01l04-03 · A function that returns a value and an error

**Lesson:** [Your First Program and go run](https://learnsome.tech/learn/go-course/m01l04) (lesson 1.4, module 1: The Go Toolchain and Fundamentals) · Free  
**Check:** Graded

## Goal

You can choose correctly between running, building and installing a program, write functions that return a value alongside an error, format output with the verbs you will use daily, and set an exit code that a pipeline can act on.

In the lesson: This is the shape of almost every function you will write in Go. Two packages imported, then a function taking two integers and returning two results: the answer, and an error. There is no exception to throw, so the second result carries the bad news, and a nil error is how the function says nothing went wrong. Look at the parameter list: two parameters of the same type share one type name. Then call it twice, once with sensible numbers and once with a total of zero. The first call gives a percentage and a nil error, which prints inside angle brackets. The second gives the zero value and a message. Printing both together is a teaching shortcut; real code checks the error and branches.

## Files

- [`starter/README.md`](starter/README.md)
- [`starter/go.mod`](starter/go.mod)
- [`starter/main.go`](starter/main.go): the listing from the lesson
- [`starter/run.sh`](starter/run.sh): the command the lesson ran
- [`starter/verification.json`](starter/verification.json)
- [`expected.txt`](expected.txt): the output the check compares with
- [`check.json`](check.json): how `./check` runs and checks this lab

## Steps

1. Go to the starter: `cd labs/m01l04/m01l04-03/starter`
2. Read `main.go` the way the lesson builds it:
   - Lines 1–6: two packages
   - Lines 7–13: two results
   - Lines 14–21: call it twice
3. Notes from the lesson:
   - Line 8: Two parameters of one type collapse to one declaration
   - Line 12: A nil error is how a function says nothing went wrong
4. Run it: `go run .`.
5. Check it from the repository root: `./check m01l04-03`.

## Expected output

```text
82 <nil>
0 total must be positive
```

## How to check

`./check m01l04-03` copies `starter/` into a scratch directory and runs `go run .` there, the way the site's lab sandbox does: that directory is the working directory and `HOME`, `LANG=C.UTF-8`, `TZ=UTC`, a limit of 10 seconds and 256 KiB of output per stream.

It passes when the output matches `expected.txt` by the site's rules, within the limits. Standard output is compared line by line; spaces at the end of a line and blank lines at the end do not count. If that differs, standard output followed by standard error is compared with Python traceback frames and blank lines set aside, so a lesson that shows an error passes when your program prints the same error. A pass here is a pass on the site.

---

[Open the lesson on LearnSome.tech](https://learnsome.tech/learn/go-course/m01l04) · [All labs of this lesson](../README.md) · [Course README](../../../README.md)
