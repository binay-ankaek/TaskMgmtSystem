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

# Build the binary from cmd/main.go
RUN go build -o app ./cmd

# ------------------------
# Stage 2: Production
# ------------------------
FROM alpine:latest
WORKDIR /root/

# Copy binary from build stage
COPY --from=build /app/app .

EXPOSE 8080
CMD ["./app"]
