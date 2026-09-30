# m06l02-02 · A health handler

**Lesson:** [Serving HTTP with net/http Server](https://learnsome.tech/learn/go-course/m06l02) (lesson 6.2, module 6: Building CLIs and Services) · Pro  
**Check:** Graded

## Goal

You can expose a small HTTP handler, return useful status codes, and configure a server with a timeout.

In the lesson: The health function receives the writer and request required by the handler contract. It sets a content type and writes a small response. Main registers the path on a new multiplexer and prints true to prove the route exists. Starting the listener is a separate concern, which keeps this example easy to test with an in memory request and recorder.

## Files

- [`starter/README.md`](starter/README.md)
- [`starter/go.mod`](starter/go.mod)
- [`starter/main.go`](starter/main.go): the listing from the lesson
- [`starter/run.sh`](starter/run.sh): the command the lesson ran
- [`starter/verification.json`](starter/verification.json)
- [`expected.txt`](expected.txt): the output the check compares with
- [`check.json`](check.json): how `./check` runs and checks this lab

## Steps

1. Go to the starter: `cd labs/m06l02/m06l02-02/starter`
2. Read `main.go`.
3. Run it: `go run .`.
4. Check it from the repository root: `./check m06l02-02`.

## Expected output

```text
true
```

## How to check

`./check m06l02-02` copies `starter/` into a scratch directory and runs `go run .` there, the way the site's lab sandbox does: that directory is the working directory and `HOME`, `LANG=C.UTF-8`, `TZ=UTC`, a limit of 10 seconds and 256 KiB of output per stream.

It passes when the output matches `expected.txt` by the site's rules, within the limits. Standard output is compared line by line; spaces at the end of a line and blank lines at the end do not count. If that differs, standard output followed by standard error is compared with Python traceback frames and blank lines set aside, so a lesson that shows an error passes when your program prints the same error. A pass here is a pass on the site.

---

[Open the lesson on LearnSome.tech](https://learnsome.tech/learn/go-course/m06l02) · [All labs of this lesson](../README.md) · [Course README](../../../README.md)
