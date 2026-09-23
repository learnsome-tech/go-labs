# m06l04-01: Static binaries simplify delivery

You can build a portable Go binary with reproducible flags and inspect the result before shipping it.

A Go program can compile into one self contained executable, which removes a runtime language installation from the target image. For a portable Linux build, disable cgo when your dependencies allow it and set the target operating system and architecture explicitly. Keep those choices in a build script or pipeline step so a release can be reproduced rather than remembered.

- Go can compile one self contained executable
- Disable cgo for portable builds
- Set the target operating system and architecture
- Reproducible flags belong in the build script

Notes or exercise panel; no executable listing.
