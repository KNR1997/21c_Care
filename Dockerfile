# ---------- Build Stage ----------
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Install git (needed for some Go dependencies)
RUN apk add --no-cache git

# Copy go mod files first for better caching
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app ./main.go


# ---------- Runtime Stage ----------
FROM alpine:3.20

WORKDIR /app

# Install CA certificates (needed for HTTPS like Google OIDC)
RUN apk add --no-cache ca-certificates

# Copy compiled binary
COPY --from=builder /app/app .

# Expose port
EXPOSE 7788

# Run binary
CMD ["./app"]
