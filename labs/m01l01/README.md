# m01l01 · What Go Is and Why It Matters

Module 1: The Go Toolchain and Fundamentals · lesson 1.1 · Free · [Open the lesson](https://learnsome.tech/learn/go-course/m01l01)

**Goal:** You can explain why platform and DevOps teams write their tooling in Go, compile and run a program with the go command, and describe exactly what the compiler hands you at the end of a build.

## Labs

| Lab | What it is | Check |
| --- | --- | --- |
| [m01l01-02](m01l01-02/) | A first program, top to bottom | Graded |
| [m01l01-03](m01l01-03/) | The build step, and what it produces | Runs, not graded |
| [m01l01-05](m01l01-05/) | The compiler refuses an unused import | Runs, not graded |

## Exercises

Open exercises from the lesson, to try on your own. They have no answer files: work them out, and use the labs above as reference.

### Try it yourself

1. Write a program that prints your name and the words all systems nominal on two lines.
2. Build it with an output name of your choosing, then run the binary.
3. Add an import you do not use, build again, and read the error carefully.

> **Hint:** The go command for producing a named binary is go build with the dash o flag.

## Check yourself

- What two things must a Go file have to be built into an executable?
- What does a successful build print, and why is that worth knowing?
- Name two things the compiler links into the binary for you.
- Why is an unused import an error in Go rather than a warning?
- What has to be installed on a host before it can run your Go binary?

---

[Course README](../../README.md) · [Concurrent Systems Programming with Go on LearnSome.tech](https://learnsome.tech/courses/go-course)
