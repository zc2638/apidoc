FROM golang:1.17 as builder

ENV GOPROXY=https://goproxy.cn,https://goproxy.io,direct

WORKDIR /work

RUN wget https://files-cdn.cnblogs.com/files/xiaochina/simsun.zip && \
    apt-get update && \
    apt-get install -y unzip && \
    unzip simsun.zip

ADD . .
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o /bin/apidoc github.com/zc2638/apidoc/cmd/apidoc

FROM alpine:3.10
MAINTAINER zc <zc2638@qq.com>

ENV LANG=C.utf8

RUN apk update && \
    apk add --no-cache wkhtmltopdf ttf-freefont fontconfig && \
    apk del --purge build-base libgfortran libpng-dev freetype-dev python3-dev && \
    rm -rvf /var/cache/apk/*

COPY --from=builder /work/simsun.ttc /usr/share/fonts/simsun.ttc
COPY --from=builder /bin/apidoc /bin/apidoc

ENTRYPOINT ["apidoc"]
