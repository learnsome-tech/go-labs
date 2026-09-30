# m07l05-02 · A final unit test shape

**Lesson:** [Containerizing and Testing the Final Build](https://learnsome.tech/learn/go-course/m07l05) (lesson 7.5, module 7: Project: End to End Health Checker) · Pro  
**Check:** Read along

## Goal

You can test the finished checker, build a small runtime image, and connect its exit status to an orchestrator.

In the lesson: The final test protects the smallest useful result contract. It constructs a typed result and asserts the health decision that a report and exit policy will consume. Add table rows for failures and timeouts, then run the complete package tests before building the release artifact.

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

There is nothing to check: `./check m07l05-02` says so and moves on.

---

[Open the lesson on LearnSome.tech](https://learnsome.tech/learn/go-course/m07l05) · [All labs of this lesson](../README.md) · [Course README](../../../README.md)
