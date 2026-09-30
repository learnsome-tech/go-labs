# m05l02-02 · Copying a stream

**Lesson:** [Streams and the io Package](https://learnsome.tech/learn/go-course/m05l02) (lesson 5.2, module 5: The Standard Library for Tooling) · Pro  
**Check:** Graded

## Goal

You can copy streams, read bounded input, and treat files and network bodies through common interfaces.

In the lesson: The source buffer is a Reader and the target buffer is a Writer. io Copy moves bytes between them and returns the count plus an error. The function does not know that both ends are memory buffers; the same call can connect a file to standard output or an HTTP body to a report file. The example prints ten bytes, a nil error, and the copied text.

## Files

- [`starter/README.md`](starter/README.md)
- [`starter/go.mod`](starter/go.mod)
- [`starter/main.go`](starter/main.go): the listing from the lesson
- [`starter/run.sh`](starter/run.sh): the command the lesson ran
- [`starter/verification.json`](starter/verification.json)
- [`expected.txt`](expected.txt): the output the check compares with
- [`check.json`](check.json): how `./check` runs and checks this lab

## Steps

1. Go to the starter: `cd labs/m05l02/m05l02-02/starter`
2. Read `main.go`.
3. Run it: `go run .`.
4. Check it from the repository root: `./check m05l02-02`.

## Expected output

```text
10 <nil> logs ready
```

## How to check

`./check m05l02-02` copies `starter/` into a scratch directory and runs `go run .` there, the way the site's lab sandbox does: that directory is the working directory and `HOME`, `LANG=C.UTF-8`, `TZ=UTC`, a limit of 10 seconds and 256 KiB of output per stream.

It passes when the output matches `expected.txt` by the site's rules, within the limits. Standard output is compared line by line; spaces at the end of a line and blank lines at the end do not count. If that differs, standard output followed by standard error is compared with Python traceback frames and blank lines set aside, so a lesson that shows an error passes when your program prints the same error. A pass here is a pass on the site.

---

[Open the lesson on LearnSome.tech](https://learnsome.tech/learn/go-course/m05l02) · [All labs of this lesson](../README.md) · [Course README](../../../README.md)
