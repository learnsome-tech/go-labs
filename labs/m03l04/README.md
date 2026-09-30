# m03l04 · The Standard testing Package

Module 3: Errors and Testing · lesson 3.4 · Pro · [Open the lesson](https://learnsome.tech/learn/go-course/m03l04)

**Goal:** You can write focused Go tests with the standard testing package and read failures as feedback about behaviour.

## Labs

| Lab | What it is | Check |
| --- | --- | --- |
| [m03l04-02](m03l04-02/) | A first test function | Read along |
| [m03l04-03](m03l04-03/) | A test can report several checks | Read along |
| [m03l04-05](m03l04-05/) | Testing an error result | Read along |

## Exercises

Open exercises from the lesson, to try on your own. They have no answer files: work them out, and use the labs above as reference.

### Try it yourself

1. Write a function that returns an error for an empty host.
2. Create a test for a valid host and a test for the empty case.
3. Run go test and read the failure after deliberately changing one expectation.

> **Hint:** Assert both the returned value and whether the error is nil.

## Check yourself

- How does go test discover a test file?
- When should a test use fatal reporting?
- Why should unit tests avoid real networks?

---

[Course README](../../README.md) · [Concurrent Systems Programming with Go on LearnSome.tech](https://learnsome.tech/courses/go-course)
