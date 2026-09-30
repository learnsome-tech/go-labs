# m05l01-05 · Missing, or broken?

**Lesson:** [Files and the os Package](https://learnsome.tech/learn/go-course/m05l01) (lesson 5.1, module 5: The Standard Library for Tooling) · Pro  
**Check:** Graded

## Goal

You can read arguments and environment, read and write files safely, tell a missing file from a broken one, and set the exit code your pipeline branches on.

In the lesson: Here is the distinction that separates a careful tool from a noisy one. Stat asks the filesystem about a path without opening it, and the error it returns is not a single thing. A missing file and a permission denied are completely different situations: one means create it, the other means stop and tell somebody. So you never compare the error to a string. You ask errors dot Is whether it matches the sentinel value o s dot ErrNotExist, and you leave a default branch for everything else, which gives you three outcomes rather than two. Run it over one file that exists and one that does not, and you get exactly the two answers you would expect.

## Files

- [`starter/README.md`](starter/README.md)
- [`starter/go.mod`](starter/go.mod)
- [`starter/main.go`](starter/main.go): the listing from the lesson
- [`starter/probe.conf`](starter/probe.conf)
- [`starter/run.sh`](starter/run.sh): the command the lesson ran
- [`starter/verification.json`](starter/verification.json)
- [`expected.txt`](expected.txt): the output the check compares with
- [`check.json`](check.json): how `./check` runs and checks this lab

## Steps

1. Go to the starter: `cd labs/m05l01/m05l01-05/starter`
2. Read `main.go` the way the lesson builds it:
   - Lines 1–11: Stat asks the filesystem
   - Lines 12–18: three outcomes
   - Lines 19–21: one file that exists
3. Notes from the lesson:
   - Line 15: errors.Is unwraps, so it works through a wrapped error too
4. Run it: `go run .`.
5. Check it from the repository root: `./check m05l01-05`.

## Expected output

```text
probe.conf: present, 23 bytes
gone.conf: absent
```

## How to check

`./check m05l01-05` copies `starter/` into a scratch directory and runs `go run .` there, the way the site's lab sandbox does: that directory is the working directory and `HOME`, `LANG=C.UTF-8`, `TZ=UTC`, a limit of 10 seconds and 256 KiB of output per stream.

It passes when the output matches `expected.txt` by the site's rules, within the limits. Standard output is compared line by line; spaces at the end of a line and blank lines at the end do not count. If that differs, standard output followed by standard error is compared with Python traceback frames and blank lines set aside, so a lesson that shows an error passes when your program prints the same error. A pass here is a pass on the site.

---

[Open the lesson on LearnSome.tech](https://learnsome.tech/learn/go-course/m05l01) · [All labs of this lesson](../README.md) · [Course README](../../../README.md)
