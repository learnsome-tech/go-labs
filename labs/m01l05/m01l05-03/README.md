# m01l05-03 · Every type has a zero value, and it is useful

**Lesson:** [Variables, Typing, and Zero Values](https://learnsome.tech/learn/go-course/m01l05) (lesson 1.5, module 1: The Go Toolchain and Fundamentals) · Pro  
**Check:** Graded

## Goal

You can declare variables both ways and say which to use where, predict the zero value of any type and design configuration structs around it, define constant sets with iota, and convert between numeric types without guessing.

In the lesson: Declare something without a value in Go and it is not undefined, and it is not a null you have to guard against. It is the zero value for its type, and the zero values were chosen to be useful: zero for numbers, the empty string, false for booleans, and nil for slices, maps and pointers. Print them all and notice the second line especially. That slice is nil, and its length is still zero rather than an error, so you can ask a nil slice how long it is and even append to it. A nil map can be read from safely and only panics when you write to it. And a struct with no values is a struct whose every field is its own zero.

## Files

- [`starter/README.md`](starter/README.md)
- [`starter/go.mod`](starter/go.mod)
- [`starter/main.go`](starter/main.go): the listing from the lesson
- [`starter/run.sh`](starter/run.sh): the command the lesson ran
- [`starter/verification.json`](starter/verification.json)
- [`expected.txt`](expected.txt): the output the check compares with
- [`check.json`](check.json): how `./check` runs and checks this lab

## Steps

1. Go to the starter: `cd labs/m01l05/m01l05-03/starter`
2. Read `main.go`.
3. Run it: `go run .`.
4. Check it from the repository root: `./check m01l05-03`.

## Expected output

```text
0 "" false
true 0 true 0
{Retries:0 Timeout: Strict:false}
```

## How to check

`./check m01l05-03` copies `starter/` into a scratch directory and runs `go run .` there, the way the site's lab sandbox does: that directory is the working directory and `HOME`, `LANG=C.UTF-8`, `TZ=UTC`, a limit of 10 seconds and 256 KiB of output per stream.

It passes when the output matches `expected.txt` by the site's rules, within the limits. Standard output is compared line by line; spaces at the end of a line and blank lines at the end do not count. If that differs, standard output followed by standard error is compared with Python traceback frames and blank lines set aside, so a lesson that shows an error passes when your program prints the same error. A pass here is a pass on the site.

---

[Open the lesson on LearnSome.tech](https://learnsome.tech/learn/go-course/m01l05) · [All labs of this lesson](../README.md) · [Course README](../../../README.md)
