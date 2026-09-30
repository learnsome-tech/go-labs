# m07l02-04 · A result carries identity

**Lesson:** [Executing Concurrent Checks](https://learnsome.tech/learn/go-course/m07l02) (lesson 7.2, module 7: Project: End to End Health Checker) · Pro  
**Check:** Graded

## Goal

You can run checks concurrently, collect one result per target, and preserve a useful report order.

In the lesson: Result keeps the target identity beside the health decision. The final report can therefore explain which target passed and which target failed without consulting a separate map. The checker workers will create these values, and the coordinator will collect them for JSON output in a later lesson.

## Files

- [`starter/README.md`](starter/README.md)
- [`starter/go.mod`](starter/go.mod)
- [`starter/main.go`](starter/main.go): the listing from the lesson
- [`starter/run.sh`](starter/run.sh): the command the lesson ran
- [`starter/verification.json`](starter/verification.json)
- [`expected.txt`](expected.txt): the output the check compares with
- [`check.json`](check.json): how `./check` runs and checks this lab

## Steps

1. Go to the starter: `cd labs/m07l02/m07l02-04/starter`
2. Read `main.go`.
3. Run it: `go run .`.
4. Check it from the repository root: `./check m07l02-04`.

## Expected output

```text
api true
web false
```

## How to check

`./check m07l02-04` copies `starter/` into a scratch directory and runs `go run .` there, the way the site's lab sandbox does: that directory is the working directory and `HOME`, `LANG=C.UTF-8`, `TZ=UTC`, a limit of 10 seconds and 256 KiB of output per stream.

It passes when the output matches `expected.txt` by the site's rules, within the limits. Standard output is compared line by line; spaces at the end of a line and blank lines at the end do not count. If that differs, standard output followed by standard error is compared with Python traceback frames and blank lines set aside, so a lesson that shows an error passes when your program prints the same error. A pass here is a pass on the site.

---

[Open the lesson on LearnSome.tech](https://learnsome.tech/learn/go-course/m07l02) · [All labs of this lesson](../README.md) · [Course README](../../../README.md)
