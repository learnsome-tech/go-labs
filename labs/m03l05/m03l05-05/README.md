# m03l05-05: Table testing a parser contract

You can express many input cases in one table test and get a useful name for every failure.

This table keeps both success and failure expectations beside their inputs. The empty row expects an error, while the service row expects a port value. The assertion first compares whether an error exists, then compares the value only for a successful case. Each subtest carries the row name, so a failure points to empty or service immediately. This is the same compact pattern you can use for configuration and command line validation.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
