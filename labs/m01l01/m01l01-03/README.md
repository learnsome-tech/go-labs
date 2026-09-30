# m01l01-03 · The build step, and what it produces

**Lesson:** [What Go Is and Why It Matters](https://learnsome.tech/learn/go-course/m01l01) (lesson 1.1, module 1: The Go Toolchain and Fundamentals) · Free  
**Check:** Runs, not graded

## Goal

You can explain why platform and DevOps teams write their tooling in Go, compile and run a program with the go command, and describe exactly what the compiler hands you at the end of a build.

In the lesson: Now the same program from the shell. First ask the toolchain which version is installed, which also tells you the operating system and processor it will build for by default. Then build with an output name of your choosing. That command prints nothing at all, which in Go tooling means it worked. What it leaves behind is a file: machine code for this platform, with the Go runtime and every package you imported already inside it. Run the binary directly and you get the same two lines. Notice what you did not do: there was no interpreter, no virtual machine and no list of dependencies to install on the machine that runs it.

## Files

- [`starter/README.md`](starter/README.md)
- [`starter/go.mod`](starter/go.mod)
- [`starter/main.go`](starter/main.go): the listing from the lesson
- [`starter/run.sh`](starter/run.sh): the command the lesson ran
- [`starter/verification.json`](starter/verification.json)
- [`check.json`](check.json): how `./check` runs and checks this lab

## Steps

1. Go to the starter: `cd labs/m01l01/m01l01-03/starter`
2. Read `main.go`.
3. Run it: `go run .`.
4. Check it from the repository root: `./check m01l01-03`.

## How to check

`./check m01l01-03` copies `starter/` into a scratch directory and runs `go run .` there, the way the site's lab sandbox does: that directory is the working directory and `HOME`, `LANG=C.UTF-8`, `TZ=UTC`, a limit of 10 seconds and 256 KiB of output per stream.

It runs without a pass or fail: the lesson recorded its output with `go version; go build -o probe .; ./probe`, not the way the site runs it, so the output is not compared. `./check` shows the output and the exit code.

---

[Open the lesson on LearnSome.tech](https://learnsome.tech/learn/go-course/m01l01) · [All labs of this lesson](../README.md) · [Course README](../../../README.md)
