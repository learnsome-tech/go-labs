# m02l04 · Interfaces and Why They Are Implicit

Module 2: Types, Structs, and Interfaces · lesson 2.4 · Pro · [Open the lesson](https://learnsome.tech/learn/go-course/m02l04)

**Goal:** You can define small interfaces, satisfy them without declarations, and use that flexibility to test and compose tooling code.

## Labs

| Lab | What it is | Check |
| --- | --- | --- |
| [m02l04-02](m02l04-02/) | A type satisfies an interface silently | Graded |
| [m02l04-03](m02l04-03/) | Two concrete types, one function | Graded |
| [m02l04-05](m02l04-05/) | A tiny interface makes testing easy | Graded |

## Exercises

Open exercises from the lesson, to try on your own. They have no answer files: work them out, and use the labs above as reference.

### Try it yourself

1. Define a Probe interface with a Check method that returns a string.
2. Create a real looking probe and a fixed probe that both satisfy it.
3. Write one function that accepts the interface and prints both results.

> **Hint:** Do not add an implements declaration; matching methods are enough.

## Check yourself

- How does a type satisfy an interface in Go?
- Where should a small interface usually be declared?
- Why are small interfaces useful in tests?

---

[Course README](../../README.md) · [Concurrent Systems Programming with Go on LearnSome.tech](https://learnsome.tech/courses/go-course)
