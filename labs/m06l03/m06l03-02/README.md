# m06l03-02 · Waiting for SIGTERM

**Lesson:** [Graceful Shutdown on SIGTERM](https://learnsome.tech/learn/go-course/m06l03) (lesson 6.3, module 6: Building CLIs and Services) · Pro  
**Check:** Runs, not graded

## Goal

You can stop accepting traffic on SIGTERM, cancel active work, and let a Go service finish cleanly.

In the lesson: The signal channel receives SIGTERM from the operating system. Main announces ready, then blocks until the signal arrives. The stopping line is the start of shutdown work, such as calling server Shutdown and cancelling a root context. A buffered channel prevents the signal sender from waiting while main is between setup and receive.

## Files

- [`starter/README.md`](starter/README.md)
- [`starter/go.mod`](starter/go.mod)
- [`starter/main.go`](starter/main.go): the listing from the lesson
- [`starter/run.sh`](starter/run.sh): the command the lesson ran
- [`starter/verification.json`](starter/verification.json)
- [`check.json`](check.json): how `./check` runs and checks this lab

## Steps

1. Go to the starter: `cd labs/m06l03/m06l03-02/starter`
2. Read `main.go`.
3. Run it: `go run .`.
4. Check it from the repository root: `./check m06l03-02`.

## What the lesson recorded

Shown for reference; the check does not compare it.

```text
ready
```

## How to check

`./check m06l03-02` copies `starter/` into a scratch directory and runs `go run .` there, the way the site's lab sandbox does: that directory is the working directory and `HOME`, `LANG=C.UTF-8`, `TZ=UTC`, a limit of 10 seconds and 256 KiB of output per stream.

It runs without a pass or fail: the recorded output depends on the machine it ran on, so the site runs it without a pass or fail. `./check` shows the output and the exit code.

---

[Open the lesson on LearnSome.tech](https://learnsome.tech/learn/go-course/m06l03) · [All labs of this lesson](../README.md) · [Course README](../../../README.md)
