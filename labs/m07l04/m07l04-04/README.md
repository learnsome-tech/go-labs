# m07l04-04 · Indented output for an operator

**Lesson:** [Exporting Results as JSON](https://learnsome.tech/learn/go-course/m07l04) (lesson 7.4, module 7: Project: End to End Health Checker) · Pro  
**Check:** Graded

## Goal

You can encode health results as stable JSON for pipelines, dashboards, and later automation.

In the lesson: An encoder can write directly to standard output and indent the document for a human operator. The report remains valid JSON, but its lines are easier to inspect in a terminal or saved artifact. For a pipeline where bytes matter, leave indentation off and keep diagnostics on standard error.

## Files

- [`starter/README.md`](starter/README.md)
- [`starter/go.mod`](starter/go.mod)
- [`starter/main.go`](starter/main.go): the listing from the lesson
- [`starter/run.sh`](starter/run.sh): the command the lesson ran
- [`starter/verification.json`](starter/verification.json)
- [`expected.txt`](expected.txt): the output the check compares with
- [`check.json`](check.json): how `./check` runs and checks this lab

## Steps

1. Go to the starter: `cd labs/m07l04/m07l04-04/starter`
2. Read `main.go`.
3. Run it: `go run .`.
4. Check it from the repository root: `./check m07l04-04`.

## Expected output

```text
{
  "healthy": true
}
```

## How to check

`./check m07l04-04` copies `starter/` into a scratch directory and runs `go run .` there, the way the site's lab sandbox does: that directory is the working directory and `HOME`, `LANG=C.UTF-8`, `TZ=UTC`, a limit of 10 seconds and 256 KiB of output per stream.

It passes when the output matches `expected.txt` by the site's rules, within the limits. Standard output is compared line by line; spaces at the end of a line and blank lines at the end do not count. If that differs, standard output followed by standard error is compared with Python traceback frames and blank lines set aside, so a lesson that shows an error passes when your program prints the same error. A pass here is a pass on the site.

---

[Open the lesson on LearnSome.tech](https://learnsome.tech/learn/go-course/m07l04) · [All labs of this lesson](../README.md) · [Course README](../../../README.md)
