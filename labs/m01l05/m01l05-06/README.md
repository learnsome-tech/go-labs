# m01l05-06 · Conversions are explicit, always

**Lesson:** [Variables, Typing, and Zero Values](https://learnsome.tech/learn/go-course/m01l05) (lesson 1.5, module 1: The Go Toolchain and Fundamentals) · Pro  
**Check:** Graded

## Goal

You can declare variables both ways and say which to use where, predict the zero value of any type and design configuration structs around it, define constant sets with iota, and convert between numeric types without guessing.

In the lesson: Last piece of the module, and the one that catches everybody arriving from a dynamic language. Go will not mix numeric types for you. A whole number and a number with a decimal point cannot be multiplied together, even though the machine could do it perfectly well, and there is no silent widening. You convert, in writing, every time. Run those three lines and read the middle one closely: converting the decimal to a whole number truncates it towards zero, so a factor of one point two five becomes one, and the multiplication quietly does nothing. That is not a bug in Go, it is arithmetic you asked for, and it is exactly the kind of mistake that produces a capacity report nobody trusts. Convert deliberately, and convert late.

## Files

- [`starter/README.md`](starter/README.md)
- [`starter/go.mod`](starter/go.mod)
- [`starter/main.go`](starter/main.go): the listing from the lesson
- [`starter/run.sh`](starter/run.sh): the command the lesson ran
- [`starter/verification.json`](starter/verification.json)
- [`expected.txt`](expected.txt): the output the check compares with
- [`check.json`](check.json): how `./check` runs and checks this lab

## Steps

1. Go to the starter: `cd labs/m01l05/m01l05-06/starter`
2. Read `main.go`.
3. Run it: `go run .`.
4. Check it from the repository root: `./check m01l05-06`.

## Expected output

```text
1.025
164
250
```

## How to check

`./check m01l05-06` copies `starter/` into a scratch directory and runs `go run .` there, the way the site's lab sandbox does: that directory is the working directory and `HOME`, `LANG=C.UTF-8`, `TZ=UTC`, a limit of 10 seconds and 256 KiB of output per stream.

It passes when the output matches `expected.txt` by the site's rules, within the limits. Standard output is compared line by line; spaces at the end of a line and blank lines at the end do not count. If that differs, standard output followed by standard error is compared with Python traceback frames and blank lines set aside, so a lesson that shows an error passes when your program prints the same error. A pass here is a pass on the site.

---

[Open the lesson on LearnSome.tech](https://learnsome.tech/learn/go-course/m01l05) · [All labs of this lesson](../README.md) · [Course README](../../../README.md)
