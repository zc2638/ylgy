FROM golang:1.19 as builder
ENV GOPROXY=https://goproxy.cn,https://goproxy.io,direct
ENV GO111MODULE=on
ENV GOCACHE=/go/pkg/.cache/go-build

WORKDIR /work
ADD . .
RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o /usr/local/bin/ylgy github.com/zc2638/ylgy/cmd/ylgy

FROM alpine:3.6 as alpine
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories && \
    apk update && \
    apk add -U --no-cache ca-certificates tzdata

FROM alpine:3.6
MAINTAINER zc
LABEL maintainer="zc" \
    email="zc2638@qq.com"

ENV TZ="Asia/Shanghai"

COPY --from=alpine /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/local/bin/ylgy /usr/local/bin/ylgy

RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo $TZ > /etc/timezone

WORKDIR /work
CMD ["ylgy"]