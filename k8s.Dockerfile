## 第一層低底
FROM golang:1.18.0-alpine AS build

RUN apk add git

# 複製原始碼
COPY . /go/src/order
WORKDIR /go/src/order

# 進行編譯(名稱為：runner)
RUN go build -o runner

# 運行golang 的基底
FROM alpine:3.9.5

# 設定容器時區(美東)
# RUN apk update \
#     && apk add tzdata \
#     && cp /usr/share/zoneinfo/Asia/Taipei /etc/localtime \
#     && apk del tzdata

COPY --from=build /go/src/order/runner /app/runner

WORKDIR /app

ENTRYPOINT [ "./runner" ]