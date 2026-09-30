# m04l01-04 · Workers receive independent values

**Lesson:** [Goroutines and Concurrent Execution](https://learnsome.tech/learn/go-course/m04l01) (lesson 4.1, module 4: Concurrency and Context) · Pro  
**Check:** Runs, not graded

## Goal

You can start goroutines, wait for their work, and recognise the ownership questions concurrency introduces.

In the lesson: Each loop value is passed as an argument to the goroutine, so the worker owns its own copy of the target name. The wait group counts both jobs and main waits for both to finish. Their scheduling order is not a contract, even though this small run may print the same order repeatedly. When output order matters, collect results through a coordination mechanism instead of relying on timing.

## Files

- [`starter/README.md`](starter/README.md)
- [`starter/go.mod`](starter/go.mod)
- [`starter/main.go`](starter/main.go): the listing from the lesson
- [`starter/run.sh`](starter/run.sh): the command the lesson ran
- [`starter/verification.json`](starter/verification.json)
- [`check.json`](check.json): how `./check` runs and checks this lab

## Steps

1. Go to the starter: `cd labs/m04l01/m04l01-04/starter`
2. Read `main.go`.
3. Run it: `go run .`.
4. Check it from the repository root: `./check m04l01-04`.

## What the lesson recorded

Shown for reference; the check does not compare it.

```text
api
worker
```

## How to check

`./check m04l01-04` copies `starter/` into a scratch directory and runs `go run .` there, the way the site's lab sandbox does: that directory is the working directory and `HOME`, `LANG=C.UTF-8`, `TZ=UTC`, a limit of 10 seconds and 256 KiB of output per stream.

It runs without a pass or fail: the recorded output depends on the machine it ran on, so the site runs it without a pass or fail. `./check` shows the output and the exit code.

---

[Open the lesson on LearnSome.tech](https://learnsome.tech/learn/go-course/m04l01) · [All labs of this lesson](../README.md) · [Course README](../../../README.md)
