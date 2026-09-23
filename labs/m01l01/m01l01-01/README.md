# m01l01-01: Why platform teams reach for Go

You can explain why platform and DevOps teams write their tooling in Go, compile and run a program with the go command, and describe exactly what the compiler hands you at the end of a build.

Docker, Kubernetes, Terraform, Prometheus and the cloud command line tools you use every day are all written in Go, and that is not a coincidence. Operations tooling has an awkward shape: it must run on a build agent, inside a container, on somebody's laptop and on a server you cannot install anything on. Go answers that by compiling to a single self contained executable with no runtime to install first. It also compiles fast enough that you stay in the loop of change something, check it, change it again. Those two properties, plus concurrency that is part of the language rather than a library you import, are why this is the language platform teams keep landing on.

- One self contained binary: copy it to a host and it runs
- Builds take seconds, so the edit and check loop stays tight
- Concurrency is in the language, not a library bolted on
- The tools you already operate are written in it

Notes or exercise panel; no executable listing.
