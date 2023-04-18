# STAGE-1: build stage
FROM golang:1.17-alpine3.15 AS build-env
RUN apk add build-base
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    go build -o main .

# STAGE-2: deploy stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/templates /app/templates
COPY --from=build-env /src/main /app/
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
RUN chown -R appuser:appgroup /app
USER appuser

ENTRYPOINT ./main