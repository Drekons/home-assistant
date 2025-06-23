# Build stage
FROM golang:1.24.2-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/server
RUN CGO_ENABLED=0 GOOS=linux go build -o cli ./cmd/cli

# Run stage
FROM alpine:latest

WORKDIR /root/

COPY --from=build /app/main .
COPY --from=build /app/cli .
COPY --from=build /app/.env .

EXPOSE 8000

CMD ["./main"]