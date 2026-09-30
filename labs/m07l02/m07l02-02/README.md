# m07l02-02 · Collecting concurrent results

**Lesson:** [Executing Concurrent Checks](https://learnsome.tech/learn/go-course/m07l02) (lesson 7.2, module 7: Project: End to End Health Checker) · Pro  
**Check:** Runs, not graded

## Goal

You can run checks concurrently, collect one result per target, and preserve a useful report order.

In the lesson: The buffered result channel can hold both worker results while the coordinator waits. Each worker receives its own target name, sends a report, and marks its work done. After the wait group reaches zero, the coordinator closes the channel and ranges over the complete stream. In a production report, sort or index results if output order must match configuration order.

## Files

- [`starter/README.md`](starter/README.md)
- [`starter/go.mod`](starter/go.mod)
- [`starter/main.go`](starter/main.go): the listing from the lesson
- [`starter/run.sh`](starter/run.sh): the command the lesson ran
- [`starter/verification.json`](starter/verification.json)
- [`check.json`](check.json): how `./check` runs and checks this lab

## Steps

1. Go to the starter: `cd labs/m07l02/m07l02-02/starter`
2. Read `main.go`.
3. Run it: `go run .`.
4. Check it from the repository root: `./check m07l02-02`.

## What the lesson recorded

Shown for reference; the check does not compare it.

```text
api checked
web checked
```

## How to check

`./check m07l02-02` copies `starter/` into a scratch directory and runs `go run .` there, the way the site's lab sandbox does: that directory is the working directory and `HOME`, `LANG=C.UTF-8`, `TZ=UTC`, a limit of 10 seconds and 256 KiB of output per stream.

It runs without a pass or fail: the recorded output depends on the machine it ran on, so the site runs it without a pass or fail. `./check` shows the output and the exit code.

---

[Open the lesson on LearnSome.tech](https://learnsome.tech/learn/go-course/m07l02) · [All labs of this lesson](../README.md) · [Course README](../../../README.md)
