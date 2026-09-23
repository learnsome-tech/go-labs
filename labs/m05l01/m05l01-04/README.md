# m05l01-04: Opening a file, and closing it

You can read arguments and environment, read and write files safely, tell a missing file from a broken one, and set the exit code your pipeline branches on.

When the file is a log, or anything you would rather not hold in memory, you open it instead. Open gives you two things: a file handle and an error, and until you have checked the error the handle is not yours to use. The defer line is the habit to build. It schedules Close for the moment this function returns, whichever path it takes, so a later early return cannot leak the descriptor. On a long running agent that leak is the bug that wakes you at three in the morning. Then we finish by asking the open file about itself: name, size in bytes and the permission mode, printed the way a listing would show it.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
