# m06l04-04 · Inspecting build metadata

**Lesson:** [Building Static Binaries](https://learnsome.tech/learn/go-course/m06l04) (lesson 6.4, module 6: Building CLIs and Services) · Pro  
**Check:** Graded

## Goal

You can build a portable Go binary with reproducible flags and inspect the result before shipping it.

In the lesson: The runtime package exposes the operating system and architecture selected for this build. On the checking machine the example reports its local pair. A cross compiled release would report Linux and the target architecture when it runs there. Keep this metadata in diagnostics when operators need to confirm which artifact reached a host.

## Files

- [`starter/README.md`](starter/README.md)
- [`starter/go.mod`](starter/go.mod)
- [`starter/main.go`](starter/main.go): the listing from the lesson
- [`starter/run.sh`](starter/run.sh): the command the lesson ran
- [`starter/verification.json`](starter/verification.json)
- [`expected.txt`](expected.txt): the output the check compares with
- [`check.json`](check.json): how `./check` runs and checks this lab

## Steps

1. Go to the starter: `cd labs/m06l04/m06l04-04/starter`
2. Read `main.go`.
3. Run it: `go run .`.
4. Check it from the repository root: `./check m06l04-04`.

## Expected output

```text
linux
amd64
```

## How to check

`./check m06l04-04` copies `starter/` into a scratch directory and runs `go run .` there, the way the site's lab sandbox does: that directory is the working directory and `HOME`, `LANG=C.UTF-8`, `TZ=UTC`, a limit of 10 seconds and 256 KiB of output per stream.

It passes when the output matches `expected.txt` by the site's rules, within the limits. Standard output is compared line by line; spaces at the end of a line and blank lines at the end do not count. If that differs, standard output followed by standard error is compared with Python traceback frames and blank lines set aside, so a lesson that shows an error passes when your program prints the same error. A pass here is a pass on the site.

---

[Open the lesson on LearnSome.tech](https://learnsome.tech/learn/go-course/m06l04) · [All labs of this lesson](../README.md) · [Course README](../../../README.md)
