<img src="https://learnsome.tech/logo.png" width="48" alt="LearnSome.tech">

# Concurrent Systems Programming with Go

7 modules, 35 lessons: The Go Toolchain and Fundamentals; Types, Structs, and Interfaces; Errors and Testing; Concurrency and Context; The Standard Library for Tooling; Building CLIs and Services; Project: End to End Health Checker.

## Watch and read

- **Course page**: [https://learnsome.tech/courses/go-course](https://learnsome.tech/courses/go-course)
- **Video player**: [https://learnsome.tech/courses/go-course/watch](https://learnsome.tech/courses/go-course/watch)
- **Handbook PDF**: [https://learnsome.tech/handbooks/go/book.pdf](https://learnsome.tech/handbooks/go/book.pdf)
- **On-site handbook**: [https://learnsome.tech/courses/go-course/book](https://learnsome.tech/courses/go-course/book)

## What is in this repository

This repository contains code artifacts, exercises and reference files for the lessons in this course.
35 lessons include a `labs/<lessonId>/` folder.
Each folder is named after the lesson identifier (e.g. `labs/m01l01/`) and contains the
artifact files shown in the course video, an `EXERCISES.md` with hands-on tasks, and
sub-directories named by artifact reference (e.g. `m01l01-02/`).

## Lessons

| # | Lesson | Watch | Labs | Handbook |
|---|--------|-------|------|----------|
| | **The Go Toolchain and Fundamentals** | | | |
| 1 | What Go Is and Why It Matters | [▶](https://learnsome.tech/courses/go-course/watch?lesson=m01l01) | [labs/m01l01/](labs/m01l01/) | [§](https://learnsome.tech/courses/go-course/book#lesson-1-1) |
| 2 | Installing the Go Toolchain | [▶](https://learnsome.tech/courses/go-course/watch?lesson=m01l02) | [labs/m01l02/](labs/m01l02/) | [§](https://learnsome.tech/courses/go-course/book#lesson-1-2) |
| 3 | Go Workspaces and Modules | [▶](https://learnsome.tech/courses/go-course/watch?lesson=m01l03) | [labs/m01l03/](labs/m01l03/) | [§](https://learnsome.tech/courses/go-course/book#lesson-1-3) |
| 4 | Your First Program and go run | [▶](https://learnsome.tech/courses/go-course/watch?lesson=m01l04) | [labs/m01l04/](labs/m01l04/) | [§](https://learnsome.tech/courses/go-course/book#lesson-1-4) |
| 5 | Variables, Typing, and Zero Values | [▶](https://learnsome.tech/courses/go-course/watch?lesson=m01l05) | [labs/m01l05/](labs/m01l05/) | [§](https://learnsome.tech/courses/go-course/book#lesson-1-5) |
| | **Types, Structs, and Interfaces** | | | |
| 6 | Pointers and Memory | [▶](https://learnsome.tech/courses/go-course/watch?lesson=m02l01) | [labs/m02l01/](labs/m02l01/) | [§](https://learnsome.tech/courses/go-course/book#lesson-2-1) |
| 7 | Arrays, Slices, and Maps | [▶](https://learnsome.tech/courses/go-course/watch?lesson=m02l02) | [labs/m02l02/](labs/m02l02/) | [§](https://learnsome.tech/courses/go-course/book#lesson-2-2) |
| 8 | Structs and Methods | [▶](https://learnsome.tech/courses/go-course/watch?lesson=m02l03) | [labs/m02l03/](labs/m02l03/) | [§](https://learnsome.tech/courses/go-course/book#lesson-2-3) |
| 9 | Interfaces and Why They Are Implicit | [▶](https://learnsome.tech/courses/go-course/watch?lesson=m02l04) | [labs/m02l04/](labs/m02l04/) | [§](https://learnsome.tech/courses/go-course/book#lesson-2-4) |
| 10 | Type Assertions and the Empty Interface | [▶](https://learnsome.tech/courses/go-course/watch?lesson=m02l05) | [labs/m02l05/](labs/m02l05/) | [§](https://learnsome.tech/courses/go-course/book#lesson-2-5) |
| | **Errors and Testing** | | | |
| 11 | Errors as Values | [▶](https://learnsome.tech/courses/go-course/watch?lesson=m03l01) | [labs/m03l01/](labs/m03l01/) | [§](https://learnsome.tech/courses/go-course/book#lesson-3-1) |
| 12 | Wrapping Errors and errors.Is | [▶](https://learnsome.tech/courses/go-course/watch?lesson=m03l02) | [labs/m03l02/](labs/m03l02/) | [§](https://learnsome.tech/courses/go-course/book#lesson-3-2) |
| 13 | Panics and Recover | [▶](https://learnsome.tech/courses/go-course/watch?lesson=m03l03) | [labs/m03l03/](labs/m03l03/) | [§](https://learnsome.tech/courses/go-course/book#lesson-3-3) |
| 14 | The Standard testing Package | [▶](https://learnsome.tech/courses/go-course/watch?lesson=m03l04) | [labs/m03l04/](labs/m03l04/) | [§](https://learnsome.tech/courses/go-course/book#lesson-3-4) |
| 15 | Writing Table Driven Tests | [▶](https://learnsome.tech/courses/go-course/watch?lesson=m03l05) | [labs/m03l05/](labs/m03l05/) | [§](https://learnsome.tech/courses/go-course/book#lesson-3-5) |
| | **Concurrency and Context** | | | |
| 16 | Goroutines and Concurrent Execution | [▶](https://learnsome.tech/courses/go-course/watch?lesson=m04l01) | [labs/m04l01/](labs/m04l01/) | [§](https://learnsome.tech/courses/go-course/book#lesson-4-1) |
| 17 | Channels and Synchronization | [▶](https://learnsome.tech/courses/go-course/watch?lesson=m04l02) | [labs/m04l02/](labs/m04l02/) | [§](https://learnsome.tech/courses/go-course/book#lesson-4-2) |
| 18 | The select Statement | [▶](https://learnsome.tech/courses/go-course/watch?lesson=m04l03) | [labs/m04l03/](labs/m04l03/) | [§](https://learnsome.tech/courses/go-course/book#lesson-4-3) |
| 19 | The context Package | [▶](https://learnsome.tech/courses/go-course/watch?lesson=m04l04) | [labs/m04l04/](labs/m04l04/) | [§](https://learnsome.tech/courses/go-course/book#lesson-4-4) |
| 20 | Context Cancellation and Timeouts | [▶](https://learnsome.tech/courses/go-course/watch?lesson=m04l05) | [labs/m04l05/](labs/m04l05/) | [§](https://learnsome.tech/courses/go-course/book#lesson-4-5) |
| | **The Standard Library for Tooling** | | | |
| 21 | Files and the os Package | [▶](https://learnsome.tech/courses/go-course/watch?lesson=m05l01) | [labs/m05l01/](labs/m05l01/) | [§](https://learnsome.tech/courses/go-course/book#lesson-5-1) |
| 22 | Streams and the io Package | [▶](https://learnsome.tech/courses/go-course/watch?lesson=m05l02) | [labs/m05l02/](labs/m05l02/) | [§](https://learnsome.tech/courses/go-course/book#lesson-5-2) |
| 23 | The time Package | [▶](https://learnsome.tech/courses/go-course/watch?lesson=m05l03) | [labs/m05l03/](labs/m05l03/) | [§](https://learnsome.tech/courses/go-course/book#lesson-5-3) |
| 24 | Parsing JSON with encoding/json | [▶](https://learnsome.tech/courses/go-course/watch?lesson=m05l04) | [labs/m05l04/](labs/m05l04/) | [§](https://learnsome.tech/courses/go-course/book#lesson-5-4) |
| 25 | Calling APIs with net/http Client | [▶](https://learnsome.tech/courses/go-course/watch?lesson=m05l05) | [labs/m05l05/](labs/m05l05/) | [§](https://learnsome.tech/courses/go-course/book#lesson-5-5) |
| | **Building CLIs and Services** | | | |
| 26 | Parsing Arguments with flag | [▶](https://learnsome.tech/courses/go-course/watch?lesson=m06l01) | [labs/m06l01/](labs/m06l01/) | [§](https://learnsome.tech/courses/go-course/book#lesson-6-1) |
| 27 | Serving HTTP with net/http Server | [▶](https://learnsome.tech/courses/go-course/watch?lesson=m06l02) | [labs/m06l02/](labs/m06l02/) | [§](https://learnsome.tech/courses/go-course/book#lesson-6-2) |
| 28 | Graceful Shutdown on SIGTERM | [▶](https://learnsome.tech/courses/go-course/watch?lesson=m06l03) | [labs/m06l03/](labs/m06l03/) | [§](https://learnsome.tech/courses/go-course/book#lesson-6-3) |
| 29 | Building Static Binaries | [▶](https://learnsome.tech/courses/go-course/watch?lesson=m06l04) | [labs/m06l04/](labs/m06l04/) | [§](https://learnsome.tech/courses/go-course/book#lesson-6-4) |
| 30 | Shipping in Scratch and Distroless Images | [▶](https://learnsome.tech/courses/go-course/watch?lesson=m06l05) | [labs/m06l05/](labs/m06l05/) | [§](https://learnsome.tech/courses/go-course/book#lesson-6-5) |
| | **Project: End to End Health Checker** | | | |
| 31 | Project Setup and Configuration | [▶](https://learnsome.tech/courses/go-course/watch?lesson=m07l01) | [labs/m07l01/](labs/m07l01/) | [§](https://learnsome.tech/courses/go-course/book#lesson-7-1) |
| 32 | Executing Concurrent Checks | [▶](https://learnsome.tech/courses/go-course/watch?lesson=m07l02) | [labs/m07l02/](labs/m07l02/) | [§](https://learnsome.tech/courses/go-course/book#lesson-7-2) |
| 33 | Handling Timeouts and Errors | [▶](https://learnsome.tech/courses/go-course/watch?lesson=m07l03) | [labs/m07l03/](labs/m07l03/) | [§](https://learnsome.tech/courses/go-course/book#lesson-7-3) |
| 34 | Exporting Results as JSON | [▶](https://learnsome.tech/courses/go-course/watch?lesson=m07l04) | [labs/m07l04/](labs/m07l04/) | [§](https://learnsome.tech/courses/go-course/book#lesson-7-4) |
| 35 | Containerizing and Testing the Final Build | [▶](https://learnsome.tech/courses/go-course/watch?lesson=m07l05) | [labs/m07l05/](labs/m07l05/) | [§](https://learnsome.tech/courses/go-course/book#lesson-7-5) |

## Exercises

Each lesson folder contains an `EXERCISES.md` with hands-on tasks drawn directly from the course material.
Open the file for a lesson to see the tasks and, where provided, hints.

---

© LearnSome.tech · support@iwantto.learnsome.tech
