# golang镜像。
FROM golang:1.25.5-alpine

# cd 
WORKDIR /luckDice

# 复制项目代码到 Docker所有容器老巢/luckDice房间（容器内“房子/房间”结构）
COPY . .

# 下载依赖
RUN go mod tidy

# go build -output to luckDicerServer from main.go
RUN go build -o luckDiceServer main.go

# 暴露端口
EXPOSE 50051

# 启动服务
CMD ["./luckDiceServer"]
