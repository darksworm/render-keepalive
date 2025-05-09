# Build stage
FROM golang:1.24.2-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum (if they exist)
COPY . .

# Build the Go app with static linking
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# Final stage
FROM scratch

# Copy the binary from the builder stage
COPY --from=builder /app/app /app

# Set the entrypoint
ENTRYPOINT ["/app"]
