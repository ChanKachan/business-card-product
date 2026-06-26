# Сборка
FROM golang:1.26.1 as builder
WORKDIR /app
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY ./ /app
RUN GO111MODULE=auto CGO_ENABLED=0 GOOS=linux GOPROXY=https://proxy.golang.org go build -o businessCardProduct cmd/main.go

# Запуск
FROM alpine:3.20.3
WORKDIR /app
COPY --from=builder /app/businessCardProduct .
COPY --from=builder /app/migrations ./migrations
ENTRYPOINT [ "./businessCardProduct" ]