# —— 构建阶段 ——
FROM golang:1.23-alpine AS builder-user
WORKDIR /src

# 1. 拷贝根目录 go.mod/go.sum
COPY go.mod go.sum ./
RUN go mod download

# 2. 拷贝整个 microservices 目录
COPY microservices ./microservices

# 3. 编译 user-service
WORKDIR /src/microservices/user-service
RUN CGO_ENABLED=0 GOOS=linux go build -o user-service

# —— 运行阶段 ——
FROM alpine:3.18
WORKDIR /app

COPY --from=builder-user /src/microservices/user-service/user-service ./user-service
RUN chmod +x ./user-service

EXPOSE 8081
ENTRYPOINT ["./user-service"]
