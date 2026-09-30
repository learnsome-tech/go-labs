# m01l05 · Variables, Typing, and Zero Values

Module 1: The Go Toolchain and Fundamentals · lesson 1.5 · Pro · [Open the lesson](https://learnsome.tech/learn/go-course/m01l05)

**Goal:** You can declare variables both ways and say which to use where, predict the zero value of any type and design configuration structs around it, define constant sets with iota, and convert between numeric types without guessing.

## Labs

| Lab | What it is | Check |
| --- | --- | --- |
| [m01l05-02](m01l05-02/) | Both declaration forms, and what was inferred | Graded |
| [m01l05-03](m01l05-03/) | Every type has a zero value, and it is useful | Graded |
| [m01l05-05](m01l05-05/) | Constants and iota for a set of states | Graded |
| [m01l05-06](m01l05-06/) | Conversions are explicit, always | Graded |

## Exercises

Open exercises from the lesson, to try on your own. They have no answer files: work them out, and use the labs above as reference.

### Try it yourself

1. Declare one variable of each basic type with no value and print them all.
2. Define a state type with a constant block where the zero value means unknown.
3. Compute a percentage from two whole numbers, correct to two decimal places.

> **Hint:** Convert both operands before dividing, or integer division will throw away the fraction.

## Check yourself

- Where can you not use the short declaration form, and what do you use instead?
- What is the zero value of a slice, and what can you still do with it?
- Why should the first constant in a state block usually mean unknown?
- What does converting a decimal number to a whole number do to the fraction?
- How do you tell a deliberate zero from an absent value in a config struct?

---

[Course README](../../README.md) · [Concurrent Systems Programming with Go on LearnSome.tech](https://learnsome.tech/courses/go-course)
