# m07l01-04 · Validating a target

**Lesson:** [Project Setup and Configuration](https://learnsome.tech/learn/go-course/m07l01) (lesson 7.1, module 7: Project: End to End Health Checker) · Pro  
**Check:** Graded

## Goal

You can start the health checker with typed configuration, clear defaults, and a small target model.

In the lesson: The validator checks the complete target before any request starts. A missing URL returns a clear error, and main prints it as the configuration failure. Later lessons can assume that every target has the fields required for a check, which keeps worker code focused on network behaviour.

## Files

- [`starter/README.md`](starter/README.md)
- [`starter/go.mod`](starter/go.mod)
- [`starter/main.go`](starter/main.go): the listing from the lesson
- [`starter/run.sh`](starter/run.sh): the command the lesson ran
- [`starter/verification.json`](starter/verification.json)
- [`expected.txt`](expected.txt): the output the check compares with
- [`check.json`](check.json): how `./check` runs and checks this lab

## Steps

1. Go to the starter: `cd labs/m07l01/m07l01-04/starter`
2. Read `main.go`.
3. Run it: `go run .`.
4. Check it from the repository root: `./check m07l01-04`.

## Expected output

```text
target needs name and URL
```

## How to check

`./check m07l01-04` copies `starter/` into a scratch directory and runs `go run .` there, the way the site's lab sandbox does: that directory is the working directory and `HOME`, `LANG=C.UTF-8`, `TZ=UTC`, a limit of 10 seconds and 256 KiB of output per stream.

It passes when the output matches `expected.txt` by the site's rules, within the limits. Standard output is compared line by line; spaces at the end of a line and blank lines at the end do not count. If that differs, standard output followed by standard error is compared with Python traceback frames and blank lines set aside, so a lesson that shows an error passes when your program prints the same error. A pass here is a pass on the site.

---

[Open the lesson on LearnSome.tech](https://learnsome.tech/learn/go-course/m07l01) · [All labs of this lesson](../README.md) · [Course README](../../../README.md)
