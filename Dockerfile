FROM golang:alpine3.16 AS builder
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://mirrors.aliyun.com/goproxy/,direct
WORKDIR /build
COPY . .
RUN go build -o app .
FROM scratch
COPY --from=builder /build/app /
ENTRYPOINT ["/app"]