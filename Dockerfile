FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
# 此处给一个镜像站环境变量,可选的
ENV GOPROXY=https://goproxy.cn,direct
RUN go mod download

COPY . .

# 确保是静态的,方便丢进 distroless 里
ENV CGO_ENABLED=false
ENV GOOS=linux
RUN  go build -ldflags="-w -s" -o lookup-cli

FROM gcr.io/distroless/base-debian10

WORKDIR /

# 此处只要构建成品
COPY --from=builder /app/lookup-cli /lookup-cli

USER nonroot:nonroot
ENTRYPOINT ["/lookup-cli"]
