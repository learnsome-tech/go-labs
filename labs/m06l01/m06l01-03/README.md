# m06l01-03 · Getting it wrong: usage and exit code two

**Lesson:** [Parsing Arguments with flag](https://learnsome.tech/learn/go-course/m06l01) (lesson 6.1, module 6: Building CLIs and Services) · Pro  
**Check:** Runs, not graded

## Goal

You can build a command line tool with the standard flag package that takes typed options, reads positional arguments, falls back to environment variables, and fails with a usage message and the right exit code.

In the lesson: Callers get flags wrong, and what happens then is part of your tool's interface. The package already does the sensible thing: an unknown flag stops parsing, prints a message, prints the usage block, and exits with status two. That usage block is generated from the declarations, which is why a decent help string next to each option pays for itself. You can replace the default heading by assigning your own function to Usage, and look at where it writes: standard error, not standard output, so a pipeline consuming your real output is never polluted by help text. Let us pass a flag that does not exist and read the whole thing, exit status and all.

## Files

- [`starter/README.md`](starter/README.md)
- [`starter/go.mod`](starter/go.mod)
- [`starter/main.go`](starter/main.go): the listing from the lesson
- [`starter/run.sh`](starter/run.sh): the command the lesson ran
- [`starter/verification.json`](starter/verification.json)
- [`check.json`](check.json): how `./check` runs and checks this lab

## Steps

1. Go to the starter: `cd labs/m06l01/m06l01-03/starter`
2. Read `main.go` the way the lesson builds it:
   - Lines 1–7: an unknown flag stops parsing
   - Lines 8–11: generated from the declarations
   - Lines 12–15: assigning your own function to Usage
   - Lines 16–18: standard error, not standard output
3. Notes from the lesson:
   - Line 13: Usage writes to standard error, never to standard output
   - Line 14: PrintDefaults lists every flag, its type and its default value
4. Run it: `go run .`.
5. Check it from the repository root: `./check m06l01-03`.

## What the lesson recorded

Shown for reference; the check does not compare it.

```text
flag provided but not defined: -bogus
usage: probe [flags] host...
  -addr string
    	address to probe (default "127.0.0.1:8080")
  -tries int
    	attempts before giving up (default 3)
exit status 2
```

## How to check

`./check m06l01-03` copies `starter/` into a scratch directory and runs `go run .` there, the way the site's lab sandbox does: that directory is the working directory and `HOME`, `LANG=C.UTF-8`, `TZ=UTC`, a limit of 10 seconds and 256 KiB of output per stream.

It runs without a pass or fail: the recorded output depends on the machine it ran on, so the site runs it without a pass or fail. `./check` shows the output and the exit code.

---

[Open the lesson on LearnSome.tech](https://learnsome.tech/learn/go-course/m06l01) · [All labs of this lesson](../README.md) · [Course README](../../../README.md)
