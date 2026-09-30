# m04l02-04 · Ranging over a closed channel

**Lesson:** [Channels and Synchronization](https://learnsome.tech/learn/go-course/m04l02) (lesson 4.2, module 4: Concurrency and Context) · Pro  
**Check:** Graded

## Goal

You can send values through channels, close them deliberately, and use channels as ownership boundaries.

In the lesson: This buffered channel receives two values before the sender closes it. The range loop reads both queued values, then ends when the channel is empty and closed. The final line proves that close marks completion without losing data already in the buffer. Keep close on the producer side, and let consumers range when a stream of work has a natural end.

## Files

- [`starter/README.md`](starter/README.md)
- [`starter/go.mod`](starter/go.mod)
- [`starter/main.go`](starter/main.go): the listing from the lesson
- [`starter/run.sh`](starter/run.sh): the command the lesson ran
- [`starter/verification.json`](starter/verification.json)
- [`expected.txt`](expected.txt): the output the check compares with
- [`check.json`](check.json): how `./check` runs and checks this lab

## Steps

1. Go to the starter: `cd labs/m04l02/m04l02-04/starter`
2. Read `main.go`.
3. Run it: `go run .`.
4. Check it from the repository root: `./check m04l02-04`.

## Expected output

```text
3
5
done
```

## How to check

`./check m04l02-04` copies `starter/` into a scratch directory and runs `go run .` there, the way the site's lab sandbox does: that directory is the working directory and `HOME`, `LANG=C.UTF-8`, `TZ=UTC`, a limit of 10 seconds and 256 KiB of output per stream.

It passes when the output matches `expected.txt` by the site's rules, within the limits. Standard output is compared line by line; spaces at the end of a line and blank lines at the end do not count. If that differs, standard output followed by standard error is compared with Python traceback frames and blank lines set aside, so a lesson that shows an error passes when your program prints the same error. A pass here is a pass on the site.

---

[Open the lesson on LearnSome.tech](https://learnsome.tech/learn/go-course/m04l02) · [All labs of this lesson](../README.md) · [Course README](../../../README.md)
