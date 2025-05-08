FROM golang:1.24-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY cmd/ ./cmd/
COPY internal/ ./internal/
COPY migrations/ ./migrations/

RUN go build -o migrate cmd/db/migrate.go
RUN go build -o api cmd/api/main.go

RUN apk add --no-cache curl
HEALTHCHECK --interval=5s --timeout=5s --start-period=5s --retries=3 CMD [ "curl", "-f", "http://localhost:8080/healthz" ]

CMD ["./migrate", "./api"]
