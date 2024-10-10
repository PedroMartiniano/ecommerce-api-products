FROM cgr.dev/chainguard/go:latest as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/api

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/cmd/api/.env .
COPY --from=builder /app/internal/migrations migrations


CMD ["./main"]