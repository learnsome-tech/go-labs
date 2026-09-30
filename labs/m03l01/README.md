# m03l01 · Errors as Values

Module 3: Errors and Testing · lesson 3.1 · Pro · [Open the lesson](https://learnsome.tech/learn/go-course/m03l01)

**Goal:** You can return errors, inspect them, and keep failure handling explicit in a command line tool.

## Labs

| Lab | What it is | Check |
| --- | --- | --- |
| [m03l01-02](m03l01-02/) | Returning a value and an error | Graded |
| [m03l01-03](m03l01-03/) | Handling a failed operation | Graded |
| [m03l01-05](m03l01-05/) | A caller chooses the policy | Graded |

## Exercises

Open exercises from the lesson, to try on your own. They have no answer files: work them out, and use the labs above as reference.

### Try it yourself

1. Write a function that parses a non empty service name and returns an error otherwise.
2. Call it with a valid and an empty name.
3. Print the value on success and a contextual message on failure.

> **Hint:** Return nil for success and check the error before printing the value.

## Check yourself

- What does a Go function return when an operation fails?
- Why check an error before using its other result?
- Where should exit and retry policy live?

---

[Course README](../../README.md) · [Concurrent Systems Programming with Go on LearnSome.tech](https://learnsome.tech/courses/go-course)
