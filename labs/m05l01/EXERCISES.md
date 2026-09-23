# Exercises — Files and the os Package

Lesson `m05l01` · [Watch](https://learnsome.tech/courses/go-course/watch?lesson=m05l01)

## Exercise 1: Try it yourself

1. Write a tool that reads a path from os.Args and a default from an env var when no path is given.
2. Print the file size when it exists, print absent and exit 3 when it does not.
3. Write a summary line to a file with mode 0o600 and read it back to prove the mode stuck.

> **Hint**: os.Stat returns the error you feed to errors.Is; os.FileInfo has Size and Mode.


---

© LearnSome.tech · support@iwantto.learnsome.tech
