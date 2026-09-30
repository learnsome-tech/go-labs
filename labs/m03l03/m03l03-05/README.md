# m03l03-05 · Deferred cleanup still runs

**Lesson:** [Panics and Recover](https://learnsome.tech/learn/go-course/m03l03) (lesson 3.3, module 3: Errors and Testing) · Pro  
**Check:** Graded

## Goal

You can distinguish a programmer invariant from an expected operational error and use recover only at a deliberate boundary.

In the lesson: A panic unwinds through deferred calls in last in first out order. The recovery function runs first and prints recovered. The cleanup defer runs next, so resources registered with defer still get their chance to close. Control returns to the caller, which prints done. This ordering is useful at a worker boundary, but it should not encourage casual panics. The normal path should still return errors for conditions the caller can reasonably handle.

## Files

- [`starter/README.md`](starter/README.md)
- [`starter/go.mod`](starter/go.mod)
- [`starter/main.go`](starter/main.go): the listing from the lesson
- [`starter/run.sh`](starter/run.sh): the command the lesson ran
- [`starter/verification.json`](starter/verification.json)
- [`expected.txt`](expected.txt): the output the check compares with
- [`check.json`](check.json): how `./check` runs and checks this lab

## Steps

1. Go to the starter: `cd labs/m03l03/m03l03-05/starter`
2. Read `main.go`.
3. Run it: `go run .`.
4. Check it from the repository root: `./check m03l03-05`.

## Expected output

```text
recovered
cleanup
done
```

## How to check

`./check m03l03-05` copies `starter/` into a scratch directory and runs `go run .` there, the way the site's lab sandbox does: that directory is the working directory and `HOME`, `LANG=C.UTF-8`, `TZ=UTC`, a limit of 10 seconds and 256 KiB of output per stream.

It passes when the output matches `expected.txt` by the site's rules, within the limits. Standard output is compared line by line; spaces at the end of a line and blank lines at the end do not count. If that differs, standard output followed by standard error is compared with Python traceback frames and blank lines set aside, so a lesson that shows an error passes when your program prints the same error. A pass here is a pass on the site.

---

[Open the lesson on LearnSome.tech](https://learnsome.tech/learn/go-course/m03l03) · [All labs of this lesson](../README.md) · [Course README](../../../README.md)
