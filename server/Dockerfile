# Fetch
FROM golang:latest AS fetch-stage
WORKDIR /usr/src/matcha
COPY go.mod go.sum .
RUN go mod download

# Generate
FROM ghcr.io/a-h/templ:latest AS generate-stage
WORKDIR /usr/src/matcha
COPY --chown=65532:65532 . .
RUN ["templ", "generate"]

# Build
FROM golang:latest AS build-stage
WORKDIR /usr/src/matcha
COPY --from=generate-stage /usr/src/matcha .
RUN CGO_ENABLED=0 GOOS=linux go build -o /usr/bin/matcha

# Test
FROM build-stage AS test-stage
WORKDIR /usr/src/matcha
RUN go test -v ./...

# Deploy
FROM gcr.io/distroless/base-debian12 AS deploy-stage
COPY --from=build-stage /usr/bin/matcha /usr/bin/matcha
EXPOSE 80
USER nonroot:nonroot
CMD ["matcha"]
