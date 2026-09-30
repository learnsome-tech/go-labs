# m06l01-04 · Positional arguments are what is left over

**Lesson:** [Parsing Arguments with flag](https://learnsome.tech/learn/go-course/m06l01) (lesson 6.1, module 6: Building CLIs and Services) · Pro  
**Check:** Runs, not graded

## Goal

You can build a command line tool with the standard flag package that takes typed options, reads positional arguments, falls back to environment variables, and fails with a usage message and the right exit code.

In the lesson: Flags are the options; the positional arguments are the things you are acting on. Call Args after parsing and you get everything that was not consumed as a flag, in order, as a slice of strings. Here the tool wants at least one host, so an empty slice is a misuse of the command line: a message on standard error and exit two, matching what the package itself does for a bad flag. One habit worth forming: parse first, validate second, work third. Validation failures are cheap and belong before you open a socket or a file, not halfway through a run. Let us run it with one flag and three hosts and watch the split.

## Files

- [`starter/README.md`](starter/README.md)
- [`starter/go.mod`](starter/go.mod)
- [`starter/main.go`](starter/main.go): the listing from the lesson
- [`starter/run.sh`](starter/run.sh): the command the lesson ran
- [`starter/verification.json`](starter/verification.json)
- [`check.json`](check.json): how `./check` runs and checks this lab

## Steps

1. Go to the starter: `cd labs/m06l01/m06l01-04/starter`
2. Read `main.go` the way the lesson builds it:
   - Lines 1–7: the things you are acting on
   - Lines 8–12: Call Args after parsing
   - Lines 13–16: an empty slice is a misuse
   - Lines 17–21: parse first, validate second, work third
3. Notes from the lesson:
   - Line 12: Args holds everything Parse did not consume, in order
   - Line 15: Exit two is the conventional status for a misused command line
4. Run it: `go run .`.
5. Check it from the repository root: `./check m06l01-04`.

## What the lesson recorded

Shown for reference; the check does not compare it.

```text
hosts: 3 tries: 2
1 web-1
2 web-2
3 db-1
```

## How to check

`./check m06l01-04` copies `starter/` into a scratch directory and runs `go run .` there, the way the site's lab sandbox does: that directory is the working directory and `HOME`, `LANG=C.UTF-8`, `TZ=UTC`, a limit of 10 seconds and 256 KiB of output per stream.

It runs without a pass or fail: the recorded output depends on the machine it ran on, so the site runs it without a pass or fail. `./check` shows the output and the exit code.

---

[Open the lesson on LearnSome.tech](https://learnsome.tech/learn/go-course/m06l01) · [All labs of this lesson](../README.md) · [Course README](../../../README.md)
