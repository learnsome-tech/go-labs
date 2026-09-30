# m03l05 · Writing Table Driven Tests

Module 3: Errors and Testing · lesson 3.5 · Pro · [Open the lesson](https://learnsome.tech/learn/go-course/m03l05)

**Goal:** You can express many input cases in one table test and get a useful name for every failure.

## Labs

| Lab | What it is | Check |
| --- | --- | --- |
| [m03l05-02](m03l05-02/) | The first table test | Read along |
| [m03l05-03](m03l05-03/) | Subtests give each row a name | Read along |
| [m03l05-05](m03l05-05/) | Table testing a parser contract | Read along |

## Exercises

Open exercises from the lesson, to try on your own. They have no answer files: work them out, and use the labs above as reference.

### Try it yourself

1. Create a table for a status to health function.
2. Include success, redirect, server error, and an unknown code.
3. Run each row as a named subtest and report the expected result.

> **Hint:** Keep the expected boolean in each row and compare it inside t.Run.

## Check yourself

- What belongs in each table row?
- Why are subtests useful?
- How do tables encourage edge cases?

---

[Course README](../../README.md) · [Concurrent Systems Programming with Go on LearnSome.tech](https://learnsome.tech/courses/go-course)
