# m04l01-02 · Starting work with go

**Lesson:** [Goroutines and Concurrent Execution](https://learnsome.tech/learn/go-course/m04l01) (lesson 4.1, module 4: Concurrency and Context) · Pro  
**Check:** Graded

## Goal

You can start goroutines, wait for their work, and recognise the ownership questions concurrency introduces.

In the lesson: The anonymous function starts as a goroutine. A wait group records that one piece of work exists, and the deferred Done call tells the group when the worker returns. Main waits before printing its own line, so the process cannot exit early. The output shows the worker completing before main finishes. A wait group is useful when a caller starts a known number of independent jobs and needs to join them before returning a result.

## Files

- [`starter/README.md`](starter/README.md)
- [`starter/go.mod`](starter/go.mod)
- [`starter/main.go`](starter/main.go): the listing from the lesson
- [`starter/run.sh`](starter/run.sh): the command the lesson ran
- [`starter/verification.json`](starter/verification.json)
- [`expected.txt`](expected.txt): the output the check compares with
- [`check.json`](check.json): how `./check` runs and checks this lab

## Steps

1. Go to the starter: `cd labs/m04l01/m04l01-02/starter`
2. Read `main.go`.
3. Run it: `go run .`.
4. Check it from the repository root: `./check m04l01-02`.

## Expected output

```text
worker finished
main finished
```

## How to check

`./check m04l01-02` copies `starter/` into a scratch directory and runs `go run .` there, the way the site's lab sandbox does: that directory is the working directory and `HOME`, `LANG=C.UTF-8`, `TZ=UTC`, a limit of 10 seconds and 256 KiB of output per stream.

It passes when the output matches `expected.txt` by the site's rules, within the limits. Standard output is compared line by line; spaces at the end of a line and blank lines at the end do not count. If that differs, standard output followed by standard error is compared with Python traceback frames and blank lines set aside, so a lesson that shows an error passes when your program prints the same error. A pass here is a pass on the site.

---

[Open the lesson on LearnSome.tech](https://learnsome.tech/learn/go-course/m04l01) · [All labs of this lesson](../README.md) · [Course README](../../../README.md)
