# m04l05 · Context Cancellation and Timeouts

Module 4: Concurrency and Context · lesson 4.5 · Pro · [Open the lesson](https://learnsome.tech/learn/go-course/m04l05)

**Goal:** You can bound a concurrent operation with a timeout and return promptly when an orchestrator or caller cancels it.

## Labs

| Lab | What it is | Check |
| --- | --- | --- |
| [m04l05-02](m04l05-02/) | A timeout returns an error | Graded |
| [m04l05-04](m04l05-04/) | A caller cancels a worker | Graded |

## Check yourself

- Why should a worker select on context Done?
- What happens when a child replaces context with background?
- How should shutdown coordinate worker cleanup?

---

[Course README](../../README.md) · [Concurrent Systems Programming with Go on LearnSome.tech](https://learnsome.tech/courses/go-course)
