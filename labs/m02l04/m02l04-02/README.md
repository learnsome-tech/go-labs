# m02l04-02 · A type satisfies an interface silently

**Lesson:** [Interfaces and Why They Are Implicit](https://learnsome.tech/learn/go-course/m02l04) (lesson 2.4, module 2: Types, Structs, and Interfaces) · Pro  
**Check:** Graded

## Goal

You can define small interfaces, satisfy them without declarations, and use that flexibility to test and compose tooling code.

In the lesson: Speaker asks for one behaviour: Speak returns a string. Robot provides that method, and there is no implements line anywhere in the type declaration. The compiler checks the relationship when Robot is passed to announce, whose parameter accepts the interface. This is why implicit interfaces feel light in Go. The consumer states what it needs, and any type that already has that behaviour can be used. The concrete type stays independent of the consumer's package.

## Files

- [`starter/README.md`](starter/README.md)
- [`starter/go.mod`](starter/go.mod)
- [`starter/main.go`](starter/main.go): the listing from the lesson
- [`starter/run.sh`](starter/run.sh): the command the lesson ran
- [`starter/verification.json`](starter/verification.json)
- [`expected.txt`](expected.txt): the output the check compares with
- [`check.json`](check.json): how `./check` runs and checks this lab

## Steps

1. Go to the starter: `cd labs/m02l04/m02l04-02/starter`
2. Read `main.go`.
3. Run it: `go run .`.
4. Check it from the repository root: `./check m02l04-02`.

## Expected output

```text
ready
```

## How to check

`./check m02l04-02` copies `starter/` into a scratch directory and runs `go run .` there, the way the site's lab sandbox does: that directory is the working directory and `HOME`, `LANG=C.UTF-8`, `TZ=UTC`, a limit of 10 seconds and 256 KiB of output per stream.

It passes when the output matches `expected.txt` by the site's rules, within the limits. Standard output is compared line by line; spaces at the end of a line and blank lines at the end do not count. If that differs, standard output followed by standard error is compared with Python traceback frames and blank lines set aside, so a lesson that shows an error passes when your program prints the same error. A pass here is a pass on the site.

---

[Open the lesson on LearnSome.tech](https://learnsome.tech/learn/go-course/m02l04) · [All labs of this lesson](../README.md) · [Course README](../../../README.md)
