# m02l01-03 · Passing a pointer to a function

**Lesson:** [Pointers and Memory](https://learnsome.tech/learn/go-course/m02l01) (lesson 2.1, module 2: Types, Structs, and Interfaces) · Pro  
**Check:** Graded

## Goal

You can read pointer syntax, mutate shared state deliberately, and recognise a nil pointer before it becomes a crash.

In the lesson: A function receives a copy of each argument, including a pointer. That copy still points at the same data, so the function can change the caller's value without returning a replacement. The parameter says this plainly: addTag accepts a pointer to a string. Inside, it follows the pointer, adds a suffix, and stores the result. The caller passes its address, then prints the changed status. This pattern is useful for a configuration object or a counter that several operations must update, while keeping the shared state visible in the function signature.

## Files

- [`starter/README.md`](starter/README.md)
- [`starter/go.mod`](starter/go.mod)
- [`starter/main.go`](starter/main.go): the listing from the lesson
- [`starter/run.sh`](starter/run.sh): the command the lesson ran
- [`starter/verification.json`](starter/verification.json)
- [`expected.txt`](expected.txt): the output the check compares with
- [`check.json`](check.json): how `./check` runs and checks this lab

## Steps

1. Go to the starter: `cd labs/m02l01/m02l01-03/starter`
2. Read `main.go`.
3. Run it: `go run .`.
4. Check it from the repository root: `./check m02l01-03`.

## Expected output

```text
service-ready
```

## How to check

`./check m02l01-03` copies `starter/` into a scratch directory and runs `go run .` there, the way the site's lab sandbox does: that directory is the working directory and `HOME`, `LANG=C.UTF-8`, `TZ=UTC`, a limit of 10 seconds and 256 KiB of output per stream.

It passes when the output matches `expected.txt` by the site's rules, within the limits. Standard output is compared line by line; spaces at the end of a line and blank lines at the end do not count. If that differs, standard output followed by standard error is compared with Python traceback frames and blank lines set aside, so a lesson that shows an error passes when your program prints the same error. A pass here is a pass on the site.

---

[Open the lesson on LearnSome.tech](https://learnsome.tech/learn/go-course/m02l01) · [All labs of this lesson](../README.md) · [Course README](../../../README.md)
