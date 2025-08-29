# ------------------------
# Stage 1: Build
# ------------------------
FROM golang:1.22.5-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o app ./...

# ------------------------
# Stage 2: Production
# ------------------------
FROM alpine:latest
WORKDIR /root/
COPY --from=build /app/app .
EXPOSE 8080
CMD ["./app"]