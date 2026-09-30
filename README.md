<p>
  <a href="https://learnsome.tech/courses/go-course">
    <picture>
      <source media="(prefers-color-scheme: dark)" srcset=".github/assets/wordmark-inverse.svg">
      <img src=".github/assets/wordmark.svg" alt="LearnSome.tech" width="260">
    </picture>
  </a>
</p>

# Concurrent Systems Programming with Go

**Goroutines, Channels, Memory Profiling & High-Throughput APIs**

7 modules, 35 lessons: The Go Toolchain and Fundamentals; Types, Structs, and Interfaces; Errors and Testing; Concurrency and Context; The Standard Library for Tooling; Building CLIs and Services; Project: End to End Health Checker. Intermediate level, about 2 hours.

This repository holds the labs of the LearnSome.tech course [Concurrent Systems Programming with Go](https://learnsome.tech/courses/go-course): each lab's starter files, a README with the goal, the steps and the expected output, and `./check`, which tests your work the way the site does.

## Start

[![Open in GitHub Codespaces](https://github.com/codespaces/badge.svg)](https://codespaces.new/learnsome-tech/go-labs?quickstart=1)

- **Codespaces:** the badge opens this repository in a dev container with Python 3.14.7 and Go 1.27.1, as in the site's lab sandbox.
- **On your machine:**

  ```sh
  git clone https://github.com/learnsome-tech/go-labs.git
  cd go-labs
  ./check m01l01-02
  ```

  You need Python 3 for `./check`, and for the labs themselves Python 3.14.7 and Go 1.27.1. Other versions mostly work, but only the sandbox's versions are sure to print what the site prints. VS Code's Dev Containers extension builds the same container as Codespaces (x86-64).

## Doing a lab

1. Open the lesson on LearnSome.tech and the lab folder beside it: `labs/<lesson>/<lab>/`. The lab README has the goal, the steps and the expected output.
2. Work in the lab's `starter/` folder.
3. From the repository root, run `./check <lab>` (for example `./check m01l01-02`), or `./check <lesson>` for all labs of a lesson, or `./check --all`. `./check --list` shows every lab and how it is checked.

`./check` runs your starter the way the site's lab sandbox does: in a scratch copy that is its working directory and `HOME`, with `LANG=C.UTF-8`, `TZ=UTC`, `input.txt` on standard input, 10 seconds and 256 KiB of output per stream. It then compares the output with the site's own rules, so a pass here is a pass on the site.

| Check | What `./check` does | Labs |
| --- | --- | --- |
| Graded | Runs the program and compares its output with `expected.txt`. | 62 |
| Runs, not graded | Runs the program and shows its output; the site gives no pass or fail, and the lab README says why. | 17 |
| Read along | Nothing to run here: the site shows the listing read-only, and the lab README says honestly what it needs (Docker, a cluster, a cloud account...). | 7 |

## What is published, and what is not

Every lab's starter is the code the lesson shows on screen, which is also what the lab editor on the site opens with. Where that code is the whole program, such as a recorded shell session or a script from the video, it is published as it is: it is the lesson content. Nothing beyond the lesson is published. There are no reference solutions and no answers to the lesson exercises, and nothing the site keeps private.

Pro lessons' labs are here as starters too. LearnSome.tech runs and grades your labs in its sandbox, hosts the videos and keeps your progress; running and grading a Pro lab on the site needs Pro.

## Modules and lessons

### Module 1: The Go Toolchain and Fundamentals

| # | Lesson | Labs | Access |
| --- | --- | --- | --- |
| 1.1 | [What Go Is and Why It Matters](https://learnsome.tech/learn/go-course/m01l01) | [3 labs](labs/m01l01/) | Free |
| 1.2 | [Installing the Go Toolchain](https://learnsome.tech/learn/go-course/m01l02) | [2 labs](labs/m01l02/) | Free |
| 1.3 | [Go Workspaces and Modules](https://learnsome.tech/learn/go-course/m01l03) | [2 labs](labs/m01l03/) | Free |
| 1.4 | [Your First Program and go run](https://learnsome.tech/learn/go-course/m01l04) | [4 labs](labs/m01l04/) | Free |
| 1.5 | [Variables, Typing, and Zero Values](https://learnsome.tech/learn/go-course/m01l05) | [4 labs](labs/m01l05/) | Pro |

### Module 2: Types, Structs, and Interfaces

| # | Lesson | Labs | Access |
| --- | --- | --- | --- |
| 2.1 | [Pointers and Memory](https://learnsome.tech/learn/go-course/m02l01) | [3 labs](labs/m02l01/) | Pro |
| 2.2 | [Arrays, Slices, and Maps](https://learnsome.tech/learn/go-course/m02l02) | [3 labs](labs/m02l02/) | Pro |
| 2.3 | [Structs and Methods](https://learnsome.tech/learn/go-course/m02l03) | [3 labs](labs/m02l03/) | Pro |
| 2.4 | [Interfaces and Why They Are Implicit](https://learnsome.tech/learn/go-course/m02l04) | [3 labs](labs/m02l04/) | Pro |
| 2.5 | [Type Assertions and the Empty Interface](https://learnsome.tech/learn/go-course/m02l05) | [3 labs](labs/m02l05/) | Pro |

### Module 3: Errors and Testing

| # | Lesson | Labs | Access |
| --- | --- | --- | --- |
| 3.1 | [Errors as Values](https://learnsome.tech/learn/go-course/m03l01) | [3 labs](labs/m03l01/) | Pro |
| 3.2 | [Wrapping Errors and errors.Is](https://learnsome.tech/learn/go-course/m03l02) | [3 labs](labs/m03l02/) | Pro |
| 3.3 | [Panics and Recover](https://learnsome.tech/learn/go-course/m03l03) | [3 labs](labs/m03l03/) | Pro |
| 3.4 | [The Standard testing Package](https://learnsome.tech/learn/go-course/m03l04) | [3 labs](labs/m03l04/) | Pro |
| 3.5 | [Writing Table Driven Tests](https://learnsome.tech/learn/go-course/m03l05) | [3 labs](labs/m03l05/) | Pro |

### Module 4: Concurrency and Context

| # | Lesson | Labs | Access |
| --- | --- | --- | --- |
| 4.1 | [Goroutines and Concurrent Execution](https://learnsome.tech/learn/go-course/m04l01) | [2 labs](labs/m04l01/) | Pro |
| 4.2 | [Channels and Synchronization](https://learnsome.tech/learn/go-course/m04l02) | [2 labs](labs/m04l02/) | Pro |
| 4.3 | [The select Statement](https://learnsome.tech/learn/go-course/m04l03) | [2 labs](labs/m04l03/) | Pro |
| 4.4 | [The context Package](https://learnsome.tech/learn/go-course/m04l04) | [2 labs](labs/m04l04/) | Pro |
| 4.5 | [Context Cancellation and Timeouts](https://learnsome.tech/learn/go-course/m04l05) | [2 labs](labs/m04l05/) | Pro |

### Module 5: The Standard Library for Tooling

| # | Lesson | Labs | Access |
| --- | --- | --- | --- |
| 5.1 | [Files and the os Package](https://learnsome.tech/learn/go-course/m05l01) | [5 labs](labs/m05l01/) | Pro |
| 5.2 | [Streams and the io Package](https://learnsome.tech/learn/go-course/m05l02) | [2 labs](labs/m05l02/) | Pro |
| 5.3 | [The time Package](https://learnsome.tech/learn/go-course/m05l03) | [2 labs](labs/m05l03/) | Pro |
| 5.4 | [Parsing JSON with encoding/json](https://learnsome.tech/learn/go-course/m05l04) | [2 labs](labs/m05l04/) | Pro |
| 5.5 | [Calling APIs with net/http Client](https://learnsome.tech/learn/go-course/m05l05) | [2 labs](labs/m05l05/) | Pro |

### Module 6: Building CLIs and Services

| # | Lesson | Labs | Access |
| --- | --- | --- | --- |
| 6.1 | [Parsing Arguments with flag](https://learnsome.tech/learn/go-course/m06l01) | [4 labs](labs/m06l01/) | Pro |
| 6.2 | [Serving HTTP with net/http Server](https://learnsome.tech/learn/go-course/m06l02) | [2 labs](labs/m06l02/) | Pro |
| 6.3 | [Graceful Shutdown on SIGTERM](https://learnsome.tech/learn/go-course/m06l03) | [2 labs](labs/m06l03/) | Pro |
| 6.4 | [Building Static Binaries](https://learnsome.tech/learn/go-course/m06l04) | [1 lab](labs/m06l04/) | Pro |
| 6.5 | [Shipping in Scratch and Distroless Images](https://learnsome.tech/learn/go-course/m06l05) | – | Pro |

### Module 7: Project: End to End Health Checker

| # | Lesson | Labs | Access |
| --- | --- | --- | --- |
| 7.1 | [Project Setup and Configuration](https://learnsome.tech/learn/go-course/m07l01) | [2 labs](labs/m07l01/) | Pro |
| 7.2 | [Executing Concurrent Checks](https://learnsome.tech/learn/go-course/m07l02) | [2 labs](labs/m07l02/) | Pro |
| 7.3 | [Handling Timeouts and Errors](https://learnsome.tech/learn/go-course/m07l03) | [2 labs](labs/m07l03/) | Pro |
| 7.4 | [Exporting Results as JSON](https://learnsome.tech/learn/go-course/m07l04) | [2 labs](labs/m07l04/) | Pro |
| 7.5 | [Containerizing and Testing the Final Build](https://learnsome.tech/learn/go-course/m07l05) | [1 lab](labs/m07l05/) | Pro |

**Free** lessons are open to anyone with a free LearnSome.tech account; **Pro** lessons need a Pro membership to watch, run and grade on the site.

## Licence

- **Code** (starter files, `check` and `.learnsome/`, the dev container and the workflows) is under the [MIT licence](LICENSE).
- **Written text** (the READMEs, lab instructions, lesson text, exercises and questions) is under [CC BY-NC-SA 4.0](LICENSE-text.md): share and adapt it with attribution to LearnSome.tech, not commercially, under the same licence.
- The LearnSome.tech name and logo are not covered by either licence.

## Contributing and security

This repository is generated from the course. Report a broken lab or a content error [as an issue](../../issues/new/choose); see [CONTRIBUTING.md](CONTRIBUTING.md). Security reports go to [SECURITY.md](SECURITY.md).

© 2026 LearnSome.tech
