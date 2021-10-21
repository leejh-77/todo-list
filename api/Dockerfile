FROM golang:alpine

ENV GO111MODULE=on \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY . .

RUN go build -o main .

ENTRYPOINT ["/app/main"]