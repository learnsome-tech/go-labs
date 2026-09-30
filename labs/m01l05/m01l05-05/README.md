# m01l05-05 · Constants and iota for a set of states

**Lesson:** [Variables, Typing, and Zero Values](https://learnsome.tech/learn/go-course/m01l05) (lesson 1.5, module 1: The Go Toolchain and Fundamentals) · Pro  
**Check:** Graded

## Goal

You can declare variables both ways and say which to use where, predict the zero value of any type and design configuration structs around it, define constant sets with iota, and convert between numeric types without guessing.

In the lesson: Go has no enumerated type, so this is the idiom that replaces one. First a named type based on a whole number, which gives the compiler something to check. Then the constant block, where the counter identifier counts from zero down the block, one per line, and the type and the assignment carry down with it. Notice that the first constant is the zero value of the type, which is why unknown deserves that slot rather than healthy: an unset state should not read as fine. Constants are compile time values, so there is no memory and no lookup, and a mistyped comparison is a build failure. Compare them and print one, and remember that printing shows the number until we teach the type to describe itself, which is a method, and methods are next module.

## Files

- [`starter/README.md`](starter/README.md)
- [`starter/go.mod`](starter/go.mod)
- [`starter/main.go`](starter/main.go): the listing from the lesson
- [`starter/run.sh`](starter/run.sh): the command the lesson ran
- [`starter/verification.json`](starter/verification.json)
- [`expected.txt`](expected.txt): the output the check compares with
- [`check.json`](check.json): how `./check` runs and checks this lab

## Steps

1. Go to the starter: `cd labs/m01l05/m01l05-05/starter`
2. Read `main.go` the way the lesson builds it:
   - Lines 1–5: a named type
   - Lines 6–12: the constant block
   - Lines 13–21: compare them
3. Notes from the lesson:
   - Line 8: iota counts from zero down the block, one per line
4. Run it: `go run .`.
5. Check it from the repository root: `./check m01l05-05`.

## Expected output

```text
0 1 2 3
true true
main.state 2 of 3
```

## How to check

`./check m01l05-05` copies `starter/` into a scratch directory and runs `go run .` there, the way the site's lab sandbox does: that directory is the working directory and `HOME`, `LANG=C.UTF-8`, `TZ=UTC`, a limit of 10 seconds and 256 KiB of output per stream.

It passes when the output matches `expected.txt` by the site's rules, within the limits. Standard output is compared line by line; spaces at the end of a line and blank lines at the end do not count. If that differs, standard output followed by standard error is compared with Python traceback frames and blank lines set aside, so a lesson that shows an error passes when your program prints the same error. A pass here is a pass on the site.

---

[Open the lesson on LearnSome.tech](https://learnsome.tech/learn/go-course/m01l05) · [All labs of this lesson](../README.md) · [Course README](../../../README.md)
