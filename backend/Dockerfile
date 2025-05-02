FROM golang:1.24-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY cmd/ ./cmd/
COPY internal/ ./internal/

RUN go build -o migrate cmd/db/migrate.go
RUN go build -o api cmd/api/main.go

ENTRYPOINT ["./api"]