# m05l01-04 · Opening a file, and closing it

**Lesson:** [Files and the os Package](https://learnsome.tech/learn/go-course/m05l01) (lesson 5.1, module 5: The Standard Library for Tooling) · Pro  
**Check:** Graded

## Goal

You can read arguments and environment, read and write files safely, tell a missing file from a broken one, and set the exit code your pipeline branches on.

In the lesson: When the file is a log, or anything you would rather not hold in memory, you open it instead. Open gives you two things: a file handle and an error, and until you have checked the error the handle is not yours to use. The defer line is the habit to build. It schedules Close for the moment this function returns, whichever path it takes, so a later early return cannot leak the descriptor. On a long running agent that leak is the bug that wakes you at three in the morning. Then we finish by asking the open file about itself: name, size in bytes and the permission mode, printed the way a listing would show it.

## Files

- [`starter/README.md`](starter/README.md)
- [`starter/app.log`](starter/app.log)
- [`starter/go.mod`](starter/go.mod)
- [`starter/main.go`](starter/main.go): the listing from the lesson
- [`starter/run.sh`](starter/run.sh): the command the lesson ran
- [`starter/verification.json`](starter/verification.json)
- [`expected.txt`](expected.txt): the output the check compares with
- [`check.json`](check.json): how `./check` runs and checks this lab

## Steps

1. Go to the starter: `cd labs/m05l01/m05l01-04/starter`
2. Read `main.go` the way the lesson builds it:
   - Lines 1–13: Open gives you two things
   - Lines 14: the defer line
   - Lines 15–21: asking the open file about itself
3. Notes from the lesson:
   - Line 14: defer runs when main returns, however it returns
4. Run it: `go run .`.
5. Check it from the repository root: `./check m05l01-04`.

## Expected output

```text
app.log: 41 bytes, mode -rw-r--r--
```

## How to check

`./check m05l01-04` copies `starter/` into a scratch directory and runs `go run .` there, the way the site's lab sandbox does: that directory is the working directory and `HOME`, `LANG=C.UTF-8`, `TZ=UTC`, a limit of 10 seconds and 256 KiB of output per stream.

It passes when the output matches `expected.txt` by the site's rules, within the limits. Standard output is compared line by line; spaces at the end of a line and blank lines at the end do not count. If that differs, standard output followed by standard error is compared with Python traceback frames and blank lines set aside, so a lesson that shows an error passes when your program prints the same error. A pass here is a pass on the site.

---

[Open the lesson on LearnSome.tech](https://learnsome.tech/learn/go-course/m05l01) · [All labs of this lesson](../README.md) · [Course README](../../../README.md)
