# m05l01-02 · Arguments and environment

**Lesson:** [Files and the os Package](https://learnsome.tech/learn/go-course/m05l01) (lesson 5.1, module 5: The Standard Library for Tooling) · Pro  
**Check:** Runs, not graded

## Goal

You can read arguments and environment, read and write files safely, tell a missing file from a broken one, and set the exit code your pipeline branches on.

In the lesson: This is the front door of any tool. o s dot Args is a slice of strings, and the first entry is the program itself, so almost always you slice it off. If the caller has not given us enough to work with, we print usage to standard error and leave with a non zero status. Configuration comes from the environment next. Getenv hands back an empty string for a variable that is not set, which is fine whenever empty is a sensible default. When you need to tell empty apart from absent, which is the usual case for a feature flag, reach for LookupEnv: it returns the value and a boolean that says whether it was there at all. Run it with a variable in front of the command and watch both arrive.

## Files

- [`starter/README.md`](starter/README.md)
- [`starter/go.mod`](starter/go.mod)
- [`starter/main.go`](starter/main.go): the listing from the lesson
- [`starter/run.sh`](starter/run.sh): the command the lesson ran
- [`starter/verification.json`](starter/verification.json)
- [`check.json`](check.json): how `./check` runs and checks this lab

## Steps

1. Go to the starter: `cd labs/m05l01/m05l01-02/starter`
2. Read `main.go` the way the lesson builds it:
   - Lines 1–10: a slice of strings
   - Lines 11–14: we print usage to standard error
   - Lines 15–16: Configuration comes from the environment
   - Lines 17–22: reach for LookupEnv
3. Notes from the lesson:
   - Line 9: Args[0] is the program path, so almost always slice it off
   - Line 17: LookupEnv separates empty from absent; Getenv cannot
4. Run it: `go run .`.
5. Check it from the repository root: `./check m05l01-02`.

## What the lesson recorded

Shown for reference; the check does not compare it.

```text
args: [check web01]
action: check host: web01
timeout: 5s
debug: unset
```

## How to check

`./check m05l01-02` copies `starter/` into a scratch directory and runs `go run .` there, the way the site's lab sandbox does: that directory is the working directory and `HOME`, `LANG=C.UTF-8`, `TZ=UTC`, a limit of 10 seconds and 256 KiB of output per stream.

It runs without a pass or fail: the lesson recorded its output with `PROBE_TIMEOUT=5s go run .`, not the way the site runs it, so the output is not compared. `./check` shows the output and the exit code.

---

[Open the lesson on LearnSome.tech](https://learnsome.tech/learn/go-course/m05l01) · [All labs of this lesson](../README.md) · [Course README](../../../README.md)
