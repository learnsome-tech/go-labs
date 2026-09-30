# m01l04-04 · The printing verbs you will use every day

**Lesson:** [Your First Program and go run](https://learnsome.tech/learn/go-course/m01l04) (lesson 1.4, module 1: The Go Toolchain and Fundamentals) · Free  
**Check:** Graded

## Goal

You can choose correctly between running, building and installing a program, write functions that return a value alongside an error, format output with the verbs you will use daily, and set an exit code that a pipeline can act on.

In the lesson: Formatted printing deserves five minutes now because it saves hours later. The verbs you will actually use: one for a string, one for a whole number, one that quotes a string so you can see whether it has spaces or is empty, one for a boolean. Then the general one that prints any value in a default shape, and its plus variant, which adds field names and is the single most useful debugging tool in the language. The type verb tells you what something really is when inference has surprised you. And Sprintf returns the string instead of printing it, which is how you build a log line or a URL. Run the lot and compare each line with its format string.

## Files

- [`starter/README.md`](starter/README.md)
- [`starter/go.mod`](starter/go.mod)
- [`starter/main.go`](starter/main.go): the listing from the lesson
- [`starter/run.sh`](starter/run.sh): the command the lesson ran
- [`starter/verification.json`](starter/verification.json)
- [`expected.txt`](expected.txt): the output the check compares with
- [`check.json`](check.json): how `./check` runs and checks this lab

## Steps

1. Go to the starter: `cd labs/m01l04/m01l04-04/starter`
2. Read `main.go`.
3. Run it: `go run .`.
4. Check it from the repository root: `./check m01l04-04`.

## Expected output

```text
host db-1 port 5432
quoted "db-1" ready true
value {db-1 5432}
fields {Host:db-1 Port:5432}
type main.target
db-1:5432
```

## How to check

`./check m01l04-04` copies `starter/` into a scratch directory and runs `go run .` there, the way the site's lab sandbox does: that directory is the working directory and `HOME`, `LANG=C.UTF-8`, `TZ=UTC`, a limit of 10 seconds and 256 KiB of output per stream.

It passes when the output matches `expected.txt` by the site's rules, within the limits. Standard output is compared line by line; spaces at the end of a line and blank lines at the end do not count. If that differs, standard output followed by standard error is compared with Python traceback frames and blank lines set aside, so a lesson that shows an error passes when your program prints the same error. A pass here is a pass on the site.

---

[Open the lesson on LearnSome.tech](https://learnsome.tech/learn/go-course/m01l04) · [All labs of this lesson](../README.md) · [Course README](../../../README.md)
