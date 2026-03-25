# Stage 1: Build
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copiar go.mod e go.sum
COPY go.mod go.sum ./
RUN go mod download

# Copiar código fonte
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o api ./cmd/main.go

# Stage 2: Runtime
FROM alpine:3.18

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copiar binário
COPY --from=builder /app/api .

# Copiar migrations (opcional, se rodar migrations pela app)
COPY --from=builder /app/migrations ./migrations

EXPOSE 8080

CMD ["./api"]