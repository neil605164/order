# 第一層基底
FROM golang:1.18.0-alpine

# 安裝 git
# go install air
RUN apk add git \
    && go install github.com/cosmtrek/air@v1.29.0

# 設定容器時區(美東)
RUN apk update \
    && apk add tzdata \
    && cp /usr/share/zoneinfo/America/Puerto_Rico /etc/localtime


# docker terminal 顯示 LOG
RUN mkdir -p /app/log/ \
    && ln -sf /dev/stdout /app/log/access.log \
    && ln -sf /dev/stdout /app/log/error.log
