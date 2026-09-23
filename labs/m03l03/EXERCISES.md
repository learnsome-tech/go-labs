# Exercises — Panics and Recover

Lesson `m03l03` · [Watch](https://learnsome.tech/courses/go-course/watch?lesson=m03l03)

## Exercise 1: Try it yourself

1. Write a worker that returns an error for bad input.
2. Add a separate invariant check that panics.
3. Recover at the supervisor boundary and print whether the worker failed or was contained.

> **Hint**: Keep recovery in a deferred function at the boundary that owns the worker.


---

© LearnSome.tech · support@iwantto.learnsome.tech
