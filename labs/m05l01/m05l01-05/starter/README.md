# m05l01-05: Missing, or broken?

You can read arguments and environment, read and write files safely, tell a missing file from a broken one, and set the exit code your pipeline branches on.

Here is the distinction that separates a careful tool from a noisy one. Stat asks the filesystem about a path without opening it, and the error it returns is not a single thing. A missing file and a permission denied are completely different situations: one means create it, the other means stop and tell somebody. So you never compare the error to a string. You ask errors dot Is whether it matches the sentinel value o s dot ErrNotExist, and you leave a default branch for everything else, which gives you three outcomes rather than two. Run it over one file that exists and one that does not, and you get exactly the two answers you would expect.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
