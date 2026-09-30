# m01l01-02 · A first program, top to bottom

**Lesson:** [What Go Is and Why It Matters](https://learnsome.tech/learn/go-course/m01l01) (lesson 1.1, module 1: The Go Toolchain and Fundamentals) · Free  
**Check:** Graded

## Goal

You can explain why platform and DevOps teams write their tooling in Go, compile and run a program with the go command, and describe exactly what the compiler hands you at the end of a build.

In the lesson: Here is the whole of a Go program. Every file starts by naming the package it belongs to, and the name main is special: it marks a program rather than a library, something that can be built into an executable. Then the imports, one per line or in a block, naming the packages this file uses. Here that is f m t, the formatting package, which is where printing lives. Last comes the function named main, which is where a Go program begins. Two calls to Println, each printing one line. Let us run it and see those two lines come back.

## Files

- [`starter/README.md`](starter/README.md)
- [`starter/go.mod`](starter/go.mod)
- [`starter/main.go`](starter/main.go): the listing from the lesson
- [`starter/run.sh`](starter/run.sh): the command the lesson ran
- [`starter/verification.json`](starter/verification.json)
- [`expected.txt`](expected.txt): the output the check compares with
- [`check.json`](check.json): how `./check` runs and checks this lab

## Steps

1. Go to the starter: `cd labs/m01l01/m01l01-02/starter`
2. Read `main.go` the way the lesson builds it:
   - Lines 1: every file starts
   - Lines 2–3: then the imports
   - Lines 4–8: the function named main
3. Notes from the lesson:
   - Line 1: package main marks a program, not a library
4. Run it: `go run .`.
5. Check it from the repository root: `./check m01l01-02`.

## Expected output

```text
checking one host
all good
```

## How to check

`./check m01l01-02` copies `starter/` into a scratch directory and runs `go run .` there, the way the site's lab sandbox does: that directory is the working directory and `HOME`, `LANG=C.UTF-8`, `TZ=UTC`, a limit of 10 seconds and 256 KiB of output per stream.

It passes when the output matches `expected.txt` by the site's rules, within the limits. Standard output is compared line by line; spaces at the end of a line and blank lines at the end do not count. If that differs, standard output followed by standard error is compared with Python traceback frames and blank lines set aside, so a lesson that shows an error passes when your program prints the same error. A pass here is a pass on the site.

---

[Open the lesson on LearnSome.tech](https://learnsome.tech/learn/go-course/m01l01) · [All labs of this lesson](../README.md) · [Course README](../../../README.md)
