# m04l05-04 · A caller cancels a worker

**Lesson:** [Context Cancellation and Timeouts](https://learnsome.tech/learn/go-course/m04l05) (lesson 4.5, module 4: Concurrency and Context) · Pro  
**Check:** Graded

## Goal

You can bound a concurrent operation with a timeout and return promptly when an orchestrator or caller cancels it.

In the lesson: The worker blocks on the context Done signal instead of sleeping forever. Main starts it, cancels the context, and waits for a separate done signal so the process knows the worker has returned. The worker prints its stop message before the program exits. In a service, the same handoff lets shutdown cancel active requests and then wait for their cleanup.

## Files

- [`starter/README.md`](starter/README.md)
- [`starter/go.mod`](starter/go.mod)
- [`starter/main.go`](starter/main.go): the listing from the lesson
- [`starter/run.sh`](starter/run.sh): the command the lesson ran
- [`starter/verification.json`](starter/verification.json)
- [`expected.txt`](expected.txt): the output the check compares with
- [`check.json`](check.json): how `./check` runs and checks this lab

## Steps

1. Go to the starter: `cd labs/m04l05/m04l05-04/starter`
2. Read `main.go`.
3. Run it: `go run .`.
4. Check it from the repository root: `./check m04l05-04`.

## Expected output

```text
worker stopped
```

## How to check

`./check m04l05-04` copies `starter/` into a scratch directory and runs `go run .` there, the way the site's lab sandbox does: that directory is the working directory and `HOME`, `LANG=C.UTF-8`, `TZ=UTC`, a limit of 10 seconds and 256 KiB of output per stream.

It passes when the output matches `expected.txt` by the site's rules, within the limits. Standard output is compared line by line; spaces at the end of a line and blank lines at the end do not count. If that differs, standard output followed by standard error is compared with Python traceback frames and blank lines set aside, so a lesson that shows an error passes when your program prints the same error. A pass here is a pass on the site.

---

[Open the lesson on LearnSome.tech](https://learnsome.tech/learn/go-course/m04l05) · [All labs of this lesson](../README.md) · [Course README](../../../README.md)
