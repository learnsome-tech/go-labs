# m02l05 · Type Assertions and the Empty Interface

Module 2: Types, Structs, and Interfaces · lesson 2.5 · Pro · [Open the lesson](https://learnsome.tech/learn/go-course/m02l05)

**Goal:** You can inspect an interface value safely, use any for truly mixed data, and avoid hiding useful types behind it.

## Labs

| Lab | What it is | Check |
| --- | --- | --- |
| [m02l05-02](m02l05-02/) | A checked type assertion | Graded |
| [m02l05-03](m02l05-03/) | A type switch handles several cases | Graded |
| [m02l05-05](m02l05-05/) | Checking a typed configuration value | Graded |

## Exercises

Open exercises from the lesson, to try on your own. They have no answer files: work them out, and use the labs above as reference.

### Try it yourself

1. Write a function that accepts any and returns a readable kind label.
2. Handle strings, integers, and booleans with a type switch.
3. Return unknown for every other type and test all four paths.

> **Hint:** The switch form value dot type gives each case its concrete value.

## Check yourself

- What two things does an interface value carry?
- When is a type switch clearer than one assertion?
- Why should any stay near a data boundary?

---

[Course README](../../README.md) · [Concurrent Systems Programming with Go on LearnSome.tech](https://learnsome.tech/courses/go-course)
