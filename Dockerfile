ARG VERSION=latest
FROM golang:${VERSION} AS builder

WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o main .

# Stage 2: Development
FROM golang:${VERSION} AS dev

WORKDIR /app

# Install air for live reloading
RUN go install github.com/air-verse/air@latest

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Run air for live reloading
ENTRYPOINT ["air"]



# # Stage 3: Run
# FROM alpine:latest AS run

# WORKDIR /root/

# # Install necessary packages
# RUN apk add --no-cache libc6-compat

# # Copy the binary from the build stage
# COPY --from=builder /app/main .

# # Set environment variables (if needed)
# ENV PORT=3000

# # Expose the port
# EXPOSE 3000

# # Run the binary
# CMD ["/root/main"]
