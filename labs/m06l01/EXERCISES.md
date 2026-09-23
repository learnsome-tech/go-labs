# Exercises — Parsing Arguments with flag

Lesson `m06l01` · [Watch](https://learnsome.tech/courses/go-course/watch?lesson=m06l01)

## Exercise 1: Try it yourself

1. Add a -json flag that prints the parsed configuration as one JSON object instead of lines.
2. Add a -timeout duration flag that falls back to the PROBE_TIMEOUT environment variable.
3. Make the tool exit two with your usage block when no host argument is given.

> **Hint**: envOr returns a string; a Duration flag will parse that string for you if you pass it as the default.


---

© LearnSome.tech · support@iwantto.learnsome.tech
