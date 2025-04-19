FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o app ./cmd

FROM alpine:3.18
WORKDIR /app

# Copiamos binario y recursos
COPY --from=builder /app/app .
COPY --from=builder /app/templates/ ./templates

EXPOSE 8080

ENTRYPOINT ["./app"]
