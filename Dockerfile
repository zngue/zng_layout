FROM golang:1.23.2 as builder
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
ENV NACOS_HOST=rust_nacos
ENV NACOS_PORT=8848
ENV NACOS_NAMESPACE=develop
# 设置监听端口
ENV HTTP_PORT=16669
ENV DB_GROUP=common
# 数据库
ENV DB_DATABASE=blog
#  info debug warn error
ENV LOG_LEVEL=info
ENV LOG_URL=""
WORKDIR /build
#666
COPY . .
RUN go mod init  github.com/zngue/zng_layout
RUN go mod tidy
RUN cd ./cmd/zng_layout && GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -ldflags="-s -w" -installsuffix cgo -o appRun ./...

FROM alpine:latest as prod
RUN apk add --no-cache tzdata
ENV TZ=Asia/Shanghai
WORKDIR /go_run
COPY --from=builder /build/cmd/zng_layout/appRun .
EXPOSE  $HTTP_PORT
ENTRYPOINT ["./appRun"]