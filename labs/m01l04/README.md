# m01l04 · Your First Program and go run

Module 1: The Go Toolchain and Fundamentals · lesson 1.4 · Free · [Open the lesson](https://learnsome.tech/learn/go-course/m01l04)

**Goal:** You can choose correctly between running, building and installing a program, write functions that return a value alongside an error, format output with the verbs you will use daily, and set an exit code that a pipeline can act on.

## Labs

| Lab | What it is | Check |
| --- | --- | --- |
| [m01l04-02](m01l04-02/) | Watching the three verbs behave | Runs, not graded |
| [m01l04-03](m01l04-03/) | A function that returns a value and an error | Graded |
| [m01l04-04](m01l04-04/) | The printing verbs you will use every day | Graded |
| [m01l04-06](m01l04-06/) | Exit codes are the contract with your pipeline | Runs, not graded |

## Exercises

Open exercises from the lesson, to try on your own. They have no answer files: work them out, and use the labs above as reference.

### Try it yourself

1. Write a function returning a percentage and an error, and call it with bad input.
2. Print the same struct with the value verb and the plus value verb, and compare.
3. Make the program exit with status three when a threshold is crossed.

> **Hint:** Check the exit status of the last command in your shell to see what your program returned.

## Check yourself

- When would you use build rather than run, and what changes for the host?
- What does a nil error mean, and what does the caller usually do next?
- Which format verb shows field names, and when do you reach for it?
- Which stream should a failure reason go to, and what should the exit status be?
- What happens to deferred cleanup when a program calls exit?

---

[Course README](../../README.md) · [Concurrent Systems Programming with Go on LearnSome.tech](https://learnsome.tech/courses/go-course)
