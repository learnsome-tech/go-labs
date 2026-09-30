# m06l01-05 · Environment fallback, for the container case

**Lesson:** [Parsing Arguments with flag](https://learnsome.tech/learn/go-course/m06l01) (lesson 6.1, module 6: Building CLIs and Services) · Pro  
**Check:** Runs, not graded

## Goal

You can build a command line tool with the standard flag package that takes typed options, reads positional arguments, falls back to environment variables, and fails with a usage message and the right exit code.

In the lesson: Containers configure through the environment and people configure through flags, so a tool that lives in both worlds wants both. The pattern is one small helper: look the variable up, return it when it is set, otherwise return the built in default. Feed that helper into the flag declaration as the default value and a clean precedence order falls out for free. An explicit flag wins, because parsing overwrites the default. The environment wins when no flag was given. The compiled in default is the last resort. Use LookupEnv rather than Getenv when an empty string is a meaningful setting, because only LookupEnv tells you whether the variable was set at all. Let us run it twice: once with just the variable, then with a flag as well.

## Files

- [`starter/README.md`](starter/README.md)
- [`starter/go.mod`](starter/go.mod)
- [`starter/main.go`](starter/main.go): the listing from the lesson
- [`starter/run.sh`](starter/run.sh): the command the lesson ran
- [`starter/verification.json`](starter/verification.json)
- [`check.json`](check.json): how `./check` runs and checks this lab

## Steps

1. Go to the starter: `cd labs/m06l01/m06l01-05/starter`
2. Read `main.go` the way the lesson builds it:
   - Lines 1–7: one small helper
   - Lines 8–14: otherwise return the built in default
   - Lines 15–17: into the flag declaration as the default value
   - Lines 18–20: An explicit flag wins
3. Notes from the lesson:
   - Line 10: LookupEnv tells you whether the variable exists; Getenv cannot
   - Line 17: The environment supplies the default; an explicit flag overrides it
4. Run it: `go run .`.
5. Check it from the repository root: `./check m06l01-05`.

## What the lesson recorded

Shown for reference; the check does not compare it.

```text
addr: 10.2.0.4:9090
addr: 10.9.9.9:80
```

## How to check

`./check m06l01-05` copies `starter/` into a scratch directory and runs `go run .` there, the way the site's lab sandbox does: that directory is the working directory and `HOME`, `LANG=C.UTF-8`, `TZ=UTC`, a limit of 10 seconds and 256 KiB of output per stream.

It runs without a pass or fail: the recorded output depends on the machine it ran on, so the site runs it without a pass or fail. `./check` shows the output and the exit code.

---

[Open the lesson on LearnSome.tech](https://learnsome.tech/learn/go-course/m06l01) · [All labs of this lesson](../README.md) · [Course README](../../../README.md)
