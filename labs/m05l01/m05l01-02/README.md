# m05l01-02: Arguments and environment

You can read arguments and environment, read and write files safely, tell a missing file from a broken one, and set the exit code your pipeline branches on.

This is the front door of any tool. o s dot Args is a slice of strings, and the first entry is the program itself, so almost always you slice it off. If the caller has not given us enough to work with, we print usage to standard error and leave with a non zero status. Configuration comes from the environment next. Getenv hands back an empty string for a variable that is not set, which is fine whenever empty is a sensible default. When you need to tell empty apart from absent, which is the usual case for a feature flag, reach for LookupEnv: it returns the value and a boolean that says whether it was there at all. Run it with a variable in front of the command and watch both arrive.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
