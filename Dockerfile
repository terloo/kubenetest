FROM golang:1.19 as builder
ENV GOPROXY=https://goproxy.cn,direct
WORKDIR /build-dir
ADD ./go.mod /build-dir/go.mod
RUN go mod download
ADD . /build-dir
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o kubenetest ./cmd/

FROM alpine
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=builder /build-dir/kubenetest /kubenetest
ENTRYPOINT ["/kubenetest"]