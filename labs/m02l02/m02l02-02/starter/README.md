# m02l02-02: Slices grow when you append

You can choose arrays, slices, and maps for predictable collections and handle their zero values safely.

A slice literal looks like a short list, and append is the operation you will use to grow it. Append returns a slice because the backing array may need to move, so assign the result back to the same variable. Length tells you how many elements are present. Indexing starts at zero, and a slice expression takes a half open range, which includes the first position and stops before the last position. The example adds one item, reads the middle item, and takes the first two items as another view.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
