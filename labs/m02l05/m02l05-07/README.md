# m02l05-07: Try it yourself

You can inspect an interface value safely, use any for truly mixed data, and avoid hiding useful types behind it.

Write a function that accepts any and returns a readable kind label. Use a type switch for strings, integers, and booleans, and return unknown for every other type. Call it with all four paths and print the labels. Then ask whether the function belongs at a data boundary or in the middle of your program. If the answer is the middle, consider replacing any with a small interface or a typed struct.

- Write a function that accepts any and returns a readable kind label.
- Handle strings, integers, and booleans with a type switch.
- Return unknown for every other type and test all four paths.

Notes or exercise panel; no executable listing.
