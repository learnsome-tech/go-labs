# m02l01-02 · Taking an address and following it

**Lesson:** [Pointers and Memory](https://learnsome.tech/learn/go-course/m02l01) (lesson 2.1, module 2: Types, Structs, and Interfaces) · Pro  
**Check:** Graded

## Goal

You can read pointer syntax, mutate shared state deliberately, and recognise a nil pointer before it becomes a crash.

In the lesson: Here is the complete round trip. We start with a string named name, take its address into p, and then print the value reached through p. The star reads the value at the address. When we assign through that same star, we change the original variable, because both names point to one storage location. This is the part to remember for reviews: the pointer is not a second copy of the string. It is a route to the existing string. Run the example and the two lines prove the mutation happened in place.

## Files

- [`starter/README.md`](starter/README.md)
- [`starter/go.mod`](starter/go.mod)
- [`starter/main.go`](starter/main.go): the listing from the lesson
- [`starter/run.sh`](starter/run.sh): the command the lesson ran
- [`starter/verification.json`](starter/verification.json)
- [`expected.txt`](expected.txt): the output the check compares with
- [`check.json`](check.json): how `./check` runs and checks this lab

## Steps

1. Go to the starter: `cd labs/m02l01/m02l01-02/starter`
2. Read `main.go`.
3. Run it: `go run .`.
4. Check it from the repository root: `./check m02l01-02`.

## Expected output

```text
api
worker
```

## How to check

`./check m02l01-02` copies `starter/` into a scratch directory and runs `go run .` there, the way the site's lab sandbox does: that directory is the working directory and `HOME`, `LANG=C.UTF-8`, `TZ=UTC`, a limit of 10 seconds and 256 KiB of output per stream.

It passes when the output matches `expected.txt` by the site's rules, within the limits. Standard output is compared line by line; spaces at the end of a line and blank lines at the end do not count. If that differs, standard output followed by standard error is compared with Python traceback frames and blank lines set aside, so a lesson that shows an error passes when your program prints the same error. A pass here is a pass on the site.

---

[Open the lesson on LearnSome.tech](https://learnsome.tech/learn/go-course/m02l01) · [All labs of this lesson](../README.md) · [Course README](../../../README.md)
