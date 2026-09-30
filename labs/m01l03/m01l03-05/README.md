# m01l03-05 · Tidy and vendor, and what they are for

**Lesson:** [Go Workspaces and Modules](https://learnsome.tech/learn/go-course/m01l03) (lesson 1.3, module 1: The Go Toolchain and Fundamentals) · Free  
**Check:** Runs, not graded

## Goal

You can start a module, split a tool into packages with import paths that follow from the module path, keep dependencies tidy and vendored, and use a workspace file when a tool and a shared library sit in one checkout.

In the lesson: Two dependency commands earn their place in a pipeline. The first adds what you import and removes what you no longer import, rewriting the requirements to match the code as it is now rather than as it was. Run it before you commit; a diff on that file is a dependency review. The second copies every dependency into a directory inside the repository, so the build needs no network at all. That matters when your build agent has no route to the internet, or when you want an audit trail of exactly what went into a release. Our example imports only the standard library, so there is nothing to copy, and the build carries on regardless.

## Files

- [`starter/README.md`](starter/README.md)
- [`starter/go.mod`](starter/go.mod)
- [`starter/main.go`](starter/main.go): the listing from the lesson
- [`starter/run.sh`](starter/run.sh): the command the lesson ran
- [`starter/verification.json`](starter/verification.json)
- [`check.json`](check.json): how `./check` runs and checks this lab

## Steps

1. Go to the starter: `cd labs/m01l03/m01l03-05/starter`
2. Read `main.go`.
3. Run it: `go run .`.
4. Check it from the repository root: `./check m01l03-05`.

## How to check

`./check m01l03-05` copies `starter/` into a scratch directory and runs `go run .` there, the way the site's lab sandbox does: that directory is the working directory and `HOME`, `LANG=C.UTF-8`, `TZ=UTC`, a limit of 10 seconds and 256 KiB of output per stream.

It runs without a pass or fail: the lesson recorded its output with `go mod tidy; go mod vendor; go build ./...`, not the way the site runs it, so the output is not compared. `./check` shows the output and the exit code.

---

[Open the lesson on LearnSome.tech](https://learnsome.tech/learn/go-course/m01l03) · [All labs of this lesson](../README.md) · [Course README](../../../README.md)
