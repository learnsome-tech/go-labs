# m03l02 · Wrapping Errors and errors.Is

Module 3: Errors and Testing · lesson 3.2 · Pro · [Open the lesson](https://learnsome.tech/learn/go-course/m03l02)

**Goal:** You can wrap an underlying error with context and test its identity with errors Is.

## Labs

| Lab | What it is | Check |
| --- | --- | --- |
| [m03l02-02](m03l02-02/) | Wrapping a missing resource | Graded |
| [m03l02-03](m03l02-03/) | Choosing a branch with errors Is | Graded |
| [m03l02-05](m03l02-05/) | A wrapped error keeps its identity | Graded |

## Exercises

Open exercises from the lesson, to try on your own. They have no answer files: work them out, and use the labs above as reference.

### Try it yourself

1. Create a sentinel error for an unavailable service.
2. Wrap it in a function that adds the service name.
3. Use errors Is to print a retry message without comparing text.

> **Hint:** The wrapping verb keeps the sentinel in the error chain.

## Check yourself

- What does wrapping preserve?
- Why is errors Is safer than comparing text?
- When should a sentinel error be public?

---

[Course README](../../README.md) · [Concurrent Systems Programming with Go on LearnSome.tech](https://learnsome.tech/courses/go-course)
