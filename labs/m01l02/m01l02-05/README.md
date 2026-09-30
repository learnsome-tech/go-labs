# m01l02-05 · Vet catches what the compiler allows

**Lesson:** [Installing the Go Toolchain](https://learnsome.tech/learn/go-course/m01l02) (lesson 1.2, module 1: The Go Toolchain and Fundamentals) · Free  
**Check:** Runs, not graded

## Goal

You can install the Go toolchain the way that suits a laptop, a build agent or a container, and use the go command's own subcommands for formatting, checking and documentation instead of reaching for extra tools.

In the lesson: This program compiles perfectly. It is also wrong: the format string asks for a decimal number and the argument is a string, so at runtime you get mangled output instead of a host name. The compiler has no opinion, because printing takes any arguments. Run vet over it and you get the file, the line, the verb and the type that does not match. Vet is a collection of checks for exactly this sort of thing: format strings that do not line up with their arguments, unreachable code, locks copied by value, struct tags that will not decode. It is part of the toolchain, it runs in about a second, and it belongs in your pipeline next to the tests.

## Files

- [`starter/README.md`](starter/README.md)
- [`starter/go.mod`](starter/go.mod)
- [`starter/main.go`](starter/main.go): the listing from the lesson
- [`starter/run.sh`](starter/run.sh): the command the lesson ran
- [`starter/verification.json`](starter/verification.json)
- [`check.json`](check.json): how `./check` runs and checks this lab

## Steps

1. Go to the starter: `cd labs/m01l02/m01l02-05/starter`
2. Read `main.go`.
3. Run it: `go run .`.
4. Check it from the repository root: `./check m01l02-05`.

## What the lesson recorded

Shown for reference; the check does not compare it.

```text
# course.local/example
# [course.local/example]
./main.go:7:2: fmt.Printf format %d has arg host of wrong type string
```

## How to check

`./check m01l02-05` copies `starter/` into a scratch directory and runs `go run .` there, the way the site's lab sandbox does: that directory is the working directory and `HOME`, `LANG=C.UTF-8`, `TZ=UTC`, a limit of 10 seconds and 256 KiB of output per stream.

It runs without a pass or fail: the lesson recorded its output with `go vet ./...`, not the way the site runs it, so the output is not compared. `./check` shows the output and the exit code.

---

[Open the lesson on LearnSome.tech](https://learnsome.tech/learn/go-course/m01l02) · [All labs of this lesson](../README.md) · [Course README](../../../README.md)
