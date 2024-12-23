# Build stage
FROM golang:1.23.1-alpine AS builder
WORKDIR /app
RUN apk add --no-cache ca-certificates git
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/server ./cmd/main.go

# Final stage
FROM gcr.io/distroless/static-debian11
USER nonroot:nonroot
COPY --from=builder --chown=nonroot:nonroot /app/server /server
EXPOSE 8080
ENTRYPOINT ["/server"]