# m05l03-02 · Measuring elapsed work

**Lesson:** [The time Package](https://learnsome.tech/learn/go-course/m05l03) (lesson 5.3, module 5: The Standard Library for Tooling) · Pro  
**Check:** Graded

## Goal

You can measure elapsed work, schedule ticks, and format timestamps without confusing a duration with a wall clock time.

In the lesson: The program records a start time, waits briefly, and asks how much time has elapsed. Since returns a duration, and the comparison proves that the work took a positive amount of time. In production you would record the duration for a metric or a log line rather than print a boolean. Use elapsed time for latency and deadlines, not a formatted timestamp subtraction.

## Files

- [`starter/README.md`](starter/README.md)
- [`starter/go.mod`](starter/go.mod)
- [`starter/main.go`](starter/main.go): the listing from the lesson
- [`starter/run.sh`](starter/run.sh): the command the lesson ran
- [`starter/verification.json`](starter/verification.json)
- [`expected.txt`](expected.txt): the output the check compares with
- [`check.json`](check.json): how `./check` runs and checks this lab

## Steps

1. Go to the starter: `cd labs/m05l03/m05l03-02/starter`
2. Read `main.go`.
3. Run it: `go run .`.
4. Check it from the repository root: `./check m05l03-02`.

## Expected output

```text
true
```

## How to check

`./check m05l03-02` copies `starter/` into a scratch directory and runs `go run .` there, the way the site's lab sandbox does: that directory is the working directory and `HOME`, `LANG=C.UTF-8`, `TZ=UTC`, a limit of 10 seconds and 256 KiB of output per stream.

It passes when the output matches `expected.txt` by the site's rules, within the limits. Standard output is compared line by line; spaces at the end of a line and blank lines at the end do not count. If that differs, standard output followed by standard error is compared with Python traceback frames and blank lines set aside, so a lesson that shows an error passes when your program prints the same error. A pass here is a pass on the site.

---

[Open the lesson on LearnSome.tech](https://learnsome.tech/learn/go-course/m05l03) · [All labs of this lesson](../README.md) · [Course README](../../../README.md)
