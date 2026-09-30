# m06l03 · Graceful Shutdown on SIGTERM

Module 6: Building CLIs and Services · lesson 6.3 · Pro · [Open the lesson](https://learnsome.tech/learn/go-course/m06l03)

**Goal:** You can stop accepting traffic on SIGTERM, cancel active work, and let a Go service finish cleanly.

## Labs

| Lab | What it is | Check |
| --- | --- | --- |
| [m06l03-02](m06l03-02/) | Waiting for SIGTERM | Runs, not graded |
| [m06l03-04](m06l03-04/) | A bounded shutdown call | Graded |

## Check yourself

- What does SIGTERM mean to a service?
- Why does Shutdown need a context?
- When should main return?

---

[Course README](../../README.md) · [Concurrent Systems Programming with Go on LearnSome.tech](https://learnsome.tech/courses/go-course)
