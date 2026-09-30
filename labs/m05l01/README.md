# m05l01 · Files and the os Package

Module 5: The Standard Library for Tooling · lesson 5.1 · Pro · [Open the lesson](https://learnsome.tech/learn/go-course/m05l01)

**Goal:** You can read arguments and environment, read and write files safely, tell a missing file from a broken one, and set the exit code your pipeline branches on.

## Labs

| Lab | What it is | Check |
| --- | --- | --- |
| [m05l01-02](m05l01-02/) | Arguments and environment | Runs, not graded |
| [m05l01-03](m05l01-03/) | A small file in one call | Graded |
| [m05l01-04](m05l01-04/) | Opening a file, and closing it | Graded |
| [m05l01-05](m05l01-05/) | Missing, or broken? | Graded |
| [m05l01-06](m05l01-06/) | Exit codes are the contract with CI | Runs, not graded |

## Exercises

Open exercises from the lesson, to try on your own. They have no answer files: work them out, and use the labs above as reference.

### Try it yourself

1. Write a tool that reads a path from os.Args and a default from an env var when no path is given.
2. Print the file size when it exists, print absent and exit 3 when it does not.
3. Write a summary line to a file with mode 0o600 and read it back to prove the mode stuck.

> **Hint:** os.Stat returns the error you feed to errors.Is; os.FileInfo has Size and Mode.

## Check yourself

- Why does os.Args[1:] appear in almost every Go tool?
- When does LookupEnv tell you something Getenv cannot?
- What does defer f.Close() protect you from that a later Close call does not?
- Why is errors.Is with os.ErrNotExist better than comparing the error text?
- Which deferred work is skipped when you call os.Exit?

---

[Course README](../../README.md) · [Concurrent Systems Programming with Go on LearnSome.tech](https://learnsome.tech/courses/go-course)
