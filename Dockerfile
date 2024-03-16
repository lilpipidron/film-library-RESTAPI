FROM golang:latest AS builder

WORKDIR /app

COPY . /app

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o vk-godeveloper-task ./cmd/vk-godeveloper-task/vk-godeveloper-task.go

EXPOSE 8080

CMD ["./vk-godeveloper-task"]