# m02l02-05 · Building a result and a lookup table

**Lesson:** [Arrays, Slices, and Maps](https://learnsome.tech/learn/go-course/m02l02) (lesson 2.2, module 2: Types, Structs, and Interfaces) · Pro  
**Check:** Graded

## Goal

You can choose arrays, slices, and maps for predictable collections and handle their zero values safely.

In the lesson: This is a common tool pattern. Start with a nil slice for results and append as you discover names. Then create a map with make before the loop stores membership flags. The range loop gives us an index and a value, and the blank identifier discards the index because it is not useful here. At the end the slice preserves order, while the map answers a membership question quickly. One collection carries the report, and the other carries the lookup table.

## Files

- [`starter/README.md`](starter/README.md)
- [`starter/go.mod`](starter/go.mod)
- [`starter/main.go`](starter/main.go): the listing from the lesson
- [`starter/run.sh`](starter/run.sh): the command the lesson ran
- [`starter/verification.json`](starter/verification.json)
- [`expected.txt`](expected.txt): the output the check compares with
- [`check.json`](check.json): how `./check` runs and checks this lab

## Steps

1. Go to the starter: `cd labs/m02l02/m02l02-05/starter`
2. Read `main.go`.
3. Run it: `go run .`.
4. Check it from the repository root: `./check m02l02-05`.

## Expected output

```text
[api worker]
true false
```

## How to check

`./check m02l02-05` copies `starter/` into a scratch directory and runs `go run .` there, the way the site's lab sandbox does: that directory is the working directory and `HOME`, `LANG=C.UTF-8`, `TZ=UTC`, a limit of 10 seconds and 256 KiB of output per stream.

It passes when the output matches `expected.txt` by the site's rules, within the limits. Standard output is compared line by line; spaces at the end of a line and blank lines at the end do not count. If that differs, standard output followed by standard error is compared with Python traceback frames and blank lines set aside, so a lesson that shows an error passes when your program prints the same error. A pass here is a pass on the site.

---

[Open the lesson on LearnSome.tech](https://learnsome.tech/learn/go-course/m02l02) · [All labs of this lesson](../README.md) · [Course README](../../../README.md)
