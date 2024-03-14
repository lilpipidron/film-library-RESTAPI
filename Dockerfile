# Этап сборки приложения
FROM golang:latest AS builder

WORKDIR /app

COPY . /app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o vk-godeveloper-task ./cmd/vk-godeveloper-task/vk-godeveloper-task.go

# Этап финальной сборки образа
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/vk-godeveloper-task .
CMD ["./vk-godeveloper-task"]

VOLUME film-library:/var/lib/postgresql/data