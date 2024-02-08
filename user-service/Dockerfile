# Dockerfile for user-service
FROM golang:1.17-alpine

WORKDIR /app
COPY . .

RUN go mod download
RUN go build cmd -o user-service .

EXPOSE 8080

CMD ["./user-service"]
