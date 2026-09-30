# m02l05-05 · Checking a typed configuration value

**Lesson:** [Type Assertions and the Empty Interface](https://learnsome.tech/learn/go-course/m02l05) (lesson 2.5, module 2: Types, Structs, and Interfaces) · Pro  
**Check:** Graded

## Goal

You can inspect an interface value safely, use any for truly mixed data, and avoid hiding useful types behind it.

In the lesson: This boundary accepts any value, then narrows it immediately. The assertion requires an integer, and the second check rejects a value below one, so the function returns a boolean that tells the caller whether the setting is usable. A real configuration loader might do this after decoding a generic map, then pass the resulting integer through typed code. Notice that the invalid string never travels deeper into the program. Convert uncertainty at the edge, and keep the rest of the tool simple.

## Files

- [`starter/README.md`](starter/README.md)
- [`starter/go.mod`](starter/go.mod)
- [`starter/main.go`](starter/main.go): the listing from the lesson
- [`starter/run.sh`](starter/run.sh): the command the lesson ran
- [`starter/verification.json`](starter/verification.json)
- [`expected.txt`](expected.txt): the output the check compares with
- [`check.json`](check.json): how `./check` runs and checks this lab

## Steps

1. Go to the starter: `cd labs/m02l05/m02l05-05/starter`
2. Read `main.go`.
3. Run it: `go run .`.
4. Check it from the repository root: `./check m02l05-05`.

## Expected output

```text
timeout 5
invalid timeout
```

## How to check

`./check m02l05-05` copies `starter/` into a scratch directory and runs `go run .` there, the way the site's lab sandbox does: that directory is the working directory and `HOME`, `LANG=C.UTF-8`, `TZ=UTC`, a limit of 10 seconds and 256 KiB of output per stream.

It passes when the output matches `expected.txt` by the site's rules, within the limits. Standard output is compared line by line; spaces at the end of a line and blank lines at the end do not count. If that differs, standard output followed by standard error is compared with Python traceback frames and blank lines set aside, so a lesson that shows an error passes when your program prints the same error. A pass here is a pass on the site.

---

[Open the lesson on LearnSome.tech](https://learnsome.tech/learn/go-course/m02l05) · [All labs of this lesson](../README.md) · [Course README](../../../README.md)
