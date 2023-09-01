# 基础镜像
FROM golang:latest AS builder

# 设置工作目录
WORKDIR /app

# 复制 Go 文件到容器内部
COPY proxy/8000.go .
COPY proxy/8001.go .
COPY proxy/8002.go .

# 编译 Go 文件
RUN go run 8000.go
RUN go run 8001.go
RUN go run 8002.go

# 第二个阶段，用于配置 Nginx
FROM nginx:latest

# 复制本地的 Nginx 配置文件到容器内部
COPY conf/nginx.conf /etc/nginx/nginx.conf 
COPY conf/conf.d /etc/nginx/conf.d 
COPY log /var/log/nginx 

# 暴露 Nginx 的默认端口
EXPOSE 80

# 启动 Nginx 服务
CMD ["nginx", "-g", "daemon off;"]