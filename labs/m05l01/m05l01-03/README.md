# m05l01-03 · A small file in one call

**Lesson:** [Files and the os Package](https://learnsome.tech/learn/go-course/m05l01) (lesson 5.1, module 5: The Standard Library for Tooling) · Pro  
**Check:** Graded

## Goal

You can read arguments and environment, read and write files safely, tell a missing file from a broken one, and set the exit code your pipeline branches on.

In the lesson: For a config file, a lock file or a small report you do not need a stream at all. WriteFile takes a name, a slice of bytes and a permission mode, and creates or truncates it in a single call. That mode is worth a moment: octal six hundred means readable and writable by the owner and by nobody else, which is what you want for anything holding a token. Permissions only apply when the file is created, and the process umask can still clear bits, so check rather than assume. ReadFile is the mirror image and hands you the whole file as bytes. Converting to a string is free enough for configuration, and printing it back shows the round trip worked.

## Files

- [`starter/README.md`](starter/README.md)
- [`starter/go.mod`](starter/go.mod)
- [`starter/main.go`](starter/main.go): the listing from the lesson
- [`starter/run.sh`](starter/run.sh): the command the lesson ran
- [`starter/verification.json`](starter/verification.json)
- [`expected.txt`](expected.txt): the output the check compares with
- [`check.json`](check.json): how `./check` runs and checks this lab

## Steps

1. Go to the starter: `cd labs/m05l01/m05l01-03/starter`
2. Read `main.go` the way the lesson builds it:
   - Lines 1–14: creates or truncates it in a single call
   - Lines 15–19: ReadFile is the mirror image
   - Lines 20–22: printing it back
3. Notes from the lesson:
   - Line 10: 0o600: owner read and write only, nobody else
4. Run it: `go run .`.
5. Check it from the repository root: `./check m05l01-03`.

## Expected output

```text
read 23 bytes
interval=30s
retries=3
```

## How to check

`./check m05l01-03` copies `starter/` into a scratch directory and runs `go run .` there, the way the site's lab sandbox does: that directory is the working directory and `HOME`, `LANG=C.UTF-8`, `TZ=UTC`, a limit of 10 seconds and 256 KiB of output per stream.

It passes when the output matches `expected.txt` by the site's rules, within the limits. Standard output is compared line by line; spaces at the end of a line and blank lines at the end do not count. If that differs, standard output followed by standard error is compared with Python traceback frames and blank lines set aside, so a lesson that shows an error passes when your program prints the same error. A pass here is a pass on the site.

---

[Open the lesson on LearnSome.tech](https://learnsome.tech/learn/go-course/m05l01) · [All labs of this lesson](../README.md) · [Course README](../../../README.md)
