# m06l04-03: Portability still has boundaries

You can build a portable Go binary with reproducible flags and inspect the result before shipping it.

Static does not mean every dependency disappears. Cgo may require native libraries, and a binary built for one architecture will not run on another. Embed assets that the program needs at runtime, and test the built artifact in the same kind of image that production will use. The release boundary is the binary plus the assumptions it carries.

- Cgo dependencies may need native libraries
- Architecture must match the runtime
- Embed assets when the program needs them
- Test the binary in its target image

Notes or exercise panel; no executable listing.
