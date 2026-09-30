# m02l02 · Arrays, Slices, and Maps

Module 2: Types, Structs, and Interfaces · lesson 2.2 · Pro · [Open the lesson](https://learnsome.tech/learn/go-course/m02l02)

**Goal:** You can choose arrays, slices, and maps for predictable collections and handle their zero values safely.

## Labs

| Lab | What it is | Check |
| --- | --- | --- |
| [m02l02-02](m02l02-02/) | Slices grow when you append | Graded |
| [m02l02-03](m02l02-03/) | Maps need a comma ok check | Graded |
| [m02l02-05](m02l02-05/) | Building a result and a lookup table | Graded |

## Exercises

Open exercises from the lesson, to try on your own. They have no answer files: work them out, and use the labs above as reference.

### Try it yourself

1. Collect service names in a nil slice and append three values.
2. Build a map from service name to port with make.
3. Look up a missing service and print a clear absent message.

> **Hint:** Use the second result from a map lookup to tell absent from a stored zero.

## Check yourself

- When should a slice be preferred to an array?
- Why is comma ok useful for maps?
- What happens when you write to a nil map?

---

[Course README](../../README.md) · [Concurrent Systems Programming with Go on LearnSome.tech](https://learnsome.tech/courses/go-course)
