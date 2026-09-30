# m03l01-02 · Returning a value and an error

**Lesson:** [Errors as Values](https://learnsome.tech/learn/go-course/m03l01) (lesson 3.1, module 3: Errors and Testing) · Pro  
**Check:** Graded

## Goal

You can return errors, inspect them, and keep failure handling explicit in a command line tool.

In the lesson: Divide returns two results. A successful calculation returns the quotient and a nil error, while a zero count returns a zero value and a descriptive error. The caller checks the error immediately. Only after that check does it print the result. This ordering prevents a bad result from travelling further into the program. In a health checker, the same shape lets a probe return a response and an error, leaving the caller in charge of whether a failed probe should be retried or reported.

## Files

- [`starter/README.md`](starter/README.md)
- [`starter/go.mod`](starter/go.mod)
- [`starter/main.go`](starter/main.go): the listing from the lesson
- [`starter/run.sh`](starter/run.sh): the command the lesson ran
- [`starter/verification.json`](starter/verification.json)
- [`expected.txt`](expected.txt): the output the check compares with
- [`check.json`](check.json): how `./check` runs and checks this lab

## Steps

1. Go to the starter: `cd labs/m03l01/m03l01-02/starter`
2. Read `main.go`.
3. Run it: `go run .`.
4. Check it from the repository root: `./check m03l01-02`.

## Expected output

```text
result: 5
```

## How to check

`./check m03l01-02` copies `starter/` into a scratch directory and runs `go run .` there, the way the site's lab sandbox does: that directory is the working directory and `HOME`, `LANG=C.UTF-8`, `TZ=UTC`, a limit of 10 seconds and 256 KiB of output per stream.

It passes when the output matches `expected.txt` by the site's rules, within the limits. Standard output is compared line by line; spaces at the end of a line and blank lines at the end do not count. If that differs, standard output followed by standard error is compared with Python traceback frames and blank lines set aside, so a lesson that shows an error passes when your program prints the same error. A pass here is a pass on the site.

---

[Open the lesson on LearnSome.tech](https://learnsome.tech/learn/go-course/m03l01) · [All labs of this lesson](../README.md) · [Course README](../../../README.md)
