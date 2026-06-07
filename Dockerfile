FROM golang:1.25.1-alpine AS builder

WORKDIR /app

COPY Server/MuchToDo/go.mod Server/MuchToDo/go.sum ./
WORKDIR /app

RUN go mod download

COPY Server/MuchToDo/ .

RUN go build -o api ./cmd/api

FROM alpine:3.19

WORKDIR /app

COPY --from=builder /app/api .

EXPOSE 8080

CMD ["./api"]