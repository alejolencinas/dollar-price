# Stage 1: Build
FROM golang:1.23 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o dollar-price ./cmd/server

# Stage 2: Run
FROM gcr.io/distroless/base-debian12

WORKDIR /app
COPY --from=builder /app/dollar-price .

# Expose port 8080
EXPOSE 8080

CMD ["./dollar-price"]
