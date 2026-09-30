# m01l05-02 · Both declaration forms, and what was inferred

**Lesson:** [Variables, Typing, and Zero Values](https://learnsome.tech/learn/go-course/m01l05) (lesson 1.5, module 1: The Go Toolchain and Fundamentals) · Pro  
**Check:** Graded

## Goal

You can declare variables both ways and say which to use where, predict the zero value of any type and design configuration structs around it, define constant sets with iota, and convert between numeric types without guessing.

In the lesson: One package level variable, outside every function, declared with the var keyword because that is the only form allowed there. Then inside, the short form four times over: a string, a whole number, a number with a decimal point. Notice what inference chose when we did not say: plain int for the whole number and the wider floating point type for the decimal. When we want a wider integer we have to say so, which is the line with the explicit type. The declaration with no value at all is the interesting one, and the next segment is about exactly that. Print the types and then the values, and read the first two lines carefully: those are the compiler's decisions, made visible.

## Files

- [`starter/README.md`](starter/README.md)
- [`starter/go.mod`](starter/go.mod)
- [`starter/main.go`](starter/main.go): the listing from the lesson
- [`starter/run.sh`](starter/run.sh): the command the lesson ran
- [`starter/verification.json`](starter/verification.json)
- [`expected.txt`](expected.txt): the output the check compares with
- [`check.json`](check.json): how `./check` runs and checks this lab

## Steps

1. Go to the starter: `cd labs/m01l05/m01l05-02/starter`
2. Read `main.go` the way the lesson builds it:
   - Lines 1–5: outside every function
   - Lines 6–12: the short form
   - Lines 13–17: print the types
3. Notes from the lesson:
   - Line 11: Explicit type: inference would have chosen plain int
4. Run it: `go run .`.
5. Check it from the repository root: `./check m01l05-02`.

## Expected output

```text
string int float64
int64 bool string
db-1 5432 0.75 3 false stable
```

## How to check

`./check m01l05-02` copies `starter/` into a scratch directory and runs `go run .` there, the way the site's lab sandbox does: that directory is the working directory and `HOME`, `LANG=C.UTF-8`, `TZ=UTC`, a limit of 10 seconds and 256 KiB of output per stream.

It passes when the output matches `expected.txt` by the site's rules, within the limits. Standard output is compared line by line; spaces at the end of a line and blank lines at the end do not count. If that differs, standard output followed by standard error is compared with Python traceback frames and blank lines set aside, so a lesson that shows an error passes when your program prints the same error. A pass here is a pass on the site.

---

[Open the lesson on LearnSome.tech](https://learnsome.tech/learn/go-course/m01l05) · [All labs of this lesson](../README.md) · [Course README](../../../README.md)
