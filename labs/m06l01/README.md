# m06l01 · Parsing Arguments with flag

Module 6: Building CLIs and Services · lesson 6.1 · Pro · [Open the lesson](https://learnsome.tech/learn/go-course/m06l01)

**Goal:** You can build a command line tool with the standard flag package that takes typed options, reads positional arguments, falls back to environment variables, and fails with a usage message and the right exit code.

## Labs

| Lab | What it is | Check |
| --- | --- | --- |
| [m06l01-02](m06l01-02/) | Declaring typed options and parsing them | Runs, not graded |
| [m06l01-03](m06l01-03/) | Getting it wrong: usage and exit code two | Runs, not graded |
| [m06l01-04](m06l01-04/) | Positional arguments are what is left over | Runs, not graded |
| [m06l01-05](m06l01-05/) | Environment fallback, for the container case | Runs, not graded |

## Exercises

Open exercises from the lesson, to try on your own. They have no answer files: work them out, and use the labs above as reference.

### Try it yourself

1. Add a -json flag that prints the parsed configuration as one JSON object instead of lines.
2. Add a -timeout duration flag that falls back to the PROBE_TIMEOUT environment variable.
3. Make the tool exit two with your usage block when no host argument is given.

> **Hint:** envOr returns a string; a Duration flag will parse that string for you if you pass it as the default.

## Check yourself

- Why must you call Parse before reading the value a flag declaration returned?
- Which stream does the usage block go to, and why does that matter in a pipeline?
- What exit status does the flag package use for an unknown flag?
- How do you make an environment variable act as a flag's default without overriding an explicit flag?
- Name two things Cobra gives you that the flag package does not.

---

[Course README](../../README.md) · [Concurrent Systems Programming with Go on LearnSome.tech](https://learnsome.tech/courses/go-course)
