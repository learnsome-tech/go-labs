# m02l03-03 · Methods give the type a vocabulary

**Lesson:** [Structs and Methods](https://learnsome.tech/learn/go-course/m02l03) (lesson 2.3, module 2: Types, Structs, and Interfaces) · Pro  
**Check:** Graded

## Goal

You can model tool data with structs, attach methods, and choose value or pointer receivers intentionally.

In the lesson: A method is a function with a receiver, and the receiver gives the type a vocabulary that callers can discover. Address reads a Target and returns the host and port in the form a dialer needs. This receiver is a value receiver, so the method gets a copy and cannot change the original. That is a good default for small data values and formatting helpers. Callers read t dot Address as naturally as they read a field, while the implementation keeps the formatting rule in one place.

## Files

- [`starter/README.md`](starter/README.md)
- [`starter/go.mod`](starter/go.mod)
- [`starter/main.go`](starter/main.go): the listing from the lesson
- [`starter/run.sh`](starter/run.sh): the command the lesson ran
- [`starter/verification.json`](starter/verification.json)
- [`expected.txt`](expected.txt): the output the check compares with
- [`check.json`](check.json): how `./check` runs and checks this lab

## Steps

1. Go to the starter: `cd labs/m02l03/m02l03-03/starter`
2. Read `main.go`.
3. Run it: `go run .`.
4. Check it from the repository root: `./check m02l03-03`.

## Expected output

```text
api:8080
```

## How to check

`./check m02l03-03` copies `starter/` into a scratch directory and runs `go run .` there, the way the site's lab sandbox does: that directory is the working directory and `HOME`, `LANG=C.UTF-8`, `TZ=UTC`, a limit of 10 seconds and 256 KiB of output per stream.

It passes when the output matches `expected.txt` by the site's rules, within the limits. Standard output is compared line by line; spaces at the end of a line and blank lines at the end do not count. If that differs, standard output followed by standard error is compared with Python traceback frames and blank lines set aside, so a lesson that shows an error passes when your program prints the same error. A pass here is a pass on the site.

---

[Open the lesson on LearnSome.tech](https://learnsome.tech/learn/go-course/m02l03) · [All labs of this lesson](../README.md) · [Course README](../../../README.md)
