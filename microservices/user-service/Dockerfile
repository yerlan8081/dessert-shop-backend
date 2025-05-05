# 选择官方 Go 镜像作为构建阶段
FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN chmod +x wait-for-it.sh

# 编译为静态二进制文件
RUN go build -o user-service ./microservices/user-service/main.go

# 运行阶段
FROM alpine:latest
WORKDIR /root/

COPY --from=builder /app/user-service .

EXPOSE 8081

CMD ["./wait-for-it.sh", "db:5432", "--", "./user-service"]
