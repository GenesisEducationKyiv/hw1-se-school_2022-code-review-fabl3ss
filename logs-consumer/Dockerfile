FROM golang:latest

LABEL maintainer="logs-consumer"

WORKDIR /logs-consumer

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o main src/cmd/main.go

RUN find . -name "*.go" -type f -delete

CMD ["./main"]