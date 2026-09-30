# m01l04-06 · Exit codes are the contract with your pipeline

**Lesson:** [Your First Program and go run](https://learnsome.tech/learn/go-course/m01l04) (lesson 1.4, module 1: The Go Toolchain and Fundamentals) · Free  
**Check:** Runs, not graded

## Goal

You can choose correctly between running, building and installing a program, write functions that return a value alongside an error, format output with the verbs you will use daily, and set an exit code that a pipeline can act on.

In the lesson: Here is the part that makes a program a tool rather than a script somebody runs by hand. A pipeline does not read your output; it reads your exit status. Zero means success and anything else means failure, and that single number decides whether a deployment continues. So the summary for humans goes to standard output, the reason for failure goes to standard error, and the process exits with one. Run it and look at the two streams together. Two rules come with this: exit immediately terminates the process, so deferred cleanup does not run, and choose your codes deliberately, because somebody will eventually write a pipeline condition against them.

## Files

- [`starter/README.md`](starter/README.md)
- [`starter/go.mod`](starter/go.mod)
- [`starter/main.go`](starter/main.go): the listing from the lesson
- [`starter/run.sh`](starter/run.sh): the command the lesson ran
- [`starter/verification.json`](starter/verification.json)
- [`check.json`](check.json): how `./check` runs and checks this lab

## Steps

1. Go to the starter: `cd labs/m01l04/m01l04-06/starter`
2. Read `main.go`.
3. Run it: `go run .`.
4. Check it from the repository root: `./check m01l04-06`.

## What the lesson recorded

Shown for reference; the check does not compare it.

```text
checked three targets
2 unhealthy
```

## How to check

`./check m01l04-06` copies `starter/` into a scratch directory and runs `go run .` there, the way the site's lab sandbox does: that directory is the working directory and `HOME`, `LANG=C.UTF-8`, `TZ=UTC`, a limit of 10 seconds and 256 KiB of output per stream.

It runs without a pass or fail: the lesson recorded its output with `go build -o diskcheck . && ./diskcheck`, not the way the site runs it, so the output is not compared. `./check` shows the output and the exit code.

---

[Open the lesson on LearnSome.tech](https://learnsome.tech/learn/go-course/m01l04) · [All labs of this lesson](../README.md) · [Course README](../../../README.md)
