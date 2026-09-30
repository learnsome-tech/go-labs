# m06l01-05: Environment fallback, for the container case

You can build a command line tool with the standard flag package that takes typed options, reads positional arguments, falls back to environment variables, and fails with a usage message and the right exit code.

Containers configure through the environment and people configure through flags, so a tool that lives in both worlds wants both. The pattern is one small helper: look the variable up, return it when it is set, otherwise return the built in default. Feed that helper into the flag declaration as the default value and a clean precedence order falls out for free. An explicit flag wins, because parsing overwrites the default. The environment wins when no flag was given. The compiled in default is the last resort. Use LookupEnv rather than Getenv when an empty string is a meaningful setting, because only LookupEnv tells you whether the variable was set at all. Let us run it twice: once with just the variable, then with a flag as well.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
