# ------------------------
# Stage 1: Build
# ------------------------
FROM golang:1.22.5-alpine AS build
WORKDIR /app

# Install git for module dependencies
RUN apk add --no-cache git

# Copy module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the binary from main.go
RUN go build -o app main.go

# ------------------------
# Stage 2: Production
# ------------------------
FROM alpine:latest
WORKDIR /root/

# Copy binary from build stage
COPY --from=build /app/app .

# Expose port (change if needed)
EXPOSE 8080

# Run the app
CMD ["./app"]
