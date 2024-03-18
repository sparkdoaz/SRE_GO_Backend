# 第一階段：建立階段
# 選擇基礎的 image
FROM golang:1.22.1-alpine as builder
# 設置工作目錄
WORKDIR /app
# 複製 Go mod 和 dependence list
COPY go.mod ./
COPY go.sum ./
# 下載 dependence
RUN go mod download
# 複製資料（會根據 .dockerignore 來進行略過）
COPY . .
# 去 Build 程式叫做 main
RUN go build -o main .

# 第二階段：運行階段
# 選擇一個輕量級的基礎映像 要跟 builder 一樣
FROM golang:1.22.1-alpine
# 設置工作目錄
WORKDIR /app
# 從建立階段複製構建出來的二進制文件
COPY --from=builder /app/main /app/main
COPY --from=builder /app/.env /app/.env
# 設置環境變數
ENV CGO_ENABLED=0 \
  GOOS=linux
RUN chmod -R 777 /app && mkdir /app/log
RUN apk --no-cache add curl
# 提供描述告知會使用 3000 port
EXPOSE 3000
# 運行程式
CMD ["./main"]
