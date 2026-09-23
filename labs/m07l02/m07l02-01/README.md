# m07l02-01: One worker per target

You can run checks concurrently, collect one result per target, and preserve a useful report order.

Each health check is independent, so the project can start one worker per target. A result channel transfers ownership back to the coordinator, and a wait group tells the coordinator when every worker has finished. Keep the target name inside each result so a fast response cannot become detached from the target it describes.

- Start independent checks concurrently
- Send results through a channel
- Wait for every worker
- Keep target identity in the result

Notes or exercise panel; no executable listing.
