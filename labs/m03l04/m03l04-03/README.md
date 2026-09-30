# m03l04-03 · A test can report several checks

**Lesson:** [The Standard testing Package](https://learnsome.tech/learn/go-course/m03l04) (lesson 3.4, module 3: Errors and Testing) · Pro  
**Check:** Read along

## Goal

You can write focused Go tests with the standard testing package and read failures as feedback about behaviour.

In the lesson: Error records a failure and lets the current test continue, which is useful when several independent expectations should be reported together. This test checks both a successful response code and a server error. The helper keeps the policy in one place, and the test names the two promises it relies on. Prefer a small assertion message that tells the reader what the code should have done, rather than repeating the entire implementation.

## Files

- [`starter/README.md`](starter/README.md)
- [`starter/go.mod`](starter/go.mod)
- [`starter/main_test.go`](starter/main_test.go): the listing from the lesson
- [`starter/run.sh`](starter/run.sh): the command the lesson ran
- [`starter/verification.json`](starter/verification.json)
- [`check.json`](check.json): how `./check` runs and checks this lab

## Steps

1. Read `starter/main_test.go` alongside the lesson.
2. On a machine that has what it needs, the lesson ran it with:

   ```sh
   go test -count=1 ./...
   ```

## How to check

**Read along.** It is a Go test file: run it with `go test` in its module. The lab sandbox only runs programs (`go run`).

There is nothing to check: `./check m03l04-03` says so and moves on.

---

[Open the lesson on LearnSome.tech](https://learnsome.tech/learn/go-course/m03l04) · [All labs of this lesson](../README.md) · [Course README](../../../README.md)
