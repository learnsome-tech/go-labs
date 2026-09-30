# m01l03 · Go Workspaces and Modules

Module 1: The Go Toolchain and Fundamentals · lesson 1.3 · Free · [Open the lesson](https://learnsome.tech/learn/go-course/m01l03)

**Goal:** You can start a module, split a tool into packages with import paths that follow from the module path, keep dependencies tidy and vendored, and use a workspace file when a tool and a shared library sit in one checkout.

## Labs

| Lab | What it is | Check |
| --- | --- | --- |
| [m01l03-03](m01l03-03/) | Splitting a tool into packages | Graded |
| [m01l03-05](m01l03-05/) | Tidy and vendor, and what they are for | Runs, not graded |

## Exercises

Open exercises from the lesson, to try on your own. They have no answer files: work them out, and use the labs above as reference.

### Try it yourself

1. Start a module, then add a package under a directory called internal.
2. Export one function from it, call it from the main package, and build.
3. Try importing that internal package from a second module and read the error.

> **Hint:** The import path is your module path plus a slash plus the directory name.

## Check yourself

- What is the difference between a package, a module and a workspace?
- How does the toolchain know whether a name is visible to other packages?
- What does a directory named internal do to imports from other modules?
- When would you vendor dependencies rather than fetch them?
- Why is a workspace file usually not committed?

---

[Course README](../../README.md) · [Concurrent Systems Programming with Go on LearnSome.tech](https://learnsome.tech/courses/go-course)
