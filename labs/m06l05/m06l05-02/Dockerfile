# Concurrent Systems Programming with Go — lesson m06l05 — Shipping in Scratch and Distroless Images
# https://learnsome.tech/courses/go-course/watch?lesson=m06l05
# © LearnSome.tech
FROM golang:alpine AS build
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 go build -o /out/app .

FROM scratch
COPY --from=build /out/app /app
ENTRYPOINT ["/app"]
