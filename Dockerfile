# --- Stage 1: Build ---
# Use an official Go image as a builder.
FROM golang:1.24.5-alpine AS builder

# Set the working directory inside the container.
WORKDIR /app

# Install swag binary
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Copy go.mod and go.sum files to download dependencies.
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code.
COPY . .

# Generate swagger docs right inside the container
RUN swag init -g ./cmd/api/main.go -o ./docs

# Build the application. CGO_ENABLED=0 is important for a static binary.
# -o /app/go-filehub specifies the output path for the binary.
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/go-filehub ./cmd/api

# --- Stage 2: Final Image ---
# Use a minimal, non-root base image for security.
FROM alpine:3.19

# Set the working directory.
WORKDIR /app

# Copy the compiled binary AND the generated docs from the builder stage.
COPY --from=builder /app/go-filehub .
COPY --from=builder /app/docs ./docs

# Expose the port the app runs on.
EXPOSE 8080

# The command to run when the container starts.
CMD ["./go-filehub"]