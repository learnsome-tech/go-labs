# m02l01-05 · Checking an optional pointer

**Lesson:** [Pointers and Memory](https://learnsome.tech/learn/go-course/m02l01) (lesson 2.1, module 2: Types, Structs, and Interfaces) · Pro  
**Check:** Graded

## Goal

You can read pointer syntax, mutate shared state deliberately, and recognise a nil pointer before it becomes a crash.

In the lesson: This small helper makes the safe shape visible. The caller starts with a nil pointer, and describe checks it before trying to read anything. That path prints unset and returns. Then we create a real string and pass its address, so the second path can safely follow the pointer and print the value. In a service, this check is the difference between a missing optional setting and a process that panics during startup. Keep the guard close to the dereference so a future edit cannot accidentally move the safety boundary.

## Files

- [`starter/README.md`](starter/README.md)
- [`starter/go.mod`](starter/go.mod)
- [`starter/main.go`](starter/main.go): the listing from the lesson
- [`starter/run.sh`](starter/run.sh): the command the lesson ran
- [`starter/verification.json`](starter/verification.json)
- [`expected.txt`](expected.txt): the output the check compares with
- [`check.json`](check.json): how `./check` runs and checks this lab

## Steps

1. Go to the starter: `cd labs/m02l01/m02l01-05/starter`
2. Read `main.go`.
3. Run it: `go run .`.
4. Check it from the repository root: `./check m02l01-05`.

## Expected output

```text
unset
value: ready
```

## How to check

`./check m02l01-05` copies `starter/` into a scratch directory and runs `go run .` there, the way the site's lab sandbox does: that directory is the working directory and `HOME`, `LANG=C.UTF-8`, `TZ=UTC`, a limit of 10 seconds and 256 KiB of output per stream.

It passes when the output matches `expected.txt` by the site's rules, within the limits. Standard output is compared line by line; spaces at the end of a line and blank lines at the end do not count. If that differs, standard output followed by standard error is compared with Python traceback frames and blank lines set aside, so a lesson that shows an error passes when your program prints the same error. A pass here is a pass on the site.

---

[Open the lesson on LearnSome.tech](https://learnsome.tech/learn/go-course/m02l01) · [All labs of this lesson](../README.md) · [Course README](../../../README.md)
