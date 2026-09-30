# m02l01 · Pointers and Memory

Module 2: Types, Structs, and Interfaces · lesson 2.1 · Pro · [Open the lesson](https://learnsome.tech/learn/go-course/m02l01)

**Goal:** You can read pointer syntax, mutate shared state deliberately, and recognise a nil pointer before it becomes a crash.

## Labs

| Lab | What it is | Check |
| --- | --- | --- |
| [m02l01-02](m02l01-02/) | Taking an address and following it | Graded |
| [m02l01-03](m02l01-03/) | Passing a pointer to a function | Graded |
| [m02l01-05](m02l01-05/) | Checking an optional pointer | Graded |

## Exercises

Open exercises from the lesson, to try on your own. They have no answer files: work them out, and use the labs above as reference.

### Try it yourself

1. Write a function that receives a pointer to an integer and increments it.
2. Call it three times and print the final value.
3. Call a second function with a nil pointer and report that the setting is absent.

> **Hint:** Compare a pointer with nil before using the star operator.

## Check yourself

- What does a pointer store?
- Why can a pointer parameter mutate the caller's value?
- What must you do before dereferencing an optional pointer?

---

[Course README](../../README.md) · [Concurrent Systems Programming with Go on LearnSome.tech](https://learnsome.tech/courses/go-course)
