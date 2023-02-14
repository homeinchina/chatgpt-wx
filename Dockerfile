FROM golang:1.19-alpine

ENV api_key=""

RUN export GOPRIVATE=github.com/homeinchina/chatgpt-wx

WORKDIR /app

COPY . /app

RUN go mod download && go build -o server main.go

CMD ./server