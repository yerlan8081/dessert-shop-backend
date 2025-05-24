# —— 构建阶段 ——
FROM golang:1.23-alpine AS builder-product
WORKDIR /src

# 1. 拷贝根目录 go.mod/go.sum，拉取依赖
COPY go.mod go.sum ./
RUN go mod download

# 2. 拷贝整个 microservices 目录（保留目录结构）
COPY microservices ./microservices

# 3. 切到服务目录编译
WORKDIR /src/microservices/product-service
RUN CGO_ENABLED=0 GOOS=linux go build -o product-service

# —— 运行阶段 ——
FROM alpine:3.18
WORKDIR /app

# 只拷贝编译好的二进制
COPY --from=builder-product /src/microservices/product-service/product-service ./product-service
RUN chmod +x ./product-service

EXPOSE 8082
ENTRYPOINT ["./product-service"]
