# m03l04-05 · Testing an error result

**Lesson:** [The Standard testing Package](https://learnsome.tech/learn/go-course/m03l04) (lesson 3.4, module 3: Errors and Testing) · Pro  
**Check:** Read along

## Goal

You can write focused Go tests with the standard testing package and read failures as feedback about behaviour.

In the lesson: This example checks the error side of a result instead of only checking success. The empty input must fail, so the test asserts that the error is not nil. In your own code use a real package error value rather than inventing one in a test. The important shape is the same: give the function input at the boundary, then assert the observable result and error contract.

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

There is nothing to check: `./check m03l04-05` says so and moves on.

---

[Open the lesson on LearnSome.tech](https://learnsome.tech/learn/go-course/m03l04) · [All labs of this lesson](../README.md) · [Course README](../../../README.md)
