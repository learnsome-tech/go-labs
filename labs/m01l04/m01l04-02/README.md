# m01l04-02 · Watching the three verbs behave

**Lesson:** [Your First Program and go run](https://learnsome.tech/learn/go-course/m01l04) (lesson 1.4, module 1: The Go Toolchain and Fundamentals) · Free  
**Check:** Runs, not graded

## Goal

You can choose correctly between running, building and installing a program, write functions that return a value alongside an error, format output with the verbs you will use daily, and set an exit code that a pipeline can act on.

In the lesson: Same program, three ways. Compile and throw away first: you get the output and nothing is left behind. Then keep the result with an output name, which prints nothing and leaves an executable. Run that executable and the output is identical, but this time no toolchain is involved at all, which is exactly what happens on the host you deploy to. The last one is a trap worth meeting now: extra words after the package are passed to your program, not to the compiler, so the compiler will not complain about them and your program has to decide what they mean. We will read them properly when we get to the flag package.

## Files

- [`starter/README.md`](starter/README.md)
- [`starter/go.mod`](starter/go.mod)
- [`starter/main.go`](starter/main.go): the listing from the lesson
- [`starter/run.sh`](starter/run.sh): the command the lesson ran
- [`starter/verification.json`](starter/verification.json)
- [`check.json`](check.json): how `./check` runs and checks this lab

## Steps

1. Go to the starter: `cd labs/m01l04/m01l04-02/starter`
2. Read `main.go`.
3. Run it: `go run .`.
4. Check it from the repository root: `./check m01l04-02`.

## How to check

`./check m01l04-02` copies `starter/` into a scratch directory and runs `go run .` there, the way the site's lab sandbox does: that directory is the working directory and `HOME`, `LANG=C.UTF-8`, `TZ=UTC`, a limit of 10 seconds and 256 KiB of output per stream.

It runs without a pass or fail: the lesson recorded its output with `go run .; go build -o diskcheck .; ./diskcheck; go run . more arguments`, not the way the site runs it, so the output is not compared. `./check` shows the output and the exit code.

---

[Open the lesson on LearnSome.tech](https://learnsome.tech/learn/go-course/m01l04) · [All labs of this lesson](../README.md) · [Course README](../../../README.md)
