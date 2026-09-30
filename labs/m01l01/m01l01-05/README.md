# m01l01-05 · The compiler refuses an unused import

**Lesson:** [What Go Is and Why It Matters](https://learnsome.tech/learn/go-course/m01l01) (lesson 1.1, module 1: The Go Toolchain and Fundamentals) · Free  
**Check:** Runs, not graded

## Goal

You can explain why platform and DevOps teams write their tooling in Go, compile and run a program with the go command, and describe exactly what the compiler hands you at the end of a build.

In the lesson: One early surprise, and it says a lot about the language. This program imports two packages but only uses one. In most languages that is a warning you learn to ignore. Try to build that in Go and the compiler refuses outright: the import of o s is there and unused, so the build fails with the file, the line and the column. The reasoning is deliberate. An unused import is either a leftover from code you deleted or a mistake, and in both cases it misleads the next person reading the file. The language would rather stop you now than let dead references pile up in a tool your whole team depends on.

## Files

- [`starter/README.md`](starter/README.md)
- [`starter/go.mod`](starter/go.mod)
- [`starter/main.go`](starter/main.go): the listing from the lesson
- [`starter/run.sh`](starter/run.sh): the command the lesson ran
- [`starter/verification.json`](starter/verification.json)
- [`check.json`](check.json): how `./check` runs and checks this lab

## Steps

1. Go to the starter: `cd labs/m01l01/m01l01-05/starter`
2. Read `main.go`.
3. Run it: `go run .`.
4. Check it from the repository root: `./check m01l01-05`.

## What the lesson recorded

Shown for reference; the check does not compare it.

```text
# course.local/example
./main.go:5:2: "os" imported and not used
```

## How to check

`./check m01l01-05` copies `starter/` into a scratch directory and runs `go run .` there, the way the site's lab sandbox does: that directory is the working directory and `HOME`, `LANG=C.UTF-8`, `TZ=UTC`, a limit of 10 seconds and 256 KiB of output per stream.

It runs without a pass or fail: the lesson recorded its output with `go build ./...`, not the way the site runs it, so the output is not compared. `./check` shows the output and the exit code.

---

[Open the lesson on LearnSome.tech](https://learnsome.tech/learn/go-course/m01l01) · [All labs of this lesson](../README.md) · [Course README](../../../README.md)
