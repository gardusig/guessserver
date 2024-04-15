FROM golang:1.21-alpine
WORKDIR /app
COPY . .
RUN go mod tidy

CMD ["go", "run", "cmd/main.go"]
