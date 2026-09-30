# m06l01-02 · Declaring typed options and parsing them

**Lesson:** [Parsing Arguments with flag](https://learnsome.tech/learn/go-course/m06l01) (lesson 6.1, module 6: Building CLIs and Services) · Pro  
**Check:** Runs, not graded

## Goal

You can build a command line tool with the standard flag package that takes typed options, reads positional arguments, falls back to environment variables, and fails with a usage message and the right exit code.

In the lesson: Here is the shape of every tool built on this package. After the imports, each option is declared before parsing, and the declaration returns a pointer to the value that parsing will fill in later. We declare each option in turn: a string for the address, an integer for the retry count, a boolean for verbosity, and a duration for the deadline on each attempt. Notice the duration: the package turns two seconds or five hundred milliseconds into a real duration value for you, which is exactly the arithmetic you would hand roll in Bash. Then we call Parse, which walks the arguments and fills every pointer. Only after that are the values safe to read, so we print what we parsed. Let us run it with values for three options and let the fourth fall back to its default.

## Files

- [`starter/README.md`](starter/README.md)
- [`starter/go.mod`](starter/go.mod)
- [`starter/main.go`](starter/main.go): the listing from the lesson
- [`starter/run.sh`](starter/run.sh): the command the lesson ran
- [`starter/verification.json`](starter/verification.json)
- [`check.json`](check.json): how `./check` runs and checks this lab

## Steps

1. Go to the starter: `cd labs/m06l01/m06l01-02/starter`
2. Read `main.go` the way the lesson builds it:
   - Lines 1–8: the imports
   - Lines 9–13: declare each option in turn
   - Lines 14: we call Parse
   - Lines 15–19: we print what we parsed
3. Notes from the lesson:
   - Line 10: Declaring returns a pointer; the value is valid only after Parse
   - Line 13: Duration parses two seconds or five hundred milliseconds for you
4. Run it: `go run .`.
5. Check it from the repository root: `./check m06l01-02`.

## What the lesson recorded

Shown for reference; the check does not compare it.

```text
addr: 10.0.0.9:443
tries: 5
verbose: true
timeout: 2s
```

## How to check

`./check m06l01-02` copies `starter/` into a scratch directory and runs `go run .` there, the way the site's lab sandbox does: that directory is the working directory and `HOME`, `LANG=C.UTF-8`, `TZ=UTC`, a limit of 10 seconds and 256 KiB of output per stream.

It runs without a pass or fail: the recorded output depends on the machine it ran on, so the site runs it without a pass or fail. `./check` shows the output and the exit code.

---

[Open the lesson on LearnSome.tech](https://learnsome.tech/learn/go-course/m06l01) · [All labs of this lesson](../README.md) · [Course README](../../../README.md)
