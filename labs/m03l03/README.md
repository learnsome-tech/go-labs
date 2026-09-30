# m03l03 · Panics and Recover

Module 3: Errors and Testing · lesson 3.3 · Pro · [Open the lesson](https://learnsome.tech/learn/go-course/m03l03)

**Goal:** You can distinguish a programmer invariant from an expected operational error and use recover only at a deliberate boundary.

## Labs

| Lab | What it is | Check |
| --- | --- | --- |
| [m03l03-02](m03l03-02/) | An ordinary error keeps control | Graded |
| [m03l03-03](m03l03-03/) | Recover at a worker boundary | Graded |
| [m03l03-05](m03l03-05/) | Deferred cleanup still runs | Graded |

## Exercises

Open exercises from the lesson, to try on your own. They have no answer files: work them out, and use the labs above as reference.

### Try it yourself

1. Write a worker that returns an error for bad input.
2. Add a separate invariant check that panics.
3. Recover at the supervisor boundary and print whether the worker failed or was contained.

> **Hint:** Keep recovery in a deferred function at the boundary that owns the worker.

## Check yourself

- Which failures should be returned as errors?
- Where does recover belong?
- Why is recovered state not automatically safe?

---

[Course README](../../README.md) · [Concurrent Systems Programming with Go on LearnSome.tech](https://learnsome.tech/courses/go-course)
