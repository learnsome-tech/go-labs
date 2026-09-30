# m05l05-02: A request to a local server

You can issue an HTTP request with a context, check the response, and close the body on every path.

The request uses the standard method name and a local address that is expected to refuse the connection on this example machine. Do returns either a response or an error, and the boolean print confirms that the connection failed. Production code should pass a context, configure a client timeout, and wrap this error with the operation and target before returning it to a caller.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
