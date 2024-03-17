FROM golang:latest AS builder

WORKDIR /app

COPY . /app

RUN apt-get update
RUN apt-get -y install postgresql-client

RUN chmod +x wait-for-postgres.sh

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o vk-godeveloper-task ./cmd/vk-godeveloper-task/vk-godeveloper-task.go

EXPOSE 8080

CMD ["./vk-godeveloper-task"]