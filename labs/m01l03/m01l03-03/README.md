# m01l03-03 · Splitting a tool into packages

**Lesson:** [Go Workspaces and Modules](https://learnsome.tech/learn/go-course/m01l03) (lesson 1.3, module 1: The Go Toolchain and Fundamentals) · Free  
**Check:** Graded

## Goal

You can start a module, split a tool into packages with import paths that follow from the module path, keep dependencies tidy and vendored, and use a workspace file when a tool and a shared library sit in one checkout.

In the lesson: Here is a second package in the same module. The package clause names it, and by convention the name matches its directory. Then a type worth sharing, describing one thing our tool watches. Then one exported function that renders a target for a human. Notice how visibility works: a capital letter means other packages can use this name, and a lower case name is private to the directory. There is no keyword for public or private, only that spelling rule. The directory called internal is special too: the toolchain refuses imports of it from outside this module, which is how you publish a tool without publishing its guts. When the main package calls it, we get one line back.

## Files

- [`starter/README.md`](starter/README.md)
- [`starter/go.mod`](starter/go.mod)
- [`starter/internal/check/check.go`](starter/internal/check/check.go)
- [`starter/main.go`](starter/main.go): the listing from the lesson
- [`starter/run.sh`](starter/run.sh): the command the lesson ran
- [`starter/verification.json`](starter/verification.json)
- [`expected.txt`](expected.txt): the output the check compares with
- [`check.json`](check.json): how `./check` runs and checks this lab

## Steps

1. Go to the starter: `cd labs/m01l03/m01l03-03/starter`
2. Read `main.go` the way the lesson builds it:
   - Lines 1: the package clause
   - Lines 2–8: a type worth sharing
   - Lines 9–14: one exported function
   - Lines 15–18: lower case name
3. Notes from the lesson:
   - Line 6: Exported: capital T, so other packages can name it
   - Line 16: Unexported: lower case, invisible outside this directory
4. Run it: `go run .`.
5. Check it from the repository root: `./check m01l03-03`.

## Expected output

```text
db-1 on port 5432
```

## How to check

`./check m01l03-03` copies `starter/` into a scratch directory and runs `go run .` there, the way the site's lab sandbox does: that directory is the working directory and `HOME`, `LANG=C.UTF-8`, `TZ=UTC`, a limit of 10 seconds and 256 KiB of output per stream.

It passes when the output matches `expected.txt` by the site's rules, within the limits. Standard output is compared line by line; spaces at the end of a line and blank lines at the end do not count. If that differs, standard output followed by standard error is compared with Python traceback frames and blank lines set aside, so a lesson that shows an error passes when your program prints the same error. A pass here is a pass on the site.

---

[Open the lesson on LearnSome.tech](https://learnsome.tech/learn/go-course/m01l03) · [All labs of this lesson](../README.md) · [Course README](../../../README.md)
