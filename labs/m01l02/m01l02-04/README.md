# m01l02-04 · Formatting is settled, not argued about

**Lesson:** [Installing the Go Toolchain](https://learnsome.tech/learn/go-course/m01l02) (lesson 1.2, module 1: The Go Toolchain and Fundamentals) · Free  
**Check:** Runs, not graded

## Goal

You can install the Go toolchain the way that suits a laptop, a build agent or a container, and use the go command's own subcommands for formatting, checking and documentation instead of reaching for extra tools.

In the lesson: Go ships one formatter and no configuration for it. Tabs for indentation, one true brace style, and no discussion. The first command lists which files disagree with the canonical form, and here there is one. The second rewrites it in place. Run the listing again and there is nothing left to say, which is how this command reports success. For a team this removes a whole category of review comment, and in a pipeline the listing form is the check you want: if it prints any file name, somebody committed unformatted code, so fail the build and move on. Most editors run this on save, and you should let yours.

## Files

- [`starter/README.md`](starter/README.md)
- [`starter/go.mod`](starter/go.mod)
- [`starter/main.go`](starter/main.go): the listing from the lesson
- [`starter/run.sh`](starter/run.sh): the command the lesson ran
- [`starter/verification.json`](starter/verification.json)
- [`check.json`](check.json): how `./check` runs and checks this lab

## Steps

1. Go to the starter: `cd labs/m01l02/m01l02-04/starter`
2. Read `main.go`.
3. Run it: `go run .`.
4. Check it from the repository root: `./check m01l02-04`.

## How to check

`./check m01l02-04` copies `starter/` into a scratch directory and runs `go run .` there, the way the site's lab sandbox does: that directory is the working directory and `HOME`, `LANG=C.UTF-8`, `TZ=UTC`, a limit of 10 seconds and 256 KiB of output per stream.

It runs without a pass or fail: the lesson recorded its output with `gofmt -l .; gofmt -w main.go; gofmt -l .`, not the way the site runs it, so the output is not compared. `./check` shows the output and the exit code.

---

[Open the lesson on LearnSome.tech](https://learnsome.tech/learn/go-course/m01l02) · [All labs of this lesson](../README.md) · [Course README](../../../README.md)
