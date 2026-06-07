FROM golang:1.22-alpine

WORKDIR /app

# Copy go modules first (better caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy entire project
COPY . .

# Build the API
RUN go build -o api ./cmd/api

EXPOSE 8080

CMD ["./api"]