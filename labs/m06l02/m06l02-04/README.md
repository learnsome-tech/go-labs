# m06l02-04 · Testing a handler in memory

**Lesson:** [Serving HTTP with net/http Server](https://learnsome.tech/learn/go-course/m06l02) (lesson 6.2, module 6: Building CLIs and Services) · Pro  
**Check:** Runs, not graded

## Goal

You can expose a small HTTP handler, return useful status codes, and configure a server with a timeout.

In the lesson: The recorder and request let us exercise the handler without opening a port. After the call, the recorder exposes the status code and body that a real client would receive. In a test this same setup gives a deterministic check for status, headers, and JSON. Keep network binding out of unit tests and reserve it for an integration check.

## Files

- [`starter/README.md`](starter/README.md)
- [`starter/go.mod`](starter/go.mod)
- [`starter/main.go`](starter/main.go): the listing from the lesson
- [`starter/run.sh`](starter/run.sh): the command the lesson ran
- [`starter/verification.json`](starter/verification.json)
- [`check.json`](check.json): how `./check` runs and checks this lab

## Steps

1. Go to the starter: `cd labs/m06l02/m06l02-04/starter`
2. Read `main.go`.
3. Run it: `go run .`.
4. Check it from the repository root: `./check m06l02-04`.

## What the lesson recorded

Shown for reference; the check does not compare it.

```text
200 ok

```

## How to check

`./check m06l02-04` copies `starter/` into a scratch directory and runs `go run .` there, the way the site's lab sandbox does: that directory is the working directory and `HOME`, `LANG=C.UTF-8`, `TZ=UTC`, a limit of 10 seconds and 256 KiB of output per stream.

It runs without a pass or fail: the recorded output depends on the machine it ran on, so the site runs it without a pass or fail. `./check` shows the output and the exit code.

---

[Open the lesson on LearnSome.tech](https://learnsome.tech/learn/go-course/m06l02) · [All labs of this lesson](../README.md) · [Course README](../../../README.md)
