# m02l02-01: Three collection shapes

You can choose arrays, slices, and maps for predictable collections and handle their zero values safely.

Go gives you three collection shapes with different promises. An array has a length fixed in its type. A slice is a flexible view over an underlying array, and it is the collection you will use most often. A map pairs keys with values and is useful for lookup by name. The choice matters at an API boundary: an array promises a fixed count, a slice promises an ordered group, and a map promises lookup. Start with the promise your caller needs, not with a clever container.

- An array has a fixed length
- A slice is a flexible view over an array
- A map pairs keys with values
- Choose the smallest promise your API needs

Notes or exercise panel; no executable listing.
