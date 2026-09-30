# m03l05-03 · Subtests give each row a name

**Lesson:** [Writing Table Driven Tests](https://learnsome.tech/learn/go-course/m03l05) (lesson 3.5, module 3: Errors and Testing) · Pro  
**Check:** Read along

## Goal

You can express many input cases in one table test and get a useful name for every failure.

In the lesson: Subtests let the testing package report each row as its own named check. The map provides two cases, and the loop opens a subtest with the row name. A failing row appears in the test output with that name, which makes a large table easier to search. For tables with more fields, a slice of structs is usually clearer than a map because it preserves order and lets each row carry all of its expected values.

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

There is nothing to check: `./check m03l05-03` says so and moves on.

---

[Open the lesson on LearnSome.tech](https://learnsome.tech/learn/go-course/m03l05) · [All labs of this lesson](../README.md) · [Course README](../../../README.md)
