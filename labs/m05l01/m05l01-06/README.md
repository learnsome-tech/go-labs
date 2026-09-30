# m05l01-06 · Exit codes are the contract with CI

**Lesson:** [Files and the os Package](https://learnsome.tech/learn/go-course/m05l01) (lesson 5.1, module 5: The Standard Library for Tooling) · Pro  
**Check:** Runs, not graded

## Goal

You can read arguments and environment, read and write files safely, tell a missing file from a broken one, and set the exit code your pipeline branches on.

In the lesson: Your exit status is an interface, even though it does not look like one. A pipeline step, a Kubernetes probe and a shell script with the dash e flag all read that number and nothing else. Name them at the top of the file rather than scattering bare numbers through the code. Zero means the check passed. One means the check itself failed, which is a result, not a crash. Two here means the caller used the tool wrongly, and keeping those apart lets a pipeline retry one and never the other. One warning: o s dot Exit stops the program immediately and skips every deferred call, so flush and close before you reach it. Run it and the shell reports it.

## Files

- [`starter/README.md`](starter/README.md)
- [`starter/go.mod`](starter/go.mod)
- [`starter/main.go`](starter/main.go): the listing from the lesson
- [`starter/run.sh`](starter/run.sh): the command the lesson ran
- [`starter/verification.json`](starter/verification.json)
- [`check.json`](check.json): how `./check` runs and checks this lab

## Steps

1. Go to the starter: `cd labs/m05l01/m05l01-06/starter`
2. Read `main.go` the way the lesson builds it:
   - Lines 1–12: Name them at the top
   - Lines 13–18: the caller used the tool wrongly
   - Lines 19–22: the check itself failed
3. Notes from the lesson:
   - Line 21: os.Exit skips deferred calls: flush and close first
4. Run it: `go run .`.
5. Check it from the repository root: `./check m05l01-06`.

## What the lesson recorded

Shown for reference; the check does not compare it.

```text
checking web01
probe: host unreachable
exit status 1
exit=1
```

## How to check

`./check m05l01-06` copies `starter/` into a scratch directory and runs `go run .` there, the way the site's lab sandbox does: that directory is the working directory and `HOME`, `LANG=C.UTF-8`, `TZ=UTC`, a limit of 10 seconds and 256 KiB of output per stream.

It runs without a pass or fail: the lesson recorded its output with `go run . web01; echo "exit=$?"`, not the way the site runs it, so the output is not compared. `./check` shows the output and the exit code.

---

[Open the lesson on LearnSome.tech](https://learnsome.tech/learn/go-course/m05l01) · [All labs of this lesson](../README.md) · [Course README](../../../README.md)
