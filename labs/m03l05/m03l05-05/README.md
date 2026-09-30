# m03l05-05 · Table testing a parser contract

**Lesson:** [Writing Table Driven Tests](https://learnsome.tech/learn/go-course/m03l05) (lesson 3.5, module 3: Errors and Testing) · Pro  
**Check:** Read along

## Goal

You can express many input cases in one table test and get a useful name for every failure.

In the lesson: This table keeps both success and failure expectations beside their inputs. The empty row expects an error, while the service row expects a port value. The assertion first compares whether an error exists, then compares the value only for a successful case. Each subtest carries the row name, so a failure points to empty or service immediately. This is the same compact pattern you can use for configuration and command line validation.

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

There is nothing to check: `./check m03l05-05` says so and moves on.

---

[Open the lesson on LearnSome.tech](https://learnsome.tech/learn/go-course/m03l05) · [All labs of this lesson](../README.md) · [Course README](../../../README.md)
