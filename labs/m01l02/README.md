# m01l02 · Installing the Go Toolchain

Module 1: The Go Toolchain and Fundamentals · lesson 1.2 · Free · [Open the lesson](https://learnsome.tech/learn/go-course/m01l02)

**Goal:** You can install the Go toolchain the way that suits a laptop, a build agent or a container, and use the go command's own subcommands for formatting, checking and documentation instead of reaching for extra tools.

## Labs

| Lab | What it is | Check |
| --- | --- | --- |
| [m01l02-04](m01l02-04/) | Formatting is settled, not argued about | Runs, not graded |
| [m01l02-05](m01l02-05/) | Vet catches what the compiler allows | Runs, not graded |

## Exercises

Open exercises from the lesson, to try on your own. They have no answer files: work them out, and use the labs above as reference.

### Try it yourself

1. Print the version, then read the usage line of the test subcommand.
2. Break the formatting of a file on purpose and fix it with the formatter.
3. Write a print call whose format verb does not match its argument, then run vet.

> **Hint:** The listing flag on the formatter is minus l, and the rewrite flag is minus w.

## Check yourself

- Which install route would you pick for a build agent, and why?
- What does the formatter print when every file is already formatted?
- Name two mistakes vet finds that the compiler accepts.
- Why is the documentation subcommand better than a web search for an API question?

---

[Course README](../../README.md) · [Concurrent Systems Programming with Go on LearnSome.tech](https://learnsome.tech/courses/go-course)
