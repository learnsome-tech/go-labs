# m06l03-04 · A bounded shutdown call

**Lesson:** [Graceful Shutdown on SIGTERM](https://learnsome.tech/learn/go-course/m06l03) (lesson 6.3, module 6: Building CLIs and Services) · Pro  
**Check:** Graded

## Goal

You can stop accepting traffic on SIGTERM, cancel active work, and let a Go service finish cleanly.

In the lesson: The server has no listener in this small example, so Shutdown returns nil immediately. The important shape is the context with a finite deadline and the deferred cancel that releases its timer. In a running service, call this after SIGTERM, then wait for the shutdown result before returning from main.

## Files

- [`starter/README.md`](starter/README.md)
- [`starter/go.mod`](starter/go.mod)
- [`starter/main.go`](starter/main.go): the listing from the lesson
- [`starter/run.sh`](starter/run.sh): the command the lesson ran
- [`starter/verification.json`](starter/verification.json)
- [`expected.txt`](expected.txt): the output the check compares with
- [`check.json`](check.json): how `./check` runs and checks this lab

## Steps

1. Go to the starter: `cd labs/m06l03/m06l03-04/starter`
2. Read `main.go`.
3. Run it: `go run .`.
4. Check it from the repository root: `./check m06l03-04`.

## Expected output

```text
<nil>
```

## How to check

`./check m06l03-04` copies `starter/` into a scratch directory and runs `go run .` there, the way the site's lab sandbox does: that directory is the working directory and `HOME`, `LANG=C.UTF-8`, `TZ=UTC`, a limit of 10 seconds and 256 KiB of output per stream.

It passes when the output matches `expected.txt` by the site's rules, within the limits. Standard output is compared line by line; spaces at the end of a line and blank lines at the end do not count. If that differs, standard output followed by standard error is compared with Python traceback frames and blank lines set aside, so a lesson that shows an error passes when your program prints the same error. A pass here is a pass on the site.

---

[Open the lesson on LearnSome.tech](https://learnsome.tech/learn/go-course/m06l03) · [All labs of this lesson](../README.md) · [Course README](../../../README.md)
