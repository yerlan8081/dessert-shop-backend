FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN chmod +x wait-for-it.sh

RUN go build -o product-service ./microservices/product-service/main.go

FROM alpine:latest
WORKDIR /root/

COPY --from=builder /app/product-service .

EXPOSE 8082

CMD ["./wait-for-it.sh", "db:5432", "--", "./product-service"]
