# m02l03 · Structs and Methods

Module 2: Types, Structs, and Interfaces · lesson 2.3 · Pro · [Open the lesson](https://learnsome.tech/learn/go-course/m02l03)

**Goal:** You can model tool data with structs, attach methods, and choose value or pointer receivers intentionally.

## Labs

| Lab | What it is | Check |
| --- | --- | --- |
| [m02l03-02](m02l03-02/) | Defining and constructing a struct | Graded |
| [m02l03-03](m02l03-03/) | Methods give the type a vocabulary | Graded |
| [m02l03-05](m02l03-05/) | A method that changes state | Graded |

## Exercises

Open exercises from the lesson, to try on your own. They have no answer files: work them out, and use the labs above as reference.

### Try it yourself

1. Define a Check struct with a name and a healthy boolean.
2. Add a method that returns the name and status as text.
3. Add a pointer method that marks the check healthy.

> **Hint:** A method that changes a field needs a pointer receiver.

## Check yourself

- Why are named struct fields useful at construction sites?
- When should a method use a pointer receiver?
- What does a method receiver give a type?

---

[Course README](../../README.md) · [Concurrent Systems Programming with Go on LearnSome.tech](https://learnsome.tech/courses/go-course)
