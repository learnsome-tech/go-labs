# m05l05-04 · Testing a response with a fake round tripper

**Lesson:** [Calling APIs with net/http Client](https://learnsome.tech/learn/go-course/m05l05) (lesson 5.5, module 5: The Standard Library for Tooling) · Pro  
**Check:** Graded

## Goal

You can issue an HTTP request with a context, check the response, and close the body on every path.

In the lesson: A custom transport lets a test or a small example supply a deterministic response without a network call. The fake round tripper returns a successful status and a body that can be closed. The client still follows the normal lifecycle, so the calling code can test status handling and decoding without waiting on DNS or a remote service. In production the transport would be the real network implementation.

## Files

- [`starter/README.md`](starter/README.md)
- [`starter/go.mod`](starter/go.mod)
- [`starter/main.go`](starter/main.go): the listing from the lesson
- [`starter/run.sh`](starter/run.sh): the command the lesson ran
- [`starter/verification.json`](starter/verification.json)
- [`expected.txt`](expected.txt): the output the check compares with
- [`check.json`](check.json): how `./check` runs and checks this lab

## Steps

1. Go to the starter: `cd labs/m05l05/m05l05-04/starter`
2. Read `main.go`.
3. Run it: `go run .`.
4. Check it from the repository root: `./check m05l05-04`.

## Expected output

```text
200 <nil>
```

## How to check

`./check m05l05-04` copies `starter/` into a scratch directory and runs `go run .` there, the way the site's lab sandbox does: that directory is the working directory and `HOME`, `LANG=C.UTF-8`, `TZ=UTC`, a limit of 10 seconds and 256 KiB of output per stream.

It passes when the output matches `expected.txt` by the site's rules, within the limits. Standard output is compared line by line; spaces at the end of a line and blank lines at the end do not count. If that differs, standard output followed by standard error is compared with Python traceback frames and blank lines set aside, so a lesson that shows an error passes when your program prints the same error. A pass here is a pass on the site.

---

[Open the lesson on LearnSome.tech](https://learnsome.tech/learn/go-course/m05l05) · [All labs of this lesson](../README.md) · [Course README](../../../README.md)
