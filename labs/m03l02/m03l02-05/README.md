# m03l02-05 · A wrapped error keeps its identity

**Lesson:** [Wrapping Errors and errors.Is](https://learnsome.tech/learn/go-course/m03l02) (lesson 3.2, module 3: Errors and Testing) · Pro  
**Check:** Graded

## Goal

You can wrap an underlying error with context and test its identity with errors Is.

In the lesson: Start returns a wrapped busy error. The first print gives the operation and its cause in one line. The second branch uses errors Is to recognise that busy is retryable and prints a different policy. This is a small but important distinction for automation: the message can be enriched at every layer, while the decision still rests on a stable error identity. Keep sentinel values focused on conditions that callers truly need to distinguish.

## Files

- [`starter/README.md`](starter/README.md)
- [`starter/go.mod`](starter/go.mod)
- [`starter/main.go`](starter/main.go): the listing from the lesson
- [`starter/run.sh`](starter/run.sh): the command the lesson ran
- [`starter/verification.json`](starter/verification.json)
- [`expected.txt`](expected.txt): the output the check compares with
- [`check.json`](check.json): how `./check` runs and checks this lab

## Steps

1. Go to the starter: `cd labs/m03l02/m03l02-05/starter`
2. Read `main.go`.
3. Run it: `go run .`.
4. Check it from the repository root: `./check m03l02-05`.

## Expected output

```text
start worker: busy
retry later
```

## How to check

`./check m03l02-05` copies `starter/` into a scratch directory and runs `go run .` there, the way the site's lab sandbox does: that directory is the working directory and `HOME`, `LANG=C.UTF-8`, `TZ=UTC`, a limit of 10 seconds and 256 KiB of output per stream.

It passes when the output matches `expected.txt` by the site's rules, within the limits. Standard output is compared line by line; spaces at the end of a line and blank lines at the end do not count. If that differs, standard output followed by standard error is compared with Python traceback frames and blank lines set aside, so a lesson that shows an error passes when your program prints the same error. A pass here is a pass on the site.

---

[Open the lesson on LearnSome.tech](https://learnsome.tech/learn/go-course/m03l02) · [All labs of this lesson](../README.md) · [Course README](../../../README.md)
