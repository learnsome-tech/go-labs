# m05l02-04 · Reading a bounded prefix

**Lesson:** [Streams and the io Package](https://learnsome.tech/learn/go-course/m05l02) (lesson 5.2, module 5: The Standard Library for Tooling) · Pro  
**Check:** Graded

## Goal

You can copy streams, read bounded input, and treat files and network bodies through common interfaces.

In the lesson: The limit reader wraps a longer source and exposes only its first four bytes. ReadAll is now safe for this deliberately bounded stream, and the output proves that the remaining input was not consumed. In a tool that accepts a response body, choose the bound from the protocol or an operational limit and return an error when the input is larger than your contract allows.

## Files

- [`starter/README.md`](starter/README.md)
- [`starter/go.mod`](starter/go.mod)
- [`starter/main.go`](starter/main.go): the listing from the lesson
- [`starter/run.sh`](starter/run.sh): the command the lesson ran
- [`starter/verification.json`](starter/verification.json)
- [`expected.txt`](expected.txt): the output the check compares with
- [`check.json`](check.json): how `./check` runs and checks this lab

## Steps

1. Go to the starter: `cd labs/m05l02/m05l02-04/starter`
2. Read `main.go`.
3. Run it: `go run .`.
4. Check it from the repository root: `./check m05l02-04`.

## Expected output

```text
0123 <nil>
```

## How to check

`./check m05l02-04` copies `starter/` into a scratch directory and runs `go run .` there, the way the site's lab sandbox does: that directory is the working directory and `HOME`, `LANG=C.UTF-8`, `TZ=UTC`, a limit of 10 seconds and 256 KiB of output per stream.

It passes when the output matches `expected.txt` by the site's rules, within the limits. Standard output is compared line by line; spaces at the end of a line and blank lines at the end do not count. If that differs, standard output followed by standard error is compared with Python traceback frames and blank lines set aside, so a lesson that shows an error passes when your program prints the same error. A pass here is a pass on the site.

---

[Open the lesson on LearnSome.tech](https://learnsome.tech/learn/go-course/m05l02) · [All labs of this lesson](../README.md) · [Course README](../../../README.md)
